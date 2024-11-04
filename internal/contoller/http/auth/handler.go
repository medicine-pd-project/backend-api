package auth

import (
	"context"

	"github.com/medicine-pd-project/backend-common/logger"

	"github.com/medicine-pd-project/backend-api/internal/entity"
)

type authService interface {
	Login(ctx context.Context, log logger.Logger, req entity.LoginOperatorRequest) (entity.JWTToken, error)
}

type Handler struct {
	log logger.Logger
	srv authService
}

func NewHandler(log logger.Logger, srv authService) *Handler {
	return &Handler{
		log: log.WithFields(logger.Fields{
			logger.Module: "auth_module",
		}),
		srv: srv,
	}
}
