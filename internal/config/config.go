package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	Env  string `yaml:"env"`
	Port int    `yaml:"port"`
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
