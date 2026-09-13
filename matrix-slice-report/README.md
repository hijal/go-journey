# matrix-slice-report

Go-তে **slice-of-slices matrix (`[][]int`) + row/column aggregation** শেখার ছোট example — warehouse × week inventory report।

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

### Lines 5–11

```go
func main() {
	stock := [][]int{
		{120, 95, 140, 110},  // warehouse A
		{60, 70, 0, 55},      // warehouse B
		{200, 180, 190, 205}, // warehouse C
	}
	warehouses := []string{"A", "B", "C"}
```

- `stock := [][]int{...}` — **slice-of-slices matrix**: 3টা row (warehouse), প্রত্যেকে 4টা element (week)।
- `warehouses` — parallel-array: row-index ↔ নাম mapping (row 0 → "A"...)।

### Lines 13–19

```go
	for i, row := range stock {
		total := 0
		for _, item := range row {
			total += item
		}
		fmt.Printf("warehouse %s total: %d\n", warehouses[i], total)
	}
```

**Row-major aggregation (per-warehouse):**

- `for i, row := range stock` — index + value (inner slice)।
- Inner loop — row-এর element যোগ।
- `warehouses[i]` — index-ভিত্তিক নাম/মান pairing।

### Lines 21–27

```go
	for w := range 4 {
		weekTotal := 0
		for _, row := range stock {
			weekTotal += row[w]
		}
		fmt.Printf("week %d total: %d\n", w+1, weekTotal)
	}
```

**Column-major aggregation (per-week):**

- `for w := range 4` — Go 1.22+ **range-over-int**: `w` = 0..3 (week-column)।
- `for _, row := range stock { weekTotal += row[w] }` — প্রতিটা row-এর একই column-`w` — **column-sum**।
- `w+1` — user-friendly 1-ভিত্তিক week-নাম।

---

## Expected Output

```
warehouse A total: 465
warehouse B total: 185
warehouse C total: 775
week 1 total: 380
week 2 total: 345
week 3 total: 330
week 4 total: 370
```

## মূল শিক্ষা / Key Takeaways

1. **`[][]int`** — slice-of-slices matrix।
2. **Outer/inner loop** — row-major।
3. **Column-major sum** — stable index `row[w]`।
4. **Parallel array** — `warehouses[i]`।
5. **Range-over-int** — `for w := range 4` (Go 1.22+)।

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

### Lines 5–11

```go
func main() {
	stock := [][]int{
		{120, 95, 140, 110},  // warehouse A
		{60, 70, 0, 55},      // warehouse B
		{200, 180, 190, 205}, // warehouse C
	}
	warehouses := []string{"A", "B", "C"}
```

- `stock := [][]int{...}` — a **slice-of-slices matrix**: 3 rows (warehouses), each with 4 elements (weeks).
- `warehouses` — a parallel array: row-index ↔ name mapping (row 0 → "A"…).

### Lines 13–19

```go
	for i, row := range stock {
		total := 0
		for _, item := range row {
			total += item
		}
		fmt.Printf("warehouse %s total: %d\n", warehouses[i], total)
	}
```

**Row-major aggregation (per warehouse):**

- `for i, row := range stock` — index + value (the inner slice).
- Inner loop — sums the row's elements.
- `warehouses[i]` — index-based name/value pairing.

### Lines 21–27

```go
	for w := range 4 {
		weekTotal := 0
		for _, row := range stock {
			weekTotal += row[w]
		}
		fmt.Printf("week %d total: %d\n", w+1, weekTotal)
	}
```

**Column-major aggregation (per week):**

- `for w := range 4` — Go 1.22+ **range-over-int**: `w` = 0..3 (the week column).
- `for _, row := range stock { weekTotal += row[w] }` — the same column-`w` of every row — a **column-sum**.
- `w+1` — a user-friendly 1-based week name.

---

## Expected Output

```
warehouse A total: 465
warehouse B total: 185
warehouse C total: 775
week 1 total: 380
week 2 total: 345
week 3 total: 330
week 4 total: 370
```

## Key Takeaways

1. **`[][]int`** — a slice-of-slices matrix.
2. **Outer/inner loop** — row-major.
3. **Column-major sum** — the stable index `row[w]`.
4. **Parallel array** — `warehouses[i]`.
5. **Range-over-int** — `for w := range 4` (Go 1.22+).