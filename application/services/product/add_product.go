package product

import (
	"waiter/application/contracts"
	"waiter/domain/product/entities"
	"waiter/domain/product/enums"
	"waiter/domain/product/repositories"
)

type AddProduct struct {
	repo repositories.ProductRepository
}

func NewAddProduct(repo repositories.ProductRepository) *AddProduct {
	return &AddProduct{repo: repo}
}

func (u *AddProduct) Execute(request contracts.AddProductCommand) (bool, error) {
	product, err := entities.NewProduct(
		request.Name,
		enums.MeasureType(request.Measure),
		request.Code,
		request.Description,
	)
	if err != nil {
		return false, err
	}
	add, err := u.repo.Add(*product)
	if err != nil {
		return false, err
	}
	return add, nil
}
