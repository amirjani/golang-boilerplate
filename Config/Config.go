package Config

import "time"

type EnvironmentConfig struct {
	Api struct {
		ApiHost         string        `env:"API_HOST" env-default:"localhost:3000"`
		ReadTimeOut     time.Duration `env:"API_READ_TIMEOUT" env-default:"5s"`
		WriteTimeOut    time.Duration `env:"API_WRITE_TIMEOUT" env-default:"5s"`
		ShutdownTimeout time.Duration `env:"API_SHUT_DOWN_TIMEOUT" env-default:"5s"`
	}
	DB struct {
		User         string `env:"DB_USER" env-default:"postgres"`
		Password     string `env:"DB_PASSWORD" env-default:"postgres"`
		Host         string `env:"DB_HOST" env-default:"localhost"`
		Name         string `env:"DB_NAME" env-default:"postgres"`
		MaxIdleConns int    `env:"DB_MAX_IDLE_CONNS" env-default:"10"`
		MaxOpenConns int    `env:"DB_MAX_OPEN_CONNS" env-default:"100"`
		DisableTLS   bool   `env:"DB_DISABLE_TLS" env-default:"false"`
	}
	Redis struct {
		Host     string `env:"REDIS_HOST" env-default:"localhost"`
		Password string `env:"REDIS_PASSWORD" env-default:""`
		Port     string `env:"REDIS_PORT" env-default:"6379"`
	}
}
