package repositories

import "waiter/domain/product/entities"

type ProductRepository interface {
	Add(product *entities.Product) error
	GetByCode(code string) (*entities.Product, error)
	Edit(product *entities.Product) error
	GetByName(name string) ([]*entities.Product, error)
	GetAll() ([]*entities.Product, error)
	Get(uuid string) (*entities.Product, error)
}
