package product

import (
	"errors"
	"waiter/domain/product/entities"
	"waiter/domain/product/repositories"
)

type GetProduct struct {
	repo repositories.ProductRepository
}

func NewGetProduct(repo repositories.ProductRepository) *GetProduct {
	return &GetProduct{repo: repo}
}

func (uc *GetProduct) Execute(productId string) (*entities.Product, error) {
	if productId == "" {
		return nil, errors.New("product id can not be empty")
	}
	return uc.repo.Get(productId)
}
