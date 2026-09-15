package main

import "fmt"

type Product struct {
	SKU   string
	Name  string
	Stock int
}

type Inventory struct {
	items map[string]*Product
}

func NewInventory() *Inventory {
	return &Inventory{items: make(map[string]*Product)}
}

func (i *Inventory) Add(p *Product) {
	i.items[p.SKU] = p
}

func (i *Inventory) ReduceStock(sku string, quantity int) error {
	product, ok := i.items[sku]

	if !ok {
		return fmt.Errorf("product %s not found", sku)
	}

	if product.Stock < quantity {
		return fmt.Errorf("insufficient stock for %s: have %d, want %d",
			sku, product.Stock, quantity)
	}

	product.Stock -= quantity

	return nil
}

func main() {
	inv := NewInventory()

	inv.Add(&Product{SKU: "SKU-1", Name: "Mechanical Keyboard", Stock: 15})

	if err := inv.ReduceStock("SKU-1", 3); err != nil {
		fmt.Println("order error:", err)
	}
	fmt.Printf("Remaining stock: %d\n", inv.items["SKU-1"].Stock)
}
