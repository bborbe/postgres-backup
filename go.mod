module github.com/bborbe/postgres-backup

go 1.23.5

replace (
	github.com/coreos/bbolt v1.3.10 => go.etcd.io/bbolt v1.3.10
	github.com/coreos/bbolt v1.3.11 => go.etcd.io/bbolt v1.3.11
	github.com/coreos/bbolt v1.3.6 => go.etcd.io/bbolt v1.3.6
	github.com/coreos/bbolt v1.3.7 => go.etcd.io/bbolt v1.3.7
	github.com/coreos/bbolt v1.3.8 => go.etcd.io/bbolt v1.3.8
	github.com/coreos/bbolt v1.3.9 => go.etcd.io/bbolt v1.3.9
)

require (
	github.com/actgardner/gogen-avro/v9 v9.2.0
	github.com/bborbe/assert v0.0.0-20181116222016-22a6c6341415
	github.com/bborbe/cron v1.2.2
	github.com/bborbe/flagenv v0.0.0-20181019084341-2956c4545608
	github.com/bborbe/io v0.0.0-20180829202151-54b762caaee8
	github.com/bborbe/lock v1.0.0
	github.com/bborbe/run v1.5.6
	github.com/golang/glog v1.2.4
	github.com/google/addlicense v1.1.1
	github.com/incu6us/goimports-reviser/v3 v3.8.2
	github.com/kisielk/errcheck v1.8.0
	github.com/maxbrunsfeld/counterfeiter/v6 v6.11.2
	golang.org/x/lint v0.0.0-20241112194109-818c5a804067
	golang.org/x/vuln v1.1.4
)

require (
	github.com/bborbe/errors v1.3.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bmatcuk/doublestar/v4 v4.8.1 // indirect
	github.com/certifi/gocertifi v0.0.0-20210507211836-431795d63e8d // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/getsentry/raven-go v0.2.0 // indirect
	github.com/incu6us/goimports-reviser v0.1.6 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.20.5 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.62.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	golang.org/x/exp v0.0.0-20250128182459-e0ece0dbea4c // indirect
	golang.org/x/mod v0.22.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/telemetry v0.0.0-20250117155846-04cd7bae618c // indirect
	golang.org/x/text v0.21.0 // indirect
	golang.org/x/tools v0.29.0 // indirect
	google.golang.org/protobuf v1.36.4 // indirect
)
