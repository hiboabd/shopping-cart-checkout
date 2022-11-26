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

func TestIsItemOnDiscount(t *testing.T) {
	fruitTea := Product{
		ProductCode: FruitTea,
		Name:        "Supermarket Fruit Tea",
		Price:       311,
	}

	discount := Discount{
		ID:       1,
		Products: []Product{fruitTea},
	}

	item := Item{
		Product:  fruitTea,
		Quantity: 2,
	}

	result := isItemOnDiscount(&item, discount.Products)
	assert.Equal(t, true, result)

	strawberries := Product{
		ProductCode: Strawberries,
		Name:        "Supermarket Strawberries",
		Price:       500,
	}

	strawberryItem := Item{
		Product:  strawberries,
		Quantity: 2,
	}

	result2 := isItemOnDiscount(&strawberryItem, discount.Products)
	assert.Equal(t, false, result2)
}

func TestBuyOneGetOneFreeWhenItemOnDiscount(t *testing.T) {
	fruitTea := Product{
		ProductCode: FruitTea,
		Name:        "Supermarket Fruit Tea",
		Price:       311,
	}

	discount := Discount{
		ID:       1,
		Products: []Product{fruitTea},
	}

	item := Item{
		Product:  fruitTea,
		Quantity: 2,
	}

	basket := Basket{&item}

	result := buyOneGetOneFree(basket, discount.Products)
	assert.Equal(t, 311, result)
}

func TestBuyOneGetOneFreeWhenItemIsNotOnDiscount(t *testing.T) {
	fruitTea := Product{
		ProductCode: FruitTea,
		Name:        "Supermarket Fruit Tea",
		Price:       311,
	}

	discount := Discount{
		ID:       1,
		Products: []Product{},
	}

	item := Item{
		Product:  fruitTea,
		Quantity: 2,
	}

	basket := Basket{&item}

	result := buyOneGetOneFree(basket, discount.Products)
	assert.Equal(t, 0, result)
}
