package entities

import (
	"testing"
	"time"
	"waiter/domain/product/value_objects"

	"github.com/google/uuid"
)

var sale = Sale{
	uuid.New(),
	value_objects.NewPrice(400, "KGS"),
	uuid.New(),
	time.Now(),
	uuid.New(),
	2,
	uuid.New(),
}

func Test_SalePrice_ShouldBePositive(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("product price must not be less than zero")
		}
	}()
	value_objects.NewPrice(-1, "KGS")
}

func Test_SaleProduct_MustNotBeNull(t *testing.T) {
	if sale.productId == uuid.Nil {
		t.Error("Sale product id should not be null")
	}
}

func Test_SaleUser_MustNotBeNull(t *testing.T) {
	if sale.userId == uuid.Nil {
		t.Error("Sale user id should not be null")
	}
}

func Test_SaleCount_ShouldNotBeNegative(t *testing.T) {
	if sale.quantity <= 0 {
		t.Error("Sale quantity should be positive")
	}
}

func Test_SaleBranch_MustNotBeNull(t *testing.T) {
	if sale.branchId == uuid.Nil {
		t.Error("Sale branch id should not be null")
	}
}

func Test_SalePrice_ShouldBeAlwaysPositive(t *testing.T) {
	if !sale.price.IsPositiveAmount() {
		t.Error("sale.price should be positive")
	}
}
