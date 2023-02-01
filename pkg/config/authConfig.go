package config

type Auth struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Secret   string `yaml:"jwt_secret"`
}
