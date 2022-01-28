package config

import "github.com/jinzhu/configor"

var Config = struct {
	BaseURL string `default:"http://localhost" yaml:"BaseURL"`
	Port    string `default:"8080" yaml:"Port"`

	DB struct {
		User     string `default:"user" yaml:"User"`
		Password string `default:"password" yaml:"Password"`
		Protocol string `default:"tcp" yaml:"Protocol"`
		Host     string `default:"db" yaml:"Host"`
		Port     string `default:"3306" yaml:"Port"`
		Name     string `default:"Url_Shortener" yaml:"Name"`
		Params   string `default:"charset=utf8mb4&parseTime=True" yaml:"Params"`
	}
}{}

func Init() {
	err := configor.Load(&Config, "config.yml")
	if err != nil {
		panic(err)
	}
}
