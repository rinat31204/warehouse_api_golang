package entities

import (
	"errors"
	"time"
	"unicode"
	"waiter/domain/product/enums"

	"github.com/google/uuid"
)

type Product struct {
	Id          uuid.UUID
	Name        string
	CreatedAt   time.Time
	Measure     enums.MeasureType
	Code        string
	Description string
}

func NewProduct(
	name string,
	measureType enums.MeasureType,
	code string,
	description string) (*Product, error) {
	e := validate(name, measureType, code)
	if e != nil {
		return nil, e
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

func validate(name string, measureType enums.MeasureType, code string) error {
	if name == "" {
		return errors.New("product name is empty")
	}
	err := checkValidCode(code)
	if err != nil {
		return err
	}
	if !enums.IsValid(measureType) {
		return errors.New("invalid measure type")
	}
	return nil
}

func checkValidCode(code string) error {
	if code == "" {
		return errors.New("product code is empty")
	}
	for _, char := range code {
		if !unicode.IsDigit(char) {
			return errors.New("invalid code")
		}
	}
	return nil
}

func (p *Product) Update(
	name string,
	measureType enums.MeasureType,
	code string,
	description string) (Product, error) {
	e := validate(name, measureType, code)
	if e != nil {
		return Product{}, e
	}
	p.Name = name
	p.Measure = measureType
	p.Code = code
	p.Description = description
	return *p, nil
}
