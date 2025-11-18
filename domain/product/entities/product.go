package entities

import (
	"errors"
	"time"
	"waiter/domain/product/enums"

	"github.com/google/uuid"
)

type Product struct {
	id          uuid.UUID
	Name        string
	createdAt   time.Time
	Measure     enums.MeasureType
	Code        string
	Description string
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

func (p *Product) Update(
	name string,
	measureType enums.MeasureType,
	code string,
	description string) (Product, error) {
	if name == "" || code == "" {
		return Product{}, errors.New("product name or code is empty")
	}
	if !enums.IsValid(measureType) {
		return Product{}, errors.New("invalid measure type")
	}
	p.Name = name
	p.Measure = measureType
	p.Code = code
	p.Description = description
	return *p, nil
}
