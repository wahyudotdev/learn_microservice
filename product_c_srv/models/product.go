package models

type Product struct {
	Id          string `json:"id" gorm:"primaryKey,column:id"`
	Name        string `json:"name" gorm:"column:name"`
	Stock       int64  `json:"stock" gorm:"column:stock"`
	Description string `json:"description" gorm:"column:description"`
}

func (Product) TableName() string {
	return "product"
}
