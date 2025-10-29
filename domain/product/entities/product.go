package entities

import (
	"errors"
	"time"
	"waiter/domain/product/enums"

	"github.com/google/uuid"
)

type Product struct {
	id          uuid.UUID
	name        string
	dateCreate  time.Time
	measure     enums.MeasureType
	code        string
	description string
}

func NewProduct(
	name string,
	measureType enums.MeasureType,
	code string,
	description string) (*Product, error) {
	if name == "" || code == "" {
		return nil, errors.New("product name or code is empty")
	}

	if !enums.IsValid(measureType) {
		return nil, errors.New("invalid measure type")
	}

	return &Product{
		uuid.New(),
		name,
		time.Now(),
		measureType,
		code,
		description,
	}, nil
}
