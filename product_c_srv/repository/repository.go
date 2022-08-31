package repository

import (
	"context"
	"product_c_srv/models"
)

type Repository interface {
	Add(ctx context.Context, data *models.Product) error
}
