# warehouse-inventory-receiver

Go-তে **pointer receiver methods + sentinel error + `errors.Is` + `%w` + `fmt.Errorf` wrap** শেখার ছোট example — inventory restock (receiver) আর order pick।

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

### Lines 3–6

```go
import (
	"errors"
	"fmt"
)
```

- `errors` — `errors.New`, `errors.Is`।
- `fmt` — `Errorf`, `Println`।

### Lines 8–12

```go
type Inventory struct {
	SKU   string
	Name  string
	Stock int
}
```

`Inventory` struct — SKU, Name, Stock — warehouse item। State-টা methods দিয়ে manage করা হয়।

### Line 14

```go
var ErrorOutOfStock = errors.New("inventory out of stock")
```

**Sentinel error** — package-level error value। `errors.Is` দিয়ে compare-এর reference।

### Lines 16–18

```go
func (i *Inventory) Add(qty int) {
	i.Stock += qty
}
```

`Add` — **pointer receiver** — `i.Stock += qty` direct mutate। Receiver-টা `*Inventory` বলে struct-টা বদলে যায়।

### Lines 20–27

```go
func (i *Inventory) Pick(qty int) error {
	if qty > i.Stock {
		return fmt.Errorf("pick %d of %s: %w", qty, i.SKU, ErrorOutOfStock)
	}

	i.Stock -= qty
	return nil
}
```

`Pick` — order fulfill:

- `qty > i.Stock` — **guard clause**: stock-এর বেশি চাইলে fail।
- Fail-এ: `fmt.Errorf("pick %d of %s: %w", qty, i.SKU, ErrorOutOfStock)` — **`%w`** দিয়ে sentinel wrap + context (কতা, কোন SKU)। `errors.Is`-এ এই chain-টা মেলে।
- Success-এ: `i.Stock -= qty`, `nil`।

### Lines 29–31

```go
func (i *Inventory) Available() int {
	return i.Stock
}
```

`Available` — **getter** (value receiver-ও চলে, কারণ শুধু read; এখানে pointer receiver — বেমানাল নয়, কিন্তু read-লেই value receiverই যথেষ্ট)।

### Lines 33–37

```go
func main() {
	inv := &Inventory{SKU: "SH-42", Name: "Running Shoes 42", Stock: 10}

	inv.Add(5)
	fmt.Println("after restock:", inv.Available())
```

- `&Inventory{...}` — pointer composite literal; initial Stock `10`।
- `Add(5)` → 15 → `after restock: 15`।

### Lines 39–42

```go
	if err := inv.Pick(12); err != nil {
		fmt.Println("order failed:", err)
		fmt.Println("is out of stock error", errors.Is(err, ErrorOutOfStock))
	}
```

- `Pick(12)` — 12 **≤ 15**, তাই এটা সফল (Stock 15→3) — error branch-টা এখানে **run হয় না** (`errors.Is` demo-ও reach হয় না)।
- Error branch-টা চালানো এবং `errors.Is` চেক-টা trigger দেখাতে গেলে — 12-এর বদলে বড় qty (যেমন 16) দিতে হতো।

### Lines 44–50

```go
	if err := inv.Pick(7); err != nil {
		fmt.Println("unexpected:", err)
		return
	}

	fmt.Println("after pickup:", inv.Available())
```

- `Pick(7)` — 7 > 3 (বাকি stock) → **fail**:
  - `unexpected: pick 7 of SH-42: inventory out of stock` — wrapped message।
  - `return` — program এখানেই শেষ, তাই `after pickup:` print হয় না।
- আউটপুটে "unexpected"-label-টা আসলে expected-ই (7 > 3)।

---

## Expected Output

```
after restock: 15
unexpected: pick 7 of SH-42: inventory out of stock
```

## মূল শিক্ষা / Key Takeaways

1. **Pointer receiver methods** — `*Inventory` দিয়ে state mutate।
2. **Sentinel error** — `errors.New` package-level value।
3. **`%w` wrapping** — context (qty/SKU) + sentinel chain।
4. **Guard clause** — `qty > i.Stock` আগেই fail।
5. **Range check first** — `Pick(12)` succeed (15 ≥ 12) → বাকি 3-এ `Pick(7)` fail।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Lines 3–6

```go
import (
	"errors"
	"fmt"
)
```

- `errors` — for `errors.New`, `errors.Is`.
- `fmt` — for `Errorf`, `Println`.

### Lines 8–12

```go
type Inventory struct {
	SKU   string
	Name  string
	Stock int
}
```

`Inventory` struct — SKU, Name, Stock — a warehouse item. State is managed through methods.

### Line 14

```go
var ErrorOutOfStock = errors.New("inventory out of stock")
```

**Sentinel error** — a package-level error value; the reference for `errors.Is` comparisons.

### Lines 16–18

```go
func (i *Inventory) Add(qty int) {
	i.Stock += qty
}
```

`Add` — **pointer receiver** — `i.Stock += qty` mutates directly. Because the receiver is `*Inventory`, the struct changes.

### Lines 20–27

```go
func (i *Inventory) Pick(qty int) error {
	if qty > i.Stock {
		return fmt.Errorf("pick %d of %s: %w", qty, i.SKU, ErrorOutOfStock)
	}

	i.Stock -= qty
	return nil
}
```

`Pick` — fulfills an order:

- `qty > i.Stock` — a **guard clause**: fails when asked for more than in stock.
- On failure: `fmt.Errorf("pick %d of %s: %w", qty, i.SKU, ErrorOutOfStock)` — wraps the sentinel with **`%w`** plus context (how many, which SKU). `errors.Is` matches this chain.
- On success: `i.Stock -= qty`, returns `nil`.

### Lines 29–31

```go
func (i *Inventory) Available() int {
	return i.Stock
}
```

`Available` — a getter (a value receiver would also work since it only reads; a pointer receiver is fine here but a value receiver is enough for reads).

### Lines 33–37

```go
func main() {
	inv := &Inventory{SKU: "SH-42", Name: "Running Shoes 42", Stock: 10}

	inv.Add(5)
	fmt.Println("after restock:", inv.Available())
```

- `&Inventory{...}` — a pointer composite literal; initial Stock `10`.
- `Add(5)` → 15 → `after restock: 15`.

### Lines 39–42

```go
	if err := inv.Pick(12); err != nil {
		fmt.Println("order failed:", err)
		fmt.Println("is out of stock error", errors.Is(err, ErrorOutOfStock))
	}
```

- `Pick(12)` — 12 **≤ 15**, so it succeeds (Stock 15→3) — this error branch is **not reached** (and neither is the `errors.Is` demo).
- To show the branch and trigger the `errors.Is` check, a larger qty (e.g. 16) would be needed.

### Lines 44–50

```go
	if err := inv.Pick(7); err != nil {
		fmt.Println("unexpected:", err)
		return
	}

	fmt.Println("after pickup:", inv.Available())
```

- `Pick(7)` — 7 > 3 (remaining stock) → **fails**:
  - `unexpected: pick 7 of SH-42: inventory out of stock` — the wrapped message.
  - `return` — the program ends here, so `after pickup:` never prints.
- The "unexpected" label is actually expected (7 > 3).

---

## Expected Output

```
after restock: 15
unexpected: pick 7 of SH-42: inventory out of stock
```

## Key Takeaways

1. **Pointer receiver methods** — mutate state via `*Inventory`.
2. **Sentinel error** — a package-level `errors.New` value.
3. **`%w` wrapping** — context (qty/SKU) on the sentinel chain.
4. **Guard clause** — fail early when `qty > i.Stock`.
5. **Range-check first** — `Pick(12)` succeeds (15 ≥ 12), leaving 3; then `Pick(7)` fails.