package main

import (
	"fmt"
	"shopping-cart-checkout/internal"
)

func main() {
	fmt.Println("Start here")
	exampleCheckout := setUpExampleData()
	coffee := internal.Product{
		ProductCode: internal.Coffee,
		Name:        "Supermarket Coffee",
		Price:       1123,
	}
	fruitTea := internal.Product{
		ProductCode: internal.FruitTea,
		Name:        "Supermarket Fruit Tea",
		Price:       311,
	}

	strawberries := internal.Product{
		ProductCode: internal.Strawberries,
		Name:        "Supermarket Strawberries",
		Price:       500,
	}

	// Scenario 1: Scanning product - coffee
	fmt.Println("Scenario 1: Scanning product - coffee")
	message := exampleCheckout.Scan(coffee)
	fmt.Printf(message + "\n")
	fmt.Println(exampleCheckout.Basket[0])

	// Scenario 2: Calculating total with discounts
	fmt.Println("Scenario 2: Calculating total with discounts")
	message = exampleCheckout.Scan(fruitTea)
	fmt.Printf(message + "\n")
	message = exampleCheckout.Scan(fruitTea)
	fmt.Printf(message + "\n")
	message = exampleCheckout.Scan(strawberries)
	fmt.Printf(message + "\n")
	// To demo that the order of scanning does not matter
	//message = exampleCheckout.Scan(coffee)
	//fmt.Printf(message + "\n")
	for _, item := range exampleCheckout.Basket {
		fmt.Println(item)
	}
	message = exampleCheckout.CalculateTotalWithDiscount()
	fmt.Printf(message + "\n")
	// Total: 11.23 + 3.11 + 3.11 + 5.00 = 22.45
	// Discount: 3.11
	// Final total: 19.34

	// Scenario 3: Bulk discount example
	fmt.Println("Scenario 3: Bulk discount example")
	message = exampleCheckout.Scan(strawberries)
	fmt.Printf(message + "\n")

	message = exampleCheckout.Scan(strawberries)
	fmt.Printf(message + "\n")

	for _, item := range exampleCheckout.Basket {
		fmt.Println(item)
	}
	message = exampleCheckout.CalculateTotalWithDiscount()
	fmt.Printf(message + "\n")
	// Total: 11.23 + 3.11 + 3.11 + 5.00 + 5.00 + 5.00 = 32.45
	// Discount: 3.11 + 1.50 = 4.61
	// Final total: 27.84
}

func setUpExampleData() internal.Checkout {
	// Create products
	fruitTea := internal.Product{
		ProductCode: internal.FruitTea,
		Name:        "Supermarket Fruit Tea",
		Price:       311,
	}

	strawberries := internal.Product{
		ProductCode: internal.Strawberries,
		Name:        "Supermarket Strawberries",
		Price:       500,
	}

	// Create discounts
	bogofDiscount := internal.Discount{
		ID:       1,
		Type:     internal.BOGOF,
		Products: []internal.Product{fruitTea},
	}

	bulkDiscount := internal.Discount{
		ID:         1,
		Type:       internal.BulkDiscount,
		Products:   []internal.Product{strawberries},
		Percentage: 0.1,
		Quantity:   3,
	}

	// Create checkout
	checkout := internal.Checkout{
		Discounts: []internal.Discount{bogofDiscount, bulkDiscount},
		Basket:    internal.Basket{},
	}
	return checkout
}
