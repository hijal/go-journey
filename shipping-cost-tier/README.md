# shipping-cost-tier

Go-তে **`if / else if / else`** chain আর `fmt.Printf` formatting দিয়ে tier-based shipping cost হিসাব শেখার ছোট example।

**📖 ভাষা নির্বাচন করুন / Choose language:**

[🇧🇩 বাংলা](#bangla) • [🇬🇧 English](#english)

---

<a name="bangla"></a>

## 🇧🇩 বাংলা সংস্করণ

### Lines 1–3

```go
package main

import "fmt"
```

- `package main` — একটা executable program।
- `fmt` — output print করার জন্য।

### Line 5

```go
func main() {
```

Program-এর entry point।

### Lines 6–7

```go
cartTotal := 1450.0
var shippingCost float64
```

- `cartTotal := 1450.0` — cart-এর মোট দাম, `float64` (type inferred)।
- `var shippingCost float64` — shipping cost-এর জন্য একটা variable, কিন্তু **initialize করা হয়নি** — শুরুতে zero value (`0.0`)।

### Lines 9–15

```go
if cartTotal > 2000 {
    shippingCost = 0.0
} else if cartTotal > 1000 {
    shippingCost = 60.0
} else {
    shippingCost = 100.0
}
```

Tier-based শিপিং cost:

- `if cartTotal > 2000` → 2000-এর বেশি = **free shipping** (`0.0`)।
- `else if cartTotal > 1000` → 1000–2000 = **60.0**।
- `else` → 1000 বা তার কম = **100.0**।

এখানে cartTotal = 1450, তাই `1450 > 1000` → `shippingCost = 60.0`।

### Line 16

```go
fmt.Printf("Cart: %.2f | Shipping: %.2f | Payable: %.2f\n", cartTotal, shippingCost, cartTotal+shippingCost)
```

`fmt.Printf` দিয়ে তিনটা মান print করে, সব `%.2f` (দশমিকের পরে ২ ঘর):

- `cartTotal` → 1450.00
- `shippingCost` → 60.00
- `cartTotal+shippingCost` → 1450+60 = 1510.00 (payable টোটাল)

---

## Expected Output

```
Cart: 1450.00 | Shipping: 60.00 | Payable: 1510.00
```

## মূল শিক্ষা / Key Takeaways

1. **`if / else if / else`** — একাধিক শর্ত ক্রমান্বয়ে যাচাই; প্রথম match-টা চলে।
2. **`var shippingCost float64`** — declare করা but initial value ছাড়া; পরে assign করা।
3. **Zero value** — `float64`-এর declare করা uninitialized variable `0.0` দিয়ে শুরু।
4. **`fmt.Printf` / `%.2f`** — float-কে দশমিক ২ ঘরে format।
5. **Tier logic** — ক্রমবর্ধমান threshold-এ ভিন্ন cost।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Lines 1–3

```go
package main

import "fmt"
```

- `package main` — an executable program.
- `fmt` — for console output.

### Line 5

```go
func main() {
```

Program entry point.

### Lines 6–7

```go
cartTotal := 1450.0
var shippingCost float64
```

- `cartTotal := 1450.0` — the cart's total price, a `float64` (type inferred).
- `var shippingCost float64` — a variable for the shipping cost, but **not initialized** — starts at the zero value (`0.0`).

### Lines 9–15

```go
if cartTotal > 2000 {
    shippingCost = 0.0
} else if cartTotal > 1000 {
    shippingCost = 60.0
} else {
    shippingCost = 100.0
}
```

Tier-based shipping cost:

- `if cartTotal > 2000` → above 2000 = **free shipping** (`0.0`).
- `else if cartTotal > 1000` → 1000–2000 = **60.0**.
- `else` → 1000 or less = **100.0**.

Here cartTotal = 1450, so `1450 > 1000` → `shippingCost = 60.0`.

### Line 16

```go
fmt.Printf("Cart: %.2f | Shipping: %.2f | Payable: %.2f\n", cartTotal, shippingCost, cartTotal+shippingCost)
```

Prints three values with `fmt.Printf`, all as `%.2f` (2 decimal places):

- `cartTotal` → 1450.00
- `shippingCost` → 60.00
- `cartTotal+shippingCost` → 1450+60 = 1510.00 (the payable total)

---

## Expected Output

```
Cart: 1450.00 | Shipping: 60.00 | Payable: 1510.00
```

## Key Takeaways

1. **`if / else if / else`** — checks multiple conditions in order; the first match runs.
2. **`var shippingCost float64`** — declared without an initial value; assigned later.
3. **Zero value** — an uninitialized declared `float64` starts at `0.0`.
4. **`fmt.Printf` / `%.2f`** — formats a float to 2 decimal places.
5. **Tier logic** — different costs at increasing thresholds.
