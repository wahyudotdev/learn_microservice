package sql_repo

import (
	"context"
	"gorm.io/gorm"
	"product_listener_srv/config"
	"product_listener_srv/models"
	"product_listener_srv/utils"
)

type SqlRepo struct {
	Db *gorm.DB
}

func New(db *gorm.DB) SqlRepo {
	return SqlRepo{
		Db: db,
	}
}

func (s SqlRepo) Add(data *models.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
	defer cancel()
	if err := s.Db.WithContext(ctx).Create(&data); err.Error != nil {
		return err.Error
	}
	searchData, _ := utils.TypeConverter[models.ProductSearch](&data)
	if err := s.Db.WithContext(ctx).Create(&searchData); err.Error != nil {
		return err.Error
	}
	return nil
}
