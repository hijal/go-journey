# parallel-health-check

Go-তে **goroutines + `sync.WaitGroup`** দিয়ে parallel operation শেখার ছোট example — health check চালানোর সময় সব server-এর check একসাথে (concurrent), waitgroup দিয়ে finish sync।

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
	"time"
)
```

- `fmt` — print।
- `sync` — `sync.WaitGroup`।
- `time` — `time.Sleep` (simulated check)।

### Line 9

```go
func main() {
```

Program-এর entry point।

### Line 10

```go
servers := []string{"web-01", "web-02", "db-01", "cache-01"}
```

৪টা server-এর slice — প্রতিটা health check হবে।

### Line 12

```go
var wg sync.WaitGroup
```

**`sync.WaitGroup`** — goroutine-গুলোর done-টা track করে main-কে wait করায়। Zero value-সহ ready (zero field struct)।

### Lines 14–21

```go
for _, server := range servers {
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(50 * time.Millisecond)
		fmt.Printf("health check %s: OK\n", server)
	}()
}
```

Loop-এ প্রতিটা server-এর জন্য একটা goroutine:

- `wg.Add(1)` — counter increment (প্রতিটা goroutine-এ ১)। 
- `go func() {...}()` — goroutine-এ anonymous function চালায়।
- ভেতরে:
  - `defer wg.Done()` — goroutine শেষের সময় counter decrement (finish signal)।
  - `time.Sleep(50ms)` — simulated network check।
  - `fmt.Printf(...)` — result print।

**Potential pitfall:** `server` loop variable capture — Go 1.22+ প্রতিটা iteration-এ fresh variable (সঠিক)। (পুরনো Go-তে capture বাগ থাকতো — shared loop var।)

**Concurrency:** ৪টা goroutine parallel — প্রতিটা 50ms sleep সমান্তরালে, মোট ~50ms (sequential-এর বদলে 4×50=200ms লাগতো)। Order non-deterministic (goroutine scheduling) — এখানে যে-কোনো sequence OK।

### Line 23

```go
wg.Wait()
```

**`wg.Wait()`** — main block করে যতক্ষণ counter 0 না হয় (সব `Done` না)। এই ছাড়া main শেষ হয়ে যেতো গোরoutines-এর আগেই (program exit)।

### Line 24

```go
fmt.Println("all health check finished")
```

সব check-এর পর: `all health check finished`।

### Line 25

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
health check web-01: OK
health check web-02: OK
health check db-01: OK
health check cache-01: OK
all health check finished
```

> Note: order change হতে পারে (goroutine scheduling-এর উপর)। সব-টা OK + `all health check finished` mandatory।

## মূল শিক্ষা / Key Takeaways

1. **`go func()`** — goroutine-এ launch — parallel execution।
2. **`sync.WaitGroup`** — `Add`/`Done`/`Wait` — goroutine-completion synchronization।
3. **`defer wg.Done()`** — finish-টা guaranteed (return path-নির্বিশেষে)।
4. **Loop var capture (Go 1.22+)** — প্রতিটা iteration-এ fresh variable।
5. **Deterministic wait** — `Wait()` main-কে অপেক্ষা করায় সব finish না হওয়া পর্যন্ত।

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
	"time"
)
```

- `fmt` — for printing.
- `sync` — for `sync.WaitGroup`.
- `time` — for `time.Sleep` (a simulated check).

### Line 9

```go
func main() {
```

Program entry point.

### Line 10

```go
servers := []string{"web-01", "web-02", "db-01", "cache-01"}
```

A slice of 4 servers — each will be health-checked.

### Line 12

```go
var wg sync.WaitGroup
```

**`sync.WaitGroup`** — tracks the goroutines' completion so main can wait. Ready at zero value (a zero-field struct).

### Lines 14–21

```go
for _, server := range servers {
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(50 * time.Millisecond)
		fmt.Printf("health check %s: OK\n", server)
	}()
}
```

In the loop, one goroutine per server:

- `wg.Add(1)` — increments the counter (1 per goroutine).
- `go func() {...}()` — runs an anonymous function in a goroutine.
- Inside:
  - `defer wg.Done()` — decrements the counter when the goroutine ends (the finish signal).
  - `time.Sleep(50ms)` — a simulated network check.
  - `fmt.Printf(...)` — prints the result.

**Potential pitfall:** the `server` loop-variable capture — since Go 1.22+ each iteration gets a fresh variable (correct). (Older Go would have a capture bug — a shared loop var.)

**Concurrency:** 4 goroutines in parallel — each sleeps 50ms at the same time, ~50ms total (instead of 4×50=200ms sequentially). The order is non-deterministic (goroutine scheduling) — any sequence is fine here.

### Line 23

```go
wg.Wait()
```

**`wg.Wait()`** — blocks main until the counter reaches 0 (all `Done`). Without this, main would finish before the goroutines (program exit).

### Line 24

```go
fmt.Println("all health check finished")
```

After all checks: `all health check finished`.

### Line 25

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
health check web-01: OK
health check web-02: OK
health check db-01: OK
health check cache-01: OK
all health check finished
```

> Note: the order may vary (depends on goroutine scheduling). The OK lines + `all health check finished` are guaranteed.

## Key Takeaways

1. **`go func()`** — launch in a goroutine — parallel execution.
2. **`sync.WaitGroup`** — `Add`/`Done`/`Wait` — goroutine-completion synchronization.
3. **`defer wg.Done()`** — the finish signal is guaranteed (regardless of return path).
4. **Loop-var capture (Go 1.22+)** — a fresh variable per iteration.
5. **Deterministic wait** — `Wait()` blocks main until all finish.