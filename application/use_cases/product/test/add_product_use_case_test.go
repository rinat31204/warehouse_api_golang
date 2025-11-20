package test

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
	command := contracts.AddProductCommand{
		Name:    "cola",
		Measure: 0,
		Code:    "123456789",
	}
	mock, addCalled := createMock(t, command)
	uc := product.NewAddProduct(mock)
	err := uc.Execute(command)
	if err != nil {
		t.Error(err)
	}
	if atomic.LoadInt32(addCalled) != 1 {
		t.Errorf("Add was called %d times", addCalled)
	}
}

func createMock(t *testing.T, command contracts.AddProductCommand) (*ProductRepositoryMock, *int32) {
	var addCalled int32
	mock := &ProductRepositoryMock{
		AddFunc: func(product *entities.Product) error {
			atomic.AddInt32(&addCalled, 1)
			if product.Name != command.Name {
				t.Errorf("product.Name should be %s, got %s", command.Name, product.Name)
			}
			if product.Code != command.Code {
				t.Errorf("product.Code should be %s, got %s", command.Code, product.Code)
			}
			if product.Measure != enums.MeasureType(command.Measure) {
				t.Errorf("product.Measure should be %d, got %d", 0, product.Measure)
			}
			return nil
		},
	}
	return mock, &addCalled
}

func TestProductName_Fail(t *testing.T) {
	_, err := entities.NewProduct(
		"",
		enums.Thing,
		"54353543",
		"")
	if err == nil {
		t.Fatal("expected error when creating product with empty, got nil")
	}
	expected := "product name is empty"
	if !strings.Contains(err.Error(), expected) {
		t.Errorf("expected error to contain %q, got %q", expected, err.Error())
	}
}

func TestProductCode_Fail(t *testing.T) {
	_, err := entities.NewProduct(
		"hsdfadfadf",
		enums.Thing,
		"",
		"")
	if err == nil {
		t.Fatal("expected error when creating product with empty, got nil")
	}
	expected := "product code is empty"
	if !strings.Contains(err.Error(), expected) {
		t.Errorf("expected error to contain %q, got %q", expected, err.Error())
	}
}

func TestProductType_Fail(t *testing.T) {
	_, err := entities.NewProduct(
		"hsdfadfadf",
		enums.MeasureType(4),
		"5342342",
		"")
	if err == nil {
		t.Fatal("expected error when creating product with empty, got nil")
	}
	expected := "invalid measure type"
	if !strings.Contains(err.Error(), expected) {
		t.Errorf("expected error to contain %q, got %q", expected, err.Error())
	}
}

func TestProductCode_OnlyNumber(t *testing.T) {
	command := contracts.AddProductCommand{
		Name:    "cola",
		Measure: 0,
		Code:    "1234rhgsd567sdf89",
	}
	mock, addCalled := createMock(t, command)
	uc := product.NewAddProduct(mock)
	err := uc.Execute(command)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	expected := "invalid code"
	if !strings.Contains(err.Error(), expected) {
		t.Errorf("expected error to contain %q, got %q", expected, err.Error())
	}
	if atomic.LoadInt32(addCalled) != 0 {
		t.Errorf("Add was called %d times", addCalled)
	}
}
