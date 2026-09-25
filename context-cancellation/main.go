package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func generateReport(ctx context.Context) error {
	for step := range 5 {
		select {
		case <-ctx.Done():
			return fmt.Errorf("report stopped at step %d: %w", step, ctx.Err())
		case <-time.After(100 * time.Millisecond):
			fmt.Println("finished step", step)
		}
	}
	return nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 250*time.Millisecond)
	defer cancel()

	if err := generateReport(ctx); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Println("timeout:", err)
			return
		}
		fmt.Println("unexpected error:", err)
	}
}
