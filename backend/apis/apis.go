package apis

import (
	"log"

	"vending_machine/drivers"
	"vending_machine/dtos"
	"vending_machine/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

var (
	appVersion = "v1.0.0"
	apiVersion = "api/v1"
	Fiber      = fiber.Config{
		ServerHeader: "vending_machine" + appVersion,
		BodyLimit:    10 * 1024 * 1024,
	}
	FiberCore = cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	})
)

func NewFiberAPI(signKey string, mysqlDB *drivers.MySQLClient, srvProduct services.ProductService, srvMoney services.CoinService) *fiber.App {
	f := fiber.New(Fiber)
	f.Use(FiberCore)
	f.Use(recover.New())

	f.Use(func(c *fiber.Ctx) error {
		log.Printf("Request: %s %s from %s", c.Method(), c.Path(), c.IP())
		return c.Next()
	})

	f.Group(apiVersion).Get("/swagger/*", fiberSwagger.WrapHandler)

	f.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(NewContext(200, "Welcome to Vending Machine API", nil))
	})

	NewHandlerProduct(f, srvProduct, srvMoney)

	return f
}

func NewContext(code int, msg string, res any) *dtos.Context {
	return &dtos.Context{
		Status:  code,
		Message: msg,
		Results: res,
	}
}
