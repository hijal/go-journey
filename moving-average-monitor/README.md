# moving-average-monitor

Go-তে **builtin `max` + sliding-window moving average** শেখার ছোট example — latency-stream-এর window-avg monitor।

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

### Line 3

```go
import "fmt"
```

`fmt` — `Printf`।

### Lines 5–7

```go
func main() {
	latencies := []float64{120, 130, 125, 140, 200, 180, 90, 95, 100}
	const window = 3
```

- `latencies` — 9টা consecutive sample (ms)।
- `const window = 3` — moving window size।

### Lines 9–18

```go
	avgs := make([]float64, 0, len(latencies))

	for end := range len(latencies) {
		start := max(0, end+1-window)
		sum := 0.0
		for _, v := range latencies[start : end+1] {
			sum += v
		}
		avgs = append(avgs, sum/float64(end+1-start))
	}
```

**Moving-average core:**

- `avgs` — pre-allocation (len 0, cap 9)।
- `for end := range len(latencies)` — Go 1.22+ `end` = 0..8।
- `start := max(0, end+1-window)` — **builtin `max` (Go 1.21+)**: window `end+1-3`-এ পিছলে 0-তে clamp — অন্যথায় leading edge-এ (t=0,1) window-টা অসম্পূর্ণ।
- `latencies[start : end+1]` — slicing: শেষ `window` তা (অথবা শুরুতে যতটা পাওয়া যায়)।
- `sum / float64(end+1-start)` — **element-count দিয়ে divide** — নোট: t=0-এ `1`, t=1-এ `2`, পরে `3` — তাই edge-এ "short average" (সাধারণ moving-window behavior)।

### Lines 20–22

```go
	for i, a := range avgs {
		fmt.Printf("t=%d: moving avg = %.2f\n", i, a)
	}
```

Print — `%.2f` দুই-দশমিক।

**Verify:**

- t=0: `[120]` → 120.00
- t=2: `(120+130+125)/3` = 125.00
- t=4: `(125+140+200)/3` = 155.00
- t=8: `(90+95+100)/3` = 95.00

---

## Expected Output

```
t=0: moving avg = 120.00
t=1: moving avg = 125.00
t=2: moving avg = 125.00
t=3: moving avg = 131.67
t=4: moving avg = 155.00
t=5: moving avg = 173.33
t=6: moving avg = 156.67
t=7: moving avg = 121.67
t=8: moving avg = 95.00
```

## মূল শিক্ষা / Key Takeaways

1. **Builtin `max`** — Go 1.21+ `max(0, ...)` clamp।
2. **Slicing window** — `latencies[start:end+1]`।
3. **Edge-short average** — leading window ছোট, sample-count দিয়ে divide।
4. **`for end := range len(...)`** — index range (Go 1.22+)।
5. **Pre-allocation** — `make([]float64, 0, n)`।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Line 3

```go
import "fmt"
```

`fmt` — for `Printf`.

### Lines 5–7

```go
func main() {
	latencies := []float64{120, 130, 125, 140, 200, 180, 90, 95, 100}
	const window = 3
```

- `latencies` — 9 consecutive samples (ms).
- `const window = 3` — the moving window size.

### Lines 9–18

```go
	avgs := make([]float64, 0, len(latencies))

	for end := range len(latencies) {
		start := max(0, end+1-window)
		sum := 0.0
		for _, v := range latencies[start : end+1] {
			sum += v
		}
		avgs = append(avgs, sum/float64(end+1-start))
	}
```

**The moving-average core:**

- `avgs` — pre-allocated (len 0, cap 9).
- `for end := range len(latencies)` — Go 1.22+ `end` = 0..8.
- `start := max(0, end+1-window)` — **builtin `max` (Go 1.21+)**: clamping the window back to 0 when `end+1-3` would go negative — otherwise the leading edge (t=0,1) would get an incomplete window.
- `latencies[start : end+1]` — slicing: the last `window` points (or however many exist at the start).
- `sum / float64(end+1-start)` — **divide by the element count** — note t=0 → `1`, t=1 → `2`, then `3` — so the edge produces a "short average" (standard moving-window behavior).

### Lines 20–22

```go
	for i, a := range avgs {
		fmt.Printf("t=%d: moving avg = %.2f\n", i, a)
	}
```

Prints — `%.2f` two decimals.

**Verify:**

- t=0: `[120]` → 120.00
- t=2: `(120+130+125)/3` = 125.00
- t=4: `(125+140+200)/3` = 155.00
- t=8: `(90+95+100)/3` = 95.00

---

## Expected Output

```
t=0: moving avg = 120.00
t=1: moving avg = 125.00
t=2: moving avg = 125.00
t=3: moving avg = 131.67
t=4: moving avg = 155.00
t=5: moving avg = 173.33
t=6: moving avg = 156.67
t=7: moving avg = 121.67
t=8: moving avg = 95.00
```

## Key Takeaways

1. **Builtin `max`** — Go 1.21+ `max(0, ...)` clamping.
2. **Sliced window** — `latencies[start:end+1]`.
3. **Edge-short average** — a small leading window, divided by the sample count.
4. **`for end := range len(...)`** — index ranging (Go 1.22+).
5. **Pre-allocation** — `make([]float64, 0, n)`.