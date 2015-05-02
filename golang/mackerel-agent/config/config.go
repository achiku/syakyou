package config

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/mackerelio/mackerel-agent/logging"
)

var configLogger = logging.GetLogger("config")

var apibase string

func getApibase() string {
	if apibase != "" {
		return apibase
	}
	return "https://mackerel.io"
}

type Config struct {
	Apibase    string
	Apikey     string
	Root       string
	Pidfile    string
	Conffile   string
	Roles      []string
	Verbose    bool
	Connection ConnectionConfig
	Plugin     map[string]PluginConfigs
	Include    string
}

type PluginConfigs map[string]PluginConfig

type PluginConfig struct {
	Command string
}

const postMetricsDequeueDelaySecondsMax = 59
const postMetricsRetryDelaySecondsMax = 3 * 60

var PostMetricsInterval = 1 * time.Minute

type ConnectionConfig struct {
	PostMetricsDequeueDelaySeconds int `toml:"post_metrics_dequeue_delay_seconds"`
	PostMetricsRetryDelaySeconds   int `toml:"post_metrics_retry_delay_seconds"`
	PostMetricsRetryMax            int `toml:"post_metrics_retry_max"`
	PostMetricsBufferSize          int `toml:"post_metrics_buffer_size"`
}

func LoadConfig(conffile string) (*Config, error) {
	config, err := loadConfigFile(conffile)

	// DefaultConfig is in each platform config
	if config.Apibase == "" {
		config.Apibase = DefaultConfig.Apibase
	}
	if config.Root == "" {
		config.Root = DefaultConfig.Root
	}
	if config.Pidfile == "" {
		config.Pidfile = DefaultConfig.Pidfile
	}
	if config.Verbose == false {
		config.Verbose = DefaultConfig.Verbose
	}
	if config.Connection.PostMetricsDequeueDelaySeconds == 0 {
		config.Connection.PostMetricsDequeueDelaySeconds = DefaultConfig.Connection.PostMetricsDequeueDelaySeconds
	}
	if config.Connection.PostMetricsDequeueDelaySeconds > postMetricsDequeueDelaySecondsMax {
		configLogger.Warningf("'post_metrics_dequese_delay_seconds' is set to %d (Maximum Value).", postMetricsDequeueDelaySecondsMax)
		config.Connection.PostMetricsDequeueDelaySeconds = postMetricsDequeueDelaySecondsMax
	}
	if config.Connection.PostMetricsRetryDelaySeconds == 0 {
		config.Connection.PostMetricsRetryDelaySeconds = DefaultConfig.Connection.PostMetricsRetryDelaySeconds
	}
	if config.Connection.PostMetricsRetryDelaySeconds > postMetricsRetryDelaySecondsMax {
		configLogger.Warningf("'post_metrics_retry_delay_seconds' is set to %d (Maximum Value).", postMetricsRetryDelaySecondsMax)
		config.Connection.PostMetricsRetryDelaySeconds = postMetricsRetryDelaySecondsMax
	}
	if config.Connection.PostMetricsRetryMax == 0 {
		config.Connection.PostMetricsRetryMax = DefaultConfig.Connection.PostMetricsRetryMax
	}
	if config.Connection.PostMetricsBufferSize == 0 {
		config.Connection.PostMetricsBufferSize = DefaultConfig.Connection.PostMetricsBufferSize
	}
	return config, err
}

func loadConfigFile(file string) (*Config, error) {
	config := &Config{}
	if _, err := toml.DecodeFile(file, config); err != nil {
		return config, err
	}

	if config.Include != "" {
		if err := includeConfigFile(config, config.Include); err != nil {
			return config, err
		}
	}

	// don't copy deprecated part of source
	return config, nil
}

func includeConfigFile(config *Config, include string) error {
	files, err := filepath.Glob(include)
	if err != nil {
		return err
	}

	for _, file := range files {
		// ???
		rolesSaved := config.Roles
		config.Roles = nil

		pluginSaved := map[string]PluginConfigs{}
		for kind, plugins := range config.Plugin {
			pluginSaved[kind] = plugins
		}

		meta, err := toml.DecodeFile(file, &config)
		if err != nil {
			return fmt.Errorf("while loading included config file %s: %s", file, err)
		}

		if meta.IsDefined("roles") == false {
			config.Roles = rolesSaved
		}

		for kind, plugins := range config.Plugin {
			for key, conf := range plugins {
				if pluginSaved[kind] == nil {
					pluginSaved[kind] = PluginConfigs{}
				}
				pluginSaved[kind][key] = conf
			}
		}
		config.Plugin = pluginSaved
	}

	return nil
}
