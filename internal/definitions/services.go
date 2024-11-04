package definitions

import (
	"github.com/sarulabs/di"

	"github.com/medicine-pd-project/backend-api/internal/infrastructure/repo/operatorrepo"
	"github.com/medicine-pd-project/backend-api/internal/usecase/authservice"
)

const (
	authService = "auth_service"
)

func getAuthService() di.Def {
	return di.Def{
		Name:  authService,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			repo, _ := ctn.Get(operatorRepo).(*operatorrepo.Repo)

			return authservice.New(repo), nil
		},
	}
}
