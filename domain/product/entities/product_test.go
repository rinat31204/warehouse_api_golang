package entities

import (
	"testing"
	"waiter/domain/product/value_objects"

	"github.com/google/uuid"
)

var product = Product{
	uuid.New(),
	"glass",
	value_objects.NewPrice(175.50, "KGS"),
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

func Test_ProductPriceCurrency_ShouldBeSimilar(t *testing.T) {
	if !product.price.EqualCurrency("KGS") {
		t.Errorf("Product price must be similar")
	}
}
