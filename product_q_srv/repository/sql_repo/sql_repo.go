package sql_repo

import (
	"context"
	"gorm.io/gorm"
	"product_q_srv/models"
)

type SqlRepo struct {
	Db *gorm.DB
}

func New(db *gorm.DB) SqlRepo {
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.ProductSearch{})

	return SqlRepo{
		Db: db,
	}
}

func (s SqlRepo) Search(ctx context.Context, query string) (*[]models.ProductSearch, error) {
	products := make([]models.ProductSearch, 0)
	q := "%" + query + "%"
	tx := s.Db.WithContext(ctx).Raw("SELECT * FROM product_search WHERE name LIKE ?", q).Scan(&products)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &products, nil
}
