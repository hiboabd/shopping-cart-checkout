package internal

type Basket []*Item

type Item struct {
	Product  Product
	Quantity int
}

func (i *Item) increaseQuantity(amount int) {
	i.Quantity = i.Quantity + amount
}
