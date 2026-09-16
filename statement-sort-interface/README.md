# statement-sort-interface

Go-তে **classic `sort.Interface` (Len/Less/Swap) vs modern `slices.SortFunc` + `cmp.Compare`** শেখার ছোট example — bank statement sorting।

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

### Lines 3–8

```go
import (
	"cmp"
	"fmt"
	"slices"
	"sort"
)
```

- `cmp` — total-order `Compare`।
- `sort` — `sort.Interface`-ভিত্তিক `sort.Sort`।
- `slices` — generic `SortFunc`।
- `fmt` — `Println`, `Printf`।

### Lines 10–19

```go
type Txn struct {
	ID     string
	Amount int64
}

type statement []Txn

func (s statement) Len() int           { return len(s) }
func (s statement) Less(i, j int) bool { return s[i].Amount < s[j].Amount }
func (s statement) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
```

**Classic `sort.Interface`** — ৩টা method:

- `Len` — size।
- `Less(i,j)` — amount ascending।
- `Swap` — element বদল।

### Line 21

```go
var _ sort.Interface = (*statement)(nil)
```

**Compile-time check** — `*statement` `sort.Interface` satisfies।

### Lines 23–29

```go
	st := statement{
		{"T-1", 250_00},
		{"T-2", -90_00},
		{"T-3", 1_000_00},
		{"T-4", -15_50},
	}
```

Amounts paisa-তে (numeric underscore): ৳250.00, ৳-90.00, ৳1000.00, ৳-15.50 — negative সহ।

### Lines 31–35

```go
	sort.Sort(st)
```

**`sort.Sort(st)`** — `Less` অনুযায়ী ascending: -90.00, -15.50, 250.00, 1000.00।

### Lines 37–39

```go
	slices.SortFunc(st, func(a, b Txn) int {
		return cmp.Compare(b.Amount, a.Amount)
	})
```

**Modern sort** — `cmp.Compare(b.Amount, a.Amount)`: a,b উল্টানো-য় compare → **descending** (b-a negative হলে a-বড়)।

### Lines 41–44

```go
	fmt.Println("descending by amount (slices.SortFunc):")
	for _, t := range st {
		fmt.Printf("  %s ৳%.2f\n", t.ID, float64(t.Amount)/100)
	}
```

৳ formatting (paisa/100)।

---

## Expected Output

```
ascending by amount (sort.Interface):
  T-2 ৳-90.00
  T-4 ৳-15.50
  T-1 ৳250.00
  T-3 ৳1000.00
descending by amount (slices.SortFunc):
  T-3 ৳1000.00
  T-1 ৳250.00
  T-4 ৳-15.50
  T-2 ৳-90.00
```

## মূল শিক্ষা / Key Takeaways

1. **`sort.Interface`** — custom type + Len/Less/Swap।
2. **`slices.SortFunc`** — generic + comparator func।
3. **`cmp.Compare(b,a)`** — descending flip।
4. **`var _ T = (*X)(nil)`** — static interface check।
5. **Paisa/`%.2f`** — currency formatting + negative।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Lines 3–8

```go
import (
	"cmp"
	"fmt"
	"slices"
	"sort"
)
```

- `cmp` — the total-order `Compare`.
- `sort` — `sort.Interface`-based `sort.Sort`.
- `slices` — generic `SortFunc`.
- `fmt` — for `Println`, `Printf`.

### Lines 10–19

```go
type Txn struct {
	ID     string
	Amount int64
}

type statement []Txn

func (s statement) Len() int           { return len(s) }
func (s statement) Less(i, j int) bool { return s[i].Amount < s[j].Amount }
func (s statement) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
```

**The classic `sort.Interface`** — 3 methods:

- `Len` — the size.
- `Less(i,j)` — ascending by amount.
- `Swap` — swaps elements.

### Line 21

```go
var _ sort.Interface = (*statement)(nil)
```

**A compile-time check** — `*statement` satisfies `sort.Interface`.

### Lines 23–29

```go
	st := statement{
		{"T-1", 250_00},
		{"T-2", -90_00},
		{"T-3", 1_000_00},
		{"T-4", -15_50},
	}
```

Amounts in paisa (numeric underscores): ৳250.00, ৳-90.00, ৳1000.00, ৳-15.50 — negatives included.

### Lines 31–35

```go
	sort.Sort(st)
```

**`sort.Sort(st)`** — ascending by `Less`: -90.00, -15.50, 250.00, 1000.00.

### Lines 37–39

```go
	slices.SortFunc(st, func(a, b Txn) int {
		return cmp.Compare(b.Amount, a.Amount)
	})
```

**The modern sort** — `cmp.Compare(b.Amount, a.Amount)` comparing in reverse → **descending** (b-a negative means a is larger).

### Lines 41–44

```go
	fmt.Println("descending by amount (slices.SortFunc):")
	for _, t := range st {
		fmt.Printf("  %s ৳%.2f\n", t.ID, float64(t.Amount)/100)
	}
```

৳ formatting (paisa/100).

---

## Expected Output

```
ascending by amount (sort.Interface):
  T-2 ৳-90.00
  T-4 ৳-15.50
  T-1 ৳250.00
  T-3 ৳1000.00
descending by amount (slices.SortFunc):
  T-3 ৳1000.00
  T-1 ৳250.00
  T-4 ৳-15.50
  T-2 ৳-90.00
```

## Key Takeaways

1. **`sort.Interface`** — a custom type with Len/Less/Swap.
2. **`slices.SortFunc`** — generic + a comparator func.
3. **`cmp.Compare(b,a)`** — the descending flip.
4. **`var _ T = (*X)(nil)`** — static interface check.
5. **Paisa/`%.2f`** — currency formatting including negatives.