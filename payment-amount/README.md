# payment-amount

Go-তে **custom type (`type PaymentAmount int64`)**, **type conversion**, আর **paisa-advanced format (`fmt.Printf`)** দিয়ে টাকার হিসাব নির্ভুল রাখা শেখার ছোট example।

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
type PaymentAmount int64
```

একটা **custom (named) type** `PaymentAmount` define করা হয় যার underlying type হলো `int64`। টাকার পরিমাণকে **paisa** (ন্যূনতম একক) হিসেবে `int64`-তে রাখা হয়।

> **কেন paisa?** Float (`float64`) দিয়ে টাকা হিসাব করলে rounding error হতে পারে (যেমন `0.1 + 0.2` আবার `0.30000000000000004` হয়)। পূর্ণসংখ্যা **cents/paisa**-তে রাখলে হিসাব নির্ভুল থাকে। `PaymentAmount` type-টা ডোমেইন নিয়ন্ত্রণ দেয় — রাস্তায় খালি `int64` না, "এটা একটা টাকার পরিমাণ"।

### Line 7

```go
func main() {
```

Program-এর entry point।

### Lines 8–9

```go
// 1 BDT = 100 paisa
// 149950 / 100 = 1499.50 BDT
```

দুটো comment — conversion-র সূত্র বোঝায়: `1 BDT = 100 paisa`। তাই `149950` paisa = `149950 / 100` = `1499.50` BDT।

### Line 10

```go
orderTotal := PaymentAmount(149950)
```

`orderTotal` declare করা হয় `PaymentAmount(149950)` দিয়ে — একটা **type conversion**: plain integer literal `149950`-কে `PaymentAmount`-এ explicit convert করা হয়। Value = 149950 paisa = 1499.50 BDT। `:=` দিয়ে type inference-এ `orderTotal`-এর type হয় `PaymentAmount`।

### Line 12

```go
var refundAmount PaymentAmount
```

`var refundAmount PaymentAmount` — নতুন variable declare করা হয় সেটা না করে। `var` দিয়ে declare করলে type-এর **zero value** নেয় — `int64`-এর জন্য `0`। অর্থাৎ refund এখন `0` paisa = 0.00 BDT।

### Line 13

```go
fmt.Println("Refund before processing", refundAmount)
```

Refund process শুরু করার আগে (zero) value print করে: `Refund before processing 0`। এটা প্রমাণ করে `refundAmount` এখনও `0`।

### Line 15

```go
refundAmount = 20000 // 200.00 BDT
```

`refundAmount`-এ `20000` assign করা হয় — আবার `// 200.00 BDT` comment-এ বোঝানো হয় 20000 paisa = 200.00 BDT। এইবার `=` (অ্যাসাইনমেন্ট) কারণ variable আগেই declare হয়ে গেছে। `20000` unttyped constant, তাই এটা free-এ `PaymentAmount`-এ convert হয়।

### Line 16

```go
remaining := orderTotal - refundAmount
```

`orderTotal - refundAmount` করে **অবশিষ্ট paisa** বের করা হয়: `149950 - 20000 = 129950`। দুটোই `PaymentAmount`, তাই subtraction-ও `PaymentAmount`। `remaining`-এর মান = 129950 paisa = 1299.50 BDT।

### Lines 18–20

```go
fmt.Printf("Order total : %d paisa (%.2f BDT)\n", orderTotal, float64(orderTotal)/100)
fmt.Printf("Refunded    : %d paisa (%.2f BDT)\n", refundAmount, float64(refundAmount)/100)
fmt.Printf("Remaining   : %d paisa (%.2f BDT)\n", remaining, float64(remaining)/100)
```

`fmt.Printf` — **formatted printing** (আগের `Println`-এর চেয়ে বেশি নিয়ন্ত্রণ):

- `%d` — decimal integer (paisa value)। `orderTotal`/`refundAmount`/`remaining`-কে `%d` দিয়ে print করা হয়, অর্থাৎ তাদের numeric value দেখায়।
- `%f` — float; `%.2f` মানে দশমিকে ২ ঘর। BDT দেখার জন্য paisa-কে `float64(...)/100` দিয়ে convert করা হয় (149950/100 = 1499.50)।
- `\n` — newline।
- Output format:
    - `Order total : 149950 paisa (1499.50 BDT)`
    - `Refunded    : 20000 paisa (200.00 BDT)`
    - `Remaining   : 129950 paisa (1299.50 BDT)`

(দৃষ্টি দাও: `Order total`, `Refunded`, `Remaining` এ তিনটা column-জাতীয় layout, যাতে একটার নিচে আরেকটা পরিমাণ চোখে ধরা পড়ে।)

### Line 21

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
Refund before processing 0
Order total : 149950 paisa (1499.50 BDT)
Refunded    : 20000 paisa (200.00 BDT)
Remaining   : 129950 paisa (1299.50 BDT)
```

## মূল শিক্ষা / Key Takeaways

1. **Integer paisa** — টাকাকে float-এর বদলে paisa (int) হিসেবে রাখলে rounding error হয় না।
2. **Custom type** — `type PaymentAmount int64` টাকার পরিমাণ-কে type-সেফ করে।
3. **Type conversion** — `PaymentAmount(149950)` explicit convert; `float64(x)/100` দিয়ে paisa→BDT।
4. **Zero value** — `var refundAmount PaymentAmount` দিয়ে declare করলে `0` দিয়ে শুরু।
5. **`fmt.Printf`** — `%d`, `%.2f`, `\n` দিয়ে নিয়ন্ত্রিত (formatted) output।
6. **Untyped constants** — `refundAmount = 20000` এ constant automatic `PaymentAmount`-এ convert হয়।

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
type PaymentAmount int64
```

Defines a **custom (named) type** `PaymentAmount` whose underlying type is `int64`. Money is stored as **paisa** (the smallest unit) in an `int64`.

> **Why paisa?** Doing money math with floats (`float64`) risks rounding errors (e.g. `0.1 + 0.2` gives `0.30000000000000004`). Keeping amounts as whole integers in cents/paisa stays exact. The `PaymentAmount` type adds domain safety — not a bare `int64`, but "this is a money amount".

### Line 7

```go
func main() {
```

Program entry point.

### Lines 8–9

```go
// 1 BDT = 100 paisa
// 149950 / 100 = 1499.50 BDT
```

Two comments explaining the conversion: `1 BDT = 100 paisa`, so `149950` paisa = `149950 / 100` = `1499.50` BDT.

### Line 10

```go
orderTotal := PaymentAmount(149950)
```

Declares `orderTotal` with `PaymentAmount(149950)` — a **type conversion**: explicitly converting the plain integer literal `149950` into `PaymentAmount`. The value is 149950 paisa = 1499.50 BDT. With `:=` type inference, `orderTotal` is `PaymentAmount`.

### Line 12

```go
var refundAmount PaymentAmount
```

`var refundAmount PaymentAmount` declares a new variable without initializing it. Declared with `var`, it takes the type's **zero value** — `0` for `int64`. So the refund is currently `0` paisa = 0.00 BDT.

### Line 13

```go
fmt.Println("Refund before processing", refundAmount)
```

Prints the value before any refund processing (`0`): `Refund before processing 0`. It shows `refundAmount` is still `0`.

### Line 15

```go
refundAmount = 20000 // 200.00 BDT
```

Assigns `20000` to `refundAmount` — the comment `// 200.00 BDT` explains that 20000 paisa = 200.00 BDT. This uses `=` (assignment) because the variable was already declared. `20000` is an untyped constant, so it converts freely to `PaymentAmount`.

### Line 16

```go
remaining := orderTotal - refundAmount
```

Computes the **remaining paisa** with `orderTotal - refundAmount`: `149950 - 20000 = 129950`. Both are `PaymentAmount`, so the subtraction yields `PaymentAmount`. `remaining` = 129950 paisa = 1299.50 BDT.

### Lines 18–20

```go
fmt.Printf("Order total : %d paisa (%.2f BDT)\n", orderTotal, float64(orderTotal)/100)
fmt.Printf("Refunded    : %d paisa (%.2f BDT)\n", refundAmount, float64(refundAmount)/100)
fmt.Printf("Remaining   : %d paisa (%.2f BDT)\n", remaining, float64(remaining)/100)
```

`fmt.Printf` — **formatted printing** (more control than `Println`):

- `%d` — a decimal integer (the paisa value). `orderTotal`, `refundAmount`, and `remaining` are printed with `%d`, showing their numeric values.
- `%f` — a float; `%.2f` means exactly 2 decimal places. To show BDT, paisa is converted with `float64(...)/100` (149950/100 = 1499.50).
- `\n` — a newline.
- Output format:
    - `Order total : 149950 paisa (1499.50 BDT)`
    - `Refunded    : 20000 paisa (200.00 BDT)`
    - `Remaining   : 129950 paisa (1299.50 BDT)`

(Notice the aligned layout of `Order total`, `Refunded`, and `Remaining` — amounts line up visually for easy comparison.)

### Line 21

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
Refund before processing 0
Order total : 149950 paisa (1499.50 BDT)
Refunded    : 20000 paisa (200.00 BDT)
Remaining   : 129950 paisa (1299.50 BDT)
```

## Key Takeaways

1. **Integer paisa** — storing money as paisa (int) instead of float avoids rounding errors.
2. **Custom type** — `type PaymentAmount int64` makes money amounts type-safe.
3. **Type conversion** — `PaymentAmount(149950)` explicit convert; `float64(x)/100` for paisa→BDT.
4. **Zero value** — declaring with `var refundAmount PaymentAmount` starts at `0`.
5. **`fmt.Printf`** — controlled (formatted) output with `%d`, `%.2f`, `\n`.
6. **Untyped constants** — with `refundAmount = 20000`, the constant auto-converts to `PaymentAmount`.
