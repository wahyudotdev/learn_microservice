package handler

import (
	"api_gateway/models"
	"api_gateway/services"
	"api_gateway/services/product"
	"api_gateway/utils"
	"context"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// AddProduct
// @Summary Add product
// @Tags Product
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param body formData models.AddProduct true "Add product"
// @Success 200 {object} models.Product
// @Router /product/add [post]
func AddProduct() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var reqBody models.AddProduct
		if err := c.BodyParser(&reqBody); err != nil {
			return c.Status(http.StatusBadRequest).JSON("bad request")
		}
		data, _ := utils.TypeConverter[product.AddProduct](&reqBody)
		result, err := services.ProductCommandService.Add(context.Background(), data)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(err.Error())
		}
		return c.JSON(result)
	}
}

// SearchProduct
// @Summary Search Product
// @Tags Product
// @Produce  json
// @Param q query string  false "Search query"
// @Success 200 {object} []models.Product
// @Failure 401 {object} models.AuthFailedResponse
// @Router /product/search [get]
func SearchProduct() fiber.Handler {
	return func(c *fiber.Ctx) error {
		q := c.Query("q")
		result, err := services.ProductQueryService.Search(context.Background(), &product.SearchProduct{
			Q: q,
		})
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(err.Error())
		}
		return c.JSON(result)
	}
}
