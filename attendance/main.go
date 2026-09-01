package main

import "fmt"

func main() {
	const regularMinutes = 8 * 60 // full work day

	workedMinutes := 537

	regular := min(workedMinutes, regularMinutes)
	overtime := max(0, workedMinutes-regularMinutes)

	hours, minutes := regular/60, regular%60
	otHours, otMinutes := overtime/60, overtime%60

	fmt.Printf("regular: %dh %dm\n", hours, minutes)
	fmt.Printf("overtime: %dh %dm\n", otHours, otMinutes)
	fmt.Println("overtime pay multiplier applies:", overtime > 0)
}
