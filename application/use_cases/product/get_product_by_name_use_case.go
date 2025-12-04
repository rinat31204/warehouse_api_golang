package product

import "waiter/domain/product/entities"

type GetByNameUseCase interface {
	GetByName(name string) ([]*entities.Product, error)
}
