package product

import (
	"sync/atomic"
	"testing"
	"waiter/application/commands"
	"waiter/domain/product/entities"
	"waiter/domain/product/enums"

	"github.com/stretchr/testify/assert"
)

func TestAddProduct_Success(t *testing.T) {
	command := commands.AddProductCommand{
		Name:    "cola",
		Measure: 0,
		Code:    "123456789",
	}

	mock, addCalled := createMock(t, command)
	uc := NewAddProduct(mock)
	err := uc.Execute(command)

	assert.Nil(t, err)
	assert.Equal(t, 1, int(*addCalled))
}

func TestProductCode_OnlyNumber(t *testing.T) {
	command := commands.AddProductCommand{
		Name:    "cola",
		Measure: 0,
		Code:    "1234rhgsd567sdf89",
	}

	mock, addCalled := createMock(t, command)
	uc := NewAddProduct(mock)
	err := uc.Execute(command)

	assert.NotNil(t, err)
	assert.Equal(t, 0, int(*addCalled))
	assert.Contains(t, err.Error(), "invalid code")
}

func TestProduct_FailCases(t *testing.T) {
	tests := []struct {
		name    string
		nameVal string
		typeVal enums.MeasureType
		codeVal string
		descVal string
	}{
		{
			name:    "empty name",
			nameVal: "",
			typeVal: enums.Thing,
			codeVal: "54353543",
		},
		{
			name:    "empty code",
			nameVal: "hsdfadfadf",
			typeVal: enums.Thing,
			codeVal: "",
		},
		{
			name:    "invalid type",
			nameVal: "hsdfadfadf",
			typeVal: enums.MeasureType(4),
			codeVal: "5342342",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := entities.NewProduct(
				tt.nameVal,
				tt.typeVal,
				tt.codeVal,
				tt.descVal,
			)

			assert.Error(t, err)
		})
	}
}

func createMock(t *testing.T, command commands.AddProductCommand) (*ProductRepositoryMock, *int32) {
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
