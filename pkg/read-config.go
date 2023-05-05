package pkg

import (
	"advocate-back/pkg/config"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Auth config.Auth `yaml:"auth_credentials"`
	Url  string      `yaml:"url"`
}

func ReadConfig() Config {
	file, err := os.ReadFile("configs/config.yml")
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
