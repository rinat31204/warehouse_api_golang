package entities

import (
	"errors"
	"time"
	"waiter/domain/product/value_objects"

	"github.com/google/uuid"
)

type Sale struct {
	id        uuid.UUID
	price     value_objects.Price
	productId uuid.UUID
	createdAt time.Time
	userId    uuid.UUID
	quantity  float64
	branchId  uuid.UUID
}

func NewSale(
	productId uuid.UUID,
	quantity float64,
	price value_objects.Price,
	userId uuid.UUID,
	branchId uuid.UUID,
) (Sale, error) {
	if productId == uuid.Nil {
		return Sale{}, errors.New("productId is required")
	}
	if quantity <= 0 {
		return Sale{}, errors.New("quantity is required")
	}
	if userId == uuid.Nil {
		return Sale{}, errors.New("userId is required")
	}
	if branchId == uuid.Nil {
		return Sale{}, errors.New("branchId is required")
	}
	if !price.IsPositiveAmount() {
		return Sale{}, errors.New("price is required")
	}
	return Sale{
		id:        uuid.New(),
		price:     price,
		productId: productId,
		quantity:  quantity,
		createdAt: time.Now(),
		userId:    userId,
		branchId:  branchId,
	}, nil
}
