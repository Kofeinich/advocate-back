package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Smtp struct {
	Host     string `yaml:"smtp_host"`
	Port     int    `yaml:"smtp_port"`
	Username string `yaml:"smtp_username"`
	Password string `yaml:"smtp_password"`
	From     string `yaml:"smtp_from"`
	To       string `yaml:"smtp_to"`
}
type Config struct {
	Smtp Smtp `yaml:"smtp_credentials"`
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
