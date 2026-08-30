package main

import (
	"fmt"
	"sync"
)

func sendWebhook(url string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	results <- fmt.Sprintf("delivered to %s", url)
}

func main() {
	endpoints := []string{
		"https://merchant-a.example.com/webhook",
		"https://merchant-b.example.com/webhook",
		"https://merchant-c.example.com/webhook",
	}

	results := make(chan string, len(endpoints))
	var wg sync.WaitGroup

	for _, endpoint := range endpoints {
		wg.Add(1)
		go sendWebhook(endpoint, results, &wg)
	}

	wg.Wait()
	close(results)

	for message := range results {
		fmt.Println(message)
	}
}
