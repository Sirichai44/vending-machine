package app

import (
	"context"
	"log"

	"vending_machine/apis"
	"vending_machine/config"
	"vending_machine/drivers"
	"vending_machine/services"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/sync/errgroup"
)

type App interface {
	Start(context.Context) error
}
type app struct {
	mysqlDB *drivers.MySQLClient
	fiber   *fiber.App
}

func NewApp(conf *config.AppConfig) (App, error) {
	// mongoDB, err := drivers.MongoDBConn(conf.DBConfig)
	// if err != nil {
	// 	return nil, err
	// }

	// srvAuth := services.NewAuthService(mongoDB.Collection(dtos.DBCollectionUsers))
	// srvList := services.NewListService(mongoDB.Collection(dtos.DBCollectionListings))
	// srvMsg := services.NewMessageService(mongoDB.Collection(dtos.DBCollectionMessages))

	// f := apis.NewFiberAPI(conf.SecretKey, mongoDB, srvAuth, srvList, srvMsg)

	mysqlDB, err := drivers.MySQLConn(conf.DBConfig)
	if err != nil {
		return nil, err
	}

	srvProduct := services.NewProductService(mysqlDB.DB)
	srvMoney := services.NewCoinService(mysqlDB.DB)

	f := apis.NewFiberAPI(conf.SecretKey, mysqlDB, srvProduct, srvMoney)

	return &app{
		mysqlDB: mysqlDB,
		fiber:   f,
	}, nil
}

func (a *app) Start(ctx context.Context) error {
	g, c := errgroup.WithContext(ctx)

	g.Go(func() error {
		path := ":8080"
		log.Printf("Server is running on %s", path)
		return a.fiber.Listen(path)
	})

	g.Go(func() error {
		<-c.Done()

		log.Println("Shutting down the server...")
		return a.fiber.ShutdownWithContext(c)
	})

	if err := g.Wait(); err != nil {
		log.Printf("Error: %v", err)
		return err
	}

	return g.Wait()
}
