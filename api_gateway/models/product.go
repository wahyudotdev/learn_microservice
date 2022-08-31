package models

type Product struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Tag         string `json:"tag" form:"tag"`
	CreatedAt   int64  `json:"created_at" form:"created_at"`
	UpdatedAt   int64  `json:"updated_at" form:"updated_at"`
	Stock       int64  `json:"stock" form:"stock"`
}

type AddProduct struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Tag         string `json:"tag" form:"tag"`
	Stock       int64  `json:"stock" form:"stock"`
}
