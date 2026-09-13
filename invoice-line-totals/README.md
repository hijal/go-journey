# invoice-line-totals

Go-তে **accumulated subtotal + `slices.Max` + float average** শেখার ছোট example — invoice line-item-এর টোটাল/সর্বোচ্চ/গড়।

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
	"slices"
)
```

- `fmt` — `Printf`।
- `slices` — `Max` (Go 1.21+)।

### Lines 8–9

```go
func main() {
	prices := []float64{24.99, 129.00, 7.25, 45.50}
```

4টা line-item দাম।

### Lines 11–14

```go
	var total float64
	for _, p := range prices {
		total += p
	}
```

**Accumulation:**

- `var total float64` — zero-value accumulator।
- `for _, p := range prices` — value-only iterate, `total += p`।
- `24.99 + 129.00 + 7.25 + 45.50 = 206.74`।

### Lines 16–17

```go
	expensive := slices.Max(prices)
	avg := total / float64(len(prices))
```

- `slices.Max(prices)` — **builtin stdlib max** (Go 1.21+) — 129.00।
- `total / float64(len(prices))` — cast দিয়ে count-দ্বারা divide — `206.74/4 = 51.685`।

### Lines 19–21

```go
	fmt.Printf("subtotal: %.2f\n", total)
	fmt.Printf("most expensive item: %.2f\n", expensive)
	fmt.Printf("average price: %.2f\n", avg)
```

`%.2f` — দুই-দশমিক। `51.685 → 51.69` (rounding)।

---

## Expected Output

```
subtotal: 206.74
most expensive item: 129.00
average price: 51.69
```

## মূল শিক্ষা / Key Takeaways

1. **Zero-value accumulator** — `var total float64`।
2. **`slices.Max`** — stdlib max (Go 1.21+)।
3. **Explicit cast division** — `total/float64(n)`।
4. **`%.2f` rounding** — currency-style display।

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
	"slices"
)
```

- `fmt` — for `Printf`.
- `slices` — for `Max` (Go 1.21+).

### Lines 8–9

```go
func main() {
	prices := []float64{24.99, 129.00, 7.25, 45.50}
```

4 line-item prices.

### Lines 11–14

```go
	var total float64
	for _, p := range prices {
		total += p
	}
```

**Accumulation:**

- `var total float64` — a zero-value accumulator.
- `for _, p := range prices` — value-only iteration, `total += p`.
- `24.99 + 129.00 + 7.25 + 45.50 = 206.74`.

### Lines 16–17

```go
	expensive := slices.Max(prices)
	avg := total / float64(len(prices))
```

- `slices.Max(prices)` — the **stdlib max** (Go 1.21+) — 129.00.
- `total / float64(len(prices))` — a cast divide by the count — `206.74/4 = 51.685`.

### Lines 19–21

```go
	fmt.Printf("subtotal: %.2f\n", total)
	fmt.Printf("most expensive item: %.2f\n", expensive)
	fmt.Printf("average price: %.2f\n", avg)
```

`%.2f` — two decimals. `51.685 → 51.69` (rounding).

---

## Expected Output

```
subtotal: 206.74
most expensive item: 129.00
average price: 51.69
```

## Key Takeaways

1. **Zero-value accumulator** — `var total float64`.
2. **`slices.Max`** — the stdlib max (Go 1.21+).
3. **Explicit cast division** — `total/float64(n)`.
4. **`%.2f` rounding** — currency-style display.