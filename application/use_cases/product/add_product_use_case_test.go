package product

import (
	"testing"
	"waiter/domain/product/entities"
	"waiter/domain/product/enums"
)

func Test_AddProduct_WhenIsSuccess(t *testing.T) {
	product, _ := entities.NewProduct(
		"cola",
		enums.Thing,
		"123456789",
		"")

	mock := &ProductRepositoryMock{
		AddFunc: func(product entities.Product) (bool, error) {
			return true, nil
		},
	}
	isAdded, _ := mock.Add(*product)
	if !isAdded {
		t.Error("add product not saved")
	}
}

func Test_UpdateProduct_WhenIsSuccess(t *testing.T) {

}

func Test_AddProduct_IsError(t *testing.T) {
	_, errorMsg := entities.NewProduct(
		"",
		enums.Thing,
		"",
		"")

	if errorMsg == nil {
		t.Error(errorMsg)
	}
}
