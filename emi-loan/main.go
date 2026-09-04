package main

import "fmt"

func main() {
	principal := 120_000.0
	flatAnnualRate := 0.10
	years := 1
	months := 12

	totalInterest := principal * flatAnnualRate * float64(years)
	emi := (principal + totalInterest) / float64(months)
	interestPart := totalInterest / float64(months)
	principalPart := emi - interestPart

	fmt.Printf("EMI: %.2f BDT/month\n\n", emi)
	fmt.Printf("%-5s %10s %12s %12s\n", "month", "principal", "interest", "remaining")

	remaining := principal

	for m := 1; m <= months; m++ {
		remaining -= principalPart
		fmt.Printf("%-5d %10.2f %12.2f %12.2f\n", m, principalPart, interestPart, remaining)
	}
}
