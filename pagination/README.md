# pagination

Go-তে **ceiling division** (`(total + limit - 1) / limit`), **`min` builtin** আর **offset-based pagination** শেখার ছোট example।

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

Console-এ output print করার জন্য `fmt` package import করা হয়।

### Lines 5–10

```go
func pageCount(total, limit int) int {
	if limit <= 0 {
		return 0
	}
	return (total + limit - 1) / limit
}
```

`pageCount` — total item থেকে page সংখ্যা বের করে:

- `if limit <= 0` — guard: limit 0 বা negative হলে **0** return (zero/negative share এড়াতে)। Division by zero avoid করা হয়।
- `return (total + limit - 1) / limit` — **ceiling division**। আগে `total + limit - 1` যোগ, তারপর `/ limit`।

**কেন `+ limit - 1`:** integer division floor করে (নিচের দিকে)। `95/10 = 9` (floor), কিন্তু আসলে 10টা page লাগবে (95টা item, প্রতিটা page-এ 10)। `95 + 10 - 1 = 104`; `104/10 = 10` (floor = 10) → সঠিক সবশেষ page-টা অন্তর্ভুক্ত হয়। 95 ইন place-এ `(95 + 9)/10 = 104/10 = 10`।

- Exact multiple হলে (যেমন `total=100, limit=10`): `(100 + 9)/10 = 109/10 = 10` — ঠিকই 10।

### Line 12

```go
func main() {
```

Program-এর entry point।

### Line 13

```go
totalProducts, limit, page := 95, 10, 3
```

Input: 95টা product, প্রতিটা page-এ 10টা, user-এর current page = 3।

### Lines 15–16

```go
pages := pageCount(totalProducts, limit)
offset := (page - 1) * limit
```

- `pages` = `pageCount(95, 10)` = 10টা page।
- `offset` = `(3 - 1) * 10` = 20 — **offset-based pagination:** page 3-এ item 21 থেকে শুরু। (page 1 → offset 0, page 2 → offset 10, page 3 → offset 20...)

### Lines 18–21

```go
fmt.Println("total pages:", pages)
fmt.Println("page", page, "show items", offset+1, "to", min(offset+limit, totalProducts))
fmt.Println("has next page:", page < pages)
fmt.Println("has prev page:", page > 1)
```

- `total pages: 10`
- `page 3 show items 21 to 30` — `offset+1` = 21 (start), `min(offset+limit, totalProducts)` = `min(20+10, 95)` = `min(30, 95)` = 30 (end)। **`min`** নিশ্চিত করে শেষ page-এ total-এর বেশি দেখায় না (যেমন last page-এ 91 to 95)।
- `has next page: true` — `3 < 10`।
- `has prev page: true` — `3 > 1`।

### Line 22

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
total pages: 10
page 3 show items 21 to 30
has next page: true
has prev page: true
```

## মূল শিক্ষা / Key Takeaways

1. **Ceiling division** — `(total + limit - 1) / limit` — সবশেষ আংশিক page-ও count করে।
2. **Offset pagination** — `offset = (page - 1) * limit` — কোন item থেকে শুরু।
3. **`min` guard** — শেষ page-এ total-এর বেশি দেখানো আটকায়।
4. **Guard clause** — `limit <= 0` → 0 (division by zero থেকে বাঁচা)।
5. **Bounds check** — `page < pages` / `page > 1` দিয়ে next/prev flag।

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

Imports the `fmt` package for console output.

### Lines 5–10

```go
func pageCount(total, limit int) int {
	if limit <= 0 {
		return 0
	}
	return (total + limit - 1) / limit
}
```

`pageCount` — derives the page count from the total items:

- `if limit <= 0` — guard: if limit is 0 or negative, return **0** (avoids a zero/negative split). Prevents division by zero.
- `return (total + limit - 1) / limit` — **ceiling division**. `total + limit - 1` is added first, then `/ limit`.

**Why `+ limit - 1`:** integer division floors (rounds down). `95/10 = 9` (floored), but you actually need 10 pages (95 items, 10 per page). `95 + 10 - 1 = 104`; `104/10 = 10` (floored = 10) → the final partial page is included. For an exact multiple (e.g. `total=100, limit=10`): `(100 + 9)/10 = 109/10 = 10` — correctly 10.

### Line 12

```go
func main() {
```

Program entry point.

### Line 13

```go
totalProducts, limit, page := 95, 10, 3
```

Input: 95 products, 10 per page, the user is on page 3.

### Lines 15–16

```go
pages := pageCount(totalProducts, limit)
offset := (page - 1) * limit
```

- `pages` = `pageCount(95, 10)` = 10 pages.
- `offset` = `(3 - 1) * 10` = 20 — **offset-based pagination:** page 3 starts at item 21. (page 1 → offset 0, page 2 → offset 10, page 3 → offset 20...)

### Lines 18–21

```go
fmt.Println("total pages:", pages)
fmt.Println("page", page, "show items", offset+1, "to", min(offset+limit, totalProducts))
fmt.Println("has next page:", page < pages)
fmt.Println("has prev page:", page > 1)
```

- `total pages: 10`
- `page 3 show items 21 to 30` — `offset+1` = 21 (start), `min(offset+limit, totalProducts)` = `min(20+10, 95)` = `min(30, 95)` = 30 (end). **`min`** ensures the last page doesn't display more than the total (e.g. last page shows 91 to 95).
- `has next page: true` — `3 < 10`.
- `has prev page: true` — `3 > 1`.

### Line 22

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
total pages: 10
page 3 show items 21 to 30
has next page: true
has prev page: true
```

## Key Takeaways

1. **Ceiling division** — `(total + limit - 1) / limit` — counts the final partial page.
2. **Offset pagination** — `offset = (page - 1) * limit` — where each page starts.
3. **`min` guard** — stops the last page from showing more than the total.
4. **Guard clause** — `limit <= 0` → 0 (avoids division by zero).
5. **Bounds check** — `page < pages` / `page > 1` for next/prev flags.