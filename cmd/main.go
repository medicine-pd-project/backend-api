package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/medicine-pd-project/backend-common/logger"
	"github.com/pkg/errors"

	"github.com/medicine-pd-project/backend-api/internal/configs"
	"github.com/medicine-pd-project/backend-api/internal/definitions"
)

func main() {
	di, err := definitions.BuildApp()
	if err != nil {
		log.Fatalf("failed to create di, error: %v", err)
	}

	logg, _ := di.Get(definitions.CustomLogger).(logger.Logger)
	cfg, _ := di.Get(definitions.Config).(configs.Config)
	server := di.Get(definitions.HTTPServer).(*echo.Echo)

	go func() {
		err = server.Start(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port))
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				logg.Fatalf("REST server started to listen port [%d]: %v", cfg.Server.Port, err)
			}
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan

	logg.Infof("start graceful shutdown, caught sig: %+v", sig)

	err = server.Shutdown(context.Background())
	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			logg.Infof("REST server shutdown: %v", err)
		}
	} else {
		logg.Infof("REST server stopped")
	}
}
