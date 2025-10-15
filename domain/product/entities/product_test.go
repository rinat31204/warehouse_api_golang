package entities

import (
	"testing"

	"github.com/google/uuid"
)

var product = Product{
	uuid.New(),
	"glass",
}

func Test_ProductName_MustNotBeNullOrEmpty(test *testing.T) {
	if product.name == "" {
		test.Errorf("Product name must not be empty")
	}
}
