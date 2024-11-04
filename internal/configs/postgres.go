package configs

import "time"

type Postgres struct {
	DSN             string        `envconfig:"POSTGRES_DSN" required:"true"` // В конфиге нет, хранится в vault
	MigrationsDir   string        `envconfig:"MIGRATIONS_DIR" required:"true"`
	MaxOpenConns    int32         `envconfig:"POSTGRES_MAX_OPEN_CONNS" default:"100"`
	MaxIdleConns    int32         `envconfig:"POSTGRES_MAX_IDLE_CONNS" default:"10"`
	ConnMaxLifetime time.Duration `envconfig:"POSTGRES_CONN_MAX_LIFETIME" default:"10m"`
}
