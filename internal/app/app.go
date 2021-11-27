package app

import (
	"github.com/lmiedzinski/cutt-ovh-backend/config"
	"github.com/lmiedzinski/cutt-ovh-backend/pkg/logger"
)

func Run(cfg *config.Config) {
	logger := logger.New(cfg.LogLevel)
	logger.Info("service starting")

}
