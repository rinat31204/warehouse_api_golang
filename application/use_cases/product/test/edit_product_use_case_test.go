package test

import (
	"strings"
	"sync/atomic"
	"testing"
	"time"
	"waiter/application/contracts"
	"waiter/application/services/product"
	"waiter/domain/product/entities"
	"waiter/domain/product/enums"

	uuid2 "github.com/google/uuid"
)

func TestEditProduct_Success(t *testing.T) {
	command := contracts.EditProductCommand{
		Id:          "123e4567-e89b-12d3-a456-426614174000",
		Name:        "fanta",
		Measure:     1,
		Code:        "987654321",
		Description: "test description",
	}
	mock, editCalled, getCalled := createEditProductMock(t, command)

	uc := product.NewEditProduct(mock)
	err := uc.Execute(command)
	if err != nil {
		t.Error(err)
	}
	if atomic.LoadInt32(editCalled) != 1 {
		t.Errorf("Add was called %d times", editCalled)
	}
	if atomic.LoadInt32(getCalled) != 1 {
		t.Errorf("Add was called %d times", getCalled)
	}

}

func createEditProductMock(t *testing.T, command contracts.EditProductCommand) (*ProductRepositoryMock, *int32, *int32) {
	var editCalled int32
	var getCalled int32
	mock := &ProductRepositoryMock{
		EditFunc: func(product *entities.Product) error {
			atomic.AddInt32(&editCalled, 1)
			if product.Name != command.Name {
				t.Errorf("product.Name should be %s, got %s", command.Name, product.Name)
			}
			if product.Code != command.Code {
				t.Errorf("product.Code should be %s, got %s", command.Code, product.Code)
			}
			if product.Measure != enums.MeasureType(command.Measure) {
				t.Errorf("product.Measure should be %d, got %d", 0, product.Measure)
			}
			if product.Description != command.Description {
				t.Errorf("product.Description should be %s, got %s", command.Description, product.Description)
			}
			return nil
		},
		GetFunc: func(uuid string) (*entities.Product, error) {
			atomic.AddInt32(&getCalled, 1)
			return &entities.Product{
				Id:          uuid2.MustParse("123e4567-e89b-12d3-a456-426614174000"),
				Name:        "cola",
				CreatedAt:   time.Date(2025, 11, 19, 0, 0, 0, 0, time.UTC),
				Measure:     enums.Thing,
				Code:        "123456789",
				Description: "",
			}, nil
		},
	}
	return mock, &editCalled, &getCalled
}

func TestEditProductName_Fail(t *testing.T) {
	command := contracts.EditProductCommand{
		Id:          "123e4567-e89b-12d3-a456-426614174000",
		Name:        "",
		Measure:     1,
		Code:        "987654321",
		Description: "test description",
	}
	mock, editCalled, getCalled := createEditProductMock(t, command)

	uc := product.NewEditProduct(mock)
	err := uc.Execute(command)

	if err == nil {
		t.Fatal("expected error when creating product with empty, got nil")
	}
	expected := "product name is empty"
	if !strings.Contains(err.Error(), expected) {
		t.Errorf("expected error to contain %q, got %q", expected, err.Error())
	}
	if atomic.LoadInt32(editCalled) != 0 {
		t.Errorf("Edit was called %d times", editCalled)
	}
	if atomic.LoadInt32(getCalled) != 1 {
		t.Errorf("Get was called %d times", getCalled)
	}
}

func TestEditProductCode_Fail(t *testing.T) {
	command := contracts.EditProductCommand{
		Id:          "123e4567-e89b-12d3-a456-426614174000",
		Name:        "sdgsdfsfd",
		Measure:     1,
		Code:        "243263fghsgs",
		Description: "test description",
	}
	mock, editCalled, getCalled := createEditProductMock(t, command)

	uc := product.NewEditProduct(mock)
	err := uc.Execute(command)

	if err == nil {
		t.Fatal("expected error when creating product with empty, got nil")
	}
	expected := "invalid code"
	if !strings.Contains(err.Error(), expected) {
		t.Errorf("expected error to contain %q, got %q", expected, err.Error())
	}
	if atomic.LoadInt32(editCalled) != 0 {
		t.Errorf("Edit was called %d times", editCalled)
	}
	if atomic.LoadInt32(getCalled) != 1 {
		t.Errorf("Get was called %d times", getCalled)
	}
}
