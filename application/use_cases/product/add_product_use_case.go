package product

import "waiter/application/contracts"

type AddProductUseCase interface {
	Execute(request contracts.AddProductCommand) (result bool, err error)
}
