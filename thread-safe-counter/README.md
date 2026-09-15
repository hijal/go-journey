# thread-safe-counter

Go-তে **`sync.Mutex` + `defer` unlock + `sync.WaitGroup`** শেখার ছোট example — concurrent-safe counter।

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

### Lines 3–6

```go
import (
	"fmt"
	"sync"
)
```

- `fmt` — `Println`।
- `sync` — `Mutex`, `WaitGroup`।

### Lines 8–11

```go
type SafeCounter struct {
	mu    sync.Mutex
	count int
}
```

`SafeCounter` — unexported `mu` (lock) + `count`। field হিসেবে mutex (embedded pointer version থেকে clear)।

### Lines 13–17

```go
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}
```

**Increment:** `Lock()` → `defer Unlock()` (সব exit-path-এ unlock নিশ্চিত) → `count++`।

### Lines 19–23

```go
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
```

**Read-ও lock-এ** — atomic রিড (মিড-`Increment`-এ ভুল মান পড়বে না)।

*(Mutex-এর নিয়ম: যে shared state write-ও করে, read-ও lock-protect করবে।)*

### Lines 25–27

```go
	counter := &SafeCounter{}
	var wg sync.WaitGroup
```

Zero-value mutex use-able; `WaitGroup` goroutine-track।

### Lines 29–35

```go
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
```

100 বারের side-effect, 100 concurrent `Increment()` — প্রতিটা goroutine-এ `wg.Add(1)` + `defer wg.Done()`। `Wait()` সব শেষে main-কে অপেক্ষা করায়।

### Line 38

```go
	fmt.Println("Total requests processed:", counter.Value())
```

`Value()` — lock-protected read → `100`।

---

## Expected Output

```
Total requests processed: 100
```

*(একই আউটপুট `-race` detector-সহ ও clean run-এ।)*

## মূল শিক্ষা / Key Takeaways

1. **`sync.Mutex`** — shared counter protect।
2. **`defer Unlock()`** — panic-ও unlock হয়।
3. **Lock-protected read** — torn-read নাই।
4. **`sync.WaitGroup`** — goroutine completion।
5. **Race-detector** — `go run -race` টেস্টে safety prove।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Lines 3–6

```go
import (
	"fmt"
	"sync"
)
```

- `fmt` — for `Println`.
- `sync` — for `Mutex`, `WaitGroup`.

### Lines 8–11

```go
type SafeCounter struct {
	mu    sync.Mutex
	count int
}
```

`SafeCounter` — an unexported `mu` (lock) + `count`. The mutex is a field (clearer than the embedded-pointer variant).

### Lines 13–17

```go
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}
```

**Increment:** `Lock()` → `defer Unlock()` (guarantees unlocking on every exit path) → `count++`.

### Lines 19–23

```go
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
```

**Reads are locked too** — atomic reads (no stale mid-`Increment` values).

*(The mutex rule: any shared state you write must also be read under the lock.)*

### Lines 25–27

```go
	counter := &SafeCounter{}
	var wg sync.WaitGroup
```

The zero-value mutex is usable; `WaitGroup` tracks goroutines.

### Lines 29–35

```go
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
```

100 goroutines each `Increment()`; each does `wg.Add(1)` + `defer wg.Done()`. `Wait()` keeps `main` until all finish.

### Line 38

```go
	fmt.Println("Total requests processed:", counter.Value())
```

`Value()` — a lock-protected read → `100`.

---

## Expected Output

```
Total requests processed: 100
```

*(Same output under the `-race` detector and a clean run.)*

## Key Takeaways

1. **`sync.Mutex`** — protects the shared counter.
2. **`defer Unlock()`** — unlocks even on panic.
3. **Lock-protected reads** — no torn reads.
4. **`sync.WaitGroup`** — goroutine completion.
5. **Race detector** — `go run -race` proves safety.