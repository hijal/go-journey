package main

import "fmt"

func newRegister(name string) func(float64) float64 {
	total := 0.0
	return func(amount float64) float64 {
		total += amount
		return total
	}
}
func main() {
	counterA := newRegister("Counter A")
	counterB := newRegister("Counter B")

	counterA(500)
	counterA(300)
	counterB(1000)

	fmt.Printf("counter-A total: %.2f\n", counterA(200))
	fmt.Printf("counter-B total: %.2f\n", counterB(450))
}
