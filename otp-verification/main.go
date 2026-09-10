package main

import "fmt"

func main() {
	serverOTP := [6]int{7, 3, 0, 9, 1, 4}
	userOTP := [6]int{7, 3, 0, 9, 1, 5}
	fmt.Println("OTP length:", len(serverOTP))

	if serverOTP == userOTP {
		fmt.Println("OTP verified: payment approved")
	} else {
		fmt.Println("OTP mismatch: payment declined")
	}
}
