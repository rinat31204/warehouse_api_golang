package product

import (
	"waiter/application/contracts"
)

type AddProductUseCase interface {
	Execute(command contracts.AddProductCommand) error
}
