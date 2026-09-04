# emi-loan

Go-তে **numeric underscore**, **`float64` calculation**, **`for` loop** আর **aligned `fmt.Printf` formatting** দিয়ে EMI loan (কিস্তি ঋণ) হিসাব শেখার ছোট example।

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

### Lines 6–9

```go
principal := 120_000.0
flatAnnualRate := 0.10
years := 1
months := 12
```

Input values:

- `principal := 120_000.0` — ঋণের মূল পরিমাণ (principal)। **`_` numeric literal** — `120_000` আর `120000` একই; underscore শুধু readability-র জন্য (হাজারের grouping)।
- `flatAnnualRate := 0.10` — ফ্ল্যাট বার্ষিক সুদের হার (10%)।
- `years := 1` — ঋণের মেয়াদ (বছর)।
- `months := 12` — প্রতি বছরে মাস-সংখ্যা 12 (আর মোট installment-ও 12)।

### Line 11

```go
totalInterest := principal * flatAnnualRate * float64(years)
```

**মোট সুদ (flat rate):** `principal × rate × years` → `120000 × 0.10 × 1 = 12000`। লক্ষ্য করো `float64(years)` — `years`-টা int, তাই float calculation-এর আগে **type conversion** (`float64`) করা হয়েছে।

### Line 12

```go
emi := (principal + totalInterest) / float64(months)
```

**EMI (মোট কিস্তি):** (প্রিন্সিপাল + মোট সুদ) ÷ মাস → `(120000 + 12000) / 12 = 11000`।

### Line 13

```go
interestPart := totalInterest / float64(months)
```

প্রতি মাসের interest অংশ: `12000 / 12 = 1000`।

### Line 14

```go
principalPart := emi - interestPart
```

প্রতি মাসের principal অংশ (EMI থেকে interest বাদ): `11000 - 1000 = 10000`।

### Line 16

```go
fmt.Printf("EMI: %.2f BDT/month\n\n", emi)
```

EMI print করে — `%.2f` (দশমিক ২ ঘর) + `\n\n` (একটা ফাঁকা line)।

### Line 17

```go
fmt.Printf("%-5s %10s %12s %12s\n", "month", "principal", "interest", "remaining")
```

Table-এর **header** — `%-5s` (left-align ৫ ঘর), `%10s`/`%12s` (right-align 10/12 ঘর)। Columns সোজা রাখতে width ব্যবহার করা হয়।

### Line 19

```go
remaining := principal
```

`remaining` শুরুতে principal (`120000`) — কত বাকি আছে ট্র্যাক করে।

### Lines 21–24

```go
for m := 1; m <= months; m++ {
    remaining -= principalPart
    fmt.Printf("%-5d %10.2f %12.2f %12.2f\n", m, principalPart, interestPart, remaining)
}
```

প্রতিটা মাসের জন্য loop (`m` 1 থেকে 12):

- `remaining -= principalPart` — প্রতি মাসে principal অংশটা omitted করে (compound assignment)।
- `fmt.Printf("%-5d %10.2f ...\n", m, ...)` — `%-5d` (left-align int), `%10.2f`/`%12.2f` (right-align float, ২ দশমিক) — aligned amortization table print করে।

প্রতি মাসে 10000 কমানো হয়, ১২ মাস শেষে `remaining` 0 হয়।

---

## Expected Output

```
EMI: 11000.00 BDT/month

month  principal     interest    remaining
1       10000.00      1000.00    110000.00
2       10000.00      1000.00    100000.00
3       10000.00      1000.00     90000.00
4       10000.00      1000.00     80000.00
5       10000.00      1000.00     70000.00
6       10000.00      1000.00     60000.00
7       10000.00      1000.00     50000.00
8       10000.00      1000.00     40000.00
9       10000.00      1000.00     30000.00
10      10000.00      1000.00     20000.00
11      10000.00      1000.00     10000.00
12      10000.00      1000.00         0.00
```

## মূল শিক্ষা / Key Takeaways

1. **Numeric underscore** — `120_000` readability-র জন্য; `120000`-এর সমান।
2. **`float64` conversion** — int-কে float calculation-এ মেশাতে `float64(x)`।
3. **EMI formula** — `(principal + interest) / months`।
4. **Formatting width** — `%-5d`, `%10.2f` দিয়ে table columns align।
5. **`for` loop + compound assignment** — `remaining -= principalPart` দিয়ে amortization।

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

### Lines 6–9

```go
principal := 120_000.0
flatAnnualRate := 0.10
years := 1
months := 12
```

Input values:

- `principal := 120_000.0` — the loan's principal amount. The **`_` numeric literal** — `120_000` and `120000` are equal; the underscore is just for readability (thousands grouping).
- `flatAnnualRate := 0.10` — the flat annual interest rate (10%).
- `years := 1` — the loan term in years.
- `months := 12` — the number of months per year (and total installments).

### Line 11

```go
totalInterest := principal * flatAnnualRate * float64(years)
```

**Total interest (flat rate):** `principal × rate × years` → `120000 × 0.10 × 1 = 12000`. Note `float64(years)` — `years` is an int, so it's **type-converted** to `float64` before the float calculation.

### Line 12

```go
emi := (principal + totalInterest) / float64(months)
```

**EMI (installment):** (principal + total interest) ÷ months → `(120000 + 12000) / 12 = 11000`.

### Line 13

```go
interestPart := totalInterest / float64(months)
```

The monthly interest part: `12000 / 12 = 1000`.

### Line 14

```go
principalPart := emi - interestPart
```

The monthly principal part (EMI minus interest): `11000 - 1000 = 10000`.

### Line 16

```go
fmt.Printf("EMI: %.2f BDT/month\n\n", emi)
```

Prints the EMI — `%.2f` (2 decimal places) + `\n\n` (a blank line).

### Line 17

```go
fmt.Printf("%-5s %10s %12s %12s\n", "month", "principal", "interest", "remaining")
```

The table **header** — `%-5s` (left-aligned, 5 wide), `%10s`/`%12s` (right-aligned, 10/12 wide). Widths keep the columns straight.

### Line 19

```go
remaining := principal
```

`remaining` starts at the principal (`120000`) — tracks how much is left.

### Lines 21–24

```go
for m := 1; m <= months; m++ {
    remaining -= principalPart
    fmt.Printf("%-5d %10.2f %12.2f %12.2f\n", m, principalPart, interestPart, remaining)
}
```

Loops for each month (`m` from 1 to 12):

- `remaining -= principalPart` — each month subtracts the principal part (compound assignment).
- `fmt.Printf("%-5d %10.2f ...\n", m, ...)` — `%-5d` (left-aligned int), `%10.2f`/`%12.2f` (right-aligned float, 2 decimals) — prints an aligned amortization table.

10000 is subtracted each month; after 12 months `remaining` hits 0.

---

## Expected Output

```
EMI: 11000.00 BDT/month

month  principal     interest    remaining
1       10000.00      1000.00    110000.00
2       10000.00      1000.00    100000.00
3       10000.00      1000.00     90000.00
4       10000.00      1000.00     80000.00
5       10000.00      1000.00     70000.00
6       10000.00      1000.00     60000.00
7       10000.00      1000.00     50000.00
8       10000.00      1000.00     40000.00
9       10000.00      1000.00     30000.00
10      10000.00      1000.00     20000.00
11      10000.00      1000.00     10000.00
12      10000.00      1000.00         0.00
```

## Key Takeaways

1. **Numeric underscore** — `120_000` is equal to `120000`; just for readability.
2. **`float64` conversion** — `float64(x)` to mix an int into float math.
3. **EMI formula** — `(principal + interest) / months`.
4. **Formatting width** — `%-5d`, `%10.2f` align table columns.
5. **`for` loop + compound assignment** — `remaining -= principalPart` for amortization.
