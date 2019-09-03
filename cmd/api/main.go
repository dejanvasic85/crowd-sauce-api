package main

import (
	"flag"

	"github.com/dejanvasic85/matchday-api/pkg/utl/config"
	"github.com/dejanvasic85/matchday-api/pkg/api"
)

func main() {
	configPath := flag.String("p", "./cmd/api/conf.local.yml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	checkErr(err)
	checkErr(api.Start(cfg))
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}