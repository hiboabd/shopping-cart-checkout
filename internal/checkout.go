package internal

import "fmt"

type Checkout struct {
	Discounts []Discount
	Basket    Basket
}

func (c *Checkout) Scan(product Product) string {
	for _, item := range c.Basket {
		if item.Product.ProductCode == product.ProductCode {
			item.increaseQuantity(1)
			return fmt.Sprintf("%s scanned", product.Name)
		}
	}
	newItem := Item{
		Product:  product,
		Quantity: 1,
	}
	c.Basket = append(c.Basket, &newItem)
	return fmt.Sprintf("%s scanned", product.Name)
}

func (c *Checkout) calculateTotal() int {
	var total int
	for _, item := range c.Basket {
		total += item.Product.Price * item.Quantity
	}
	return total
}

func (c *Checkout) calculateDiscount() int {
	var discountAmount int
	for _, discount := range c.Discounts {
		discountAmount += discount.getDiscount(c.Basket)
	}
	return discountAmount
}

func (c *Checkout) CalculateTotalWithDiscount() string {
	total := c.calculateTotal()
	discount := c.calculateDiscount()
	totalWithDiscount := total - discount
	return fmt.Sprintf("Total with discount applied: %.2f", penniesToPounds(totalWithDiscount))
}
