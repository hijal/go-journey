package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"time"
)

func newLogger() *slog.Logger {
	opts := &slog.HandlerOptions{ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey {
			return slog.Attr{}
		}
		return a
	}}
	return slog.New(slog.NewTextHandler(os.Stdout, opts))
}

func processReport(id int) {
	time.Sleep(time.Millisecond * 40)
}

func consume(ctx context.Context, log *slog.Logger, jobs <-chan int) int {
	processed := 0

	for {
		if ctx.Err() != nil {
			return processed
		}

		select {
		case <-ctx.Done():
			return processed
		case id, ok := <-jobs:
			if !ok {
				return processed
			}
			log.Info("report started", "id", id)
			processReport(id)
			processed++
			log.Info("report finished", "id", id)
		}
	}
}

func main() {
	log := newLogger()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()

	jobs := make(chan int, 10)
	
	for id := range 10 {
		jobs <- id + 1
	}
	close(jobs)

	done := make(chan int)

	go func() { done <- consume(ctx, log, jobs) }()
	processed := <-done
	log.Info("shutdown complete", "processed", processed, "left_in_queue", len(jobs), "reason", ctx.Err())
}
