package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"os"
	"time"
)

type Config struct {
	Env    string `yaml:"env" env-required:"true"`
	Server `yaml:"server" env-required:"true"`
	DB     `yaml:"db" env-required:"true"`
}

type Server struct {
	Port    string        `yaml:"port" env-required:"true"`
	Timeout time.Duration `yaml:"timeout" env-required:"true"`
}

type DB struct {
	Username string `yaml:"username" env-required:"true"`
	Host     string `yaml:"host" env-required:"true"`
	Port     string `yaml:"port" env-required:"true"`
	DBName   string `yaml:"dbname" env-required:"true"`
	SSLMode  string `yaml:"sslmode" env-required:"true"`
	Password string `env:"DB_PASSWORD" env-required:"true"`
}

func MustLoad(path string) *Config {
	if path == "" {
		panic("config file not exist")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file not exist")
	}
	cfg := &Config{}
	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		panic(err)
	}

	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	cfg.DB.Password = os.Getenv("DB_PASSWORD")
	if cfg.DB.Password == "" {
		panic("DB_PASSWORD not set")
	}

	return cfg
}
