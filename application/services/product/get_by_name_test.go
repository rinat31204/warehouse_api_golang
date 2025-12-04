package product

import (
	"errors"
	"strings"
	"testing"
	"waiter/domain/product/entities"

	"github.com/stretchr/testify/assert"
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

	uc := NewGetByName(mock)
	products, err := uc.Execute("cola")

	assert.Nil(t, err)
	assert.Contains(t, products, "cola 0.5")

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
