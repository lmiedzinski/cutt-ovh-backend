package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/lmiedzinski/cutt-ovh-backend/config"
	"github.com/lmiedzinski/cutt-ovh-backend/internal/controller/http/apiv1"
	"github.com/lmiedzinski/cutt-ovh-backend/internal/usecase"
	"github.com/lmiedzinski/cutt-ovh-backend/internal/usecase/repository"
	"github.com/lmiedzinski/cutt-ovh-backend/pkg/httpserver"
	"github.com/lmiedzinski/cutt-ovh-backend/pkg/logger"
	"github.com/lmiedzinski/cutt-ovh-backend/pkg/postgres"

	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {
	logger := logger.New(cfg.LogLevel)
	logger.Info("service starting")

	// Repositories setup
	pg, err := postgres.New(logger, cfg.PostgresConnectionString)
	if err != nil {
		logger.Fatal(err)
	}
	defer pg.Close()

	redirectUseCase := usecase.NewRedirectUseCase(repository.NewRedirectPostgresRepository(pg))

	// HTTP Server
	handler := gin.New()
	apiv1.NewRouter(handler, logger, redirectUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.PortNumber))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Info("signal: " + s.String())
	case err = <-httpServer.Notify():
		logger.Error(err)
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		logger.Error(err)
	}
}
