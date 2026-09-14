package main

import "fmt"

func main() {
	prices := map[string]int{
		"SKU-1001": 1999,
		"SKU-1002": 499,
		"SKU-1003": 12999,
	}
	fmt.Printf("total products: %d\n", len(prices))

	sku := "SKU-1002"

	if price, ok := prices[sku]; ok {
		fmt.Printf("%s costs %d cents (%.2f USD)\n", sku, price, float64(price)/100)
	} else {
		fmt.Printf("%s not found in catalog\n", sku)
	}

	sku = "SKU-9999"
	if _, ok := prices[sku]; !ok {
		fmt.Printf("%s not found, suggest alternatives\n", sku)
	}
}
