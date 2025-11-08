package entities

import (
	"errors"
	"time"
	"waiter/domain/product/value_objects"

	"github.com/google/uuid"
)

type Supply struct {
	id          uuid.UUID
	quantity    float64
	productId   uuid.UUID
	incomePrice value_objects.Price
	salePrice   value_objects.Price
	createdAt   time.Time
	branchid    uuid.UUID
	supplierId  uuid.UUID
	userId      uuid.UUID
}

func NewSupply(
	productId uuid.UUID,
	quantity float64,
	incomePrice value_objects.Price,
	branchid uuid.UUID,
	salePrice value_objects.Price,
	supplierId uuid.UUID,
	userId uuid.UUID,
) (Supply, error) {
	if quantity <= 0 {
		return Supply{}, errors.New("quantity must be positive")
	}
	if !incomePrice.IsPositiveAmount() {
		return Supply{}, errors.New("incomePrice should be positive")
	}
	if !salePrice.IsPositiveAmount() {
		return Supply{}, errors.New("salePrice should be positive")
	}
	if productId == uuid.Nil {
		return Supply{}, errors.New("productId is required")
	}
	if branchid == uuid.Nil {
		return Supply{}, errors.New("branchId is required")
	}
	return Supply{
		uuid.New(),
		quantity,
		productId,
		incomePrice,
		salePrice,
		time.Now(),
		branchid,
		supplierId,
		userId,
	}, nil
}
