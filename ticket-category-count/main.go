package main

import "fmt"

func main() {
	categories := []string{
		"billing", "shipping", "billing", "refund",
		"shipping", "billing", "login", "refund", "refund",
	}

	counts := make(map[string]int)

	for _, c := range categories {
		counts[c]++
	}

	fmt.Println("Ticket counts by category")

	for cat, n := range counts {
		fmt.Printf("  %-10s %d\n", cat, n)
	}
}
