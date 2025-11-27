package test

import (
	"testing"
	"waiter/domain/product/entities"
)

func TestGetAllProduct_Success(t *testing.T) {

}

func TestGetAllProduct_Fail(t *testing.T) {

}

func createGetAllProductMock(t *testing.T, filter ProductFilter) (*ProductRepositoryMock, error) {
	mock := &ProductRepositoryMock{
		GetAllFunc: func() ([]*entities.Product, error) {

			return []*entities.Product{}, nil
		},
	}
}
