module github.com/bborbe/postgres-backup

go 1.27.1

replace (
	github.com/coreos/bbolt v1.3.10 => go.etcd.io/bbolt v1.3.10
	github.com/coreos/bbolt v1.3.11 => go.etcd.io/bbolt v1.3.11
	github.com/coreos/bbolt v1.3.6 => go.etcd.io/bbolt v1.3.6
	github.com/coreos/bbolt v1.3.7 => go.etcd.io/bbolt v1.3.7
	github.com/coreos/bbolt v1.3.8 => go.etcd.io/bbolt v1.3.8
	github.com/coreos/bbolt v1.3.9 => go.etcd.io/bbolt v1.3.9
)

require (
	github.com/bborbe/cron v1.9.0
	github.com/bborbe/errors v1.6.1
	github.com/bborbe/io v0.0.0-20180829202151-54b762caaee8
	github.com/bborbe/lock v1.0.6
	github.com/bborbe/run v1.11.0
	github.com/bborbe/time v1.27.14
	github.com/onsi/ginkgo/v2 v2.32.1
	github.com/onsi/gomega v1.43.0
	github.com/spf13/cobra v1.10.2
)

require (
	github.com/Masterminds/semver/v3 v3.5.0 // indirect
	github.com/bborbe/assert v0.0.0-20181116222016-22a6c6341415 // indirect
	github.com/bborbe/collection v1.21.0 // indirect
	github.com/bborbe/math v1.4.11 // indirect
	github.com/bborbe/parse v1.12.0 // indirect
	github.com/bborbe/validation v1.5.2 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/getsentry/sentry-go v0.50.0 // indirect
	github.com/go-logr/logr v1.4.4 // indirect
	github.com/go-task/slim-sprig/v3 v3.0.0 // indirect
	github.com/golang/glog v1.2.5 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/pprof v0.0.0-20260903180319-d6c3cb2f37ec // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.25.0 // indirect
	github.com/prometheus/client_model v0.6.3 // indirect
	github.com/prometheus/common v0.72.0 // indirect
	github.com/prometheus/procfs v0.22.0 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/spf13/pflag v1.0.10 // indirect
	go.yaml.in/yaml/v3 v3.0.5 // indirect
	golang.org/x/mod v0.41.0 // indirect
	golang.org/x/net v0.59.0 // indirect
	golang.org/x/sync v0.23.0 // indirect
	golang.org/x/sys v0.48.0 // indirect
	golang.org/x/text v0.42.0 // indirect
	golang.org/x/tools v0.49.0 // indirect
	google.golang.org/protobuf v1.36.12 // indirect
)
