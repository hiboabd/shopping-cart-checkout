package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddProduct(t *testing.T) {
	fruitTea := Product{
		ProductCode: FruitTea,
		Name:        "Supermarket Fruit Tea",
		Price:       311,
	}

	discount := Discount{
		ID:         1,
		Type:       BOGOF,
		Percentage: 100,
		Products:   []Product{fruitTea},
	}
	assert.Equal(t, 1, len(discount.Products))

	strawberries := Product{
		ProductCode: Strawberries,
		Name:        "Supermarket Strawberries",
		Price:       500,
	}

	message := discount.addProduct(strawberries)
	assert.Equal(t, 2, len(discount.Products))
	assert.Equal(t, "Supermarket Strawberries added", message)
}
