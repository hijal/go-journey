# gateway-request-metrics

Go-তে **concurrent request metrics (Mutex-map + `atomic.Int64`)** — 600 goroutine request-এর route-wise count ও total।

**📖 ভাষা নির্বাচন করুন / Choose language:**

[🇧🇩 বাংলা](#bangla) • [🇬🇧 English](#english)

---

<a name="bangla"></a>

## 🇧🇩 বাংলা সংস্করণ

### Line 1

```go
package main
```

একটা executable program (`main` package) declare করে, যা `go run` দিয়ে চালানো যায়।

### Lines 3–9

```go
import (
	"fmt"
	"maps"
	"slices"
	"sync"
	"sync/atomic"
)
```

`maps` (clone/keys) + `slices` (sorted) + `sync` + `atomic`।

### Lines 11–15

```go
type metrics struct {
	mu       sync.Mutex
	byRoute  map[string]int
	totalReq atomic.Int64
}
```

**State structure:**
- `mu` — map access-কে সুরক্ষিত রাখে।
- `byRoute` — route → count।
- `totalReq` — **`atomic.Int64`** — lock ছাড়াই counter।

### Lines 21–26

```go
func (m *metrics) record(route string) {
	m.totalReq.Add(1)
	m.mu.Lock()
	defer m.mu.Unlock()
	m.byRoute[route]++
}
```

**record** — atomic total bump + lock-এর ভেতরে route-count।

### Lines 28–32

```go
func (m *metrics) snapshot() map[string]int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return maps.Clone(m.byRoute)
}
```

**snapshot** — externally-visible copy (`maps.Clone`) — caller map-টা mutate করলে state নষ্ট হবে না।

### Lines 41–45

```go
for i := range 600 {
	wg.Go(func() {
		m.record(routes[i%len(routes)])
	})
}
```

600 goroutine → round-robin → প্রতি route-এ **200** request।

### Lines 48–54

```go
snapshot := m.snapshot()

for _, route := range slices.Sorted(maps.Keys(snapshot)) {
	fmt.Printf("%-10s %d\n", route, snapshot[route])
}

fmt.Println("total requests:", m.totalReq.Load())
```

Sorted keys + atomic load।

---

## Expected Output

```
/cart      200
/checkout  200
/search    200
total requests: 600
```

## মূল শিক্ষা / Key Takeaways

1. **Mutex-guarded map** — grow-able state-এর standard shielding।
2. **`atomic.Int64` counter** — lock-less fast total।
3. **`maps.Clone` snapshot** — safe exported state।
4. **Mixed sync strategy** — per-কাঠামো সঠিক tool।
5. **Round-robin load** — 600/3 → 200 each।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Lines 3–9

```go
import (
	"fmt"
	"maps"
	"slices"
	"sync"
	"sync/atomic"
)
```

`maps` (clone/keys) + `slices` (sorted) + `sync` + `atomic`.

### Lines 11–15

```go
type metrics struct {
	mu       sync.Mutex
	byRoute  map[string]int
	totalReq atomic.Int64
}
```

**The state structure:**
- `mu` — protects the map access.
- `byRoute` — route → count.
- `totalReq` — an **`atomic.Int64`** — a counter without a lock.

### Lines 21–26

```go
func (m *metrics) record(route string) {
	m.totalReq.Add(1)
	m.mu.Lock()
	defer m.mu.Unlock()
	m.byRoute[route]++
}
```

**record** — atomic total bump + route-count under the lock.

### Lines 28–32

```go
func (m *metrics) snapshot() map[string]int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return maps.Clone(m.byRoute)
}
```

**snapshot** — an externally-visible copy (`maps.Clone`) — callers can't corrupt the state.

### Lines 41–45

```go
for i := range 600 {
	wg.Go(func() {
		m.record(routes[i%len(routes)])
	})
}
```

600 goroutines → round-robin → **200** requests per route.

### Lines 48–54

```go
snapshot := m.snapshot()

for _, route := range slices.Sorted(maps.Keys(snapshot)) {
	fmt.Printf("%-10s %d\n", route, snapshot[route])
}

fmt.Println("total requests:", m.totalReq.Load())
```

Sorted keys + an atomic load.

---

## Expected Output

```
/cart      200
/checkout  200
/search    200
total requests: 600
```

## Key Takeaways

1. **Mutex-guarded map** — the standard shield for growable state.
2. **`atomic.Int64` counter** — a lock-free fast total.
3. **`maps.Clone` snapshot** — safe exported state.
4. **Mixed sync strategy** — the right tool per structure.
5. **Round-robin load** — 600/3 → 200 each.