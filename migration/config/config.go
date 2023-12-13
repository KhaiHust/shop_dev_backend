package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Database struct {
		Username     string `yaml:"username"`
		Password     string `yaml:"password"`
		Host         string `yaml:"host"`
		Port         string `yaml:"port"`
		DatabaseName string `yaml:"databaseName"`
	} `yaml:"database"`
}

func LoadConfig() (*Config, error) {
	config := &Config{}
	file, err := os.ReadFile("../config.yml")
	if err != nil {
		log.Error("yamlFile.Get err   #%v ", err)
		return nil, err
	}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		log.Error("Cannot unmarshall %v", err)
		return nil, err
	}
	return config, nil
}
