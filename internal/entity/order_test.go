package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfItGetAnError_IfIdIsBlank(t *testing.T) {
	order := Order {
		ID: "",
		Price: 12.12,
		Tax: 1.00,
	}

	assert.Error(t, order.Validate(),  "id is required")
}

func TestIfItGetAnError_IfPriceIsInvalid(t *testing.T) {
	order := Order {
		ID: "1",
		Price: 0,
		Tax: 1.00,
	}

	assert.Error(t, order.Validate(),  "price must be greater than zero")
}

func TestIfItGetAnError_IfTaxIsInvalid(t *testing.T) {
	order := Order {
		ID: "1",
		Price: 12.23,
		Tax: 0,
	}

	assert.Error(t, order.Validate(),  "invalid tax")
}

func TestIfItOk_IfAllParamsSuccess(t *testing.T) {
	order := Order {
		ID: "1",
		Price: 12.23,
		Tax: 2.00,
	}

	order.CalculateFinalPrice()

	assert.NoError(t, order.Validate())
	assert.Equal(t, 14.23, order.FinalPrice)
}
