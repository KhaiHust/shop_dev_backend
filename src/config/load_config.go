package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func getViper() *viper.Viper {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigFile("config.yml")
	return v
}

func NewConfig() (*Config, error) {
	log.Info("Load configuration")
	v := getViper()

	err := v.ReadInConfig()
	if err != nil {
		log.Error(fmt.Sprintf("cannot log config: %v", err))
		return nil, err
	}
	var config Config
	err = v.Unmarshal(&config)
	return &config, nil
}
