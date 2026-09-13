package main

import "fmt"

type SlidingWindow struct {
	counts []int
	size   int
	head   int
}

func NewSlidingWindow(size int) *SlidingWindow {
	return &SlidingWindow{
		counts: make([]int, size),
		size:   size,
		head:   0,
	}
}

func (w *SlidingWindow) Advance(n int) {
	w.head = (w.head + 1) % w.size
	w.counts[w.head] = n
}

func (w *SlidingWindow) Total() int {
	total := 0
	for _, count := range w.counts {
		total += count
	}
	return total
}

func main() {
	w := NewSlidingWindow(5)
	traffic := []int{20, 25, 30, 35, 40, 90, 95}

	for tick, n := range traffic {
		w.Advance(n)
		fmt.Printf("tick %d: +%d reqs, in window = %d\n", tick, n, w.Total())

		if w.Total() > 200 {
			fmt.Printf("  -> RATE LIMITED at tick %d\n", tick)
		}
	}
}
