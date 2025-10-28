package entities

import (
	"waiter/domain/product/value_objects"

	"github.com/google/uuid"
)

type Product struct {
	id    uuid.UUID
	name  string
	price value_objects.Price
}
