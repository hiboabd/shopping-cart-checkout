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
