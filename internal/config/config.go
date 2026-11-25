package config

import (
	"log/slog"
	"time"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	AppShutdownTimeout   time.Duration `env:"APP_SHUTDOWN_TIMEOUT" envDefault:"20s"`
	AppReadTimeout       time.Duration `env:"APP_READ_TIMEOUT" envDefault:"60s"`
	AppWriteTimeout      time.Duration `env:"APP_WRITE_TIMEOUT" envDefault:"60s"`
	AppReadHeaderTimeout time.Duration `env:"APP_READ_HEADER_TIMEOUT" envDefault:"5s"`
	AppHTTPPort          int           `env:"APP_HTTP_PORT" envDefault:"80"`

	AppLogLevel   slog.Level `env:"APP_LOG_LEVEL" envDefault:"debug"`
	AppLogHandler string     `env:"APP_LOG_HANDLER" envDefault:"text"`

	PostgresHost string `env:"POSTGRES_HOST" envDefault:"postgres"`
	PostgresPort int    `env:"POSTGRES_PORT" envDefault:"5432"`
	PostgresUser string `env:"POSTGRES_USER" envDefault:"postgres_user"`
	PostgresPass string `env:"POSTGRES_PASSWORD" envDefault:"postgres_pass"`
	PostgresDB   string `env:"POSTGRES_DB" envDefault:"tickets"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
