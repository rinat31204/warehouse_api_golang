package product

import (
	"waiter/application/commands"
)

type EditProductUseCase interface {
	Execute(command commands.EditProductCommand) error
}
