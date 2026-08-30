package main

import "fmt"

type PaymentAmount int64

func main() {
	// 1 BDT = 100 paisa
	// 149950 / 100 = 1499.50 BDT
	orderTotal := PaymentAmount(149950)

	var refundAmount PaymentAmount
	fmt.Println("Refund before processing", refundAmount)

	refundAmount = 20000 // 200.00 BDT
	remaining := orderTotal - refundAmount

	fmt.Printf("Order total : %d paisa (%.2f BDT)\n", orderTotal, float64(orderTotal)/100)
	fmt.Printf("Refunded    : %d paisa (%.2f BDT)\n", refundAmount, float64(refundAmount)/100)
	fmt.Printf("Remaining   : %d paisa (%.2f BDT)\n", remaining, float64(remaining)/100)
}
