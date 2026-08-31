# retry-backoff-timing

Go-তে **linear backoff** প্যাটার্ন আর `time.Duration` constants শেখার ছোট example — retry attempt-এর সাথে কতক্ষণ অপেক্ষা করবে, সেটা bounded করে।

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

- `fmt` — output print।
- `time` — `time.Duration`, `time.Millisecond`, `time.Second`।

### Lines 8–12

```go
const (
	InitialBackoff = 200 * time.Millisecond
	MaxBackoff     = 8 * time.Second
	BaseDelaySec   = 2
)
```

তিনটা config constant:

- `InitialBackoff = 200 * time.Millisecond` — প্রথম retry-তে অপেক্ষা: 200ms।
- `MaxBackoff = 8 * time.Second` — retry-র শেষ limit: 8 সেকেন্ড। এর চেয়ে বেশি অপেক্ষা করবে না।
- `BaseDelaySec = 2` — "base delay" (unused এই example-এ, তবে নামে hint দেয় key-value config-এ unit suffix)।

**Type inference:** `InitialBackoff` = `time.Duration`; `MaxBackoff` = `time.Duration`; `BaseDelaySec` = untyped `int`।

### Lines 14–20

```go
func backoff(attempt int) time.Duration {
	d := time.Duration(attempt) * InitialBackoff
	if d > MaxBackoff {
		return MaxBackoff
	}
	return d
}
```

`backoff` — attempt-নম্বর দিয়ে linear backoff compute করে:

- `d := time.Duration(attempt) * InitialBackoff` — attempt=0 → 0ms, attempt=1 → 200ms, 2 → 400ms, 3 → 600ms, 4 → 800ms...
- `if d > MaxBackoff` — cap: 8 সেকেন্ডের বেশি গেলে 8s return করে। (এই example-এ 0-4 attempts-এ cap hit হয় না, কিন্তু বড় attempt-এর জন্য protection।)

**কেন linear:** `attempt × InitialBackoff` — প্রতিটা retry-তে সমান বাড়ে (200ms করে)। Exponential backoff-এর বদলে `2×2×2`-এর মতো। এই pattern-এ প্রথম কয়টাতে দ্রুত আবার try হয়।

### Lines 22–26

```go
func main() {
	for attempt := range 5 {
		fmt.Printf("attempt %d -> backoff %v\n", attempt+1, backoff(attempt))
	}
}
```

`range 5` — loop 0-4 (5টা attempt)। `attempt+1` দিয়ে 1-based output:

- attempt 0 → backoff 0ms
- attempt 1 → 200ms
- attempt 2 → 400ms
- attempt 3 → 600ms
- attempt 4 → 800ms

(800ms MaxBackoff-এর চেয়ে ছোট, তাই cap hit হয় না।)

### Line 27

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
attempt 1 -> backoff 0s
attempt 2 -> backoff 200ms
attempt 3 -> backoff 400ms
attempt 4 -> backoff 600ms
attempt 5 -> backoff 800ms
```

## মূল শিক্ষা / Key Takeaways

1. **Linear backoff** — `attempt × constant` — প্রতিটা retry-তে সমান বাড়ে।
2. **`time.Duration` constants** — `time.Millisecond`, `time.Second` দিয়ে readable config।
3. **Cap (MaxBackoff)** — অনিশ্চিত retry-তে অপেক্ষাকে সীমিত করা।
4. **`time.Duration(attempt)`** — int-কে Duration-এ explicit cast।
5. **`range N`** — `for attempt := range 5` (Go 1.22+) — shorthand 0 to 4।

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

- `fmt` — for output.
- `time` — for `time.Duration`, `time.Millisecond`, `time.Second`.

### Lines 8–12

```go
const (
	InitialBackoff = 200 * time.Millisecond
	MaxBackoff     = 8 * time.Second
	BaseDelaySec   = 2
)
```

Three config constants:

- `InitialBackoff = 200 * time.Millisecond` — wait for the first retry: 200ms.
- `MaxBackoff = 8 * time.Second` — the upper bound: 8 seconds. It never waits longer than this.
- `BaseDelaySec = 2` — a "base delay" (unused in this example, but demonstrates unit-suffix naming in configs).

**Type inference:** `InitialBackoff` = `time.Duration`; `MaxBackoff` = `time.Duration`; `BaseDelaySec` = untyped `int`.

### Lines 14–20

```go
func backoff(attempt int) time.Duration {
	d := time.Duration(attempt) * InitialBackoff
	if d > MaxBackoff {
		return MaxBackoff
	}
	return d
}
```

`backoff` — computes a linear backoff for the given attempt number:

- `d := time.Duration(attempt) * InitialBackoff` — attempt=0 → 0ms, 1 → 200ms, 2 → 400ms, 3 → 600ms, 4 → 800ms...
- `if d > MaxBackoff` — cap: if the computed wait exceeds 8 seconds, return 8s instead. (Cap isn't hit on attempts 0–4 here, but it's protection for larger values.)

**Why linear:** `attempt × constant` — each retry adds the same increment (200ms). Unlike exponential backoff which multiplies, the linear pattern retries faster in the early attempts.

### Lines 22–26

```go
func main() {
	for attempt := range 5 {
		fmt.Printf("attempt %d -> backoff %v\n", attempt+1, backoff(attempt))
	}
}
```

`range 5` — loops 0–4 (5 attempts). `attempt+1` gives a 1-based display:

- attempt 0 → backoff 0s
- attempt 1 → 200ms
- attempt 2 → 400ms
- attempt 3 → 600ms
- attempt 4 → 800ms

(800ms is still under MaxBackoff, so the cap doesn't trigger.)

### Line 27

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
attempt 1 -> backoff 0s
attempt 2 -> backoff 200ms
attempt 3 -> backoff 400ms
attempt 4 -> backoff 600ms
attempt 5 -> backoff 800ms
```

## Key Takeaways

1. **Linear backoff** — `attempt × constant` — adds the same increment on each retry.
2. **`time.Duration` constants** — readable config via `time.Millisecond` / `time.Second`.
3. **Cap (MaxBackoff)** — bounds the wait on uncertain retries.
4. **`time.Duration(attempt)`** — explicit cast from int to Duration.
5. **`range N`** — `for attempt := range 5` (Go 1.22+) — shorthand for 0 to 4.