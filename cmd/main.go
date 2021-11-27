package main

import (
	"github.com/lmiedzinski/cutt-ovh-backend/config"
	"github.com/lmiedzinski/cutt-ovh-backend/internal/app"
)

func main() {
	cfg := config.GetConfig()
	app.Run(cfg)
}
