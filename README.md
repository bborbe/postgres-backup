# Postgres Backup Cron

## Install

`go get github.com/bborbe/postgres-backup`

## Run Backup

One time

```
postgres-backup \
-logtostderr \
-v=2 \
-host=localhost \
-port=5432 \
-lock=/tmp/lock \ 
-username=postgres \
-password=S3CR3T \
-database=db \
-targetdir=/backup \
-name=postgres \
-one-time
```

Cron

```
postgres-backup \
-logtostderr \
-v=2 \
-host=localhost \
-port=5432 \
-lock=/tmp/lock \ 
-username=postgres \
-password=S3CR3T \
-database=db \
-targetdir=/backup \
-name=postgres \
-wait=1h
```

## Start Postgres

```
docker run \
--name db \
--env POSTGRES_DB=mydb \
--env POSTGRES_USER=myuser \
--env POSTGRES_PASSWORD=mypassword \
--env PGDATA=/var/lib/postgresql/data/pgdata \
postgres:9.6.1-alpine
```

## Backup once with Docker

```
docker run \
--env HOST=db \
--env PORT=5432 \
--env DATABASE=mydb \
--env USERNAME=myuser \
--env PASSWORD=mypassword \
--env ONE_TIME=true \
--env LOCK=/postgres_backup_cron.lock \
--env TARGETDIR=/backup \
--volume /tmp:/backup \
--link db:db \
bborbe/postgres-backup:latest \
-logtostderr \
-v=1
```

`ls /tmp/postgres_mydb_*.dump`

## Backup every hour with Docker

```
docker run \
--env HOST=db \
--env PORT=5432 \
--env DATABASE=mydb \
--env USERNAME=myuser \
--env PASSWORD=mypassword \
--env WAIT=1h \
--env ONE_TIME=false \
--env LOCK=/postgres_backup_cron.lock \
--env TARGETDIR=/backup \
--volume /tmp:/backup \
--link db:db \
bborbe/postgres-backup:latest \
-logtostderr \
-v=1
```

`ls /tmp/postgres_mydb_*.dump`
