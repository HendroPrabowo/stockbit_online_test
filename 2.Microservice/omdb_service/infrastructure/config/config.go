package config

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Init() AppConfig {
	var config AppConfig
	path, err := os.Getwd()
	if err != nil {
		logrus.Fatalf("Error Read Path : %s\n", err.Error())
	}
	viper.SetConfigName("app_env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(fmt.Sprintf("%s/env", path))
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("Error Load Config : %s\n", err.Error())
	}
	if err := viper.Unmarshal(&config); err != nil {
		logrus.Fatalf("Error Load Config : %s\n", err.Error())
	}
	return config
}
