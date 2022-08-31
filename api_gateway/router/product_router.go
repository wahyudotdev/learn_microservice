package router

import (
	"api_gateway/handler"
	"github.com/gofiber/fiber/v2"
)

func ProductRouter(router fiber.Router) {
	group := router.Group("/product")
	group.Post("/add", handler.AddProduct())
	group.Get("/search", handler.SearchProduct())
}
