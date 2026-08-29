package main

import "fmt"

func main() {
	invoiceAmounts := []float64{1200, 0, 850, -50, 3000, 400}
	var total float64

	for i, amount := range invoiceAmounts {
		if amount == 0 {
			continue
		}
		if amount < 0 {
			fmt.Println("invalid invoice at index", i, "- stopping batch")
			break
		}
		total += amount
		fmt.Println("processed invoice", i, ":", amount)
	}

	fmt.Println("Total processed:", total)
}
