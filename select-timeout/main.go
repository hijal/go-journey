package main

import (
	"fmt"
	"time"
)

func main() {
	primary := make(chan string)
	secondary := make(chan string)

	go func() {
		time.Sleep(time.Millisecond * 50)
		primary <- "primary-db result"
	}()

	go func() {
		time.Sleep(time.Millisecond * 10)
		secondary <- "secondary-db result"
	}()

	select {
	case r := <-primary:
		fmt.Println(r)
	case r := <-secondary:
		fmt.Println(r)
	case <-time.After(time.Millisecond * 100):
		fmt.Println("timeout")
	}
}
