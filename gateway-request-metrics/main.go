package main

import (
	"fmt"
	"maps"
	"slices"
	"sync"
	"sync/atomic"
)

type metrics struct {
	mu       sync.Mutex
	byRoute  map[string]int
	totalReq atomic.Int64
}

func newMetrics() *metrics {
	return &metrics{byRoute: make(map[string]int)}
}

func (m *metrics) record(route string) {
	m.totalReq.Add(1)
	m.mu.Lock()
	defer m.mu.Unlock()
	m.byRoute[route]++
}

func (m *metrics) snapshot() map[string]int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return maps.Clone(m.byRoute)
}

func main() {
	m := newMetrics()

	routes := []string{"/checkout", "/search", "/cart"}

	var wg sync.WaitGroup

	for i := range 600 {
		wg.Go(func() {
			m.record(routes[i%len(routes)])
		})
	}
	wg.Wait()

	snapshot := m.snapshot()

	for _, route := range slices.Sorted(maps.Keys(snapshot)) {
		fmt.Printf("%-10s %d\n", route, snapshot[route])
	}

	fmt.Println("total requests:", m.totalReq.Load())
}
