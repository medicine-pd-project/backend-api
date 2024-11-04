package definitions

import (
	"github.com/medicine-pd-project/backend-common/logger"
	"github.com/sarulabs/di"

	"github.com/medicine-pd-project/backend-api/internal/contoller/http/auth"
	"github.com/medicine-pd-project/backend-api/internal/usecase/authservice"
)

const (
	authHandler = "auth_handler"
)

func getAuthHandler() di.Def {
	return di.Def{
		Name:  authHandler,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			log, _ := ctn.Get(CustomLogger).(logger.Logger)
			srv, _ := ctn.Get(authService).(*authservice.Service)

			return auth.NewHandler(log, srv), nil
		},
	}
}
