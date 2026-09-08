# cart-discount

Go-তে **slice + `for range` loop**, **tag-ছাড়া `switch`** দিয়ে tier discount, আর **`fmt.Printf` formatting** শেখার ছোট example.

**ভাষা নির্বাচন করুন / Choose language:**

[বাংলা](#bangla) * [English](#english)

---

<a name="bangla"></a>

## বাংলা সংস্করণ

### Line 1

```go
package main
```

একটা executable program (`main` package) declare করে, যা `go run` দিয়ে চালানো যায়.

### Line 3

```go
import "fmt"
```

Console-এ output print করার জন্য `fmt` package import করা হয়.

### Line 5

```go
func main() {
```

Program-এর entry point.

### Line 6

```go
cart := []float64{1290, 349, 89, 2450, 599}
```

`cart` -- দামের একটা **slice** (`[]float64`). 5টা item-এর দাম টাকায় (BDT). `:=` দিয়ে declare + initialize এক ধাপে.

### Lines 8-12

```go
subtotal := 0.0
for i, price := range cart {
	subtotal += price
	fmt.Printf("item %d: %8.2f BDT\n", i, price)
}
```

- `subtotal := 0.0` -- মোট যোগফল রাখার accumulator, শুরুতে 0.
- `for i, price := range cart` -- slice-এর প্রতিটা element-এর উপর loop; `i` হলো index (0 থেকে), `price` হলো value.
- `subtotal += price` -- প্রতিটা দাম যোগ করে.
- `fmt.Printf("item %d: %8.2f BDT\n", i, price)` -- `%d` দিয়ে index, `%8.2f` দিয়ে দাম print করে: `8` মানে মোট width 8 (ডান-দিকে align, space দিয়ে pad), `.2f` মানে দশমিকের পরে 2 ঘর (টাকার জন্য).

### Line 14

```go
discountRate := 0.0
```

`discountRate` -- ছাড়ের হার (0.10 = 10%), শুরুতে 0 (কোনো ছাড় নেই ধরে নেওয়া).

### Lines 16-21

```go
switch {
case subtotal >= 5000:
	discountRate = 0.10
case subtotal >= 2000:
	discountRate = 0.05
}
```

**Tag-ছাড়া `switch`** -- প্রতিটা `case`-এ একটা boolean condition থাকে, উপর থেকে নিচে প্রথম true-টা match করে:

- `subtotal >= 5000` হলে 10% ছাড়.
- না হলে `subtotal >= 2000` হলে 5% ছাড়.
- কোনোটাই না হলে `discountRate` 0-ই থাকে (কোনো ছাড় নেই).

এখানে `subtotal` = 4777, তাই দ্বিতীয় case match করে -> 5% ছাড়. **Order গুরুত্বপূর্ণ:** বড় threshold আগে লিখতে হয়, না হলে ছোটটা আগেই match করে যেত.

### Lines 23-24

```go
discount := subtotal * discountRate
total := subtotal - discount
```

- `discount` = 4777 * 0.05 = 238.85.
- `total` (payable) = 4777 - 238.85 = 4538.15.

### Line 26

```go
fmt.Printf("\nsubtotal: %.2f, discount(%.0f%%): %.2f, payable: %.2f\n", subtotal, discountRate*100, discount, total)
```

- `\n` -- আগে একটা খালি line.
- `%.2f` -- 2 দশমিক ঘরসহ টাকা.
- `%.0f%%` -- `discountRate*100` (5) কে দশমিক ছাড়া (`%.0f`) + `%%` দিয়ে আসল `%` চিহ্ন print করে -> `5%`.

### Line 27

```go
}
```

Closing brace -- `main` function শেষ হয়.

---

## Expected Output

```
item 0:  1290.00 BDT
item 1:   349.00 BDT
item 2:    89.00 BDT
item 3:  2450.00 BDT
item 4:   599.00 BDT

subtotal: 4777.00, discount(5%): 238.85, payable: 4538.15
```

## মূল শিক্ষা / Key Takeaways

1. **Slice + `for range`** -- `i, price` দিয়ে index ও value নিয়ে loop.
2. **Tag-ছাড়া `switch`** -- `switch { case cond: ... }` দিয়ে tier/threshold logic.
3. **Case order** -- বড় threshold আগে, না হলে ভুল tier match করবে.
4. **`%8.2f` formatting** -- width + 2 দশমিক ঘর; `%%` দিয়ে `%` চিহ্ন print.
5. **Accumulator pattern** -- `subtotal := 0.0` + `+=` দিয়ে যোগফল.

---

---

<a name="english"></a>

##  English Version

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

### Line 6

```go
cart := []float64{1290, 349, 89, 2450, 599}
```

`cart` -- a **slice** (`[]float64`) of item prices in BDT. Declared and initialized in one step with `:=`.

### Lines 8-12

```go
subtotal := 0.0
for i, price := range cart {
	subtotal += price
	fmt.Printf("item %d: %8.2f BDT\n", i, price)
}
```

- `subtotal := 0.0` -- an accumulator for the running total, starting at 0.
- `for i, price := range cart` -- loops over each element; `i` is the index (from 0), `price` is the value.
- `subtotal += price` -- adds each price.
- `fmt.Printf("item %d: %8.2f BDT\n", i, price)` -- `%d` for the index, `%8.2f` for the price: `8` is the total width (right-aligned, space-padded), `.2f` is 2 decimal places (for money).

### Line 14

```go
discountRate := 0.0
```

`discountRate` -- the discount fraction (0.10 = 10%), starting at 0 (assume no discount).

### Lines 16-21

```go
switch {
case subtotal >= 5000:
	discountRate = 0.10
case subtotal >= 2000:
	discountRate = 0.05
}
```

A **`switch` with no tag** -- each `case` is a boolean condition, the first true one from top to bottom wins:

- `subtotal >= 5000` -> 10% off.
- Else `subtotal >= 2000` -> 5% off.
- Else `discountRate` stays 0 (no discount).

Here `subtotal` = 4777, so the second case matches -> 5% off. **Order matters:** the larger threshold must come first, otherwise the smaller one would match early.

### Lines 23-24

```go
discount := subtotal * discountRate
total := subtotal - discount
```

- `discount` = 4777 * 0.05 = 238.85.
- `total` (payable) = 4777 - 238.85 = 4538.15.

### Line 26

```go
fmt.Printf("\nsubtotal: %.2f, discount(%.0f%%): %.2f, payable: %.2f\n", subtotal, discountRate*100, discount, total)
```

- `\n` -- a blank line first.
- `%.2f` -- money with 2 decimals.
- `%.0f%%` -- `discountRate*100` (5) with no decimals (`%.0f`) + `%%` prints a literal `%` sign -> `5%`.

### Line 27

```go
}
```

Closing brace -- ends the `main` function.

---

## Expected Output

```
item 0:  1290.00 BDT
item 1:   349.00 BDT
item 2:    89.00 BDT
item 3:  2450.00 BDT
item 4:   599.00 BDT

subtotal: 4777.00, discount(5%): 238.85, payable: 4538.15
```

## Key Takeaways

1. **Slice + `for range`** -- loop with `i, price` for index and value.
2. **Tagless `switch`** -- `switch { case cond: ... }` for tier/threshold logic.
3. **Case order** -- larger threshold first, otherwise the wrong tier matches.
4. **`%8.2f` formatting** -- width + 2 decimals; `%%` prints a `%` sign.
5. **Accumulator pattern** -- `subtotal := 0.0` + `+=` to sum up.
