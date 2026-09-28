package main

import "fmt"

func produce(out chan<- int) {
	for i := range 10 {
		out <- i * 100
	}
	close(out)
}

func consume(in <-chan int) {
	for n := range in {
		fmt.Println("number:", n)
	}
}

func main() {
	ch := make(chan int)

	go produce(ch)
	consume(ch)
}
