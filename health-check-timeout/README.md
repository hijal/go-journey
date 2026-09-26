# health-check-timeout

Go-তে **`time.After` budget-সহ select timeout** — কোনো service 150ms-এর মধ্যে না আসলে TIMEOUT।

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
	"time"
)
```

`time` (latency, timeout)।

### Lines 8–16

```go
func pingService(latency time.Duration) <-chan time.Duration {
	out := make(chan time.Duration, 1)

	go func() {
		time.Sleep(latency)
		out <- latency
	}()
	return out
}
```

**Async ping** — goroutine-এ sleep করে latency report-দেয়। Buffered cap-1 channel — receiver ভিন্ন-থাকা সত্ত্বেও send block-হয় না।

### Lines 18–23

```go
const budget = 150 * time.Millisecond

services := map[string]time.Duration{
	"auth-api":    60 * time.Millisecond,
	"billing-api": 900 * time.Millisecond,
}
```

`budget` — health-check deadline। দুটো service — একটি দ্রুত, একটি ধীর।

### Lines 26–32

```go
	for _, name := range []string{"auth-api", "billing-api"} {
		select {
		case took := <-pingService(services[name]):
			fmt.Printf("%-12s OK (%v)\n", name, took)
		case <-time.After(budget):
			fmt.Printf("%-12s TIMEOUT (over %v)\n", name, budget)
		}
	}
```

**`select` race**: `time.After(budget)`-র timer fire-এর আগে result-আসলে → **OK**; নাহলে → **TIMEOUT**।

- `auth-api` (60ms) → 150ms-এর ভেতরে → OK।
- `billing-api` (900ms) → budget cut, TIMEOUT।

---

## Expected Output

```
auth-api     OK (60ms)
billing-api  TIMEOUT (over 150ms)
```

## মূল শিক্ষা / Key Takeaways

1. **`time.After`** — এক-time timer যেটা fire-হতেই value পাঠায়।
2. **`select` dual-case** — result বনাম deadline, জেতা first one।
3. **Buffered result chan** — leak-free async reply।
4. **Deterministic** — latency budget-কে স্পষ্টভাবে অতিক্রম/মুতি-দিলে ফলাফল স্থির।
5. **Cancellation-lite** — goroutine পুরাতন-থেকে-থেকে-ই ঘুমিয়ে/লিখে পরে শেষ।

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
	"time"
)
```

`time` (latency, timeout).

### Lines 8–16

```go
func pingService(latency time.Duration) <-chan time.Duration {
	out := make(chan time.Duration, 1)

	go func() {
		time.Sleep(latency)
		out <- latency
	}()
	return out
}
```

**An async ping** — a goroutine sleeps for the latency and reports back. Buffered cap-1 channel — the send never blocks even if the receiver is elsewhere.

### Lines 18–23

```go
const budget = 150 * time.Millisecond

services := map[string]time.Duration{
	"auth-api":    60 * time.Millisecond,
	"billing-api": 900 * time.Millisecond,
}
```

`budget` — the health-check deadline. Two services: one fast, one slow.

### Lines 26–32

```go
	for _, name := range []string{"auth-api", "billing-api"} {
		select {
		case took := <-pingService(services[name]):
			fmt.Printf("%-12s OK (%v)\n", name, took)
		case <-time.After(budget):
			fmt.Printf("%-12s TIMEOUT (over %v)\n", name, budget)
		}
	}
```

**A `select` race**: the result arriving before the `time.After(budget)` timer fires → **OK**; otherwise → **TIMEOUT**.

- `auth-api` (60ms) → within 150ms → OK.
- `billing-api` (900ms) → cut by the budget → TIMEOUT.

---

## Expected Output

```
auth-api     OK (60ms)
billing-api  TIMEOUT (over 150ms)
```

## Key Takeaways

1. **`time.After`** — a one-shot timer that sends a value once it fires.
2. **`select` dual-case** — result vs deadline, first one wins.
3. **Buffered result chan** — a leak-free async reply.
4. **Deterministic** — the outcomes are stable since latencies clearly exceed/best the budget.
5. **Cancellation-lite** — the slow goroutine just finishes late on its own.