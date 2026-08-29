package main

import "fmt"

func riskFlags(amount float64) []string {
	var flags []string

	switch {
	case amount >= 100000:
		flags = append(flags, "manual-review")
		fallthrough
	case amount >= 10000:
		flags = append(flags, "extra-verification")
		fallthrough
	case amount >= 1000:
		flags = append(flags, "log-audit-trail")
	default:
		// small amount, no flags needed
	}
	return flags
}

func main() {
	amount := 150000.0
	if amount <= 0 {
		fmt.Println("Invalid amount")
		return
	}

	fmt.Println("Amount:", amount)
	fmt.Println("Flags:", riskFlags(amount))
}
