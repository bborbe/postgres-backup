REGISTRY ?= docker.io
IMAGE ?= bborbe/postgres-backup
ifeq ($(VERSION),)
	VERSION := $(shell git fetch --tags; git describe --tags `git rev-list --tags --max-count=1`)
endif

precommit: ensure format test check addlicense
	@echo "ready to commit"

ensure:
	GO111MODULE=on go mod verify
	GO111MODULE=on go mod vendor

format:
	go get golang.org/x/tools/cmd/goimports
	find . -type f -name '*.go' -not -path './vendor/*' -exec gofmt -w "{}" +
	find . -type f -name '*.go' -not -path './vendor/*' -exec goimports -w "{}" +

test:
	GO111MODULE=on go test -cover -race $(shell go list ./... | grep -v /vendor/)

check: lint vet errcheck

lint:
	@GO111MODULE=on go get golang.org/x/lint/golint
	@golint -min_confidence 1 $(shell go list ./... | grep -v /vendor/)

vet:
	@GO111MODULE=on go vet $(shell go list ./... | grep -v /vendor/)

errcheck:
	@GO111MODULE=on go get github.com/kisielk/errcheck
	@errcheck -ignore '(Close|Write|Fprint)' $(shell go list ./... | grep -v /vendor/)

addlicense:
	@GO111MODULE=on go get github.com/google/addlicense
	@addlicense -c "Benjamin Borbe" -y 2020 -l bsd ./*.go ./model/*.go ./backup/*.go
