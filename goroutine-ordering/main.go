package main

import (
	"fmt"
	"time"
)

func sendEmail(to string) {
	fmt.Println("email sent to", to)
}

func main() {
	go sendEmail("example@example.com")

	go func() {
		fmt.Println("i am anonymous function")
	}()

	fmt.Println("this is goroutine ordering")

	time.Sleep(100 * time.Millisecond)

	go fmt.Println("you will probably never see this")
	fmt.Println("main is done")
}
