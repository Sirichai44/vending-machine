/*
How to start the app:
go version >= 1.23
docker-compose up

swagger docs:
http://localhost:8080/api/v1/swagger/index.html#/
*/

package main

import (
	"context"
	"log"

	"vending_machine/app"
	"vending_machine/config"
)

var conf *config.AppConfig

func init() {
	var err error
	conf, err = config.NewAppConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
}

func main() {
	ctx := context.Background()

	app, err := app.NewApp(conf)
	if err != nil {
		log.Fatalf("Failed to create app: %v", err)
	}

	if err := app.Start(ctx); err != nil {
		log.Fatalf("Failed to start app: %v", err)
		panic(err)
	}
}
