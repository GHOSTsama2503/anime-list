package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Environment struct {
	// app
	Port      int    `env:"PORT" envDefault:"8000"`
	JWTSecret string `env:"JWT_SECRET"`

	// api
	MaxLimit     int64 `env:"MAX_LIMIT"`
	DefaultLimit int64 `env:"DEFAULT_LIMIT"`

	// database
	DatabaseUrl string `env:"DATABASE_URL"`

	// logs
	LogsAppPath    string `env:"LOGS_APP_PATH" envDefault:"logs/app.log"`
	LogsAccessPath string `env:"LOGS_ACCESS_PATH" envDefault:"logs/access.log"`
	LogsMaxSize    int    `env:"LOGS_MAX_SIZE"`
	LogsMaxBackup  int    `env:"LOGS_MAX_BACKUP"`
	LogsMaxAge     int    `env:"LOGS_MAX_AGE"`
	LogsCompress   bool   `env:"LOGS_COMPRESS"`
	LogsLocalTime  bool   `env:"LOGS_LOCAL_TIME"`

	IsProduction bool `env:"PRODUCTION"`
}

var Env *Environment

func LoadEnv(filenames ...string) error {
	godotenv.Load(filenames...)

	environment := &Environment{}

	if err := env.Parse(environment); err != nil {
		return err
	}

	Env = environment

	return nil
}
