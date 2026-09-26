package main

import (
	"fmt"
	"sync"
)

var loadConfig = sync.OnceValue(func() map[string]string {
	fmt.Println("reading config from disk (happens once)")
	return map[string]string{"currency": "BDT"}
})

func main() {
	var wg sync.WaitGroup
	for range 3 {
		wg.Go(func() {
			fmt.Println("currency:", loadConfig()["currency"])
		})
	}
	wg.Wait()
}
