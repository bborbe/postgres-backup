FROM golang:1.11.1 AS build
COPY . /go/src/github.com/bborbe/postgres-backup
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o /postgres-backup ./src/github.com/bborbe/postgres-backup
CMD ["/bin/bash"]

FROM postgres:10
MAINTAINER Benjamin Borbe <bborbe@rocketnews.de>

ENV LOGLEVEL info
ENV HOST localhost
ENV PORT 5432
ENV DATABASE postgres
ENV USERNAME postgres
ENV PASSWORD S3CR3T
ENV TARGETDIR /backup
ENV WAIT 1h
ENV ONE_TIME false
ENV LOCK /var/run/postgres-backup.lock

RUN set -x \
	&& DEBIAN_FRONTEND=noninteractive apt-get update --quiet \
	&& DEBIAN_FRONTEND=noninteractive apt-get upgrade --quiet --yes \
	&& DEBIAN_FRONTEND=noninteractive apt-get autoremove --yes \
	&& DEBIAN_FRONTEND=noninteractive apt-get clean

VOLUME ["/backup"]

COPY  --from=build postgres-backup /
ENTRYPOINT ["/postgres-backup"]
