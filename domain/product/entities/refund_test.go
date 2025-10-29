package entities

import (
	"testing"
	"time"
	"waiter/domain/product/value_objects"

	"github.com/google/uuid"
)

var refund = Refund{
	uuid.New(),
	time.Now(),
	uuid.New(),
	uuid.New(),
	2,
	uuid.New(),
}

func Test_RefundAmount_ShouldBePositive(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("product price must not be less than zero")
		}
	}()
	value_objects.NewPrice(-1, "KGS")
}

func Test_RefundProductId_MustNotBeNullOrEmpty(t *testing.T) {
	if refund.id == uuid.Nil {
		t.Errorf("Product id must not be empty")
	}
}

func Test_RefundProductCount_MustNotBeNull(t *testing.T) {
	if refund.quantity <= 0 {
		t.Errorf("Product count must not be negative")
	}
}
