# batch-payment-grouping

Go-তে **conditional grouping via two slices + pre-allocation** শেখার ছোট example — threshold-এর ভিত্তিতে transaction-কে small/high-value-তে ভাগ করা।

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

### Lines 5–8

```go
type Transaction struct {
	ID     string
	Amount float64
}
```

`Transaction` struct — ID + Amount।

### Lines 10–17

```go
func main() {
	incoming := []Transaction{
		{"TX-001", 12.50},
		{"TX-002", 8900.00},
		{"TX-003", 45.00},
		{"TX-004", 15200.75},
		{"TX-005", 3.99},
	}
```

5টা transaction — **positional struct literal** (`{"TX-001", 12.50}` — field name ছাড়া, declare-order অনুযায়ী)। Mix: ৩টা ছোট, ২টা বড়।

### Lines 19–21

```go
	const highValueThreshold = 5000.0
	small := make([]Transaction, 0, len(incoming))
	highValue := make([]Transaction, 0)
```

- `const` — threshold `5000.0`।
- `small` — **pre-allocation**: `make(..., 0, len(incoming))` — capacity 5 reserve (max-সব small-ও হতে পারত) → `append` re-alloc ছাড়া।
- `highValue` — `make(..., 0)` — কোনো pre-allocation না (size অজানা)।

### Lines 23–29

```go
	for _, tx := range incoming {
		if tx.Amount >= highValueThreshold {
			highValue = append(highValue, tx)
		} else {
			small = append(small, tx)
		}
	}
```

One-pass **conditional grouping** — `>= 5000` → highValue, নাহলে small।

### Lines 31–34

```go
	fmt.Printf("small transactions: %d\n", len(small))
	for _, tx := range small {
		fmt.Printf("  %s: $%.2f\n", tx.ID, tx.Amount)
	}
```

- `len(small)` — count।
- Inner loop — প্রতিটা transaction print: `%s` (ID), `$%.2f` (Amount)।

### Lines 36–39

```go
	fmt.Printf("high-value(manual review): %d\n", len(highValue))
	for _, tx := range highValue {
		fmt.Printf("  %s: $%.2f\n", tx.ID, tx.Amount)
	}
```

Same pattern highValue-র জন্য — "manual review" bucket।

---

## Expected Output

```
small transactions: 3
  TX-001: $12.50
  TX-003: $45.00
  TX-005: $3.99
high-value(manual review): 2
  TX-002: $8900.00
  TX-004: $15200.75
```

## মূল শিক্ষা / Key Takeaways

1. **Conditional split** — threshold-ভিত্তিক দুটো slice-তে grouping।
2. **Pre-allocation** — known max → `make([]T, 0, cap)`।
3. **One-pass** — single `range` + `if/else`।
4. **Positional struct literal** — `{"TX-001", 12.50}`।
5. **bucket printing** — `len` count + inner loop + `%.2f`।

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

### Lines 5–8

```go
type Transaction struct {
	ID     string
	Amount float64
}
```

`Transaction` struct — ID + Amount.

### Lines 10–17

```go
func main() {
	incoming := []Transaction{
		{"TX-001", 12.50},
		{"TX-002", 8900.00},
		{"TX-003", 45.00},
		{"TX-004", 15200.75},
		{"TX-005", 3.99},
	}
```

5 transactions — **positional struct literals** (`{"TX-001", 12.50}` — no field names, filled in declaration order). A mix: 3 small, 2 large.

### Lines 19–21

```go
	const highValueThreshold = 5000.0
	small := make([]Transaction, 0, len(incoming))
	highValue := make([]Transaction, 0)
```

- `const` — the threshold `5000.0`.
- `small` — **pre-allocation**: `make(..., 0, len(incoming))` — reserves capacity 5 (all could be small) → no re-alloc on `append`.
- `highValue` — `make(..., 0)` — no pre-allocation (size unknown).

### Lines 23–29

```go
	for _, tx := range incoming {
		if tx.Amount >= highValueThreshold {
			highValue = append(highValue, tx)
		} else {
			small = append(small, tx)
		}
	}
```

One-pass **conditional grouping** — `>= 5000` into highValue, otherwise small.

### Lines 31–34

```go
	fmt.Printf("small transactions: %d\n", len(small))
	for _, tx := range small {
		fmt.Printf("  %s: $%.2f\n", tx.ID, tx.Amount)
	}
```

- `len(small)` — the count.
- Inner loop — prints each transaction: `%s` (ID), `$%.2f` (Amount).

### Lines 36–39

```go
	fmt.Printf("high-value(manual review): %d\n", len(highValue))
	for _, tx := range highValue {
		fmt.Printf("  %s: $%.2f\n", tx.ID, tx.Amount)
	}
```

The same pattern for highValue — the "manual review" bucket.

---

## Expected Output

```
small transactions: 3
  TX-001: $12.50
  TX-003: $45.00
  TX-005: $3.99
high-value(manual review): 2
  TX-002: $8900.00
  TX-004: $15200.75
```

## Key Takeaways

1. **Conditional split** — threshold-based grouping into two slices.
2. **Pre-allocation** — known max → `make([]T, 0, cap)`.
3. **One-pass** — a single `range` + `if/else`.
4. **Positional struct literal** — `{"TX-001", 12.50}`.
5. **Bucket printing** — `len` count + inner loop + `%.2f`.