package main

import (
	"log"

	"github.com/sm/contact-api/config"
	"github.com/sm/contact-api/internal/adapters/framework/left/rest"
	"github.com/sm/contact-api/internal/adapters/framework/right/db"
	"github.com/sm/contact-api/internal/application/api"
)

var appConfig = config.Config{}

func init() {
	appConfig.Read()
}


func main() {
	dbAdapter, err := db.NewAdapter(appConfig)
	if err != nil {
		log.Fatal(err)
	}

	appApi := api.NewApplication(dbAdapter)
	restAdapter := rest.NewAdapter(appApi)
	restAdapter.Run()
}