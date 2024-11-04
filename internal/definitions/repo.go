package definitions

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sarulabs/di"

	"github.com/medicine-pd-project/backend-api/internal/infrastructure/repo/operatorrepo"
)

const (
	operatorRepo = "operator_repo"
)

func getOperatorRepo() di.Def {
	return di.Def{
		Name:  operatorRepo,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			pool, _ := ctn.Get(postgres).(*pgxpool.Pool)

			return operatorrepo.New(pool), nil
		},
	}
}
