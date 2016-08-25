install:
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/postgres_backup_cron/postgres_backup_cron.go
test:
	GO15VENDOREXPERIMENT=1 go test `glide novendor`
vet:
	go tool vet .
	go tool vet .-shadow .
lint:
	golint -min_confidence 1 ./...
errcheck:
	errcheck -ignore '(Close|Write)' ./...
check: lint vet errcheck
run:
	postgres_backup_cron \
	-loglevel=debug \
	-host=localhost \
	-port=5432 \
	-lock=/tmp/lock \
	-username=postgres \
	-password=S3CR3T \
	-database=db \
	-targetdir=/tmp \
	-wait=1h
format:
	find . -name "*.go" -exec gofmt -w "{}" \;
	goimports -w=true .
prepare:
	npm install
	go get -u golang.org/x/tools/cmd/goimports
	go get -u github.com/Masterminds/glide
	go get -u github.com/golang/lint/golint
	go get -u github.com/kisielk/errcheck
	glide install
update:
	glide up
clean:
	rm -rf vendor