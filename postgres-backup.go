// Copyright (c) 2020 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/bborbe/cron"
	flag "github.com/bborbe/flagenv"
	"github.com/bborbe/lock"
	"github.com/bborbe/postgres-backup/backup"
	"github.com/bborbe/postgres-backup/model"
	"github.com/golang/glog"

	"github.com/bborbe/run"
)

const (
	defaultLockName           = "/var/run/postgres-backup.lock"
	defaultName               = "postgres"
	parameterPostgresHost     = "host"
	parameterPostgresPort     = "port"
	parameterPostgresDatabase = "database"
	parameterPostgresUser     = "username"
	parameterPostgresPassword = "password"
	parameterTargetDir        = "targetdir"
	parameterWait             = "wait"
	parameterOneTime          = "one-time"
	parameterLock             = "lock"
	parameterName             = "name"
)

var (
	hostPtr      = flag.String(parameterPostgresHost, "", "host")
	portPtr      = flag.Int(parameterPostgresPort, 5432, "port")
	databasePtr  = flag.String(parameterPostgresDatabase, "", "database")
	userPtr      = flag.String(parameterPostgresUser, "", "username")
	passwordPtr  = flag.String(parameterPostgresPassword, "", "password")
	waitPtr      = flag.Duration(parameterWait, time.Minute*60, "wait")
	oneTimePtr   = flag.Bool(parameterOneTime, false, "exit after first backup")
	targetDirPtr = flag.String(parameterTargetDir, "", "target directory")
	lockPtr      = flag.String(parameterLock, defaultLockName, "lock")
	namePtr      = flag.String(parameterName, defaultName, "name")
)

func main() {
	defer glog.Flush()
	glog.CopyStandardLogTo("info")
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	if err := do(); err != nil {
		glog.Exit(err)
	}
}

func do() error {
	lockName := *lockPtr
	l := lock.NewLock(lockName)
	if err := l.Lock(); err != nil {
		return err
	}
	defer func() {
		if err := l.Unlock(); err != nil {
			glog.Warningf("unlock failed: %v", err)
		}
	}()

	glog.V(1).Info("backup postgres cron started")
	defer glog.V(1).Info("backup postgres cron finished")

	return exec()
}

func exec() error {
	host := model.PostgresqlHost(*hostPtr)
	if len(host) == 0 {
		return fmt.Errorf("parameter %s missing", parameterPostgresHost)
	}
	port := model.PostgresqlPort(*portPtr)
	if port <= 0 {
		return fmt.Errorf("parameter %s missing", parameterPostgresPort)
	}
	user := model.PostgresqlUser(*userPtr)
	if len(user) == 0 {
		return fmt.Errorf("parameter %s missing", parameterPostgresUser)
	}
	pass := model.PostgresqlPassword(*passwordPtr)
	if len(pass) == 0 {
		return fmt.Errorf("parameter %s missing", parameterPostgresPassword)
	}
	database := model.PostgresqlDatabase(*databasePtr)
	if len(database) == 0 {
		return fmt.Errorf("parameter %s missing", parameterPostgresDatabase)
	}
	targetDir := model.TargetDirectory(*targetDirPtr)
	if len(targetDir) == 0 {
		return fmt.Errorf("parameter %s missing", parameterTargetDir)
	}
	name := model.Name(*namePtr)
	if len(name) == 0 {
		return fmt.Errorf("parameter %s missing", parameterName)
	}

	oneTime := *oneTimePtr
	wait := *waitPtr
	lockName := *lockPtr

	glog.V(1).Infof("name: %s, host: %s, port: %d, user: %s, password-length: %d, database: %s, targetDir: %s, wait: %v, oneTime: %v, lockName: %s", name, host, port, user, len(pass), database, targetDir, wait, oneTime, lockName)

	action := run.Func(func(ctx context.Context) error {
		return backup.Create(name, host, port, user, pass, database, targetDir)
	})

	var c cron.Cron
	if *oneTimePtr {
		c = cron.NewOneTimeCron(action)
	} else {
		c = cron.NewWaitCron(
			*waitPtr,
			action,
		)
	}
	return c.Run(context.Background())
}
