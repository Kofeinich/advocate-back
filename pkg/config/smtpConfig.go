package config

type Smtp struct {
	Host     string `yaml:"smtp_host"`
	Port     int    `yaml:"smtp_port"`
	Username string `yaml:"smtp_username"`
	Password string `yaml:"smtp_password"`
	From     string `yaml:"smtp_from"`
	To       string `yaml:"smtp_to"`
}
