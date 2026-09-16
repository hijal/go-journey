# transaction-stringer

Go-তে **`String()` method + `fmt.Stringer` + width-aligned formatting (`%-7s`)** শেখার ছোট example — transaction report।

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

`fmt` — `Println`, `Sprintf`।

### Lines 5–9

```go
type Transaction struct {
	ID     string
	Amount int64
	Status string
}
```

Transaction model — Amount cents-এ (`int64`, বড় মান-ও overflow-safe)।

### Lines 11–13

```go
func (t Transaction) String() string {
	return fmt.Sprintf("txn[%s] %-7s ৳%.2f", t.ID, t.Status, float64(t.Amount)/100)
}
```

**Stringer method:**

- `fmt.Println(t)` যখন কোন type-এর `String() string` থাকে, সেটা auto-use করে।
- `%-7s` — **left-align width 7**: `SETTLED`/`PENDING` ঠিক-ঠিক, `FAILED` (6 chars) প্যাড → ৳ position align।
- `float64(t.Amount)/100` — cents → ৳ (`%.2f`)।
- Value receiver — read-only formatting।

### Lines 15–24

```go
	txns := []Transaction{
		{ID: "T-1001", Amount: 249900, Status: "SETTLED"},
		{ID: "T-1002", Amount: 15000, Status: "PENDING"},
		{ID: "T-1003", Amount: 8900, Status: "FAILED"},
	}

	for _, t := range txns {
		fmt.Println(t)
	}
```

৩টা transaction; `fmt.Println(t)` প্রতিটা `String()` render → column-aligned table।

---

## Expected Output

```
txn[T-1001] SETTLED ৳2499.00
txn[T-1002] PENDING ৳150.00
txn[T-1003] FAILED ৳89.00
```

## মূল শিক্ষা / Key Takeaways

1. **`fmt.Stringer`** — `String() string` auto-called।
2. **`%-7s`** — left-aligned width formatting।
3. **Value receiver String** — immutable formatting।
4. **`int64` cents + `/100`** — precision-safe currency।
5. **Table layout** — aligned columns via width spec।

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

`fmt` — for `Println`, `Sprintf`.

### Lines 5–9

```go
type Transaction struct {
	ID     string
	Amount int64
	Status string
}
```

The transaction model — `Amount` in cents (`int64`, overflow-safe even for big values).

### Lines 11–13

```go
func (t Transaction) String() string {
	return fmt.Sprintf("txn[%s] %-7s ৳%.2f", t.ID, t.Status, float64(t.Amount)/100)
}
```

**The Stringer method:**

- `fmt.Println(t)` auto-uses `String() string` when a type has one.
- `%-7s` — **left-align width 7**: `SETTLED`/`PENDING` fit exactly; `FAILED` (6 chars) gets padded so the ৳ column aligns.
- `float64(t.Amount)/100` — cents → ৳ (`%.2f`).
- A value receiver — read-only formatting.

### Lines 15–24

```go
	txns := []Transaction{
		{ID: "T-1001", Amount: 249900, Status: "SETTLED"},
		{ID: "T-1002", Amount: 15000, Status: "PENDING"},
		{ID: "T-1003", Amount: 8900, Status: "FAILED"},
	}

	for _, t := range txns {
		fmt.Println(t)
	}
```

3 transactions; each `fmt.Println(t)` invokes `String()`, producing an aligned table.

---

## Expected Output

```
txn[T-1001] SETTLED ৳2499.00
txn[T-1002] PENDING ৳150.00
txn[T-1003] FAILED ৳89.00
```

## Key Takeaways

1. **`fmt.Stringer`** — `String() string` is auto-invoked.
2. **`%-7s`** — left-aligned width formatting.
3. **Value-receiver String** — immutable formatting.
4. **`int64` cents + `/100`** — precision-safe currency.
5. **Table layout** — aligned columns via width specifier.