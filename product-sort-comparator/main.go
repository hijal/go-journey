package main

import (
	"cmp"
	"fmt"
	"slices"
)

type Product struct {
	name  string
	price float64
	sold  int
}

func main() {
	catalog := []Product{
		{name: "keyboard", price: 1500, sold: 320},
		{name: "mouse", price: 750, sold: 950},
		{name: "monitor", price: 12500, sold: 120},
		{name: "usb hub", price: 750, sold: 410},
	}

	slices.SortFunc(catalog, func(a, b Product) int {
		if c := cmp.Compare(a.price, b.price); c != 0 {
			return c
		}
		return cmp.Compare(a.name, b.name)
	})

	for _, p := range catalog {
		fmt.Printf("%-10s %8.2f\n", p.name, p.price)
	}

	best := slices.MaxFunc(catalog, func(a, b Product) int {
		return cmp.Compare(a.sold, b.sold)
	})

	fmt.Println("best seller:", best.name)
}
