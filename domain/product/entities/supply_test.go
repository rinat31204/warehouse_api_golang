package entities

import (
	"testing"
	"time"
	"waiter/domain/product/value_objects"

	"github.com/google/uuid"
)

var supply = Supply{
	uuid.New(),
	35,
	uuid.New(),
	value_objects.NewPrice(85, "KGS"),
	value_objects.NewPrice(90, "KGS"),
	time.Now(),
	uuid.New(),
	uuid.New(),
}

func Test_SupplyCount_ShouldBePositive(t *testing.T) {
	if supply.quantity <= 0 {
		t.Error("supply.quantity should be positive")
	}
}

func Test_SupplyProduct_MustNotBeNull(t *testing.T) {
	if supply.productId == uuid.Nil {
		t.Error("supply.productId should not be null")
	}
}
func Test_SupplyIncomePrice_MustNotBeNegative(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("product price must not be less than zero")
		}
	}()
	value_objects.NewPrice(-1, "KGS")
}

func Test_SupplyIncomePrice_ShouldBeAlwaysPositive(t *testing.T) {
	if !supply.incomePrice.IsPositiveAmount() {
		t.Error("supply.price should be positive")
	}
}

func Test_SupplySalePrice_ShouldBeAlwaysPositive(t *testing.T) {
	if !supply.salePrice.IsPositiveAmount() {
		t.Error("supply.price should be positive")
	}
}
