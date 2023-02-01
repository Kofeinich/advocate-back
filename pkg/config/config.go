package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Smtp Smtp `yaml:"smtp_credentials"`
	Auth Auth `yaml:"auth_credentials"`
}

func ReadConfig() Config {
	file, err := os.ReadFile("configs/app-config.yml")
	if err != nil {
		log.Fatalln(err)
		return Config{}
	}
	var cfg Config
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		log.Fatalln(err)
		return Config{}
	}
	return cfg
}

var AppConfig = ReadConfig()
