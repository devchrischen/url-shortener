package config

import "github.com/jinzhu/configor"

var Config = struct {
	BaseURL string `env:"BASE_URL" yaml:"BaseURL" default:"http://localhost"`
	Port    string `env:"PORT" yaml:"Port" default:"8080"`

	DB struct {
		User     string `env:"DB_HOST" yaml:"User" default:"user"`
		Password string `env:"DB_PASSWORD" yaml:"Password" default:"password"`
		Protocol string `env:"DB_PROTOCOL" yaml:"Protocol" default:"tcp"`
		Host     string `env:"DB_HOST" yaml:"Host" default:"localhost" `
		Port     string `env:"DB_PORT" yaml:"Port" default:"3306"`
		Name     string `env:"DB_NAME" yaml:"Name" default:"Url_Shortener"`
		Params   string `env:"DB_PARAMS" yaml:"Params" default:"charset=utf8mb4&parseTime=True"`
	}

	Redis struct {
		Host string `env:"REDIS_HOST" yaml:"Host" default:"localhost"`
		Port string `env:"REDIS_PORT" yaml:"Port" default:"6379"`
	}
}{}

func Init() {
	err := configor.Load(&Config, "config.yml")
	if err != nil {
		panic(err)
	}
}
