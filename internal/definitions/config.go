package definitions

import (
	"github.com/sarulabs/di"

	"github.com/medicine-pd-project/backend-api/internal/configs"
)

const Config = "Config"

func getConfig() di.Def {
	return di.Def{
		Name:  Config,
		Scope: di.App,
		Build: func(ctn di.Container) (any, error) {
			return configs.Setup()
		},
	}
}
