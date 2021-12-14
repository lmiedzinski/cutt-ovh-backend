package main

import (
	"github.com/lmiedzinski/cutt-ovh-backend/config"
	"github.com/lmiedzinski/cutt-ovh-backend/pkg/app"
)

func main() {
	cfg := config.GetConfig()
	app.Run(cfg)
}
