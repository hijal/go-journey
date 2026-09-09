# customer-tier-discount

Go-তে **function-as-value** আর **conditional function assignment** দিয়ে tier-based discount শেখার ছোট example — `var applyDiscount func(float64) float64` পরে `if/else`-এ setup।

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

Console-এ output print করার জন্য `fmt` package import করা হয়।

### Line 5

```go
func main() {
```

Program-এর entry point।

### Lines 6–7

```go
customerTier := "gold"
listPrice := 4000.0
```

Input: customer-এর tier ("gold"), list price 4000.0।

### Line 9

```go
var applyDiscount func(float64) float64
```

**Function-type variable** — `applyDiscount` একটা variable যার type `func(float64) float64` (একটা float নেয়, একটা float return করে)। এখনো **unassigned (nil)** — মানুষের জন্য placeholder।

### Lines 11–23

```go
if customerTier == "gold" {
	applyDiscount = func(f float64) float64 {
		return f * 0.85
	}
} else if customerTier == "silver" {
	applyDiscount = func(f float64) float64 {
		return f * 0.92
	}
} else {
	applyDiscount = func(f float64) float64 {
		return f
	}
}
```

**Conditional function assignment** — tier-এর উপর ভিত্তি করে `applyDiscount`-এ conflict function assign:

- gold → 15% off (`f * 0.85`)
- silver → 8% off (`f * 0.92`)
- else → কোনো discount (identity, `f`)

প্রতিটা একটা **anonymous function** (inline literal) ফ্যাক্টর-এর মতো assignment-মাত্র। Selectivity function-names নয় — function-objectএ।

### Line 25

```go
finalPrice := applyDiscount(listPrice)
```

Function-টা **call** — `applyDiscount(4000.0)` = `4000 * 0.85` = 3400.0 (যদি gold)।

### Line 27

```go
fmt.Printf("tier: %s, final price: %.2f\n", customerTier, finalPrice)
```

`%.2f` — 2 decimal-এ: `tier: gold, final price: 3400.00`।

### Line 28

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
tier: gold, final price: 3400.00
```

## মূল শিক্ষা / Key Takeaways

1. **Function as value** — `var applyDiscount func(float64) float64` — variable-এ function store।
2. **Conditional assignment** — `if/else` দিয়ে কোন function assign হবে।
3. **Strategy pattern** — runtime-এ behavior select করা (strategies-এর মতো)।
4. **`%.2f`** — 2-decimal float format।
5. **Multiplier** — percent-কে decimal fraction-এ (`0.85` = 15% off)।

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

Imports the `fmt` package for console output.

### Line 5

```go
func main() {
```

Program entry point.

### Lines 6–7

```go
customerTier := "gold"
listPrice := 4000.0
```

Input: the customer's tier ("gold"), the list price 4000.0.

### Line 9

```go
var applyDiscount func(float64) float64
```

**Function-type variable** — `applyDiscount` is a variable of type `func(float64) float64` (takes a float, returns a float). It's still **unassigned (nil)** — a placeholder waiting for a function.

### Lines 11–23

```go
if customerTier == "gold" {
	applyDiscount = func(f float64) float64 {
		return f * 0.85
	}
} else if customerTier == "silver" {
	applyDiscount = func(f float64) float64 {
		return f * 0.92
	}
} else {
	applyDiscount = func(f float64) float64 {
		return f
	}
}
```

**Conditional function assignment** — assigns a function to `applyDiscount` based on the tier:

- gold → 15% off (`f * 0.85`)
- silver → 8% off (`f * 0.92`)
- else → no discount (identity, `f`)

Each is an **anonymous function** assigned wholesale. The selection is done via function objects, not function names.

### Line 25

```go
finalPrice := applyDiscount(listPrice)
```

Calls the function — `applyDiscount(4000.0)` = `4000 × 0.85` = 3400.0 (for gold).

### Line 27

```go
fmt.Printf("tier: %s, final price: %.2f\n", customerTier, finalPrice)
```

`%.2f` — 2 decimals: `tier: gold, final price: 3400.00`.

### Line 28

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
tier: gold, final price: 3400.00
```

## Key Takeaways

1. **Function as value** — `var applyDiscount func(float64) float64` — storing a function in a variable.
2. **Conditional assignment** — `if/else` decides which function gets assigned.
3. **Strategy pattern** — selecting behavior at runtime (strategies).
4. **`%.2f`** — 2-decimal float formatting.
5. **Multiplier** — percent as a decimal fraction (`0.85` = 15% off).