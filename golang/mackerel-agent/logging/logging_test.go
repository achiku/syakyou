package logging

import "testing"

func TestGetLogger(t *testing.T) {
	var logger = GetLogger("tag")
	if logger.tag != "tag" {
		t.Errorf("tag should be tag but %v", logger.tag)
	}
}

func TestConfigureLoggers(t *testing.T) {
	ConfigureLoggers("INFO")
	if logLevelConfigs["root"] != info {
		t.Errorf("tag should be tag but %v", logLevelConfigs["root"])
	}
}
