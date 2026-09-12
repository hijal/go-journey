package main

import "fmt"

type Product struct {
	SKU   string
	Name  string
	Price float64
}

func applyDiscountBad(p Product) {
	p.Price *= 0.9
}

func applyDiscount(p *Product, pct float64) {
	p.Price *= 1 - pct
}

func main() {
	laptop := Product{
		SKU:   "LAP123",
		Name:  "Laptop Pro 14",
		Price: 90000.0,
	}

	applyDiscountBad(laptop)
	fmt.Println("After bad discount:", laptop)

	applyDiscount(&laptop, 0.1)
	fmt.Println("After good discount:", laptop)

	backup := laptop
	backup.Price = 0
	fmt.Println("Backup:", backup)
	fmt.Println("Laptop untouched:", laptop)
}
