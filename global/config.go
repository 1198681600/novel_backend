package global

import (
	"fmt"
	"github.com/spf13/viper"
)

func initConfig() {
	initViper()
}

func initViper() {
	viper.SetConfigFile("./config/config.toml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read config file: %w", err))
	}
}
