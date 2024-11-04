package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/medicine-pd-project/backend-common/logger"

	"github.com/medicine-pd-project/backend-api/internal/entity"
)

func (h *Handler) Login(ctx echo.Context) error {
	log := h.log.WithFields(logger.Fields{
		logger.Action: "Login",
	})

	var req entity.LoginOperatorRequest

	err := ctx.Bind(&req)
	if err != nil {
		log.Errorf("failed to bind request: %v", err)

		// TODO: error remaper middleware
		return ctx.JSON(http.StatusUnprocessableEntity, echo.Map{
			"error": "failed to bind request",
		})
	}

	token, err := h.srv.Login(ctx.Request().Context(), log, req)
	if err != nil {
		log.Errorf("failed to login: %v", err)

		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error": "failed to login",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}
