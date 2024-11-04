package definitions

import (
	"github.com/sarulabs/di"

	"github.com/medicine-pd-project/backend-api/internal/contoller/http/auth"
	"github.com/medicine-pd-project/backend-api/internal/contoller/http/server"
)

const (
	HTTPServer = "http_server"
)

func getHTTPRouter() di.Def {
	return di.Def{
		Name:  HTTPServer,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			authHandlers, _ := ctn.Get(authHandler).(*auth.Handler)

			return server.NewServer(&server.Config{
				AuthHandlers: authHandlers,
			}), nil
		},
	}
}
