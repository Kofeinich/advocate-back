package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Auth     Auth     `yaml:"auth_credentials"`
	Postgres Postgres `yaml:"postgres_config"`
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
