package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func heartbeat(ctx context.Context, every time.Duration) {
	ticker := time.NewTicker(every)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("heartbeat: stopping")
			return
		case <-ticker.C:
			fmt.Println("heartbeat: service alive")
		}
	}
}

func consumer(ctx context.Context, name string, queue <-chan string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s: stopping\n", name)
			return
		case msg, ok := <-queue:
			if !ok {
				return
			}
			fmt.Printf("%s: handled %s\n", name, msg)
			time.Sleep(40 * time.Millisecond)
		}
	}
}

func main() {
	ctx, stop := context.WithTimeout(context.Background(), 250*time.Millisecond)
	defer stop()

	var wg sync.WaitGroup

	queue := make(chan string)

	wg.Go(func() { heartbeat(ctx, 100*time.Millisecond) })
	wg.Go(func() { consumer(ctx, "worker-1", queue) })

	wg.Go(func() {
		for i := 1; ; i++ {
			select {
			case <-ctx.Done():
				return
			case queue <- fmt.Sprintf("msg-%d", i):
				time.Sleep(60 * time.Millisecond)
			}
		}
	})

	wg.Wait()
	fmt.Println("shutdown complete, reason:", context.Cause(ctx))
}
