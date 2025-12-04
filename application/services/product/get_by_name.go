package product

import (
	"errors"
	"waiter/domain/product/entities"
	"waiter/domain/product/repositories"
)

type GetByName struct {
	repo repositories.ProductRepository
}

func NewGetByName(repo repositories.ProductRepository) GetByName {
	return GetByName{repo: repo}
}

func (u GetByName) Execute(name string) ([]*entities.Product, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}
	return u.repo.GetByName(name)
}
