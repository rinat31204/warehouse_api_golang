package product

import (
	"waiter/application/commands"
	"waiter/domain/product/enums"
	"waiter/domain/product/repositories"
)

type EditProduct struct {
	repo repositories.ProductRepository
}

func NewEditProduct(repo repositories.ProductRepository) *EditProduct {
	return &EditProduct{
		repo: repo,
	}
}

func (p *EditProduct) Execute(command commands.EditProductCommand) error {
	product, err := p.repo.Get(command.Id)
	if err != nil {
		return err
	}

	updatedProduct, err := product.Update(
		command.Name,
		enums.MeasureType(command.Measure),
		command.Code,
		command.Description)
	if err != nil {
		return err
	}

	err = p.repo.Edit(updatedProduct)
	if err != nil {
		return err
	}
	return nil
}
