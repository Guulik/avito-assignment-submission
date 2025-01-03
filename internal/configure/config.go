package configure

import (
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env      string        `yaml:"env"`
	Port     int           `yaml:"port"`
	Timeout  time.Duration `yaml:"timeout"`
	Postgres Postgres
	Redis    Redis
}

type Postgres struct {
	Host     string `yaml:"host"`
	SQLPort  int    `yaml:"SQLPort"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"DBName"`
	SslMode  string `yaml:"sslMode"`
	Driver   string `yaml:"driver"`
}

type Redis struct {
	Address    string        `yaml:"address"`
	Password   string        `yaml:"password"`
	DB         int           `yaml:"DB"`
	TTLMinutes time.Duration `yaml:"TTLMinutes"`
}

func MustLoadConfig() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config path is empty")
	}

	return MustLoadPath(configPath)
}

func MustLoadPath(configPath string) *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	const key = "CONFIG_PATH"

	if v := os.Getenv(key); v != "" {
		return v
	}

	return "./internal/configure/local.yaml"
}
