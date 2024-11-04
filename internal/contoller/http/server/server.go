package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/medicine-pd-project/backend-api/internal/contoller/http/auth"
)

type Config struct {
	AuthHandlers *auth.Handler
}

func NewServer(cfg *Config) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.POST(loginUrl, cfg.AuthHandlers.Login)

	/*jtwMiddleware := echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(entity.JWTClaims)
		},
		SigningKey: []byte(entity.JWTKey),
	})

	e.Use(jtwMiddleware)*/

	return e
}
