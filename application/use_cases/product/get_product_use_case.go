package product

import "waiter/domain/product/entities"

type GetProductUseCase interface {
	Execute(productId string) (*entities.Product, error)
}
