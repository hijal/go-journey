package main

import (
	"fmt"
	"strings"
)

func matches(query, product string) bool {
	product = strings.ToLower(product)
	for _, word := range strings.Fields(strings.ToLower(query)) {
		if !strings.Contains(product, word) {
			return false
		}
	}
	return true
}

func main() {
	products := []string{
		"Wireless Mouse Pro",
		"USB-C Charging Cable",
		"Mechanical Keyboard",
	}

	query := "  wireless   MOUSE "

	for _, p := range products {
		if matches(query, p) {
			fmt.Println("HIT:", p)
		}
	}
}
