# status-code-analytics

Go-তে **map frequency-count + `slices.Sorted(maps.Keys())` deterministic report + linear max-scan** শেখার ছোট example — HTTP status-code analytics।

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
	"maps"
	"slices"
)
```

- `fmt` — `Printf`।
- `maps` — `Keys` (Go 1.21+)।
- `slices` — `Sorted` (Go 1.21+)।

### Lines 9–12

```go
func main() {
	statusCodes := []int{
		200, 404, 200, 500, 200, 404, 301, 200, 500, 200, 403, 404,
	}
```

12টা log-এক্সট্রিপ: 200×5, 404×3, 500×2, 301×1, 403×1।

### Lines 14–18

```go
	counts := make(map[int]int)

	for _, code := range statusCodes {
		counts[code]++
	}
```

**Frequency idiom** — `counts[code]++` (zero-value `0`)।

### Lines 20–22

```go
	for _, code := range slices.Sorted(maps.Keys(counts)) {
		fmt.Printf("HTTP %d: %d hits\n", code, counts[code])
	}
```

**Deterministic sorted iteration:**

- `maps.Keys(counts)` — keys-এর iterator।
- `slices.Sorted(...)` — sorted slice-তে convert (map loop random-order-এর বদলায়)।
- Print 200, 301, 403, 404, 500 — stable report।

*(না-সর্ট করলে map-এর random order-এ হাজির হতো।)*

### Lines 23–29

```go
	mostCode, mostHits := 0, 0

	for code, hits := range counts {
		if hits > mostHits {
			mostCode, mostHits = code, hits
		}
	}
	fmt.Printf("Most frequent: HTTP %d (%d hits)\n", mostCode, mostHits)
```

**Linear max-scan:**

- `mostCode, mostHits := 0, 0` — initializer।
- `if hits > mostHits` — update (strict `>` — টাই হলে প্রথমটা থাকে)।
- Result: HTTP 200 (5)।

---

## Expected Output

```
HTTP 200: 5 hits
HTTP 301: 1 hits
HTTP 403: 1 hits
HTTP 404: 3 hits
HTTP 500: 2 hits
Most frequent: HTTP 200 (5 hits)
```

## মূল শিক্ষা / Key Takeaways

1. **Zero-value count** — `counts[code]++`।
2. **Sorted-keys iteration** — `slices.Sorted(maps.Keys(m))`।
3. **Deterministic report** — random map order এড়ানো।
4. **Linear max-scan** — O(n) `> ` update।
5. **Key map** — `map[int]int`।

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
	"maps"
	"slices"
)
```

- `fmt` — for `Printf`.
- `maps` — for `Keys` (Go 1.21+).
- `slices` — for `Sorted` (Go 1.21+).

### Lines 9–12

```go
func main() {
	statusCodes := []int{
		200, 404, 200, 500, 200, 404, 301, 200, 500, 200, 403, 404,
	}
```

12 log strips: 200×5, 404×3, 500×2, 301×1, 403×1.

### Lines 14–18

```go
	counts := make(map[int]int)

	for _, code := range statusCodes {
		counts[code]++
	}
```

**The frequency idiom** — `counts[code]++` (zero-value `0`).

### Lines 20–22

```go
	for _, code := range slices.Sorted(maps.Keys(counts)) {
		fmt.Printf("HTTP %d: %d hits\n", code, counts[code])
	}
```

**Deterministic sorted iteration:**

- `maps.Keys(counts)` — an iterator over the keys.
- `slices.Sorted(...)` — converts to a sorted slice (instead of the map's random loop order).
- Prints 200, 301, 403, 404, 500 — a stable report.

*(Without sorting, the map would surface them in random order.)*

### Lines 23–29

```go
	mostCode, mostHits := 0, 0

	for code, hits := range counts {
		if hits > mostHits {
			mostCode, mostHits = code, hits
		}
	}
	fmt.Printf("Most frequent: HTTP %d (%d hits)\n", mostCode, mostHits)
```

**A linear max-scan:**

- `mostCode, mostHits := 0, 0` — an initializer.
- `if hits > mostHits` — update (strict `>` keeps the first on ties).
- Result: HTTP 200 (5).

---

## Expected Output

```
HTTP 200: 5 hits
HTTP 301: 1 hits
HTTP 403: 1 hits
HTTP 404: 3 hits
HTTP 500: 2 hits
Most frequent: HTTP 200 (5 hits)
```

## Key Takeaways

1. **Zero-value counting** — `counts[code]++`.
2. **Sorted-keys iteration** — `slices.Sorted(maps.Keys(m))`.
3. **Deterministic report** — avoiding the random map order.
4. **Linear max-scan** — an O(n) `>` update.
5. **Key map** — `map[int]int`.