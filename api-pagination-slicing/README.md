# api-pagination-slicing

Go-তে **API-style pagination via slicing** শেখার ছোট example — `start = (page-1)*perPage`, bounds-guard, clamp, `all[start:end]` subslice return।

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

`fmt` — `Println`।

### Lines 5–17

```go
func pageItems(all []int, page, perPage int) []int {
	start := (page - 1) * perPage
	if start >= len(all) || start < 0 {
		return []int{}
	}

	end := start + perPage
	if end > len(all) {
		end = len(all)
	}

	return all[start:end]
}
```

**Pagination helper:**

- `start := (page - 1) * perPage` — 1-ভিত্তিক page → 0-ভিত্তিক offset: page 1, perPage 10 → `0`; page 3 → `20`।
- `if start >= len(all) || start < 0` — **bounds-guard**: পেজ last-item-এর বাইরে (page 9 → 80 ≥ 25) বা negative page → `[]int{}` (খালি slice)। Slice out-of-range panic ঠেকানো হয়।
- `end := start + perPage` — chapter-এর শেষ index।
- `if end > len(all) { end = len(all) }` — **clamp**: শেষ পেজ ছোট হলে (page 3 → 20+10=30 → 25-এ কাটা)।
- `return all[start:end]` — **subslice**: নতুন copy নয় — মূল `all`-এর view (অপশনাল: শিখো যে এটা read-এ লক্ষণীয় নয়, লিখলে মূল catalog-ও আঙুলে লাগত)।

### Lines 19–24

```go
func main() {
	catalog := []int{}

	for id := range 25 {
		catalog = append(catalog, 1000+id)
	}
```

- `catalog` — খালি slice literal।
- `for id := range 25` — Go 1.22+ **range-over-int**: `id` = 0..24 → append `1000..1024` (25 item)।

### Lines 26–32

```go
	page1 := pageItems(catalog, 1, 10)
	page3 := pageItems(catalog, 3, 10)
	page9 := pageItems(catalog, 9, 10)

	fmt.Println("page 1:", page1)
	fmt.Println("page 3:", page3)
	fmt.Println("page 9:", page9, "(empty, out of range)")
```

- page 1 → offset 0, items 0–9।
- page 3 → offset 20, end clamp 25 → items 20–24 (5 item)।
- page 9 → offset 80 ≥ 25 → `[]int{}`।

---

## Expected Output

```
page 1: [1000 1001 1002 1003 1004 1005 1006 1007 1008 1009]
page 3: [1020 1021 1022 1023 1024]
page 9: [] (empty, out of range)
```

## মূল শিক্ষা / Key Takeaways

1. **Offset formula** — `(page-1)*perPage`।
2. **Bounds-guard** — `start >= len || start < 0` → empty।
3. **End clamp** — `end > len(all)` → `len(all)`।
4. **Subslice return** — `all[start:end]`।
5. **Range-over-int** — `for id := range 25` (Go 1.22+)।

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

`fmt` — for `Println`.

### Lines 5–17

```go
func pageItems(all []int, page, perPage int) []int {
	start := (page - 1) * perPage
	if start >= len(all) || start < 0 {
		return []int{}
	}

	end := start + perPage
	if end > len(all) {
		end = len(all)
	}

	return all[start:end]
}
```

**Pagination helper:**

- `start := (page - 1) * perPage` — 1-based page → 0-based offset: page 1, perPage 10 → `0`; page 3 → `20`.
- `if start >= len(all) || start < 0` — **bounds-guard**: a page past the last item (page 9 → 80 ≥ 25) or a negative page → `[]int{}`. Prevents slice out-of-range panics.
- `end := start + perPage` — the last index of this page.
- `if end > len(all) { end = len(all) }` — **clamp**: when the final page is shorter (page 3 → 20+10=30 → cut to 25).
- `return all[start:end]` — a **subslice**: not a copy — a view over the original `all` (worth knowing: for reads it's fine, a write here would also touch the catalog).

### Lines 19–24

```go
func main() {
	catalog := []int{}

	for id := range 25 {
		catalog = append(catalog, 1000+id)
	}
```

- `catalog` — an empty slice literal.
- `for id := range 25` — Go 1.22+ **range-over-int**: `id` = 0..24 → append `1000..1024` (25 items).

### Lines 26–32

```go
	page1 := pageItems(catalog, 1, 10)
	page3 := pageItems(catalog, 3, 10)
	page9 := pageItems(catalog, 9, 10)

	fmt.Println("page 1:", page1)
	fmt.Println("page 3:", page3)
	fmt.Println("page 9:", page9, "(empty, out of range)")
```

- page 1 → offset 0, items 0–9.
- page 3 → offset 20, end clamp 25 → items 20–24 (5 items).
- page 9 → offset 80 ≥ 25 → `[]int{}`.

---

## Expected Output

```
page 1: [1000 1001 1002 1003 1004 1005 1006 1007 1008 1009]
page 3: [1020 1021 1022 1023 1024]
page 9: [] (empty, out of range)
```

## Key Takeaways

1. **Offset formula** — `(page-1)*perPage`.
2. **Bounds-guard** — `start >= len || start < 0` → empty.
3. **End clamp** — `end > len(all)` → `len(all)`.
4. **Subslice return** — `all[start:end]`.
5. **Range-over-int** — `for id := range 25` (Go 1.22+).