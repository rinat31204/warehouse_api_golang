package test

import (
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
	mock, editCalled, getCalled := createForEditMock(t, command)

	uc := product.NewEditProduct(mock)
	isEdited, err := uc.Execute(command)
	if err != nil {
		t.Error(err)
	}
	if !isEdited {
		t.Error("Product was not edited")
	}
	if atomic.LoadInt32(editCalled) != 1 {
		t.Errorf("Add was called %d times", editCalled)
	}
	if atomic.LoadInt32(getCalled) != 1 {
		t.Errorf("Add was called %d times", getCalled)
	}

}

func createForEditMock(t *testing.T, command contracts.EditProductCommand) (*ProductRepositoryMock, *int32, *int32) {
	var editCalled int32
	var getCalled int32
	mock := &ProductRepositoryMock{
		EditFunc: func(product entities.Product) (bool, error) {
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
			return true, nil
		},
		GetFunc: func(uuid string) (entities.Product, error) {
			atomic.AddInt32(&getCalled, 1)
			return entities.Product{
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

func TestEditProduct_Fail(t *testing.T) {

}
