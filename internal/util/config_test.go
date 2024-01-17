package util

import (
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestSetupDefaultConfigPath(t *testing.T) {
	err := SetupDefaultConfigPath()

	if err != nil {
		t.Error("Error setting up default config path:", err)
	}
}

func TestSetupCustomConfigPath(t *testing.T) {
	os.Setenv("APP_ENV_PATH", "./app.test.env")
	defer os.Unsetenv("APP_ENV_PATH")

	SetupCustomConfigPath("./app.test.env")

	if viper.ConfigFileUsed() != "./app.test.env" {
		t.Error("Error setting up custom config path")
	}

	config, err := LoadConfig()
	if err != nil {
		t.Error("Error loading config:", err)
	}

	if config.BaseURL != "test_base_url" {
		t.Error("Error loading config:", config.BaseURL)
	}

	if config.Port != "test_port" {
		t.Error("Error loading config:", config.Port)
	}

	if config.DbPath != "test_url_shortener_db_path" {
		t.Error("Error loading config:", config.DbPath)
	}
}

func TestLoadConfig(t *testing.T) {
	os.Setenv("APP_ENV_PATH", "./app.test.env")
	defer os.Unsetenv("APP_ENV_PATH")

	config, err := LoadConfig()
	if err != nil {
		t.Error("Error loading config:", err)
	}

	if config.BaseURL != "test_base_url" {
		t.Error("Error loading config:", config.BaseURL)
	}

	if config.Port != "test_port" {
		t.Error("Error loading config:", config.Port)
	}

	if config.DbPath != "test_url_shortener_db_path" {
		t.Error("Error loading config:", config.DbPath)
	}
}
