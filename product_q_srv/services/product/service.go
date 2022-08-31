package product

import (
	"context"
	"product_q_srv/repository"
	"product_q_srv/utils"
)

type Service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s Service) Search(ctx context.Context, product *SearchProduct) (*SearchProductResult, error) {
	data, err := s.repository.Search(ctx, product.Page, product.Q)
	if err != nil {
		return nil, err
	}
	result, err := utils.ListConverter[*ProductShowcase](data)
	if err != nil {
		return nil, err
	}
	return &SearchProductResult{
		Products: result,
	}, nil
}

func (s Service) mustEmbedUnimplementedProductQueryServiceServer() {
	panic("implement me")
}
