# payment-fee-calculator

Go-তে **function**, **`const`**, **`float64`** আর `var` declaration দিয়ে সহজ fee (ফি) হিসাব শেখার ছোট example।

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
const feePercentage = 2.5
```

একটা **constant** `feePercentage` যার মান `2.5`। এটা হলো transaction-এর উপর নেওয়া **ফি-র শতকরা হার** (2.5%)। Constant হওয়ায় মানটা program-জুড়ে অপরিবর্তনীয় রয়ে যায়।

### Lines 7–9

```go
func calculateFee(amount float64) float64 {
    return amount * feePercentage / 100
}
```

একটা function `calculateFee` যা একটা `float64` (`amount`) নেয় এবং একটা `float64` return করে।

- `amount * feePercentage / 100` — ফি বের করার সূত্র। যেমন `4500 * 2.5 / 100 = 112.5`।
- `feePercentage`-কে `100` দিয়ে ভাগ করা হয় কারণ এটা শতকরা (percentage) — 2.5% মানে 2.5/100।

### Line 11

```go
func main() {
```

Program-এর entry point।

### Line 12

```go
var transactionAmount float64 = 4500.0
```

আগের folder-গুলোর `:=` shortcut-এর বদলে এখানে `var` keyword দিয়ে explicitly declare করা হয় — `transactionAmount` নামে, type `float64`, মান `4500.0`। (দুটো পদ্ধতিই ঠিক; `var` আরও স্পষ্ট।) এটা transaction-এর মূল পরিমাণ।

### Line 13

```go
fee := calculateFee(transactionAmount)
```

`calculateFee` function call করে `transactionAmount`-এর উপর ফি বের করা হয় এবং `:=` দিয়ে `fee`-তে রাখা হয়। type inference-এ `fee`-এর type হয় `float64`।

### Line 14

```go
total := transactionAmount + fee
```

মূল পরিমাণের সাথে ফি যোগ করে `total` (মোট দেয়) বের করা হয়।

### Lines 15–18

```go
fmt.Println("Amount:", transactionAmount)
fmt.Println("Fee:", fee)
fmt.Println("Total:", total)
```

তিনটা মান print করে:

- `Amount: 4500`
- `Fee: 112.5`
- `Total: 4612.5`

---

## Expected Output

```
Amount: 4500
Fee: 112.5
Total: 4612.5
```

## মূল শিক্ষা / Key Takeaways

1. **`const`** — fixed, অপরিবর্তনীয় মান (যেমন ফি-র হার)।
2. **Function with return** — `func calculateFee(amount float64) float64` — input নেয়, output return করে।
3. **`float64`** — দশমিক (decimal) সংখ্যা type।
4. **`var` declaration** — `var x float64 = ...` দিয়ে type সহ explicit declare।
5. **Float arithmetic** — percentage-দের `100` দিয়ে ভাগ করলে মান সঠিক হয়।

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
const feePercentage = 2.5
```

A **constant** `feePercentage` with the value `2.5`. This is the **fee percentage rate** (2.5%) charged on a transaction. Since it's a constant, the value stays immutable across the program.

### Lines 7–9

```go
func calculateFee(amount float64) float64 {
    return amount * feePercentage / 100
}
```

A function `calculateFee` that takes a single `float64` (`amount`) and returns a `float64`.

- `amount * feePercentage / 100` — the formula to compute the fee. E.g. `4500 * 2.5 / 100 = 112.5`.
- We divide `feePercentage` by `100` because it's a percentage — 2.5% means 2.5/100.

### Line 11

```go
func main() {
```

Program entry point.

### Line 12

```go
var transactionAmount float64 = 4500.0
```

Instead of the `:=` shortcut used in earlier folders, this uses the `var` keyword to declare explicitly — `transactionAmount`, type `float64`, value `4500.0`. (Both approaches are valid; `var` is more explicit.) This is the base transaction amount.

### Line 13

```go
fee := calculateFee(transactionAmount)
```

Calls `calculateFee` to compute the fee on `transactionAmount` and stores it in `fee` with `:=`. By type inference, `fee` is `float64`.

### Line 14

```go
total := transactionAmount + fee
```

Adds the fee to the base amount to get `total` (the total payable).

### Lines 15–18

```go
fmt.Println("Amount:", transactionAmount)
fmt.Println("Fee:", fee)
fmt.Println("Total:", total)
```

Prints the three values:

- `Amount: 4500`
- `Fee: 112.5`
- `Total: 4612.5`

---

## Expected Output

```
Amount: 4500
Fee: 112.5
Total: 4612.5
```

## Key Takeaways

1. **`const`** — a fixed, immutable value (e.g. a fee rate).
2. **Function with return** — `func calculateFee(amount float64) float64` — takes input, returns output.
3. **`float64`** — the decimal (floating-point) number type.
4. **`var` declaration** — `var x float64 = ...` declares with an explicit type.
5. **Float arithmetic** — dividing a percentage by `100` gives the correct value.
