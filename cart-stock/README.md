# cart-stock

Go-তে identifier-এর **case-sensitivity** এবং `if/else` শর্ত বুঝতে সহায়ক ছোট example।

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

Program-এর entry point; `{` দিয়ে function body শুরু হয়।

### Lines 6–7

```go
quantity := 3
Quantity := 10
```

দুটো variable declare করা হয়। `quantity` (ছোট হাতের) — কাস্টমারের cart-এ চাওয়া item সংখ্যা → `3`। `Quantity` (বড় হাতের) — warehouse-এ থাকা item সংখ্যা → `10`।

মূল বিষয়: Go-র identifier-গুলো **case-sensitive**, তাই `quantity` আর `Quantity` দুটো **সম্পূর্ণ ভিন্ন variable**। এটা deliberate — বোঝানো হচ্ছে যে বড় হাতের আর ছোট হাতের নাম আলাদা নাম হিসেবে গণ্য হয়।

### Lines 9–10

```go
fmt.Println("Cart requested quantity:", quantity)
fmt.Println("Warehouse stock quantity:", Quantity)
```

দুটো value print করে। Output:

```
Cart requested quantity: 3
Warehouse stock quantity: 10
```

### Lines 12–16

```go
if quantity <= Quantity {
    fmt.Println("Order can be fulfilled")
} else {
    fmt.Println("Not enough stock")
}
```

`if` শর্তটা পরীক্ষা করে `quantity <= Quantity` (মানে `3 <= 10`?)। যেহেতু এটা `true`, প্রথম block চলে এবং `Order can be fulfilled` print করে। যদি false হতো, তাহলে `else` block-এ গিয়ে `Not enough stock` print করত।

### Line 17

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
Cart requested quantity: 3
Warehouse stock quantity: 10
Order can be fulfilled
```

## মূল শিক্ষা / Key Takeaways

1. **Case-sensitive identifiers** — `quantity` আর `Quantity` আলাদা variable।
2. **Naming clarity** — real code-এ কেবল case দিয়ে পার্থক্য করা confusing, তাই এড়িয়ে চলা উচিত।
3. **`if/else`** — শর্ত true হলে `if` block, না হলে `else` block চলে।
4. **Uppercase vs lowercase** — বড় হাতের দিয়ে শুরু মানে exported; কিন্তু এখানে দুটোই local variable।

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

Program entry point; `{` opens the function body.

### Lines 6–7

```go
quantity := 3
Quantity := 10
```

Two variables are declared. `quantity` (lowercase) — the number of items the customer requested in their cart → `3`. `Quantity` (uppercase) — the number of items available in the warehouse → `10`.

The key point: Go identifiers are **case-sensitive**, so `quantity` and `Quantity` are two **completely different variables**. This is deliberate — showing that uppercase vs lowercase names are treated as separate.

### Lines 9–10

```go
fmt.Println("Cart requested quantity:", quantity)
fmt.Println("Warehouse stock quantity:", Quantity)
```

Prints both values. Output:

```
Cart requested quantity: 3
Warehouse stock quantity: 10
```

### Lines 12–16

```go
if quantity <= Quantity {
    fmt.Println("Order can be fulfilled")
} else {
    fmt.Println("Not enough stock")
}
```

`if` checks the condition `quantity <= Quantity` (is `3 <= 10`?). Since it's `true`, the first block runs and prints `Order can be fulfilled`. If it were false, the `else` block would print `Not enough stock`.

### Line 17

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
Cart requested quantity: 3
Warehouse stock quantity: 10
Order can be fulfilled
```

## Key Takeaways

1. **Case-sensitive identifiers** — `quantity` and `Quantity` are different identifiers.
2. **Naming clarity** — Don't rely only on case to differentiate names; it's confusing.
3. **`if/else`** — Runs the `if` block when the condition is true, otherwise the `else` block.
4. **Uppercase vs lowercase** — Uppercase-start would mean exported; but here both are local variables.
