package config

import (
	"net/url"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Port         string        `env:"PORT"`
		WriteTimeout time.Duration `env:"WRITE_TIMEOUT"`
		ReadTimeout  time.Duration `env:"READ_TIMEOUT"`
		CORSorigins  string        `env:"CORS_ORIGINS"`
		// log levels:
		// debug, info, warn, error, fatal, panic
		LogLevel string `env:"LOG_LEVEL"`
		Database database
	}
	database struct {
		Host           string `env:"POSTGRES_HOST"`
		Port           string `env:"POSTGRES_PORT"`
		User           string `env:"POSTGRES_USER"`
		Pass           string `env:"POSTGRES_PASSWORD"`
		Name           string `env:"POSTGRES_DB"`
		WithMigrations bool   `env:"WITH_MIGRATIONS"`
	}
)

func LoadConfigs() (Config, error) {
	var config Config
	err := cleanenv.ReadEnv(&config)
	return config, err
}

func (d database) URL() string {
	url := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(d.User, d.Pass),
		Host:   d.Host + ":" + d.Port,
		Path:   d.Name,
	}
	return url.String() + "?sslmode=disable"
}
