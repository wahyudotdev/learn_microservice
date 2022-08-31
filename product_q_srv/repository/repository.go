package repository

import (
	"context"
	"product_q_srv/models"
)

type Repository interface {
	Search(ctx context.Context, page int64, query string) (*[]models.ProductSearch, error)
}
