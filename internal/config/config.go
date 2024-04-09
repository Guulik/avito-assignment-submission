package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	Env      string `yaml:"env"`
	Port     int    `yaml:"port"`
	Postgres Postgres
}

type Postgres struct {
	Host     string `yaml:"host"`
	SQLPort  string `yaml:"SQLPort"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"DBName"`
	SslMode  string `yaml:"sslMode"`
	Driver   string `yaml:"driver"`
}

func MustLoad() *Config {
	configPath := fetchConfigPath()

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}
	var config Config

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &config
}

func fetchConfigPath() string {
	path := "local.yaml"
	return path
}
