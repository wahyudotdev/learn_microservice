package repository

import (
	"context"
	"product_q_srv/models"
)

type Repository interface {
	Search(ctx context.Context, query string) (*[]models.ProductSearch, error)
}
