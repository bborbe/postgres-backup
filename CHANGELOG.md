# Changelog

All notable changes to this project will be documented in this file.

## Unreleased

- fix: `make build` refuses to stamp a version onto a tree that is not that version's tag (`check-version-tag`, escape hatch `ALLOW_UNTAGGED_BUILD=1`). `VERSION` defaults to the newest tag repo-wide, so an operator-run build from an untagged or older tree silently republishes under the newest tag. The guard compares `git describe --exact-match HEAD` against `$(VERSION)` and exits non-zero on mismatch.

## 2.1.3

- chore: update Go to 1.27.0 and github.com/bborbe/cron to v1.8.28, github.com/bborbe/errors to v1.6.0, github.com/bborbe/lock to v1.0.5, github.com/bborbe/run to v1.10.1, github.com/bborbe/time to v1.27.11

## 2.1.2

- chore: add .maintainer.yaml to enable automated Go version and dependency updates
- fix: modernize postgres-backup — cobra CLI (replaces flagenv), log/slog (replaces glog), bborbe/errors wrapping with context threading, main_test gexec compile check, tools.go → tools.env, postgres:17 base image
- chore: update Go to 1.26.6

## 2.1.1

- go mod update

## 2.1.0

- Update Postgres to 11.7
- Update Golang to 1.13.9
- Use gomod instead gopkg

## 2.0.1

- Change build

## 2.0.0

- Add Jenkinsfile
- Use deps instead glide
- Move main to root
