package entities

import (
	"testing"
	"time"
	"waiter/domain/product/enums"
	"waiter/domain/product/value_objects"

	"github.com/google/uuid"
)

var product = Product{
	uuid.New(),
	"glass",
	time.Now(),
	enums.Thing,
	"122500053883513",
	"",
}

func Test_ProductName_MustNotBeNullOrEmpty(t *testing.T) {
	if product.name == "" {
		t.Errorf("Product name must not be empty")
	}
}

func Test_ProductPrice_MustNotBeLessThanZero(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("product price must not be less than zero")
		}
	}()
	value_objects.NewPrice(-1, "KGS")
}

func Test_ProductCode_MustNotBeNullOrEmpty(t *testing.T) {
	if product.code == "" {
		t.Errorf("Product code must not be empty")
	}
}

func Test_ProductDescription_LimitOffset(t *testing.T) {
	if len(product.description) > 500 {
		t.Errorf("Product description length must be less than 500")
	}
}
