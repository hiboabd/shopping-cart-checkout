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
		Percentage: 1.0,
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
		Type:     BOGOF,
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
		Type:     BOGOF,
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

func TestBulkDiscountWhenDoNotHaveRequiredQuantity(t *testing.T) {
	strawberries := Product{
		ProductCode: Strawberries,
		Name:        "Supermarket Strawberries",
		Price:       500,
	}

	discount := Discount{
		ID:         1,
		Type:       BulkDiscount,
		Products:   []Product{strawberries},
		Percentage: 0.1,
		Quantity:   3,
	}

	item := Item{
		Product:  strawberries,
		Quantity: 2,
	}

	basket := Basket{&item}

	result := bulkDiscount(basket, &discount)
	assert.Equal(t, 0, result)
}

func TestBulkDiscountWhenDoHaveRequiredQuantity(t *testing.T) {
	strawberries := Product{
		ProductCode: Strawberries,
		Name:        "Supermarket Strawberries",
		Price:       500,
	}

	discount := Discount{
		ID:         1,
		Type:       BulkDiscount,
		Products:   []Product{strawberries},
		Percentage: 0.1,
		Quantity:   3,
	}

	item := Item{
		Product:  strawberries,
		Quantity: 4,
	}

	basket := Basket{&item}

	result := bulkDiscount(basket, &discount)
	assert.Equal(t, 200, result)
}

func TestBulkDiscountWhenProductNotOnDiscount(t *testing.T) {
	strawberries := Product{
		ProductCode: Strawberries,
		Name:        "Supermarket Strawberries",
		Price:       500,
	}

	discount := Discount{
		ID:         1,
		Type:       BulkDiscount,
		Products:   []Product{},
		Percentage: 0.1,
		Quantity:   3,
	}

	item := Item{
		Product:  strawberries,
		Quantity: 4,
	}

	basket := Basket{&item}

	result := bulkDiscount(basket, &discount)
	assert.Equal(t, 0, result)
}

func TestGetDiscountBulkDiscount(t *testing.T) {
	strawberries := Product{
		ProductCode: Strawberries,
		Name:        "Supermarket Strawberries",
		Price:       500,
	}

	discount := Discount{
		ID:         1,
		Type:       BulkDiscount,
		Products:   []Product{strawberries},
		Percentage: 0.1,
		Quantity:   3,
	}

	item := Item{
		Product:  strawberries,
		Quantity: 4,
	}

	basket := Basket{&item}

	result := discount.getDiscount(basket)
	assert.Equal(t, 200, result)
}

func TestGetDiscountBuyOneGetOneFree(t *testing.T) {
	fruitTea := Product{
		ProductCode: FruitTea,
		Name:        "Supermarket Fruit Tea",
		Price:       311,
	}

	discount := Discount{
		ID:       1,
		Type:     BOGOF,
		Products: []Product{fruitTea},
	}

	item := Item{
		Product:  fruitTea,
		Quantity: 2,
	}

	basket := Basket{&item}

	result := discount.getDiscount(basket)
	assert.Equal(t, 311, result)
}

func TestGetDiscountDiscountWithNoType(t *testing.T) {
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

	result := discount.getDiscount(basket)
	assert.Equal(t, 0, result)
}
