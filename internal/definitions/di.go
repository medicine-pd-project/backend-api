package definitions

import (
	"github.com/pkg/errors"
	"github.com/sarulabs/di"
)

func BuildApp() (di.Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create builder")
	}

	err = builder.Add([]di.Def{
		// Config
		getConfig(),

		// Logger
		getLogger(),

		// Repo
		getOperatorRepo(),

		// Services
		getAuthService(),

		// Handlers
		getAuthHandler(),

		// Router
		getHTTPRouter(),
	}...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create definitions")
	}

	return builder.Build(), nil
}
