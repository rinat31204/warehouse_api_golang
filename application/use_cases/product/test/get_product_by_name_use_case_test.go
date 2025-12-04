package test

import (
	"errors"
	"strings"
	"testing"
	product2 "waiter/application/services/product"
	"waiter/domain/product/entities"
)

func TestGetProductByName_Success(t *testing.T) {
	mock := &ProductRepositoryMock{
		GetByNameFunc: func(name string) ([]*entities.Product, error) {
			if name == "" {
				return nil, errors.New("name is empty")
			}
			return []*entities.Product{{Name: name}, {Name: "cola 0.5"}, {Name: "CokaCola"}}, nil
		},
	}

	uc := product2.NewGetByName(mock)
	products, err := uc.Execute("cola")
	if err != nil {
		t.Errorf("GetProductByName error : %v", err)
	}
	if !containsProductWithName(products, "cola") {
		t.Error("GetProductByName error not same product")
	}
}

func containsProductWithName(products []*entities.Product, searchTerm string) bool {
	if products != nil || len(products) > 0 {
		for _, product := range products {
			if !strings.Contains(strings.ToLower(product.Name), strings.ToLower(searchTerm)) {
				return false
			}
		}
		return true
	}
	return false
}

func TestGetProductByName_Fail(t *testing.T) {

}

func TestGetProductById_NotFound(t *testing.T) {

}
