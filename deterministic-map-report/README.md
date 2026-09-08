# deterministic-map-report

Go-তে **map iteration randomness**, **key sort করে deterministic report** (`slices.Sort`), আর **`make` + `Printf` alignment** শেখার ছোট example.

**ভাষা নির্বাচন করুন / Choose language:**

[বাংলা](#bangla) * [English](#english)

---

<a name="bangla"></a>

## বাংলা সংস্করণ

### Lines 1-6

```go
package main

import (
	"fmt"
	"slices"
)
```

- `package main` -- একটা executable program.
- `fmt` -- output print করার জন্য.
- `slices` -- slice sort করার standard library package (Go 1.21+).

### Lines 9-14

```go
revenue := map[string]int{
	electronics: 850_000,
	fashion:     420_000,
	books:       95_000,
	grocery:     610_000,
}
```

`revenue` -- category -> টাকার একটা **map**. Key `string`, value `int`.

- `850_000` -- underscore দিয়ে সংখ্যা লেখা যায় readability-র জন্য (Go ignore করে; 850000-এর সমান).

### Line 15

```go
categories := make([]string, 0, len(revenue))
```

`categories` -- map-এর key-গুলো রাখার slice. `make([]string, 0, len(revenue))` দিয়ে length 0 কিন্তু capacity `len(revenue)` (= 4) নিয়ে বানানো হয়, যেন `append` করতে reallocation না লাগে.

### Lines 17-19

```go
for cat := range revenue {
	categories = append(categories, cat)
}
```

`for ... range` দিয়ে map-এর সব key তুলে slice-এ `append` করা হয়.

### Line 20

```go
slices.Sort(categories)
```

Key-গুলো alphabetically sort করা হয়: books, electronics, fashion, grocery.

**কেন sort দরকার:** Go-তে map iteration order **random** -- প্রতিবার `go run` করলে ভিন্ন order আসতে পারে. Report/invoice-এ stable output চাইলে key sort করে তারপর loop করতে হয় (deterministic order).

### Lines 22-26

```go
total := 0
for _, cat := range categories {
	fmt.Printf("%-12s %10d BDT\n", cat, revenue[cat])
	total += revenue[cat]
}
```

- `total := 0` -- মোট revenue-এর accumulator.
- `for _, cat := range categories` -- sorted key-এর উপর loop (`_` দিয়ে index ignore করা হয়).
- `fmt.Printf("%-12s %10d BDT\n", cat, revenue[cat])` -- `%-12s` মানে string-টা বাম-দিকে align করে 12 ঘর জায়গা নেবে, `%10d` মানে সংখ্যা ডান-দিকে align করে 10 ঘর -- ফলে column সোজা থাকে.
- `total += revenue[cat]` -- প্রতিটা category যোগ হয় (95000 + 850000 + 420000 + 610000 = 1975000).

### Line 27

```go
fmt.Printf("%-12s %10d BDT\n", "TOTAL", total)
```

একই format-এ শেষে TOTAL line print করে.

### Line 28

```go
}
```

Closing brace -- `main` function শেষ হয়.

---

## Expected Output

```
books             95000 BDT
electronics      850000 BDT
fashion          420000 BDT
grocery          610000 BDT
TOTAL           1975000 BDT
```

## মূল শিক্ষা / Key Takeaways

1. **Map iteration random** -- Go-তে map-এর order প্রতিবার বদলাতে পারে.
2. **Deterministic report** -- key slice-এ তুলে `slices.Sort` করে loop করলে stable output.
3. **`make` with capacity** -- `make([]string, 0, len(m))` দিয়ে append-efficient slice.
4. **`%-12s` / `%10d`** -- `-` মানে left-align, সংখ্যা ছাড়া মানে right-align; column সোজা রাখে.
5. **Underscore literal** -- `850_000` = 850000, শুধু পড়তে সুবিধা.

---

---

<a name="english"></a>

##  English Version

### Lines 1-6

```go
package main

import (
	"fmt"
	"slices"
)
```

- `package main` -- an executable program.
- `fmt` -- for console output.
- `slices` -- the standard library package for sorting slices (Go 1.21+).

### Lines 9-14

```go
revenue := map[string]int{
	electronics: 850_000,
	fashion:     420_000,
	books:       95_000,
	grocery:     610_000,
}
```

`revenue` -- a **map** from category to money. Key is `string`, value is `int`.

- `850_000` -- underscores are allowed in number literals for readability (Go ignores them; same as 850000).

### Line 15

```go
categories := make([]string, 0, len(revenue))
```

`categories` -- a slice to hold the map keys. `make([]string, 0, len(revenue))` creates it with length 0 but capacity `len(revenue)` (= 4), so `append` needs no reallocation.

### Lines 17-19

```go
for cat := range revenue {
	categories = append(categories, cat)
}
```

Collects every map key into the slice with `for ... range` + `append`.

### Line 20

```go
slices.Sort(categories)
```

Sorts the keys alphabetically: books, electronics, fashion, grocery.

**Why sort:** map iteration order in Go is **random** -- each `go run` may print a different order. For a stable report/invoice, sort the keys first and then loop (deterministic order).

### Lines 22-26

```go
total := 0
for _, cat := range categories {
	fmt.Printf("%-12s %10d BDT\n", cat, revenue[cat])
	total += revenue[cat]
}
```

- `total := 0` -- an accumulator for the total revenue.
- `for _, cat := range categories` -- loops over the sorted keys (`_` discards the index).
- `fmt.Printf("%-12s %10d BDT\n", cat, revenue[cat])` -- `%-12s` left-aligns the string in 12 columns, `%10d` right-aligns the number in 10 columns -- so columns line up.
- `total += revenue[cat]` -- adds each category (95000 + 850000 + 420000 + 610000 = 1975000).

### Line 27

```go
fmt.Printf("%-12s %10d BDT\n", "TOTAL", total)
```

Prints the final TOTAL line in the same format.

### Line 28

```go
}
```

Closing brace -- ends the `main` function.

---

## Expected Output

```
books             95000 BDT
electronics      850000 BDT
fashion          420000 BDT
grocery          610000 BDT
TOTAL           1975000 BDT
```

## Key Takeaways

1. **Map iteration is random** -- Go may return keys in a different order each run.
2. **Deterministic report** -- collect keys into a slice, `slices.Sort`, then loop for stable output.
3. **`make` with capacity** -- `make([]string, 0, len(m))` makes appends efficient.
4. **`%-12s` / `%10d`** -- `-` means left-align, otherwise right-align; keeps columns straight.
5. **Underscore literal** -- `850_000` = 850000, just easier to read.
