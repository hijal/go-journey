package main

import (
	"fmt"
	"slices"
)

func main() {
	revenue := map[string]int{
		"electronics": 850_000,
		"fashion":     420_000,
		"books":       95_000,
		"grocery":     610_000,
	}
	categories := make([]string, 0, len(revenue))

	for cat := range revenue {
		categories = append(categories, cat)
	}
	slices.Sort(categories)

	total := 0
	for _, cat := range categories {
		fmt.Printf("%-12s %10d BDT\n", cat, revenue[cat])
		total += revenue[cat]
	}
	fmt.Printf("%-12s %10d BDT\n", "TOTAL", total)
}
