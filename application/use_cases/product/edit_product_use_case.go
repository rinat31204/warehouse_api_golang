package product

import (
	"waiter/application/contracts"
)

type EditProductUseCase interface {
	Execute(command contracts.EditProductCommand) error
}
