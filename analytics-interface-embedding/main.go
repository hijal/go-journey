package main

import "fmt"

type Counter interface {
	Incr(name string)
}

type Gauge interface {
	Set(name string, value float64)
}

type Reporter interface {
	Counter
	Gauge
}

type memReporter struct {
	counters map[string]int
	gauges   map[string]float64
}

func newMemReporter() *memReporter {
	return &memReporter{
		counters: map[string]int{},
		gauges:   map[string]float64{},
	}
}

func (m *memReporter) Incr(name string)           { m.counters[name]++ }
func (m *memReporter) Set(name string, v float64) { m.gauges[name] = v }
func (m *memReporter) Snapshot() (map[string]int, map[string]float64) {
	return m.counters, m.gauges
}

func recordCheckoutAttempt(c Counter) { c.Incr("checkout_attempts") }

func main() {
	rep := newMemReporter()

	recordCheckoutAttempt(rep)
	recordCheckoutAttempt(rep)
	recordCheckoutAttempt(rep)

	var g Gauge = rep
	g.Set("active_carts", 12.5)

	var _ Reporter = rep

	counters, gauges := rep.Snapshot()

	fmt.Println("counters:", counters)
	fmt.Println("gauges:  ", gauges)
}
