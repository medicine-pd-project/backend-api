package definitions

import (
	"github.com/medicine-pd-project/backend-common/logger/logrus"
	"github.com/sarulabs/di"
)

const CustomLogger = "custom_logger"

func getLogger() di.Def {
	return di.Def{
		Name:  CustomLogger,
		Scope: di.App,
		Build: func(ctn di.Container) (any, error) {

			return logrus.New(), nil
		},
	}
}
