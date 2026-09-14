# currency-conversion-map

Go-তে **global map + function-method + error paths + spread side-effect** শেখার ছোট example — currency converter।

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

`fmt` — `Printf`, `Println`, `Errorf`।

### Lines 5–10

```go
var rates = map[string]float64{
	"USD": 1.0,
	"EUR": 0.92,
	"BDT": 117.50,
	"GBP": 0.79,
}
```

**Package-level map** — USD-based rate। (৪টা currency।)

### Lines 12–26

```go
func convert(amount float64, from, to string) (float64, error) {
	fromRate, ok := rates[from]

	if !ok {
		return 0, fmt.Errorf("unsupported currency %q", from)
	}

	toRate, ok := rates[to]

	if !ok {
		return 0, fmt.Errorf("unsupported currency %q", to)
	}

	return amount / fromRate * toRate, nil
}
```

**Convert function:**

- **Comma-ok guard দিয়ে error-return**: wildcard `%q` quoting (`unsupported currency "XAU"`)।
- সূত্র: `amount / fromRate * toRate` — base-currency-normalize।
- Named-return-ছাড়া `(float64, error)`।

### Lines 28–32

```go
func applySpread(r map[string]float64, pct float64) {
	for code := range r {
		r[code] *= 1 + pct/100
	}
}
```

**Map side-effect:**

- Map **reference-type** — function-এ pass করলে caller-এর map-ই mutate হয় (copy নয়)।
- 2% spread: `× 1.02`।

### Lines 34–41

```go
func main() {
	bdt, err := convert(100, "USD", "BDT")

	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("100 USD = %.2f BDT\n", bdt)
```

Present-path: 100 USD → BDT: `100/1.0 × 117.5 = 11750.00`।

### Lines 43–45

```go
	if _, err := convert(50, "USD", "XAU"); err != nil {
		fmt.Println("error:", err)
	}
```

**Absent-path**: `XAU` নাই → `unsupported currency "XAU"`।

### Lines 47–49

```go
	applySpread(rates, 2)

	fmt.Printf("After 2%% spread, USD->BDT rate is %.2f\n", rates["BDT"])
```

- `rates` এখন mutated (BDT → 119.85) — spread তারপরেও পুরো map-এ।
- `%%` — literal percent।

---

## Expected Output

```
100 USD = 11750.00 BDT
error: unsupported currency "XAU"
After 2% spread, USD->BDT rate is 119.85
```

## মূল শিক্ষা / Key Takeaways

1. **Package-level map** — global conversion state।
2. **Comma-ok error guard** — missing-currency returns error।
3. **Reference-type side-effect** — function-এ pass = mutate।
4. **`fmt.Errorf` + `%q`** — quoting in error।
5. **`%%` escape** — literal percent in `Printf`।

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

`fmt` — for `Printf`, `Println`, `Errorf`.

### Lines 5–10

```go
var rates = map[string]float64{
	"USD": 1.0,
	"EUR": 0.92,
	"BDT": 117.50,
	"GBP": 0.79,
}
```

A **package-level map** — USD-based rates (4 currencies).

### Lines 12–26

```go
func convert(amount float64, from, to string) (float64, error) {
	fromRate, ok := rates[from]

	if !ok {
		return 0, fmt.Errorf("unsupported currency %q", from)
	}

	toRate, ok := rates[to]

	if !ok {
		return 0, fmt.Errorf("unsupported currency %q", to)
	}

	return amount / fromRate * toRate, nil
}
```

**The convert function:**

- **Comma-ok guard returning an error** — with the `%q` wildcard quoting (`unsupported currency "XAU"`).
- Formula: `amount / fromRate * toRate` — base-currency normalization.
- `(float64, error)` without named returns.

### Lines 28–32

```go
func applySpread(r map[string]float64, pct float64) {
	for code := range r {
		r[code] *= 1 + pct/100
	}
}
```

**Map side-effect:**

- Maps are **reference types** — passing one to a function mutates the caller's map (it's not copied).
- A 2% spread: `× 1.02`.

### Lines 34–41

```go
func main() {
	bdt, err := convert(100, "USD", "BDT")

	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("100 USD = %.2f BDT\n", bdt)
```

The present path: 100 USD → BDT: `100/1.0 × 117.5 = 11750.00`.

### Lines 43–45

```go
	if _, err := convert(50, "USD", "XAU"); err != nil {
		fmt.Println("error:", err)
	}
```

**The absent path**: `XAU` doesn't exist → `unsupported currency "XAU"`.

### Lines 47–49

```go
	applySpread(rates, 2)

	fmt.Printf("After 2%% spread, USD->BDT rate is %.2f\n", rates["BDT"])
```

- `rates` is now mutated (BDT → 119.85) — the spread applied across the whole map.
- `%%` — a literal percent.

---

## Expected Output

```
100 USD = 11750.00 BDT
error: unsupported currency "XAU"
After 2% spread, USD->BDT rate is 119.85
```

## Key Takeaways

1. **Package-level map** — global conversion state.
2. **Comma-ok error guard** — a missing currency returns an error.
3. **Reference-type side-effect** — passing = mutating.
4. **`fmt.Errorf` + `%q`** — quoting inside an error.
5. **`%%` escape** — a literal percent in `Printf`.