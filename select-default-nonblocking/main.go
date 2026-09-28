package main

import "fmt"

func main() {
	alerts := make(chan string, 1)

	for _, msg := range []string{"cpu high", "disk full"} {
		select {
		case alerts <- msg:
			fmt.Println("queued:", msg)
		default:
			fmt.Println("dropped (queue full):", msg)
		}
	}
}
