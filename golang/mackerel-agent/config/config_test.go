package config

import (
	"io/ioutil"
	"testing"
)

var sampleConfig = `
apikey = "abcde"

[connection]
post_metrics_retry_delay_seconds = 600
post_metrics_retry_max = 5

[plugin.metrics.mysql]
command = "ruby /path/to/your/plugin/mysql.rb"
`

func TestLoadConfig(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "")
	if err != nil {
		t.Errorf("should not raise error: %v", err)
	}
	if err = ioutil.WriteFile(tmpFile.Name(), []byte(sampleConfig), 0644); err != nil {
		t.Errorf("should not raise error: %v", err)
	}

	config, err := LoadConfig(tmpFile.Name())
	if err != nil {
		t.Errorf("shoud not raise error: %v", err)
	}

	if config.Apibase != "https://mackerel.io" {
		t.Errorf("should be https://mackerel.io (arg value should be used)")
	}

	if config.Apikey != "abcde" {
		t.Errorf("shoudl be abcde (config value should be used)")
	}

	if config.Connection.PostMetricsDequeueDelaySeconds != 30 {
		t.Errorf("should be 30 (default value should be used)")
	}

	if config.Connection.PostMetricsRetryDelaySeconds != 180 {
		t.Errorf("should be 180 (max retry delay seconds is 180)")
	}

	if config.Connection.PostMetricsRetryMax != 5 {
		t.Errorf("should be 5 (config value should be used)")
	}
}
