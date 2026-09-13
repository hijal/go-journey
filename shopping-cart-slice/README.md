# shopping-cart-slice

Go-তে **slice basics** শেখার ছোট example — `[]float64` slice, `len()`, index access (first/last), `range`-এ accumulate।

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

`fmt` — `Println`, `Printf`।

### Lines 5–6

```go
func main() {
	cart := []float64{24.99, 19.99, 5.49, 12.99}
```

`[]float64{...}` — **slice literal** — ৪টা item-এর দাম। (Array-র থেকে আলাদা: size-নামি না, dynamic টাই pick — এখানে 4 element initialized।)

### Lines 7–9

```go
	fmt.Println("items in cart:", len(cart))
	fmt.Println("first item costs:", cart[0])
	fmt.Println("last item costs:", cart[len(cart)-1])
```

- `len(cart)` — element-সংখ্যা → `4`।
- `cart[0]` — **first** element (zero-based) → `24.99`।
- `cart[len(cart)-1]` — **last** element index-logic: `len - 1` → `12.99`।

### Lines 11–15

```go
	total := 0.0
	for _, item := range cart {
		total += item
	}
	fmt.Printf("total cost: %.2f\n", total)
```

- `total := 0.0` — accumulator।
- `for _, item := range cart` — value-only iteration।
- `total += item` — accumulate: `24.99+19.99+5.49+12.99 = 63.46`।
- `%.2f` — দু-দশমিক।

---

## Expected Output

```
items in cart: 4
first item costs: 24.99
last item costs: 12.99
total cost: 63.46
```

## মূল শিক্ষা / Key Takeaways

1. **Slice literal** — `[]float64{...}`।
2. **`len()`** — element-সংখ্যা।
3. **Index access** — `cart[0]`, `cart[len-1]` first/last।
4. **`range` accumulation** — value-only iteration + sum।
5. **`%.2f`** — float formatting।

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

`fmt` — for `Println`, `Printf`.

### Lines 5–6

```go
func main() {
	cart := []float64{24.99, 19.99, 5.49, 12.99}
```

`[]float64{...}` — a **slice literal** — the prices of 4 items. (Unlike an array, this type has no fixed size; here 4 elements are initialized.)

### Lines 7–9

```go
	fmt.Println("items in cart:", len(cart))
	fmt.Println("first item costs:", cart[0])
	fmt.Println("last item costs:", cart[len(cart)-1])
```

- `len(cart)` — the number of elements → `4`.
- `cart[0]` — the **first** element (zero-based) → `24.99`.
- `cart[len(cart)-1]` — the **last** element via index logic: `len - 1` → `12.99`.

### Lines 11–15

```go
	total := 0.0
	for _, item := range cart {
		total += item
	}
	fmt.Printf("total cost: %.2f\n", total)
```

- `total := 0.0` — an accumulator.
- `for _, item := range cart` — value-only iteration.
- `total += item` — accumulate: `24.99+19.99+5.49+12.99 = 63.46`.
- `%.2f` — two decimals.

---

## Expected Output

```
items in cart: 4
first item costs: 24.99
last item costs: 12.99
total cost: 63.46
```

## Key Takeaways

1. **Slice literal** — `[]float64{...}`.
2. **`len()`** — the number of elements.
3. **Index access** — `cart[0]`, `cart[len-1]` for first/last.
4. **`range` accumulation** — value-only iteration + sum.
5. **`%.2f`** — float formatting.