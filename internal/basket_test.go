package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIncreaseQuantity(t *testing.T) {
	fruitTea := Product{
		ProductCode: FruitTea,
		Name:        "Supermarket Fruit Tea",
		Price:       311,
	}

	item := Item{
		Product:  fruitTea,
		Quantity: 2,
	}
	assert.Equal(t, 2, item.Quantity)

	item.increaseQuantity(1)
	assert.Equal(t, 3, item.Quantity)
}
