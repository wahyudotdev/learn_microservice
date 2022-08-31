package repository

import (
	"product_listener_srv/models"
)

type Repository interface {
	Add(data *models.Product) error
}
