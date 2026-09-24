# atomic-counter

Go-তে **`sync/atomic` typed atomic (`atomic.Int64`)** — shared counter-এর lock-free safe increment।

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

### Lines 3–7

```go
import (
	"fmt"
	"sync"
	"sync/atomic"
)
```

`sync` (`WaitGroup`) + **`sync/atomic`** (lock-free counter ops)।

### Lines 9–10

```go
func main() {
	var wg sync.WaitGroup
	var pageViews atomic.Int64
```

**`atomic.Int64`** — typed atomic; zero-value 0; সব goroutine শেয়ার করে।

### Lines 13–18

```go
	for range 1000 {
		wg.Go(func() {
			pageViews.Add(1)
		})
	}
	wg.Wait()
```

**`pageViews.Add(1)`** — **atomic increment** — এক command-এ read+update; mutex ছাড়াই safe (lost-update impossible)। Race detector-ও চুপ।

### Lines 19

```go
	fmt.Println("page views:", pageViews.Load())
```

**`Load()`** — atomic read (plain read-ও ok, but `Load` explicit-safe)।

### Mutex vs Atomic

| approach | safe? | `-race` | pros |
|---|---|---|---|
| plain `++` | ✗ | `DATA RACE` | — |
| `sync.Mutex` | ✓ | clean | flexible (multi-op) |
| **`atomic`** | ✓ | clean | lock-free, fastest |

> নিয়ম: single counter/flag → **atomic**; জটিল multi-statement state → **mutex**।

---

## Expected Output

```
page views: 1000
```

## মূল শিক্ষা / Key Takeaways

1. **`atomic.Int64`** — typed atomic variable।
2. **`Add(1)`** — atomic increment (lost-update impossible)।
3. **`Load()`** — explicit atomic read।
4. **Lock-free** — mutex sleep-cost নাই।
5. **`-race` clean** — shared state correct।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Lines 3–7

```go
import (
	"fmt"
	"sync"
	"sync/atomic"
)
```

`sync` (`WaitGroup`) + **`sync/atomic`** (lock-free counter ops).

### Lines 9–10

```go
func main() {
	var wg sync.WaitGroup
	var pageViews atomic.Int64
```

**`atomic.Int64`** — a typed atomic; zero value 0; shared by all the goroutines.

### Lines 13–18

```go
	for range 1000 {
		wg.Go(func() {
			pageViews.Add(1)
		})
	}
	wg.Wait()
```

**`pageViews.Add(1)`** — an **atomic increment** — read+update in one command; safe without a mutex (lost updates are impossible). Even the race detector stays quiet.

### Lines 19

```go
	fmt.Println("page views:", pageViews.Load())
```

**`Load()`** — an atomic read (a plain read works too, but `Load` is explicitly safe).

### Mutex vs Atomic

| approach | safe? | `-race` | pros |
|---|---|---|---|
| plain `++` | ✗ | `DATA RACE` | — |
| `sync.Mutex` | ✓ | clean | flexible (multi-op) |
| **`atomic`** | ✓ | clean | lock-free, fastest |

> Rule: a single counter/flag → **atomic**; complex multi-statement state → **mutex**.

---

## Expected Output

```
page views: 1000
```

## Key Takeaways

1. **`atomic.Int64`** — a typed atomic variable.
2. **`Add(1)`** — atomic increment (lost updates impossible).
3. **`Load()`** — explicit atomic read.
4. **Lock-free** — no mutex sleep cost.
5. **`-race` clean** — correct shared state.