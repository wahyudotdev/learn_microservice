package services

import (
	"api_gateway/config"
	"api_gateway/services/product"
)

var ProductCommandService = func() product.ProductCommandServiceClient {
	return product.NewProductCommandServiceClient(config.ProductCService)
}()

var ProductQueryService = func() product.ProductQueryServiceClient {
	return product.NewProductQueryServiceClient(config.ProductQService)
}()
