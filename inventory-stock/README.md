# inventory-stock

Go-তে **unsigned integer (`uint32`) underflow** আর **signed `int` দিয়ে safe comparison** বোঝার ছোট example — uint-এ negative result গেলে কী happens, আর কীভাবে তার থেকে বাঁচা যায়।

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

### Line 6

```go
var stock uint32 = 5
```

`stock` নামে একটা **unsigned integer** variable — type `uint32`, মান `5`। Unsigned মানে এটা **negative হতে পারে না** (0 থেকে 4294967295 অবধি)। Inventory-র stock যেহেতু negative হওয়ার কথা না, তাই `uint32` বাছাই করা হয়েছে — কিন্তু এই সিদ্ধান্তটাই নিচে bug তৈরি করে।

### Line 8

```go
sold := 8
```

`sold` — বিক্রি হওয়া পরিমাণ, `:=` দিয়ে declare; value `8` plain int। খেয়াল করো: এটা **signed** (`int`), unsigned `uint32`-র মতো নয়। `stock` (5) থেকে বেশি।

### Line 9

```go
newStock := stock - uint32(sold)
```

ভুলটা এখানে: `stock - uint32(sold)` = `5 - 8`। `stock` `uint32` হওয়ায় result-ও `uint32`-তে রাখা হয়, কিন্তু `5 - 8 = -3` — negative value unsigned type-এ fit হয় না। তাই **underflow** ঘটে: value ঘুরে গিয়ে giant number-এ পরিণত হয় (`2³² − 3` = 4294967293)। কোনো error হয় না, কোনো warning-ও হয় না — নীরবে ভুল result।

### Line 10

```go
fmt.Println("Stock after (buggy) sale:", newStock)
```

Buggy result print করে: `Stock after (buggy) sale: 4294967293`। স্পষ্টতই ভুল — stock বাড়ল স্টক কমলে না!

### Line 12

```go
safeStock := int(stock) - sold
```

সঠিক উপায়: `int(stock)` দিয়ে দুই পাশকে **signed** `int`-এ নিয়ে আসা হয়, তারপর subtraction। `5 - 8 = -3` — এবার `int` negative ধরতে পারে, তাই `safeStock` = `-3`।

### Lines 14–18

```go
if safeStock < 0 {
    fmt.Println("Rejected: cannot sell more than available stock")
} else {
    fmt.Println("New stock:", safeStock)
}
```

`safeStock < 0` check — negative হলে sale reject করা হয়: `Rejected: cannot sell more than available stock` print হয়। (যথেষ্ট stock থাকলে `else` branch-এ নতুন stock print হতো।)

> **মূল শিক্ষা:** unsigned type-এ numeric operation করার আগে ভাবতে হবে result negative হতে পারে কিনা। নাহলে silent underflow-এ bug ধরা পড়বে না। Negative সম্ভাবনা থাকলে signed type ব্যবহার করো, অথবা compare-এর আগে explicit-type-conversion করো।

### Line 19

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
Stock after (buggy) sale: 4294967293
Rejected: cannot sell more than available stock
```

## মূল শিক্ষা / Key Takeaways

1. **Unsigned underflow** — `uint`-এ negative result গেলে value ঘুরে giant number হয়, error ছাড়াই।
2. **Silent bug** — `uint32` subtraction compile-time-এ ধরা পড়ে না; human-এ বোঝা যাচ্ছে না।
3. **`int(stock)` conversion** — unsigned থেকে signed-এ এনে negative সম্ভাবনা ধরার capacity।
4. **Check before acting** — `safeStock < 0` guard দিয়ে invalid operation reject।
5. **Type choice matters** — inventory-র মতো counter-এও signed নিতে হয় যদি subtraction/negative থাকতে পারে।

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

### Line 6

```go
var stock uint32 = 5
```

An **unsigned integer** variable — type `uint32`, value `5`. Unsigned means it **cannot be negative** (0 to 4294967295). Stock in inventory shouldn't be negative, so `uint32` was chosen — but that very decision creates the bug below.

### Line 8

```go
sold := 8
```

`sold` — the amount sold, declared with `:=`; value `8` as a plain int. Note: it's **signed** (`int`), unlike the unsigned `uint32`. It's more than `stock` (5).

### Line 9

```go
newStock := stock - uint32(sold)
```

The bug: `stock - uint32(sold)` = `5 - 8`. Since `stock` is `uint32`, the result is also stored as `uint32`, but `5 - 8 = -3` — a negative value that doesn't fit an unsigned type. So an **underflow** occurs: the value wraps around to a giant number (`2³² − 3` = 4294967293). No error, no warning — a silent wrong result.

### Line 10

```go
fmt.Println("Stock after (buggy) sale:", newStock)
```

Prints the buggy result: `Stock after (buggy) sale: 4294967293`. Clearly wrong — the stock went up instead of down!

### Line 12

```go
safeStock := int(stock) - sold
```

The correct approach: bring both sides to **signed** `int` with `int(stock)`, then subtract. `5 - 8 = -3` — `int` can hold negatives, so `safeStock` = `-3`.

### Lines 14–18

```go
if safeStock < 0 {
    fmt.Println("Rejected: cannot sell more than available stock")
} else {
    fmt.Println("New stock:", safeStock)
}
```

Checks `safeStock < 0` — if negative, the sale is rejected: it prints `Rejected: cannot sell more than available stock`. (With enough stock, the `else` branch would print the new stock.)

> **Key lesson:** before doing arithmetic on an unsigned type, ask whether the result could be negative. Otherwise silent underflow hides the bug. If negativity is possible, use a signed type, or convert explicitly before comparing.

### Line 19

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
Stock after (buggy) sale: 4294967293
Rejected: cannot sell more than available stock
```

## Key Takeaways

1. **Unsigned underflow** — a negative result in a `uint` wraps around to a giant number, silently.
2. **Silent bug** — `uint32` subtraction isn't caught at compile time; hard to spot by eye.
3. **`int(stock)` conversion** — casting unsigned to signed lets it represent negatives.
4. **Check before acting** — guard with `safeStock < 0` to reject invalid operations.
5. **Type choice matters** — even counters like inventory may need signed types if subtraction/negatives are possible.
