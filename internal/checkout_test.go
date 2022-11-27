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

func TestCalculateTotal(t *testing.T) {
	fruitTea := Product{
		ProductCode: FruitTea,
		Name:        "Supermarket Fruit Tea",
		Price:       311,
	}

	strawberries := Product{
		ProductCode: Strawberries,
		Name:        "Supermarket Strawberries",
		Price:       500,
	}

	item := Item{
		Product:  fruitTea,
		Quantity: 2,
	}

	itemStrawberries := Item{
		strawberries,
		1,
	}

	basket := Basket{&item, &itemStrawberries}

	checkout := Checkout{
		Discounts: []Discount{},
		Basket:    basket,
	}

	total := checkout.calculateTotal()

	assert.Equal(t, 1122, total)
}

func TestCalculateDiscount(t *testing.T) {
	fruitTea := Product{
		ProductCode: FruitTea,
		Name:        "Supermarket Fruit Tea",
		Price:       311,
	}

	strawberries := Product{
		ProductCode: Strawberries,
		Name:        "Supermarket Strawberries",
		Price:       500,
	}

	item := Item{
		Product:  fruitTea,
		Quantity: 2,
	}

	itemStrawberries := Item{
		strawberries,
		4,
	}

	basket := Basket{&item, &itemStrawberries}

	bogofDiscount := Discount{
		ID:       1,
		Type:     BOGOF,
		Products: []Product{fruitTea},
	}

	bulkDiscount := Discount{
		ID:         1,
		Type:       BulkDiscount,
		Products:   []Product{strawberries},
		Percentage: 0.1,
		Quantity:   3,
	}

	checkout := Checkout{
		Discounts: []Discount{bogofDiscount, bulkDiscount},
		Basket:    basket,
	}

	total := checkout.calculateDiscount()

	assert.Equal(t, 511, total)
}

func TestCalculateTotalWithDiscount(t *testing.T) {
	fruitTea := Product{
		ProductCode: FruitTea,
		Name:        "Supermarket Fruit Tea",
		Price:       311,
	}

	strawberries := Product{
		ProductCode: Strawberries,
		Name:        "Supermarket Strawberries",
		Price:       500,
	}

	item := Item{
		Product:  fruitTea,
		Quantity: 2,
	}

	itemStrawberries := Item{
		strawberries,
		4,
	}

	basket := Basket{&item, &itemStrawberries}

	bogofDiscount := Discount{
		ID:       1,
		Type:     BOGOF,
		Products: []Product{fruitTea},
	}

	bulkDiscount := Discount{
		ID:         1,
		Type:       BulkDiscount,
		Products:   []Product{strawberries},
		Percentage: 0.1,
		Quantity:   3,
	}

	checkout := Checkout{
		Discounts: []Discount{bogofDiscount, bulkDiscount},
		Basket:    basket,
	}

	result := checkout.calculateTotalWithDiscount()

	assert.Equal(t, "Total with discount applied: 21.11", result)
}
