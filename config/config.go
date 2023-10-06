package config

import (
	"flag"
	"net/url"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HttpAddr     string        `env:"HTTP_ADDR" env-default:":8080"`
		HttpNetwork  string        `env:"HTTP_NETWORK" env-default:"tcp"`
		WriteTimeout time.Duration `env:"WRITE_TIMEOUT" env-default:"10s"`
		ReadTimeout  time.Duration `env:"READ_TIMEOUT" env-default:"10s"`
		LogLevel     string        `env:"LOG_LEVEL" env-default:"debug"`

		Database database
		Flags    flags
	}
	database struct {
		Host string `env:"POSTGRES_HOST" env-required:"true"`
		Port string `env:"POSTGRES_PORT" env-default:"5432"`
		User string `env:"POSTGRES_USER" env-required:"true"`
		Pass string `env:"POSTGRES_PASSWORD" env-required:"true"`
		Name string `env:"POSTGRES_DB" env-required:"true"`
	}

	flags struct {
		ConfigPath     string
		WithMigrations bool
	}
)

func LoadConfigs() (Config, error) {
	var cfg Config
	cfg.Flags = loadFlags()

	if cfg.Flags.ConfigPath != "" {
		if err := cleanenv.ReadConfig(cfg.Flags.ConfigPath, &cfg); err != nil {
			return cfg, err
		}
	} else {
		if err := cleanenv.ReadEnv(&cfg); err != nil {
			return cfg, err
		}
	}

	return cfg, nil
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

func loadFlags() flags {
	var f flags

	flag.StringVar(&f.ConfigPath, "config", "", "path to config file")
	flag.BoolVar(&f.WithMigrations, "with-migrations", false, "run migrations")

	flag.Parse()

	return f
}
