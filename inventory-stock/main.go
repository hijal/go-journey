package main

import "fmt"

func main() {
	var stock uint32 = 5

	sold := 8
	newStock := stock - uint32(sold)
	fmt.Println("Stock after (buggy) sale:", newStock)

	safeStock := int(stock) - sold

	if safeStock < 0 {
		fmt.Println("Rejected: cannot sell more than available stock")
	} else {
		fmt.Println("New stock:", safeStock)
	}
}
