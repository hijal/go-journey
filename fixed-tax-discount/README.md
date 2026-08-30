# fixed-tax-discount

Go-তে **multiple `const`**, **`fmt.Printf`** এবং float calculation দিয়ে fixed VAT (tax) + discount হিসাব শেখার ছোট example।

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

### Lines 5–7

```go
const vatRate = 0.15
const discount = 0.10
const currency = "BDT"
```

তিনটা **constant** define করা হয়:

- `vatRate = 0.15` — VAT (কর)-এর হার 15% (0.15)।
- `discount = 0.10` — ছাড়/ডিসকাউন্ট 10% (0.10)।
- `currency = "BDT"` — টাকার currency code (বাংলাদেশি টাকা)।

Constant হওয়ায় এগুলো program-জুড়ে fixed থাকে।

### Line 9

```go
func main() {
```

Program-এর entry point।

### Line 10

```go
subtotal := 1100.0
```

`subtotal` — একটা `float64` variable (subtotal = মূল দাম), `1100.0`। `:=` দিয়ে type inference হয় float64।

### Line 12

```go
tax := subtotal * vatRate
```

**Tax (কর) হিসাব:** `subtotal * vatRate` → `1100 * 0.15 = 165`। `tax`-এ রাখা হয়।

### Line 13

```go
afterDiscount := subtotal * (1 - discount)
```

**Discount-এর পরে দাম:** `subtotal * (1 - discount)` → `1100 * (1 - 0.10) = 1100 * 0.90 = 990`। `(1 - discount)`-এর মানে হলো মূল দামের 90%।

### Line 14

```go
total := afterDiscount + tax
```

মোট: discount-এর পরে দাম + tax → `990 + 165 = 1155`।

### Line 16

```go
fmt.Printf("%s %.2f\n", currency, total)
```

**`fmt.Printf`** — format দিয়ে print করে, `fmt.Println`-এর মতো নয়। Format string-এ placeholder:

- `%s` — string; এখানে `currency` ("BDT") বসে।
- `%.2f` — float; `total`-কে **দশমিকের পরে ২ ঘর** দিয়ে বসে (1155.00)।
- `\n` — নতুন line।

ফলে output: `BDT 1155.00`।

> **`fmt.Printf` vs `fmt.Println`:** Printf দেয় formatting control (`%s`, `%d`, `%.2f`, `%v` ইত্যাদি); Println সহজে space-সহ value print করে।

---

## Expected Output

```
BDT 1155.00
```

## মূল শিক্ষা / Key Takeaways

1. **Multiple `const`** — একাধিক fixed মান একসাথে define।
2. **`fmt.Printf`** — format specifier (`%s`, `%.2f`) দিয়ে নিয়ন্ত্রিত output।
3. **`%.2f`** — float-কে দশমিকের ২ ঘর পর্যন্ত round করে।
4. **Discount math** — `amount * (1 - discount)` দিয়ে ছাড়-করা দাম।
5. **Float arithmetic** — tax = rate × amount।

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

### Lines 5–7

```go
const vatRate = 0.15
const discount = 0.10
const currency = "BDT"
```

Three **constants** are defined:

- `vatRate = 0.15` — VAT (tax) rate of 15% (0.15).
- `discount = 0.10` — a discount of 10% (0.10).
- `currency = "BDT"` — the currency code (Bangladeshi Taka).

Being constants, these stay fixed for the whole program.

### Line 9

```go
func main() {
```

Program entry point.

### Line 10

```go
subtotal := 1100.0
```

`subtotal` — a `float64` variable representing the base price, `1100.0`. With `:=`, the type is inferred as float64.

### Line 12

```go
tax := subtotal * vatRate
```

**Tax calculation:** `subtotal * vatRate` → `1100 * 0.15 = 165`. Stored in `tax`.

### Line 13

```go
afterDiscount := subtotal * (1 - discount)
```

**Price after discount:** `subtotal * (1 - discount)` → `1100 * (1 - 0.10) = 1100 * 0.90 = 990`. `(1 - discount)` means 90% of the base price.

### Line 14

```go
total := afterDiscount + tax
```

Total: price after discount + tax → `990 + 165 = 1155`.

### Line 16

```go
fmt.Printf("%s %.2f\n", currency, total)
```

**`fmt.Printf`** — prints with a format string, unlike `fmt.Println`. Placeholders:

- `%s` — string; here `currency` ("BDT").
- `%.2f` — float; `total` formatted with **2 decimal places** (1155.00).
- `\n` — newline.

Result: `BDT 1155.00`.

> **`fmt.Printf` vs `fmt.Println`:** Printf gives formatting control (`%s`, `%d`, `%.2f`, `%v`, etc.); Println simply prints values separated by spaces.

---

## Expected Output

```
BDT 1155.00
```

## Key Takeaways

1. **Multiple `const`** — define several fixed values together.
2. **`fmt.Printf`** — controlled output using format specifiers (`%s`, `%.2f`).
3. **`%.2f`** — rounds a float to 2 decimal places.
4. **Discount math** — `amount * (1 - discount)` computes the discounted price.
5. **Float arithmetic** — tax = rate × amount.
