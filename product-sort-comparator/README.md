# product-sort-comparator

Go-তে **`slices.SortFunc`** (custom comparator), **`cmp.Compare`** আর **composite (multi-key) sorting** শেখার ছোট example — struct slice-কে price, তারপর name দিয়ে sort।

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

- `cmp` — `cmp.Compare` (Go 1.21+)।
- `fmt` — print।
- `slices` — `slices.SortFunc`, `slices.MaxFunc`।

### Lines 9–13

```go
type Product struct {
	name  string
	price float64
	sold  int
}
```

`Product` struct — ৩টা field: name, price, sold (sales count)।

### Line 15

```go
func main() {
```

Program-এর entry point।

### Lines 16–21

```go
catalog := []Product{
	{name: "keyboard", price: 1500, sold: 320},
	{name: "mouse", price: 750, sold: 950},
	{name: "monitor", price: 12500, sold: 120},
	{name: "usb hub", price: 750, sold: 410},
}
```

৪টা product-এর slice। Note: mouse আর usb hub-এর price **same (750)**।

### Lines 23–28

```go
slices.SortFunc(catalog, func(a, b Product) int {
	if c := cmp.Compare(a.price, b.price); c != 0 {
		return c
	}
	return cmp.Compare(a.name, b.name)
})
```

**`slices.SortFunc(catalog, cmp)`** — `catalog`-কে custom comparator-এ sort করে (in-place):

- comparator: একটা function যা 2টা `Product` নেয় এবং int return করে:
  - negative → a আসবে আগে
  - zero → equal
  - positive → b আসবে আগে
- **Multi-key sort:** প্রথমে `cmp.Compare(a.price, b.price)` — price compare। `c != 0` হলে return (price আলাদা → sort দেয়া)। equal হলে (750 vs 750) **fallback**: `cmp.Compare(a.name, b.name)` — lexicographic name sort।

**Result order (price asc, tie → name asc):**
- `mouse` (750), `usb hub` (750) — equal price → name: "mouse" < "usb hub"
- `keyboard` (1500)
- `monitor` (12500)

### Lines 30–32

```go
for _, p := range catalog {
	fmt.Printf("%-10s %8.2f\n", p.name, p.price)
}
```

**`%-10s`** — left-aligned 10-width string; **`%8.2f`** — right-aligned 8-width 2-decimal float:

```
mouse        750.00
usb hub      750.00
keyboard    1500.00
monitor    12500.00
```

### Lines 34–38

```go
best := slices.MaxFunc(catalog, func(a, b Product) int {
	return cmp.Compare(a.sold, b.sold)
})
fmt.Println("best seller:", best.name)
```

**`slices.MaxFunc`** — comparator-এর অনুযায়ী max element:

- comparator: `cmp.Compare(a.sold, b.sold)` — sold-count অনুযায়ী।
- সবচেয়ে বেশি sold: mouse (950)।
- Output: `best seller: mouse`।

### Line 39

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
mouse        750.00
usb hub      750.00
keyboard    1500.00
monitor    12500.00
best seller: mouse
```

## মূল শিক্ষা / Key Takeaways

1. **`slices.SortFunc`** — custom comparator-এ in-place sort।
2. **`cmp.Compare`** — ordered types-এর জন্য generic comparison (−1/0/+1)।
3. **Multi-key sort** — primary (price) compare; tie-তে secondary (name)।
4. **`slices.MaxFunc`** — comparator-defined max।
5. **`Printf` alignment** — `%-10s`, `%8.2f` column formatting।

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

- `cmp` — for `cmp.Compare` (Go 1.21+).
- `fmt` — for printing.
- `slices` — for `slices.SortFunc`, `slices.MaxFunc`.

### Lines 9–13

```go
type Product struct {
	name  string
	price float64
	sold  int
}
```

The `Product` struct — 3 fields: name, price, sold (sales count).

### Line 15

```go
func main() {
```

Program entry point.

### Lines 16–21

```go
catalog := []Product{
	{name: "keyboard", price: 1500, sold: 320},
	{name: "mouse", price: 750, sold: 950},
	{name: "monitor", price: 12500, sold: 120},
	{name: "usb hub", price: 750, sold: 410},
}
```

A slice of 4 products. Note that mouse and usb hub share the **same price (750)**.

### Lines 23–28

```go
slices.SortFunc(catalog, func(a, b Product) int {
	if c := cmp.Compare(a.price, b.price); c != 0 {
		return c
	}
	return cmp.Compare(a.name, b.name)
})
```

**`slices.SortFunc(catalog, cmp)`** — sorts `catalog` in place with a custom comparator:

- The comparator: a function taking two `Product`s and returning an int:
  - negative → `a` comes first
  - zero → equal
  - positive → `b` comes first
- **Multi-key sort:** first `cmp.Compare(a.price, b.price)` — price comparison. If `c != 0`, return (prices differ → order decided). Otherwise (750 vs 750) **fall back** to `cmp.Compare(a.name, b.name)` — lexicographic order by name.

**Result order (price asc, tie → name asc):**
- `mouse` (750), `usb hub` (750) — equal price → name: "mouse" < "usb hub"
- `keyboard` (1500)
- `monitor` (12500)

### Lines 30–32

```go
for _, p := range catalog {
	fmt.Printf("%-10s %8.2f\n", p.name, p.price)
}
```

**`%-10s`** — left-aligned 10-width string; **`%8.2f`** — right-aligned 8-width 2-decimal float:

```
mouse        750.00
usb hub      750.00
keyboard    1500.00
monitor    12500.00
```

### Lines 34–38

```go
best := slices.MaxFunc(catalog, func(a, b Product) int {
	return cmp.Compare(a.sold, b.sold)
})
fmt.Println("best seller:", best.name)
```

**`slices.MaxFunc`** — the max element per the comparator:

- Comparator: `cmp.Compare(a.sold, b.sold)` — ordered by sold count.
- The largest is mouse (950).
- Output: `best seller: mouse`.

### Line 39

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
mouse        750.00
usb hub      750.00
keyboard    1500.00
monitor    12500.00
best seller: mouse
```

## Key Takeaways

1. **`slices.SortFunc`** — in-place sort with a custom comparator.
2. **`cmp.Compare`** — generic comparison for ordered types (−1/0/+1).
3. **Multi-key sort** — primary (price) compare; tie → secondary (name).
4. **`slices.MaxFunc`** — comparator-defined max.
5. **`Printf` alignment** — `%-10s`, `%8.2f` column formatting.