package util

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/spf13/viper"
)

type Config struct {
    BaseURL string `mapstructure:"BASE_URL"`
    Port    string `mapstructure:"PORT"`
    DbPath  string `mapstructure:"URL_SHORTENER_DB_PATH"`
}

func SetupCustomConfigPath(customConfigPath string) {
	viper.SetConfigFile(customConfigPath)
}

func SetupDefaultConfigPath() error {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename))
	filePath := dir + "/app.env"

	// Check if the config file exists
	_, fileErr := os.Stat(filePath)
	if fileErr != nil {
		return fileErr
	}

	viper.SetConfigFile(filePath)
	return nil
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigType("env")
	viper.SetConfigName("app")

    customConfigPath, isCustomPathSet := os.LookupEnv("APP_ENV_PATH")
    if isCustomPathSet {
		SetupCustomConfigPath(customConfigPath)
    } else {
		err := SetupDefaultConfigPath()
		if err != nil {
			fmt.Println("Error setting up default config path:", err)
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
