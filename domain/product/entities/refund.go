package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Refund struct {
	id        uuid.UUID
	createdAt time.Time
	productId uuid.UUID
	branchId  uuid.UUID
	quantity  float64
	userId    uuid.UUID
}

func NewRefund(
	productId uuid.UUID,
	quantity float64,
	branchId uuid.UUID,
	userId uuid.UUID) (*Refund, error) {

	if productId == uuid.Nil {
		return nil, errors.New(`productId must not be nil`)
	}
	if branchId == uuid.Nil {
		return nil, errors.New(`branchId must not be nil`)
	}
	if quantity <= 0 {
		return nil, errors.New(`quantity must > 0`)
	}
	if userId == uuid.Nil {
		return nil, errors.New(`userId must not be nil`)
	}
	return &Refund{
		uuid.New(),
		time.Now(),
		productId,
		branchId,
		quantity,
		userId}, nil
}
