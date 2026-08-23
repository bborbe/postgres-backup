// Copyright (c) 2020 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package model

import (
	"fmt"
	"log/slog"
	"os"
	"time"
)

type BackupFilename string

func BuildBackupfileName(name Name, targetDirectory TargetDirectory, database PostgresqlDatabase, date time.Time) BackupFilename {
	return BackupFilename(fmt.Sprintf("%s/%s_%s_%s.dump", targetDirectory, name, database, date.Format("2006-01-02")))
}

func (b BackupFilename) Delete() error {
	return os.Remove(b.String())
}

func (b BackupFilename) String() string {
	return string(b)
}

func (b BackupFilename) Exists() bool {
	fileInfo, err := os.Stat(b.String())
	if err != nil {
		slog.Debug("file not exists", "file", b)
		return false
	}
	if fileInfo.Size() == 0 {
		slog.Debug("file empty", "file", b)
		return false
	}
	slog.Debug("file exists and not empty", "file", b)
	return true
}
