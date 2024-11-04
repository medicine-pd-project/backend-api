package definitions

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sarulabs/di"

	"github.com/medicine-pd-project/backend-api/internal/configs"
	"github.com/medicine-pd-project/backend-api/internal/infrastructure/repo/database"
)

const postgres = "postgres"

func getPostgres() di.Def {
	return di.Def{
		Name:  postgres,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg, _ := ctn.Get(Config).(configs.Config)

			return database.Init(cfg.PostgresConfig)
		},
		Close: func(obj interface{}) error {
			pool, _ := obj.(*pgxpool.Pool)

			pool.Close()

			return nil
		},
	}
}
