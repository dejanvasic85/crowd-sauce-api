package api

import (
	"fmt"
	"github.com/dejanvasic85/matchday-api/pkg/utl/config"
	"github.com/dejanvasic85/matchday-api/pkg/utl/database"
	"github.com/dejanvasic85/matchday-api/pkg/utl/server"
)

func Start(cfg *config.Configuration) error {
	fmt.Println("Opening Database connection...")
	_, err := database.New(cfg.DB.PSN, cfg.DB.Timeout, cfg.DB.LogQueries)
	if err != nil {
		return err
	}

	e := server.New()

	server.Start(e, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})

	return nil
}