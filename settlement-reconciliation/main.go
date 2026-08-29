package main

import (
	"errors"
	"fmt"
)

func connectToBank(attempt int) error {
	if attempt < 3 {
		return errors.New("failed to connect to bank")
	}
	return nil
}

func main() {
	var err error
	attempt := 1
retry:
	err = connectToBank(attempt)

	if err != nil && attempt < 5 {
		fmt.Println("Attempt", attempt, "failed:", err)
		attempt++
		goto retry
	}

	if err != nil {
		fmt.Println("giving up after", attempt, "attempts")
		return
	}

	fmt.Println("Connected on attempt", attempt)

	batches := [][]float64{
		{500, 300, 200},
		{700, -1, 100},
		{900, 400},
	}
scan:
	for batchIndex, batch := range batches {
		for _, entry := range batch {
			if entry < 0 {
				fmt.Println("mismatched entry in batch", batchIndex, ":", entry)
				break scan
			}
			fmt.Println("verified entry:", entry)
		}
	}
}
