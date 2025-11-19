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

func (u *AddProduct) Execute(command contracts.AddProductCommand) (bool, error) {
	product, err := entities.NewProduct(
		command.Name,
		enums.MeasureType(command.Measure),
		command.Code,
		command.Description,
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
