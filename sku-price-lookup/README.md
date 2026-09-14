# sku-price-lookup

Go-তে **map + comma-ok lookup idiom** শেখার ছোট example — SKU → price (cents) ক্যাটালগ lookup, present/absent দুটো path।

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

`fmt` — `Printf`।

### Lines 5–10

```go
func main() {
	prices := map[string]int{
		"SKU-1001": 1999,
		"SKU-1002": 499,
		"SKU-1003": 12999,
	}
```

**Map literal** — `string → int` (cents-integer — float-মানি এড়ায়)। ৩টা SKU।

### Line 11

```go
	fmt.Printf("total products: %d\n", len(prices))
```

`len(map)` — entries-সংখ্যা → 3।

### Lines 13–19

```go
	sku := "SKU-1002"

	if price, ok := prices[sku]; ok {
		fmt.Printf("%s costs %d cents (%.2f USD)\n", sku, price, float64(price)/100)
	} else {
		fmt.Printf("%s not found in catalog\n", sku)
	}
```

**Comma-ok lookup (present path):**

- `if price, ok := prices[sku]; ok` — **initializer + comma-ok**।
- `ok` = key-টা map-এ আছে কিনা; `price` = value (missing-এ zero)।
- Found → print: `float64(price)/100` — cents→USD।

*(কেন comma-ok? শুধু `prices[sku]` দিলে missing-এ 0 return হতো — 0-দাম বা not-found আলাদা করা যেত না।)*

### Lines 21–24

```go
	sku = "SKU-9999"
	if _, ok := prices[sku]; !ok {
		fmt.Printf("%s not found, suggest alternatives\n", sku)
	}
```

**Absent path:**

- `_, ok` — value-টা discard, শুধু presence-check।
- `!ok` → not-found flow (suggest alternatives)।

---

## Expected Output

```
total products: 3
SKU-1002 costs 499 cents (4.99 USD)
SKU-9999 not found, suggest alternatives
```

## মূল শিক্ষা / Key Takeaways

1. **Map literal** — `map[string]int{...}`।
2. **Comma-ok** — `v, ok := m[k]`।
3. **`if` initializer** — `if v, ok := ...; ok` guard।
4. **Value-discard** — `_, ok` presence-only।
5. **Implicit-currency** — cents-integer + cast।

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

`fmt` — for `Printf`.

### Lines 5–10

```go
func main() {
	prices := map[string]int{
		"SKU-1001": 1999,
		"SKU-1002": 499,
		"SKU-1003": 12999,
	}
```

**A map literal** — `string → int` (integer cents — avoids float money). 3 SKUs.

### Line 11

```go
	fmt.Printf("total products: %d\n", len(prices))
```

`len(map)` — the entry count → 3.

### Lines 13–19

```go
	sku := "SKU-1002"

	if price, ok := prices[sku]; ok {
		fmt.Printf("%s costs %d cents (%.2f USD)\n", sku, price, float64(price)/100)
	} else {
		fmt.Printf("%s not found in catalog\n", sku)
	}
```

**Comma-ok lookup (present path):**

- `if price, ok := prices[sku]; ok` — an **initializer + comma-ok**.
- `ok` = whether the key exists; `price` = the value (zero when missing).
- Found → prints: `float64(price)/100` — cents → USD.

*(Why comma-ok? A bare `prices[sku]` would return 0 when missing — you couldn't tell a 0-price from a not-found.)*

### Lines 21–24

```go
	sku = "SKU-9999"
	if _, ok := prices[sku]; !ok {
		fmt.Printf("%s not found, suggest alternatives\n", sku)
	}
```

**The absent path:**

- `_, ok` — discards the value, presence check only.
- `!ok` → the not-found flow (suggest alternatives).

---

## Expected Output

```
total products: 3
SKU-1002 costs 499 cents (4.99 USD)
SKU-9999 not found, suggest alternatives
```

## Key Takeaways

1. **Map literal** — `map[string]int{...}`.
2. **Comma-ok** — `v, ok := m[k]`.
3. **`if` initializer** — the `if v, ok := ...; ok` guard.
4. **Value-discard** — `_, ok` presence-only.
5. **Implicit currency** — integer cents + a cast.**