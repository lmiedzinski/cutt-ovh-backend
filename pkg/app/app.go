package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/lmiedzinski/cutt-ovh-backend/config"
	"github.com/lmiedzinski/cutt-ovh-backend/internal/httpserver"
	"github.com/lmiedzinski/cutt-ovh-backend/internal/logger"
	"github.com/lmiedzinski/cutt-ovh-backend/internal/postgres"
	"github.com/lmiedzinski/cutt-ovh-backend/pkg/domain/redirect"
	"github.com/lmiedzinski/cutt-ovh-backend/pkg/transport/http/apiv1"

	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {
	logger := logger.New(cfg.LogLevel)
	logger.Info("service starting")

	// Repositories setup
	pg, err := postgres.New(cfg.PostgresConnectionString)
	if err != nil {
		logger.Fatal(err)
	}
	defer pg.Close()
	redirectRepository := redirect.NewRedirectPostgresRepository(pg)

	// Services setup
	redirectService := redirect.NewRedirectService(redirectRepository)

	// HTTP Server
	handler := gin.New()
	redirectHandler := redirect.NewRedirectHttpHandler(redirectService, logger)
	apiv1.AddHttpRouter(handler, logger, redirectHandler)
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
