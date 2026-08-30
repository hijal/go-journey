# checkout-cart

Go-তে **float arithmetic**, **type conversion (`float64(int)`)** আর ঘুরে shopping-cart-এর subtotal/total হিসাব শেখার ছোট example।

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

### Lines 6–8

```go
itemPrice := 249.40
quantity := 3
shippingFee := 60.0
```

তিনটা input value:

- `itemPrice := 249.40` — প্রতি-item দাম; `249.40` float literal হওয়ায় type-টা `float64`।
- `quantity := 3` — কতগুলো item; পূর্ণসংখ্যা, তাই type `int`।
- `shippingFee := 60.0` — shipping charge; `60.0` float, type `float64`।

### Line 10

```go
subtotal := itemPrice * float64(quantity)
```

**subtotal = itemPrice × quantity**। সমস্যা: `itemPrice` `float64`, কিন্তু `quantity` `int` — Go-তে mixed-type arithmetic **compile করে না**। তাই `float64(quantity)` দিয়ে explicit **type conversion** করা হয়: `3` → `3.0`। তারপর `249.40 × 3` = `748.2` (শুধু decimal values-এ float multiply-র সামান্য rounding থাকতে পারে, কিন্তু print-এ সেটা দেখায় না)।

### Line 11

```go
total := subtotal + shippingFee
```

**total = subtotal + shippingFee**। দুটোই `float64`, তাই সরাসরি যোগ: `748.2 + 60.0` = `808.2`।

### Lines 13–15

```go
fmt.Println("Subtotal:", subtotal)
fmt.Println("Shipping fee:", shippingFee)
fmt.Println("Total payable:", total)
```

তিনটা value print করে:

- `Subtotal: 748.2`
- `Shipping fee: 60`
- `Total payable: 808.2`

### Line 16

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
Subtotal: 748.2
Shipping fee: 60
Total payable: 808.2
```

## মূল শিক্ষা / Key Takeaways

1. **Mixed-type arithmetic যায় না** — `float64` আর `int` একসাথে গুণ/যোগ করা যায় না।
2. **`float64(quantity)`** — int-কে explicit float-এ convert করে arithmetic-এ ব্যবহার।
3. **Float arithmetic** — subtotal/total হিসাব করলে decimal রাখা হয়।
4. **`:=` type inference** — literal দেখে type নির্ধারণ (int vs float64)।
5. **Real-world pattern** — cart checkout-এর মতো হিসাব-এ units match করানো জরুরি।

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

### Lines 6–8

```go
itemPrice := 249.40
quantity := 3
shippingFee := 60.0
```

Three input values:

- `itemPrice := 249.40` — price per item; a float literal, so its type is `float64`.
- `quantity := 3` — number of items; a whole number, so type `int`.
- `shippingFee := 60.0` — the shipping charge; a float, type `float64`.

### Line 10

```go
subtotal := itemPrice * float64(quantity)
```

**subtotal = itemPrice × quantity.** The catch: `itemPrice` is `float64` but `quantity` is `int` — Go **won't compile** mixed-type arithmetic. So `float64(quantity)` does an explicit **type conversion**: `3` → `3.0`. Then `249.40 × 3` = `748.2` (float multiplication of decimals can have tiny rounding, but it doesn't show in the printed value).

### Line 11

```go
total := subtotal + shippingFee
```

**total = subtotal + shippingFee.** Both are `float64`, so they add directly: `748.2 + 60.0` = `808.2`.

### Lines 13–15

```go
fmt.Println("Subtotal:", subtotal)
fmt.Println("Shipping fee:", shippingFee)
fmt.Println("Total payable:", total)
```

Prints the three values:

- `Subtotal: 748.2`
- `Shipping fee: 60`
- `Total payable: 808.2`

### Line 16

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
Subtotal: 748.2
Shipping fee: 60
Total payable: 808.2
```

## Key Takeaways

1. **No mixed-type arithmetic** — you can't multiply/add a `float64` and an `int` directly.
2. **`float64(quantity)`** — explicit conversion so the arithmetic works.
3. **Float arithmetic** — decimals are preserved in subtotal/total math.
4. **`:=` type inference** — the literal decides the type (int vs float64).
5. **Real-world pattern** — units must match in math like cart checkout.