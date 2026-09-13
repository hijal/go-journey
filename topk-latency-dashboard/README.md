# topk-latency-dashboard

Go-তে **top-K extraction + percentile from a sorted slice** শেখার ছোট example — API latency monitoring।

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
	"math/rand"
	"slices"
)
```

- `fmt` — `Print`, `Println` (no-fmt line-printing)।
- `math/rand` — random latency।
- `slices` — `Sort`।

### Lines 9–14

```go
func main() {
	latencies := make([]int, 50)

	for i := range latencies {
		latencies[i] = 20 + rand.Intn(400)
	}
```

- `make([]int, 50)` — 50-স্লট slice (ইতিমধ্যে zero-filled)।
- `for i := range latencies` — index-ভিত্তিক র্যান্ডম fill।
- `20 + rand.Intn(400)` — value range `20..419` ms — base 20 + 0-399 jitter।

### Lines 16–17

```go
	slices.Sort(latencies)
	slowest := latencies[len(latencies)-3:]
```

- `Sort` — ascending।
- `latencies[len-3:]` — **last-3 slice** = সর্টেড-এর সবচেয়ে ধীর ৩টা (top-K)। সর্ট করা slice-এ terminal window-ই TOP/K।

> **শিখবেন:** sorting + slicing = সহজ "top-k" এবং percentile। Sort না-করলে এক-একটা O(n) অতিরিক্ত scan দরকার হতো।

### Lines 18–22

```go
	fmt.Print("slowest 3 requests (ms): ")
	for _, latency := range slowest {
		fmt.Print(latency, " ")
	}
	fmt.Println()
```

- `Print` — newline ছাড়া (same line)।
- Iterate slowest — space-separated, তারপর `Println()` — newline finish।

### Lines 24–25

```go
	p95 := latencies[int(float64(len(latencies))*0.95)-1]
	fmt.Println("p95 latency (ms):", p95)
```

- `float64(50)*0.95 = 47.5` → `int()` → `47` → `-1` → **index 46** (47th element)।
- Sorted array-র middle-of-data point-টা **95th percentile**-এর approximation — data-র 95% এই মানের চেয়ে ছোট/সমান।
- (Exact-statistics-এই n*0.95 rounding variant; এটা pedagogue-র সহজ approximation।)

---

## Expected Output

(সম্ভাব্য — random-এর জন্য ভিন্ন হতে পারে)

```
slowest 3 requests (ms): 407 413 418 
p95 latency (ms): 365
```

## মূল শিক্ষা / Key Takeaways

1. **Sort + tail-slice** — `slices.Sort` + `[len-3:]` = top-K।
2. **Percentile approximation** — `int(len*0.95)-1` index।
3. **Range-fill** — `make` + index loop।
4. **`Print` vs `Println`** — same-line output।
5. **Sort before slice** — O(n·log n) একবার, পরে O(1) index-অ্যাক্সেস।

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
	"math/rand"
	"slices"
)
```

- `fmt` — for `Print`, `Println`.
- `math/rand` — random latencies.
- `slices` — for `Sort`.

### Lines 9–14

```go
func main() {
	latencies := make([]int, 50)

	for i := range latencies {
		latencies[i] = 20 + rand.Intn(400)
	}
```

- `make([]int, 50)` — a 50-slot slice (already zero-filled).
- `for i := range latencies` — index-based random fill.
- `20 + rand.Intn(400)` — value range `20..419` ms — base 20 + jitter 0–399.

### Lines 16–17

```go
	slices.Sort(latencies)
	slowest := latencies[len(latencies)-3:]
```

- `Sort` — ascending.
- `latencies[len-3:]` — the **last-3 slice** = the 3 slowest of the sorted data (top-K). On a sorted slice, a tail window *is* the top-K.

> **Lesson:** sorting + slicing = easy "top-k" and percentiles. Without sorting, each would need an O(n) scan.

### Lines 18–22

```go
	fmt.Print("slowest 3 requests (ms): ")
	for _, latency := range slowest {
		fmt.Print(latency, " ")
	}
	fmt.Println()
```

- `Print` — no newline (same line).
- Iterate slowest — space-separated, then a `Println()` to finish the line.

### Lines 24–25

```go
	p95 := latencies[int(float64(len(latencies))*0.95)-1]
	fmt.Println("p95 latency (ms):", p95)
```

- `float64(50)*0.95 = 47.5` → `int()` → `47` → `-1` → **index 46** (the 47th element).
- This midpoint-of-data point on the sorted array is an **approximation** of the 95th percentile — 95% of the data is ≤ this value.
- (Exact stats would use a rounding variant of n·0.95; this is the simple pedagogical version.)

---

## Expected Output

(possible — varies with randomness)

```
slowest 3 requests (ms): 407 413 418 
p95 latency (ms): 365
```

## Key Takeaways

1. **Sort + tail-slice** — `slices.Sort` + `[len-3:]` = top-K.
2. **Percentile approximation** — the `int(len*0.95)-1` index.
3. **Range-fill** — `make` + an index loop.
4. **`Print` vs `Println`** — same-line output.
5. **Sort before slice** — one O(n·log n), then O(1) index access.