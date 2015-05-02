package config

import (
	"os"
	"path/filepath"
)

var mackerelRoot = filepath.Join(os.Getenv("HOME"), "Library", "mackerel-agent")

var DefaultConfig = &Config{
	ApiBase:  getApibase(),
	Root:     mackerelRoot,
	Pidfile:  filepath.Join(mackerelRoot, "pid"),
	Conffile: filepath.Join(mackerelRoot, "mackerel-agent.conf"),
	Roles:    []string{},
	Verbose:  false,
	Connection: ConnectionConfig{
		PostMetricsDequeueDelaySeconds: 30,
		PostMetricsRetryDelaySeconds:   60,
		PostMetricsRetryMax:            60,
		PostMetricsBufferSize:          6 * 60,
	},
}
