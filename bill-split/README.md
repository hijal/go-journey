# bill-split

Go-তে **explicit float64 constants**, percentage-based calculations, `float64(People)` type conversion আর `%.2f` formatting দিয়ে বিল ভাগ শেখার ছোট example।

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

### Lines 5–9

```go
const (
	ChargeRate = 0.10
	VatRate    = 0.15
	People     = 7
)
```

তিনটা **constant**:

- `ChargeRate = 0.10` — 10% service charge
- `VatRate = 0.15` — 15% VAT
- `People = 7` — মোট কত জন ভাগ করবে

**Exported** (Uppercase) — অন্যান্য package-এ access-যোগ্য হবে। `ChargeRate`/`VatRate` untyped float; `People` untyped int (default `int`)। Percentage-কে decimal fraction হিসেবে রাখা (0.10 = 10%)।

### Line 11

```go
func main() {
```

Program-এর entry point।

### Line 12

```go
subtotal := 2450.0
```

`subtotal` — বিলের আগের যোগফল: 2450.0 (float literal — `.0` দিয়ে float বোঝানো হয়েছে)।

### Lines 13–15

```go
serviceCharge := subtotal * ChargeRate
vat := (subtotal + serviceCharge) * VatRate
total := subtotal + serviceCharge + vat
```

Step-by-step বিল হিসাব:

- `serviceCharge` = 2450 × 0.10 = 245.00
- `vat` = (2450 + 245) × 0.15 = 2695 × 0.15 = 404.25
- `total` = 2450 + 245 + 404.25 = 3099.25

**লক্ষ্য করো:** VAT গণনা হয় `subtotal + serviceCharge`-র **উপরে** (service charge-এর উপরও VAT)। Percentage-based billing-standard।

### Line 17

```go
perHead := total / float64(People)
```

`total / People` — কিন্তু `total` float, `People` int। **Type mismatch** — তাই `float64(People)` দিয়ে `People`-কে float-এ **cast** করা হয়েছে। Result = 3099.25 / 7 = 442.75। Go implicit int↔float conversion-কে **allow করে না** — explicit cast লাগে।

### Lines 19–23

```go
fmt.Printf("subtotal: %.2f\n", subtotal)
fmt.Printf("service charge: %.2f\n", serviceCharge)
fmt.Printf("vat: %.2f\n", vat)
fmt.Printf("total: %.2f\n", total)
fmt.Printf("per head (%d people): %.2f\n", People, perHead)
```

`fmt.Printf` + `%.2f` — float-কে **২ দশমিক**-এ format করে:

- `subtotal: 2450.00`
- `service charge: 245.00`
- `vat: 404.25`
- `total: 3099.25`
- `per head (7 people): 442.75`

`%d` int-এর জন্য, `%.2f` float-এর জন্য (২ digit decimal)।

### Line 24

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
subtotal: 2450.00
service charge: 245.00
vat: 404.25
total: 3099.25
per head (7 people): 442.75
```

## মূল শিক্ষা / Key Takeaways

1. **Percentage as fraction** — 10% → `0.10`, 15% → `.15` (decimal fraction)।
2. **Explicit cast** — `float64(People)`; Go int↔float auto-convert করে না।
3. **`%.2f` formatting** — ২ decimal place-এ float।
4. **Exported constants** — Uppercase নাম অন্য package-এ visible।
5. **VAT stacking** — VAT প্রায়ই subtotal+charge-এর উপর।

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

### Lines 5–9

```go
const (
	ChargeRate = 0.10
	VatRate    = 0.15
	People     = 7
)
```

Three **constants**:

- `ChargeRate = 0.10` — 10% service charge
- `VatRate = 0.15` — 15% VAT
- `People = 7` — how many people split it

**Exported** (uppercase) — accessible from other packages where relevant. `ChargeRate`/`VatRate` are untyped floats; `People` is an untyped int (defaults to `int`). Percentages are stored as decimal fractions (0.10 = 10%).

### Line 11

```go
func main() {
```

Program entry point.

### Line 12

```go
subtotal := 2450.0
```

`subtotal` — the bill subtotal: 2450.0 (a float literal — the `.0` marks it float).

### Lines 13–15

```go
serviceCharge := subtotal * ChargeRate
vat := (subtotal + serviceCharge) * VatRate
total := subtotal + serviceCharge + vat
```

Step-by-step bill math:

- `serviceCharge` = 2450 × 0.10 = 245.00
- `vat` = (2450 + 245) × 0.15 = 2695 × 0.15 = 404.25
- `total` = 2450 + 245 + 404.25 = 3099.25

Note: VAT is computed **on top of** `subtotal + serviceCharge` (VAT applies to the service charge too) — standard billing math.

### Line 17

```go
perHead := total / float64(People)
```

`total / People` — but `total` is float, `People` is int. A **type mismatch**, so `People` is **cast** to float with `float64(People)`. Result = 3099.25 / 7 = 442.75. Go doesn't **allow** implicit int↔float conversion — an explicit cast is required.

### Lines 19–23

```go
fmt.Printf("subtotal: %.2f\n", subtotal)
fmt.Printf("service charge: %.2f\n", serviceCharge)
fmt.Printf("vat: %.2f\n", vat)
fmt.Printf("total: %.2f\n", total)
fmt.Printf("per head (%d people): %.2f\n", People, perHead)
```

`fmt.Printf` + `%.2f` — formats floats to **2 decimals**:

- `subtotal: 2450.00`
- `service charge: 245.00`
- `vat: 404.25`
- `total: 3099.25`
- `per head (7 people): 442.75`

`%d` for ints, `%.2f` for floats (2 decimal digits).

### Line 24

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
subtotal: 2450.00
service charge: 245.00
vat: 404.25
total: 3099.25
per head (7 people): 442.75
```

## Key Takeaways

1. **Percentage as fraction** — 10% → `0.10`, 15% → `0.15` (decimal fraction).
2. **Explicit cast** — `float64(People)`; Go doesn't auto-convert int↔float.
3. **`%.2f` formatting** — floats rounded to 2 decimals.
4. **Exported constants** — uppercase names are visible to other packages.
5. **VAT stacking** — VAT is often charged on subtotal + charge.