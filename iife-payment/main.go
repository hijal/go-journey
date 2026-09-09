package main

import "fmt"

func main() {
	amount := 1250.75
	currency := "BDT"

	summary := func(amt float64, cur string) string {
		if amt <= 0 {
			return "REJECT: invalid amount"
		}
		return fmt.Sprintf("%s %.2f accepted", cur, amt)
	}(amount, currency)

	fmt.Println(summary)
}
