package product

import (
	"errors"
	"sync/atomic"
	"testing"
	"time"
	"waiter/domain/product/entities"
	"waiter/domain/product/enums"

	uuid2 "github.com/google/uuid"
)

func TestGetProduct_Success(t *testing.T) {
	const id = "123e4567-e89b-12d3-a456-426614174000"
	mock, getCalled := createGetProductMock(t, id)

	uc := NewGetProduct(mock)
	p, err := uc.Execute(id)

	if err != nil {
		t.Fatal(err)
	}

	if atomic.LoadInt32(getCalled) != 1 {
		t.Errorf("GetProduct called %d times, expected 1", getCalled)
	}
	if p.Name != "cola" {
		t.Errorf("GetProduct returned wrong product name, expected cola, got %s", p.Name)
	}
	if p.Code != "123456789" {
		t.Errorf("GetProduct returned wrong product code, expected 123456789, got %s", p.Code)
	}
	if p.Id != uuid2.MustParse("123e4567-e89b-12d3-a456-426614174000") {
		t.Errorf("GetProduct returned wrong product id")
	}
}

func TestGetProduct_Fail(t *testing.T) {
	mock, getCalled := createGetProductMock(t, "123e4567-e89b-12d3-a456-426614174000")

	uc := NewGetProduct(mock)
	p, err := uc.Execute("")

	if err == nil {
		t.Fatal("GetProduct should have returned an error")
	}

	called := atomic.LoadInt32(getCalled)
	if called != 0 {
		t.Errorf("GetProduct called %d times, expected 0", called)
	}

	if p != nil {
		t.Errorf("Expected product to be nil, got %+v", p)
	}
}

func createGetProductMock(t *testing.T, productId string) (*ProductRepositoryMock, *int32) {
	var getCalled int32
	mock := &ProductRepositoryMock{
		GetFunc: func(uuid string) (*entities.Product, error) {
			getCalled++
			if productId == "" {
				return nil, errors.New("productId is required")
			}
			return &entities.Product{
				Id:        uuid2.MustParse("123e4567-e89b-12d3-a456-426614174000"),
				Name:      "cola",
				CreatedAt: time.Now(),
				Measure:   enums.Thing,
				Code:      "123456789",
			}, nil
		},
	}
	return mock, &getCalled
}
