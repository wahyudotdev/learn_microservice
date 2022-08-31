package product

import (
	"context"
	"product_c_srv/models"
	"product_c_srv/repository"
	"product_c_srv/utils"
)

type Service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s Service) Add(ctx context.Context, product *AddProduct) (*EmptyResponse, error) {
	data, _ := utils.TypeConverter[models.Product](&product)
	err := s.repository.Add(ctx, data)
	if err != nil {
		return nil, err
	}
	return &EmptyResponse{Message: "add product success"}, nil
}

func (Service) mustEmbedUnimplementedProductCommandServiceServer() {
	panic("implement me")
}
