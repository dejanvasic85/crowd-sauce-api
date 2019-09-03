package api

import (
	"fmt"
	"github.com/dejanvasic85/matchday-api/pkg/utl/config"
	"github.com/dejanvasic85/matchday-api/pkg/utl/database"
)

func Start(cfg *config.Configuration) error {
	fmt.Println("Opening Database connection...")
	_, err := database.New(cfg.DB.PSN, cfg.DB.Timeout, cfg.DB.LogQueries)
	if err != nil {
		return err
	}

	return nil
}