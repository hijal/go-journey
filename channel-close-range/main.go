package main

import "fmt"

func main() {
	events := make(chan string, 2)

	events <- "user.signup"
	events <- "user.login"

	close(events)

	for e := range events {
		fmt.Println("event:", e)
	}

	v, ok := <-events
	fmt.Printf("v=%q ok=%v\n", v, ok)
}
