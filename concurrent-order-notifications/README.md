# concurrent-order-notifications

Go-তে **`wg.Go` দিয়ে parallel notification fan-out** — checkout-পরবর্তী email/sms/push একসাথে চলে, মোট সময় ~100ms-এ সম্পন্ন।

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

`sleep/duration` (সময়) + `WaitGroup` (সিঙ্ক)।

### Lines 9–12

```go
func notify(channel string, orderID int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("notification sent via %s for order %d\n", channel, orderID)
}
```

Simulated notification — 100ms-এর external-call talk-time।

### Lines 14–23

```go
func main() {
	start := time.Now()

	var wg sync.WaitGroup

	for _, channel := range []string{"email", "sms", "push"} {
		wg.Go(func() {
			notify(channel, 5001)
		})
	}
```

**Parallel fan-out** — প্রতিটি channel নিজ নিজ goroutine-এ; `time.Sleep` এখন side-by-side। Loop var `channel` — Go 1.22+ এ per-iteration safe (capture bug নেই)।

### Lines 25–27

```go
	wg.Wait()

	fmt.Println("checkout finished in about", time.Since(start).Round(50*time.Millisecond))
```

সব শেষে **মোট সময়** — `Round(50ms)`-এ curated:

- **Sequential** হলে: 3 × 100ms = 300ms।
- **কনকারেন্ট এখানে**: ~100ms (সবগুলো টা সমান্তরাল)।

---

## Expected Output (order varies)

```
notification sent via push for order 5001
notification sent via email for order 5001
notification sent via sms for order 5001
checkout finished in about 100ms
```

## মূল শিক্ষা / Key Takeaways

1. **Fan-out** — independent কাজগুলো parallel goroutine-ই দাও।
2. **Elapsed-time win** — 300ms → 100ms।
3. **Output order nondeterministic** — scheduling-এর উপর।
4. **Loop var safe** — Go 1.22++-তে closure capture bug gone।
5. **`wg.Wait`** — main সবগুলো শেষ না হওয়া অবধি অপেক্ষা।

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

`sleep/duration` (time) + `WaitGroup` (sync).

### Lines 9–12

```go
func notify(channel string, orderID int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("notification sent via %s for order %d\n", channel, orderID)
}
```

A simulated notification — a 100ms external-talking time.

### Lines 14–23

```go
func main() {
	start := time.Now()

	var wg sync.WaitGroup

	for _, channel := range []string{"email", "sms", "push"} {
		wg.Go(func() {
			notify(channel, 5001)
		})
	}
```

**Parallel fan-out** — each channel runs in its own goroutine; the `time.Sleep` now happens side-by-side. Loop var `channel` — per-iteration safe since Go 1.22 (capture bug gone).

### Lines 25–27

```go
	wg.Wait()

	fmt.Println("checkout finished in about", time.Since(start).Round(50*time.Millisecond))
```

The **total elapsed time**, curated via `Round(50ms)`:

- **Sequential** would take: 3 × 100ms = 300ms.
- **Concurrent here**: ~100ms (all run in parallel).

---

## Expected Output (order varies)

```
notification sent via push for order 5001
notification sent via email for order 5001
notification sent via sms for order 5001
checkout finished in about 100ms
```

## Key Takeaways

1. **Fan-out** — give independent work to parallel goroutines.
2. **Elapsed-time win** — 300ms → 100ms.
3. **Output order is nondeterministic** — depends on scheduling.
4. **Loop var safe** — the closure capture bug is gone since Go 1.22.
5. **`wg.Wait`** — main waits until all are done.