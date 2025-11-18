package product

import (
	"strings"
	"sync/atomic"
	"testing"
	"waiter/application/contracts"
	"waiter/application/services/product"
	"waiter/domain/product/entities"
	"waiter/domain/product/enums"
)

func TestAddProduct_Success(t *testing.T) {
	request := contracts.AddProductCommand{
		Name:    "cola",
		Measure: 0,
		Code:    "123456789",
	}
	mock, addCalled := createMock(t, request)
	p := product.NewAddProduct(mock)
	isAdded, err := p.Execute(request)
	if err != nil {
		t.Error(err)
	}
	if !isAdded {
		t.Error("Product was not added")
	}
	if atomic.LoadInt32(addCalled) != 1 {
		t.Errorf("Add was called %d times", addCalled)
	}
}

func createMock(t *testing.T, request contracts.AddProductCommand) (*ProductRepositoryMock, *int32) {
	var addCalled int32
	mock := &ProductRepositoryMock{
		AddFunc: func(product entities.Product) (bool, error) {
			atomic.AddInt32(&addCalled, 1)
			if product.Name != request.Name {
				t.Errorf("product.Name should be %s, got %s", request.Name, product.Name)
			}
			if product.Code != request.Code {
				t.Errorf("product.Code should be %s, got %s", request.Code, product.Code)
			}
			if product.Measure != enums.MeasureType(request.Measure) {
				t.Errorf("product.Measure should be %d, got %d", 0, product.Measure)
			}
			return true, nil
		},
	}
	return mock, &addCalled
}

func Test_AddProduct_IsError(t *testing.T) {
	_, err := entities.NewProduct(
		"",
		enums.Thing,
		"",
		"")
	if err == nil {
		t.Fatal("expected error when creating product with empty name and code, got nil")
	}
	expected := "product name or code is empty"
	if !strings.Contains(err.Error(), expected) {
		t.Errorf("expected error to contain %q, got %q", expected, err.Error())
	}
}
