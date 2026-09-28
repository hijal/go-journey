package main

import "fmt"

func main() {
	jobs := make(chan string, 3)

	jobs <- "resize-image"
	jobs <- "send-invoice"

	fmt.Println("len:", len(jobs), "cap:", cap(jobs))

	fmt.Println(<-jobs)
	fmt.Println(<-jobs)

	fmt.Println("len:", len(jobs))
}
