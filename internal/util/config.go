package util

import (
	"embed"
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

//go:embed internal/util/app.env
var defaultConfig embed.FS

type Config struct {
    BaseURL string `mapstructure:"BASE_URL"`
    Port    string `mapstructure:"PORT"`
    DbPath  string `mapstructure:"URL_SHORTENER_DB_PATH"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigName("app")
	viper.SetConfigType("env")

    customConfigPath, isCustomPathSet := os.LookupEnv("APP_ENV_PATH")
    if isCustomPathSet {
		fmt.Println("Loading custom config file from", customConfigPath)
        viper.SetConfigFile(customConfigPath)
    } else {
        configFile, fileErr := defaultConfig.ReadFile("internal/util/app.env")
        if fileErr != nil {
            return Config{}, fileErr
        }

        if err := viper.ReadConfig(bytes.NewReader(configFile)); err != nil {
            return Config{}, err
        }
    }

    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)
        return Config{}, err
    }

    if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("Error unmarshalling config file:", err)
        return Config{}, err
    }

	return config, nil
}

func GetConfig() Config {
    c, err := LoadConfig()
    if err != nil {
        log.Fatal("Cannot load config:", err)
    }
    return c
}
