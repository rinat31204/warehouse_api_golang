package test

import (
	"testing"
	"time"
	"waiter/domain/product/entities"
	"waiter/domain/product/enums"

	uuid2 "github.com/google/uuid"
)

func TestGetAllProduct_Success(t *testing.T) {

}

func TestGetAllProduct_Fail(t *testing.T) {

}

func createGetAllProductMock(t *testing.T, filter ProductFilter) (*ProductRepositoryMock, error) {
	mock := &ProductRepositoryMock{
		GetAllFunc: func() ([]*entities.Product, error) {
			return []*entities.Product{
				{
					Id:        uuid2.MustParse("123e4567-e89b-12d3-a456-426614174000"),
					Name:      "cola",
					CreatedAt: time.Now(),
					Measure:   enums.Thing,
					Code:      "123456789",
				},
				{
					Id:        uuid2.MustParse("123e4567-e89b-12d3-a456-426614174111"),
					Name:      "fanta",
					CreatedAt: time.Now(),
					Measure:   enums.Liter,
					Code:      "987654321",
				},
			}, nil
		},
	}
	return mock, nil
}
