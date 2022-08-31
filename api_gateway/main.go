package main

import (
	"api_gateway/config"
	_ "api_gateway/docs"
	"api_gateway/router"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"log"
)

// @title API
// @version 1.0
// @description Public API
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api
func main() {
	app := fiber.New()
	app.Use(config.Cors)
	app.Get("/health", healthCheck())
	app.Get("/*", swagger.HandlerDefault)
	group := app.Group("/api")
	router.ProductRouter(group)
	log.Fatalln(app.Listen(config.Port))
}

func healthCheck() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("api-gateway is running")
	}
}
