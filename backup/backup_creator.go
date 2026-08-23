// Copyright (c) 2020 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package backup

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/bborbe/errors"
	"github.com/bborbe/io/util"

	"github.com/bborbe/postgres-backup/model"
)

// Create backs up the database via pg_dump. The now clock is injected so the
// backup filename is testable.
func Create(
	ctx context.Context,
	now func() time.Time,
	name model.Name,
	host model.PostgresqlHost,
	port model.PostgresqlPort,
	user model.PostgresqlUser,
	pass model.PostgresqlPassword,
	database model.PostgresqlDatabase,
	targetDirectory model.TargetDirectory,
) error {
	//pg_dump -Z 9 -h ${POSTGRES_HOST} -p ${POSTGRES_PORT} -U ${POSTGRES_USER} -F c -b -v -f ${BACKUP_NAME} ${POSTGRES_DB}
	backupfile := model.BuildBackupfileName(name, targetDirectory, database, now())

	if backupfile.Exists() {
		slog.Info("backup already exists, skipping", "backup", backupfile)
		return nil
	}

	if err := writePasswordFile(ctx, host, port, user, pass); err != nil {
		return errors.Wrapf(ctx, err, "write password file")
	}

	slog.Info("pg_dump started")
	if err := runCommand(ctx, "pg_dump", targetDirectory, "-Z", "9", "-h", host.String(), "-p", port.String(), "-U", user.String(), "-F", "c", "-b", "-v", "-f", backupfile.String(), database.String()); err != nil {
		slog.Info("pg_dump failed, deleting incomplete backup", "error", err)
		if err := backupfile.Delete(); err != nil {
			slog.Warn("delete incomplete backup failed", "error", err)
		}
		return errors.Wrapf(ctx, err, "pg_dump")
	}
	slog.Info("pg_dump finished")
	return nil
}

func writePasswordFile(ctx context.Context, host model.PostgresqlHost, port model.PostgresqlPort, user model.PostgresqlUser, pass model.PostgresqlPassword) error {
	content := fmt.Sprintf("%s:%d:*:%s:%s\n", host, port, user, pass)
	path, err := util.NormalizePath("~/.pgpass")
	if err != nil {
		return errors.Wrapf(ctx, err, "normalize pgpass path")
	}
	if err := os.WriteFile(path, []byte(content), 0600); err != nil {
		return errors.Wrapf(ctx, err, "write pgpass file %s", path)
	}
	return nil
}

func runCommand(ctx context.Context, command string, cwd model.TargetDirectory, args ...string) error {
	debug := fmt.Sprintf("%s %s", command, strings.Join(args, " "))
	slog.Debug("execute", "command", debug)
	cmd := exec.Command(command, args...)
	if cwd != "" {
		cmd.Dir = cwd.String()
	}
	if err := cmd.Start(); err != nil {
		return errors.Wrapf(ctx, err, "start %s", debug)
	}
	slog.Debug("command started", "command", debug)
	if err := cmd.Wait(); err != nil {
		return errors.Wrapf(ctx, err, "%s failed", debug)
	}
	slog.Debug("command finished", "command", command)
	return nil
}
