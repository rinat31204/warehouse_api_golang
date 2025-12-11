package product

import (
	"waiter/application/commands"
)

type AddProductUseCase interface {
	Execute(command commands.AddProductCommand) error
}
