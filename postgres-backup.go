// Copyright (c) 2020 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/bborbe/cron"
	"github.com/bborbe/errors"
	"github.com/bborbe/lock"
	"github.com/bborbe/run"
	libtime "github.com/bborbe/time"
	"github.com/spf13/cobra"

	"github.com/bborbe/postgres-backup/backup"
	"github.com/bborbe/postgres-backup/model"
)

const (
	defaultLockName = "/var/run/postgres-backup.lock"
	defaultName     = "postgres"
)

type backupConfig struct {
	host      string
	port      int
	database  string
	username  string
	password  string
	targetDir string
	wait      time.Duration
	oneTime   bool
	lockName  string
	name      string
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-sigCh
		cancel()
	}()

	if err := Run(ctx, os.Args[1:]); err != nil {
		slog.Error("postgres-backup failed", "error", err)
		os.Exit(1)
	}
}

// Run parses the CLI arguments and runs the backup loop until the context is
// cancelled or a fatal error occurs.
func Run(ctx context.Context, args []string) error {
	config := backupConfig{}

	rootCmd := &cobra.Command{
		Use:          "postgres-backup",
		Short:        "Backup a PostgreSQL database on a schedule",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, _ []string) error {
			slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo})))
			runtime.GOMAXPROCS(runtime.NumCPU())
			return do(ctx, config)
		},
	}
	rootCmd.Flags().StringVar(&config.host, "host", "", "PostgreSQL host")
	rootCmd.Flags().IntVar(&config.port, "port", 5432, "PostgreSQL port")
	rootCmd.Flags().StringVar(&config.database, "database", "", "PostgreSQL database")
	rootCmd.Flags().StringVar(&config.username, "username", "", "PostgreSQL username")
	rootCmd.Flags().StringVar(&config.password, "password", "", "PostgreSQL password")
	rootCmd.Flags().StringVar(&config.targetDir, "targetdir", "", "target directory")
	rootCmd.Flags().DurationVar(&config.wait, "wait", time.Hour, "wait between backups")
	rootCmd.Flags().BoolVar(&config.oneTime, "one-time", false, "exit after first backup")
	rootCmd.Flags().StringVar(&config.lockName, "lock", defaultLockName, "lock file path")
	rootCmd.Flags().StringVar(&config.name, "name", defaultName, "backup name")

	rootCmd.SetArgs(args)
	return rootCmd.ExecuteContext(ctx)
}

func do(ctx context.Context, config backupConfig) error {
	l := lock.NewLock(config.lockName)
	if err := l.Lock(); err != nil {
		return errors.Wrapf(ctx, err, "acquire lock %s", config.lockName)
	}
	defer func() {
		if err := l.Unlock(); err != nil {
			slog.Warn("unlock failed", "error", err)
		}
	}()

	slog.Info("backup postgres cron started")
	defer slog.Info("backup postgres cron finished")

	return exec(ctx, config)
}

func exec(ctx context.Context, config backupConfig) error {
	host := model.PostgresqlHost(config.host)
	if len(host) == 0 {
		return errors.Errorf(ctx, "parameter %s missing", "host")
	}
	port := model.PostgresqlPort(config.port)
	if port <= 0 {
		return errors.Errorf(ctx, "parameter %s missing", "port")
	}
	user := model.PostgresqlUser(config.username)
	if len(user) == 0 {
		return errors.Errorf(ctx, "parameter %s missing", "username")
	}
	pass := model.PostgresqlPassword(config.password)
	if len(pass) == 0 {
		return errors.Errorf(ctx, "parameter %s missing", "password")
	}
	database := model.PostgresqlDatabase(config.database)
	if len(database) == 0 {
		return errors.Errorf(ctx, "parameter %s missing", "database")
	}
	targetDir := model.TargetDirectory(config.targetDir)
	if len(targetDir) == 0 {
		return errors.Errorf(ctx, "parameter %s missing", "targetdir")
	}
	name := model.Name(config.name)
	if len(name) == 0 {
		return errors.Errorf(ctx, "parameter %s missing", "name")
	}

	slog.Info("backup postgres configuration",
		"name", name,
		"host", host,
		"port", port,
		"user", user,
		"passwordLength", len(pass),
		"database", database,
		"targetDir", targetDir,
		"wait", config.wait,
		"oneTime", config.oneTime,
		"lockName", config.lockName,
	)

	action := run.Func(func(ctx context.Context) error {
		return backup.Create(name, host, port, user, pass, database, targetDir)
	})

	var c cron.Cron
	if config.oneTime {
		c = cron.NewOneTimeCron(action)
	} else {
		c = cron.NewWaitCron(
			libtime.Duration(config.wait),
			action,
		)
	}
	return c.Run(ctx)
}
