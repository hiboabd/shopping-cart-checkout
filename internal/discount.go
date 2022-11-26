package internal

import (
	"fmt"
)

const (
	BOGOF        string = "Buy One Get One Free"
	BulkDiscount string = "Bulk Discount"
)

type Discount struct {
	ID         int
	Type       string
	Percentage int
	Products   []Product
	Quantity   int
}

func (d *Discount) addProduct(product Product) string {
	d.Products = append(d.Products, product)
	return fmt.Sprintf("%s added", product.Name)
}

func buyOneGetOneFree(basket Basket, discountProducts []Product) int {
	var discountAmount int
	for _, item := range basket {
		if isItemOnDiscount(item, discountProducts) {
			adjustedQuantity := item.Quantity
			if !isDivisibleBy(item.Quantity, 2) {
				adjustedQuantity = item.Quantity - 1
			}
			discountAmount += adjustedQuantity * item.Product.Price / 2
		}
	}

	return discountAmount
}

func isItemOnDiscount(item *Item, discountProducts []Product) bool {
	for _, discountProduct := range discountProducts {
		if item.Product.ProductCode == discountProduct.ProductCode {
			return true
		}
	}
	return false
}
