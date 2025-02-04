REGISTRY ?= docker.io
IMAGE ?= bborbe/postgres-backup
ifeq ($(VERSION),)
	VERSION := $(shell git fetch --tags; git describe --tags `git rev-list --tags --max-count=1`)
endif

default: precommit

precommit: ensure format generate test check addlicense
	@echo "ready to commit"

ensure:
	go mod tidy
	go mod verify
	go mod vendor

format:
	go run -mod=vendor github.com/incu6us/goimports-reviser/v3 -project-name github.com/bborbe/run -format -excludes vendor ./...

generate:
	rm -rf mocks avro
	go generate -mod=vendor ./...

test:
	go test -mod=vendor -p=$${GO_TEST_PARALLEL:-1} -cover -race $(shell go list -mod=vendor ./... | grep -v /vendor/)

check: vet errcheck vulncheck

vet:
	go vet -mod=vendor $(shell go list -mod=vendor ./... | grep -v /vendor/)

errcheck:
	go run -mod=vendor github.com/kisielk/errcheck -ignore '(Close|Write|Fprint)' $(shell go list -mod=vendor ./... | grep -v /vendor/)

addlicense:
	go run -mod=vendor github.com/google/addlicense -c "Benjamin Borbe" -y $$(date +'%Y') -l bsd $$(find . -name "*.go" -not -path './vendor/*')

vulncheck:
	go run -mod=vendor golang.org/x/vuln/cmd/govulncheck $(shell go list -mod=vendor ./... | grep -v /vendor/)

build:
	imagebuilder -t $(REGISTRY)/$(IMAGE):$(VERSION) -f Dockerfile:Dockerfile .

upload:
	docker push $(REGISTRY)/$(IMAGE):$(VERSION)

clean:
	docker rmi $(REGISTRY)/$(IMAGE):$(VERSION) || true
