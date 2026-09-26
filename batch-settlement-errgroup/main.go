package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

type task struct {
	account string
	delay   time.Duration
}

var errAccountFrozen = errors.New("account frozen")

func settle(ctx context.Context, account string, d time.Duration) error {
	select {
	case <-ctx.Done():
		return fmt.Errorf("settle %s: %w", account, ctx.Err())
	case <-time.After(d):
	}

	if account == "ACC-203" {
		return fmt.Errorf("settle %s: %w", account, errAccountFrozen)
	}

	fmt.Println("settled", account)
	return nil
}

func main() {
	tasks := []task{
		{"ACC-201", 50 * time.Millisecond},
		{"ACC-202", 80 * time.Millisecond},
		{"ACC-203", 100 * time.Millisecond},
		{"ACC-204", 400 * time.Millisecond},
		{"ACC-205", 450 * time.Millisecond},
	}

	g, ctx := errgroup.WithContext(context.Background())

	g.SetLimit(3)

	for _, t := range tasks {
		g.Go(func() error {
			return settle(ctx, t.account, t.delay)
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println("batch failed:", err)
		fmt.Println("frozen account involved?", errors.Is(err, errAccountFrozen))
		return
	}
	fmt.Println("all accounts settled")
}
