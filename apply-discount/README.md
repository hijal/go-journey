# apply-discount

Go-তে **pass-by-value** behavior শেখার ছোট example -- function parameter-এ local copy তৈরি হয়, caller-এর variable পরিবর্তন হয় না.

**ভাষা নির্বাচন করুন / Choose language:**

[বাংলা](#bangla) * [English](#english)

---

<a name="bangla"></a>

## বাংলা সংস্করণ

### Lines 1-3

```go
package main

import "fmt"
```

- `package main` -- একটা executable program.
- `fmt` -- output print করার জন্য.

### Lines 5-9

```go
func applyDiscount(price, percent float64) float64 {
    price = price - (price * percent / 100)
    fmt.Printf(" inside function: %.2f\n", price)
    return price
}
```

`applyDiscount` function দুটো `float64` নেয়:

- `price, percent` -- original price আর discount percent.
- `price = price - (price * percent / 100)` -- discounted price compute করে, **local `price` variable-এ assign**.
- `return price` -- final discounted value ফেরত দেয়.

### Lines 11-18

```go
func main() {
    listPrice := 1000.0
    fmt.Printf("before call: %.2f\n", listPrice)

    final := applyDiscount(listPrice, 15)
    fmt.Printf("after call:  %.2f\n", listPrice)
    fmt.Printf("final price: %.2f\n", final)
}
```

- `listPrice := 1000.0` -- main-এ original value.
- `applyDiscount(listPrice, 15)` -- listPrice-র **একটা copy** যায় function-এ.
- function-এর ভিতরে local `price` change হয় (1000 -> 850), কিন্তু caller-এর `listPrice` অপরিবর্তিত থাকে (1000).
- `final` ধরে 850.

---

## Expected Output

```
before call: 1000.00
 inside function: 850.00
after call:  1000.00
final price: 850.00
```

## মূল শিক্ষা / Key Takeaways

1. **Pass-by-value** -- Go-তে সব value-typed arguments copy হয়ে যায় function-এ.
2. **Local scope** -- function parameter reassign করলে caller-এর variable প্রভাবিত হয় না.
3. **Return value** -- modified value ফেরত পেতে হলে explicit return করতে হবে.
4. **Float formatting** -- `%.2f` দুই দশমিক ঘরে output দেয়.

---

---

<a name="english"></a>

## English Version

### Lines 1-3

```go
package main

import "fmt"
```

- `package main` -- an executable program.
- `fmt` -- for console output.

### Lines 5-9

```go
func applyDiscount(price, percent float64) float64 {
    price = price - (price * percent / 100)
    fmt.Printf(" inside function: %.2f\n", price)
    return price
}
```

`applyDiscount` takes two `float64` arguments:

- `price, percent` -- the original price and the discount percent.
- `price = price - (price * percent / 100)` -- computes the discounted price and **assigns it to the local `price` variable**.
- `return price` -- returns the final discounted value.

### Lines 11-18

```go
func main() {
    listPrice := 1000.0
    fmt.Printf("before call: %.2f\n", listPrice)

    final := applyDiscount(listPrice, 15)
    fmt.Printf("after call:  %.2f\n", listPrice)
    fmt.Printf("final price: %.2f\n", final)
}
```

- `listPrice := 1000.0` -- the original value in main.
- `applyDiscount(listPrice, 15)` -- a **copy of listPrice** is passed in.
- the function changes its local `price` (1000 -> 850), but the caller's `listPrice` is unchanged (1000).
- `final` captures 850.

---

## Expected Output

```
before call: 1000.00
 inside function: 850.00
after call:  1000.00
final price: 850.00
```

## Key Takeaways

1. **Pass-by-value** -- all value-typed arguments in Go are copied into the function.
2. **Local scope** -- reassigning a function parameter does not affect the caller's variable.
3. **Return value** -- to surface a modified value, the function must return it explicitly.
4. **Float formatting** -- `%.2f` formats a float to 2 decimal places.
