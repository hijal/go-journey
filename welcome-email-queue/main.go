package main

import "fmt"

func main() {
	emails := make(chan string, 10)
	done := make(chan struct{})

	go func() {
		defer close(done)
		for add := range emails {
			fmt.Println("sending welcome email to", add)
		}
		fmt.Println("queue drained, worker exiting")
	}()

	signups := []string{"rina@shop.test", "tanvir@shop.test", "nadia@shop.test"}

	for _, add := range signups {
		emails <- add
	}
	close(emails)

	<-done
	fmt.Println("enqueued", len(signups), "emails; shutting down")
}
