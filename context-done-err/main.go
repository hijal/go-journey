package main

import (
	"context"
	"fmt"
	"time"
)

func fetchRate(ctx context.Context) (float64, error) {
	result := make(chan float64, 1)
	go func() {
		time.Sleep(time.Millisecond * 300)
		result <- 119.55
	}()

	select {
	case r := <-result:
		return r, nil

	case <-ctx.Done():
		return 0, fmt.Errorf("fetch rate: %w", ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)

	defer cancel()

	rate, err := fetchRate(ctx)

	if err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Println("rate:", rate)
}
