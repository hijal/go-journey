package main

import (
	"cmp"
	"fmt"
	"slices"
	"sync"
	"time"
)

type job struct {
	id   int
	file string
}

type result struct {
	id       int
	workerID int
	thumb    string
}

func worker(id int, jobs <-chan job, results chan<- result) {
	for j := range jobs {
		time.Sleep(50 * time.Millisecond)
		results <- result{id: j.id, workerID: id, thumb: j.file + ".thumb.jpg"}
	}
}

func main() {
	const workers = 4

	start := time.Now()

	jobs := make(chan job)
	results := make(chan result, 12)

	var wg sync.WaitGroup

	for w := range workers {
		wg.Go(func() { worker(w+1, jobs, results) })
	}

	go func() {
		defer close(jobs)
		for i := range 12 {
			jobs <- job{id: i + 1, file: fmt.Sprintf("upload-%02d", i+1)}
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	var done []result
	for r := range results {
		done = append(done, r)
	}

	slices.SortFunc(done, func(a, b result) int {
		return cmp.Compare(a.id, b.id)
	})

	for _, r := range done {
		fmt.Printf("job %2d -> %s (worker %d)\n", r.id, r.thumb, r.workerID)
	}

	fmt.Printf("processed %d uploads with %d workers in %v\n", len(done), workers, time.Since(start).Round(50*time.Millisecond))
}
