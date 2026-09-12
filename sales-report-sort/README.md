# sales-report-sort

Go-তে **`slices.SortFunc` + `cmp.Compare` multi-key sort + `slices.MaxFunc` + anonymous struct** শেখার ছোট example — sales report (top-by-amount, by-region-then-amount)।

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

### Lines 3–7

```go
import (
	"cmp"
	"fmt"
	"slices"
)
```

- `cmp` — `cmp.Compare` (generic comparison, Go 1.21+)।
- `fmt` — `Printf`।
- `slices` — `SortFunc`, `MaxFunc` (Go 1.21+)।

### Lines 9–13

```go
type Sale struct {
	Region string
	Amount float64
	Rep    string
}
```

`Sale` struct — region, amount, rep — sales record।

### Lines 15–22

```go
func main() {
	sales := []Sale{
		{Region: "Dhaka", Amount: 1200, Rep: "Ayesha"},
		{Region: "Chattogram", Amount: 800, Rep: "Rakib"},
		{Region: "Dhaka", Amount: 450, Rep: "Tanvir"},
		{Region: "Sylhet", Amount: 990, Rep: "Mitu"},
		{Region: "Chattogram", Amount: 310, Rep: "Farhan"},
	}
```

৫টা sales record।

### Lines 24–31

```go
	slices.SortFunc(sales, func(a, b Sale) int {
		return cmp.Compare(b.Amount, a.Amount)
	})

	fmt.Println("Top sales (desc):")
	for _, s := range sales {
		fmt.Printf("  %-11s %-8s %8.2f\n", s.Region, s.Rep, s.Amount)
	}
```

**Descending sort:**

- `slices.SortFunc(sales, less)` — আপনার-provided comparator `func(a, b) int`।
- `cmp.Compare(b.Amount, a.Amount)` — arguments swapped → **descending** (1200 → 310)।
- `%-11s` — left-align 11-char; `%-8s`; `%8.2f` — right-align, 2 decimals।

**Output:** Ayesha 1200, Mitu 990, Rakib 800, Tanvir 450, Farhan 310 ✓।

### Lines 33–43

```go
	slices.SortFunc(sales, func(a, b Sale) int {
		if c := cmp.Compare(a.Region, b.Region); c != 0 {
			return c
		}
		return cmp.Compare(b.Amount, a.Amount)
	})

	fmt.Println("By region, then amount:")
```

**Multi-key sort:**

- Primary: `cmp.Compare(a.Region, b.Region)` — region ascending (Chattogram < Dhaka < Sylhet)।
- `if c != 0 { return c }` — region আলাদা হলে এখানেই সিদ্ধান্ত নেওয়া হয়।
- Secondary: `cmp.Compare(b.Amount, a.Amount)` — একই region-এ amount descending।

**Output:** Chattogram (800, 310) → Dhaka (1200, 450) → Sylhet (990) — region asc + amount desc।

### Lines 45–53

```go
	var summary struct {
		Orders  int
		Revenue float64
		Best    Sale
	}

	summary.Best = slices.MaxFunc(sales, func(a, b Sale) int {
		return cmp.Compare(a.Amount, b.Amount)
	})
```

**Anonymous struct:** `var summary struct {...}` — inline-defined type (নামহীন) — ছোট aggregator-এর জন্য functions-বিহীন।

`slices.MaxFunc` — comparator-এর basis-এ সর্বোচ্চ: `cmp.Compare(a.Amount, b.Amount)` → সর্বোচ্চ Amount-এর Sale। Ayesha (1200)।

### Lines 55–60

```go
	for _, s := range sales {
		summary.Orders++
		summary.Revenue += s.Amount
	}
	fmt.Printf("Summary: %d orders, revenue %.2f, best rep = %s\n",
		summary.Orders, summary.Revenue, summary.Best.Rep)
```

- Loop: count orders (5) + accumulate revenue (3750)।
- `summary.Best.Rep` — Ayesha।

---

## Expected Output

```
Top sales (desc):
  Dhaka       Ayesha    1200.00
  Sylhet      Mitu       990.00
  Chattogram  Rakib      800.00
  Dhaka       Tanvir     450.00
  Chattogram  Farhan     310.00
By region, then amount:
  Chattogram  Rakib      800.00
  Chattogram  Farhan     310.00
  Dhaka       Ayesha    1200.00
  Dhaka       Tanvir     450.00
  Sylhet      Mitu       990.00
Summary: 5 orders, revenue 3750.00, best rep = Ayesha
```

## মূল শিক্ষা / Key Takeaways

1. **`slices.SortFunc`** — custom comparator দিয়ে sort (reverse order arguments দিয়ে descending)।
2. **`cmp.Compare`** — generic ordering।
3. **Multi-key sort** — primary key আলাদা হলে early return, নাহলে secondary।
4. **`slices.MaxFunc`** — comparator-ভিত্তিক max।
5. **Anonymous struct** — `var summary struct{...}` inline type।
6. **Format padding** — `%-11s`, `%8.2f`।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Lines 3–7

```go
import (
	"cmp"
	"fmt"
	"slices"
)
```

- `cmp` — `cmp.Compare` (generic comparison, Go 1.21+).
- `fmt` — for `Printf`.
- `slices` — `SortFunc`, `MaxFunc` (Go 1.21+).

### Lines 9–13

```go
type Sale struct {
	Region string
	Amount float64
	Rep    string
}
```

`Sale` struct — region, amount, rep — a sales record.

### Lines 15–22

```go
func main() {
	sales := []Sale{
		{Region: "Dhaka", Amount: 1200, Rep: "Ayesha"},
		{Region: "Chattogram", Amount: 800, Rep: "Rakib"},
		{Region: "Dhaka", Amount: 450, Rep: "Tanvir"},
		{Region: "Sylhet", Amount: 990, Rep: "Mitu"},
		{Region: "Chattogram", Amount: 310, Rep: "Farhan"},
	}
```

Five sales records.

### Lines 24–31

```go
	slices.SortFunc(sales, func(a, b Sale) int {
		return cmp.Compare(b.Amount, a.Amount)
	})

	fmt.Println("Top sales (desc):")
	for _, s := range sales {
		fmt.Printf("  %-11s %-8s %8.2f\n", s.Region, s.Rep, s.Amount)
	}
```

**Descending sort:**

- `slices.SortFunc(sales, less)` — your own `func(a, b) int` comparator.
- `cmp.Compare(b.Amount, a.Amount)` — swapped arguments → **descending** (1200 → 310).
- `%-11s` — left-aligned in 11 chars; `%-8s`; `%8.2f` — right-aligned, 2 decimals.

**Output:** Ayesha 1200, Mitu 990, Rakib 800, Tanvir 450, Farhan 310 ✓.

### Lines 33–43

```go
	slices.SortFunc(sales, func(a, b Sale) int {
		if c := cmp.Compare(a.Region, b.Region); c != 0 {
			return c
		}
		return cmp.Compare(b.Amount, a.Amount)
	})

	fmt.Println("By region, then amount:")
```

**Multi-key sort:**

- Primary: `cmp.Compare(a.Region, b.Region)` — region ascending (Chattogram < Dhaka < Sylhet).
- `if c != 0 { return c }` — if regions differ, decide here.
- Secondary: `cmp.Compare(b.Amount, a.Amount)` — within the same region, amount descending.

**Output:** Chattogram (800, 310) → Dhaka (1200, 450) → Sylhet (990) — region asc + amount desc.

### Lines 45–53

```go
	var summary struct {
		Orders  int
		Revenue float64
		Best    Sale
	}

	summary.Best = slices.MaxFunc(sales, func(a, b Sale) int {
		return cmp.Compare(a.Amount, b.Amount)
	})
```

**Anonymous struct:** `var summary struct{...}` — an inline-defined type (nameless) — a function-free small aggregator.

`slices.MaxFunc` — finds the max keyed by the comparator: `cmp.Compare(a.Amount, b.Amount)` → the Sale with the highest Amount. Ayesha (1200).

### Lines 55–60

```go
	for _, s := range sales {
		summary.Orders++
		summary.Revenue += s.Amount
	}
	fmt.Printf("Summary: %d orders, revenue %.2f, best rep = %s\n",
		summary.Orders, summary.Revenue, summary.Best.Rep)
```

- Loop: counts orders (5) and accumulates revenue (3750).
- `summary.Best.Rep` — Ayesha.

---

## Expected Output

```
Top sales (desc):
  Dhaka       Ayesha    1200.00
  Sylhet      Mitu       990.00
  Chattogram  Rakib      800.00
  Dhaka       Tanvir     450.00
  Chattogram  Farhan     310.00
By region, then amount:
  Chattogram  Rakib      800.00
  Chattogram  Farhan     310.00
  Dhaka       Ayesha    1200.00
  Dhaka       Tanvir     450.00
  Sylhet      Mitu       990.00
Summary: 5 orders, revenue 3750.00, best rep = Ayesha
```

## Key Takeaways

1. **`slices.SortFunc`** — custom-comparator sorting (descending via swapped arguments).
2. **`cmp.Compare`** — generic ordering.
3. **Multi-key sort** — early-return on the primary key, secondary otherwise.
4. **`slices.MaxFunc`** — comparator-based max.
5. **Anonymous struct** — `var summary struct{...}` inline type.
6. **Format padding** — `%-11s`, `%8.2f`.