package main

import (
	"cmp"
	"errors"
	"fmt"
	"slices"
	"sync"
	"time"
)

var errCorrupt = errors.New("corrupt image")

type thumbJob struct {
	id   int
	file string
}

type thumbResult struct {
	id   int
	path string
	err  error
}

func makeThumbnail(j thumbJob) (string, error) {
	time.Sleep(50 * time.Millisecond)
	if j.file == "" {
		return "", fmt.Errorf("thumbnail for product %d: %w", j.id, errCorrupt)
	}

	return "thumbs/" + j.file, nil
}

func worker(jobs <-chan thumbJob, results chan<- thumbResult, wg *sync.WaitGroup) {
	defer wg.Done()

	for j := range jobs {
		path, err := makeThumbnail(j)
		results <- thumbResult{id: j.id, path: path, err: err}
	}
}

func main() {
	files := []string{"shoe.jpg", "bag.jpg", "", "watch.jpg", "belt.jpg", "cap.jpg"}
	const numWorkers = 3

	jobs := make(chan thumbJob)
	results := make(chan thumbResult)

	var wg sync.WaitGroup

	for range numWorkers {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	go func() {
		defer close(jobs)

		for i, f := range files {
			jobs <- thumbJob{id: i + 1, file: f}
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	start := time.Now()

	var collected []thumbResult

	for r := range results {
		collected = append(collected, r)
	}

	elapsed := time.Since(start)

	slices.SortFunc(collected, func(a, b thumbResult) int {
		return cmp.Compare(a.id, b.id)
	})

	var failed int

	for _, r := range collected {
		if r.err != nil {
			failed++
			fmt.Println("product", r.id, "error:", r.err)
			continue
		}
		fmt.Println("product", r.id, "->", r.path)
	}
	fmt.Printf("done: %d ok, %d failed, faster than sequential: %v\n", len(collected)-failed, failed, elapsed < 250*time.Millisecond)
}
