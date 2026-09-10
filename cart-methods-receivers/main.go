package main

import "fmt"

type Item struct {
	ID    string
	Price float64
	Qty   int
}

type Cart struct {
	Items []Item
}

func (c *Cart) AddItem(item Item) {
	c.Items = append(c.Items, item)
}

func (c *Cart) RemoveItem(itemID string) {
	for i, item := range c.Items {
		if item.ID == itemID {
			c.Items = append(c.Items[:i], c.Items[i+1:]...)
			return
		}
	}
}

func (c Cart) TotalPrice() float64 {
	var total float64
	for _, item := range c.Items {
		total += item.Price * float64(item.Qty)
	}
	return total
}

func main() {
	cart := &Cart{}

	cart.AddItem(Item{ID: "apple", Price: 1.20, Qty: 5})
	cart.AddItem(Item{ID: "banana", Price: 0.80, Qty: 10})

	fmt.Printf("Total: $%.2f\n", cart.TotalPrice())

	cart.RemoveItem("apple")

	fmt.Printf("total after remove: $%.2f\n", cart.TotalPrice())
}
