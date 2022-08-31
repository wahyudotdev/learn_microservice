package sqlite_repo

import (
	"context"
	"gorm.io/gorm"
	"product_q_srv/models"
)

type SqliteRepo struct {
	Db *gorm.DB
}

func New(db *gorm.DB) SqliteRepo {
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.ProductSearch{})

	return SqliteRepo{
		Db: db,
	}
}

func (s SqliteRepo) Search(ctx context.Context, page int64, query string) (*[]models.ProductSearch, error) {
	products := make([]models.ProductSearch, 0)
	tx := s.Db.WithContext(ctx).Raw("SELECT * FROM product_search").Scan(&products)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &products, nil
}
