package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScanEmptyBasket(t *testing.T) {
	fruitTea := Product{
		ProductCode: FruitTea,
		Name:        "Supermarket Fruit Tea",
		Price:       311,
	}

	checkout := Checkout{
		Discounts: []Discount{},
		Basket:    Basket{},
	}

	assert.Equal(t, 0, len(checkout.Basket))

	message := checkout.scan(fruitTea)
	assert.Equal(t, "Supermarket Fruit Tea scanned", message)
	assert.Equal(t, 1, len(checkout.Basket))
	assert.Equal(t, 1, checkout.Basket[0].Quantity)
	assert.Equal(t, FruitTea, checkout.Basket[0].Product.ProductCode)
	assert.Equal(t, "Supermarket Fruit Tea", checkout.Basket[0].Product.Name)
	assert.Equal(t, 311, checkout.Basket[0].Product.Price)
}

func TestScanBasketWithExistingItem(t *testing.T) {
	fruitTea := Product{
		ProductCode: FruitTea,
		Name:        "Supermarket Fruit Tea",
		Price:       311,
	}

	item := Item{
		Product:  fruitTea,
		Quantity: 2,
	}

	basket := Basket{&item}

	checkout := Checkout{
		Discounts: []Discount{},
		Basket:    basket,
	}

	assert.Equal(t, 1, len(checkout.Basket))

	strawberries := Product{
		ProductCode: Strawberries,
		Name:        "Supermarket Strawberries",
		Price:       500,
	}

	firstProductMessage := checkout.scan(fruitTea)

	assert.Equal(t, "Supermarket Fruit Tea scanned", firstProductMessage)
	assert.Equal(t, 1, len(checkout.Basket))
	assert.Equal(t, 3, checkout.Basket[0].Quantity)
	assert.Equal(t, FruitTea, checkout.Basket[0].Product.ProductCode)
	assert.Equal(t, "Supermarket Fruit Tea", checkout.Basket[0].Product.Name)
	assert.Equal(t, 311, checkout.Basket[0].Product.Price)

	secondProductMessage := checkout.scan(strawberries)
	assert.Equal(t, "Supermarket Strawberries scanned", secondProductMessage)
	assert.Equal(t, 2, len(checkout.Basket))
	assert.Equal(t, 1, checkout.Basket[1].Quantity)
	assert.Equal(t, Strawberries, checkout.Basket[1].Product.ProductCode)
	assert.Equal(t, "Supermarket Strawberries", checkout.Basket[1].Product.Name)
	assert.Equal(t, 500, checkout.Basket[1].Product.Price)

}
