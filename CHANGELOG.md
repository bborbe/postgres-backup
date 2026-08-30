# Changelog

All notable changes to this project will be documented in this file.

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
