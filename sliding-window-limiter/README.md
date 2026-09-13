# sliding-window-limiter

Go-তে **circular/ring-buffer sliding window + `%` wrap** শেখার ছোট example — fixed-size window-এ ট্রাফিক ট্র্যাক করে threshold-cross এ `total > 200` limit।

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

### Lines 5–9

```go
type SlidingWindow struct {
	counts []int
	size   int
	head   int
}
```

- `counts` — fixed-capacity **ring buffer** (স্লাইস)।
- `size` — window length (slots-সংখ্যা)।
- `head` — current slot-pointer।

### Lines 11–17

```go
func NewSlidingWindow(size int) *SlidingWindow {
	return &SlidingWindow{
		counts: make([]int, size),
		size:   size,
		head:   0,
	}
}
```

Constructor: zero-filled `counts` + pointer struct-return।

### Lines 19–22

```go
func (w *SlidingWindow) Advance(n int) {
	w.head = (w.head + 1) % w.size
	w.counts[w.head] = n
}
```

**Advance — circular write:**

- `(w.head + 1) % w.size` — **প্রায় বৃত্ত**: size-এ পৌঁছালে 0-তে ফেরত (ring wrap)।
- `w.counts[w.head] = n` — পুরনো ভ্যালু overwrite — `size`-এর সবচেয়ে **পুরনো tick-ই বেরিয়ে যায়**।

### Lines 24–30

```go
func (w *SlidingWindow) Total() int {
	total := 0
	for _, count := range w.counts {
		total += count
	}
	return total
}
```

Window-এর সব slot যোগ — অর্থাৎ **শেষ `size` টিক-এর সমষ্টি** (sliding-window sum)।

### Lines 32–38

```go
func main() {
	w := NewSlidingWindow(5)
	traffic := []int{20, 25, 30, 35, 40, 90, 95}
```

- Window size 5।
- 7টা tick-এর request-count, শেষ 2টা স্পাইক (90, 95)।

### Lines 36–42

```go
	for tick, n := range traffic {
		w.Advance(n)
		fmt.Printf("tick %d: +%d reqs, in window = %d\n", tick, n, w.Total())

		if w.Total() > 200 {
			fmt.Printf("  -> RATE LIMITED at tick %d\n", tick)
		}
	}
```

- প্রতিটা tick Advance + Total।
- `> 200` cross → `RATE LIMITED` flag।

**Walkthrough (size 5):**

- tick 0–4: window 20→150 (তাই পুরোটা ধরতে পারছে)।
- tick 5: +90, window = 90+25+30+35+40 = **220** → limit → ওভার 200-tick-তেই ৫টা পুরনো slot-j-এর সমষ্টি।
- tick 6: +95, window = 90+95+30+35+40 = **290** → আবার limit।

---

## Expected Output

```
tick 0: +20 reqs, in window = 20
tick 1: +25 reqs, in window = 45
tick 2: +30 reqs, in window = 75
tick 3: +35 reqs, in window = 110
tick 4: +40 reqs, in window = 150
tick 5: +90 reqs, in window = 220
  -> RATE LIMITED at tick 5
tick 6: +95 reqs, in window = 290
  -> RATE LIMITED at tick 6
```

## মূল শিক্ষা / Key Takeaways

1. **Ring buffer** — fixed `counts` + `% size` wrap।
2. **Old-value eviction** — head advance-এ পুরনো slot overwrite।
3. **Sliding-window sum** — `Total()` = শেষ size-টিক sum।
4. **Threshold limit** — `> 200` → rate-limited flag।
5. **Constructor `*T`** — `NewSlidingWindow`।

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

### Lines 5–9

```go
type SlidingWindow struct {
	counts []int
	size   int
	head   int
}
```

- `counts` — a fixed-capacity **ring buffer** (slice).
- `size` — the window length (number of slots).
- `head` — the current slot pointer.

### Lines 11–17

```go
func NewSlidingWindow(size int) *SlidingWindow {
	return &SlidingWindow{
		counts: make([]int, size),
		size:   size,
		head:   0,
	}
}
```

A constructor: zero-filled `counts` + a pointer-struct return.

### Lines 19–22

```go
func (w *SlidingWindow) Advance(n int) {
	w.head = (w.head + 1) % w.size
	w.counts[w.head] = n
}
```

**Advance — circular write:**

- `(w.head + 1) % w.size` — **nearly circular**: reaching `size` wraps back to 0 (ring wrap).
- `w.counts[w.head] = n` — overwrites the old value — the `size`-oldest tick falls off.

### Lines 24–30

```go
func (w *SlidingWindow) Total() int {
	total := 0
	for _, count := range w.counts {
		total += count
	}
	return total
}
```

Sums all window slots — i.e. the **sum of the last `size` ticks** (a sliding-window sum).

### Lines 32–38

```go
func main() {
	w := NewSlidingWindow(5)
	traffic := []int{20, 25, 30, 35, 40, 90, 95}
```

- Window size 5.
- 7 ticks of request counts, the last 2 spiking (90, 95).

### Lines 36–42

```go
	for tick, n := range traffic {
		w.Advance(n)
		fmt.Printf("tick %d: +%d reqs, in window = %d\n", tick, n, w.Total())

		if w.Total() > 200 {
			fmt.Printf("  -> RATE LIMITED at tick %d\n", tick)
		}
	}
```

- Each tick: Advance + Total.
- Crossing `> 200` → the `RATE LIMITED` flag.

**Walkthrough (size 5):**

- ticks 0–4: window 20 → 150 (it fits).
- tick 5: +90, window = 90+25+30+35+40 = **220** → limit crossed.
- tick 6: +95, window = 90+95+30+35+40 = **290** → again limited.

---

## Expected Output

```
tick 0: +20 reqs, in window = 20
tick 1: +25 reqs, in window = 45
tick 2: +30 reqs, in window = 75
tick 3: +35 reqs, in window = 110
tick 4: +40 reqs, in window = 150
tick 5: +90 reqs, in window = 220
  -> RATE LIMITED at tick 5
tick 6: +95 reqs, in window = 290
  -> RATE LIMITED at tick 6
```

## Key Takeaways

1. **Ring buffer** — fixed `counts` + `% size` wrap.
2. **Old-value eviction** — head advance overwrites the oldest slot.
3. **Sliding-window sum** — `Total()` = the last-`size`-ticks sum.
4. **Threshold limit** — `> 200` → the rate-limited flag.
5. **Constructor `*T`** — `NewSlidingWindow`.