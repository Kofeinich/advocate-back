package config

type Postgres struct {
	Username string `yaml:"username"`
	DB       string `yaml:"db"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}
