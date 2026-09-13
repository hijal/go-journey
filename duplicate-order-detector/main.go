package main

import (
	"fmt"
	"slices"
)

func main() {
	orders := []string{
		"ORD-101", "ORD-102", "ORD-101", "ORD-103", "ORD-102", "ORD-101",
	}

	seen := make(map[string]bool)
	duplicates := []string{}

	for _, id := range orders {
		if seen[id] && !slices.Contains(duplicates, id) {
			duplicates = append(duplicates, id)
		}
		seen[id] = true
	}

	slices.Sort(duplicates)
	fmt.Println("duplicate order ids:", duplicates)
}
