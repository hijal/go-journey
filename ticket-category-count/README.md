# ticket-category-count

Go-তে **map frequency-count (`counts[c]++` zero-value idiom)** শেখার ছোট example — ticket category-ভিত্তিক গোনা।

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

`fmt` — `Printf`, `Println`।

### Lines 5–9

```go
func main() {
	categories := []string{
		"billing", "shipping", "billing", "refund",
		"shipping", "billing", "login", "refund", "refund",
	}
```

9টা ticket — categories: billing ×3, shipping ×2, refund ×3, login ×1।

### Line 11

```go
	counts := make(map[string]int)
```

**Frequency map** — `category → count`।

### Lines 13–15

```go
	for _, c := range categories {
		counts[c]++
	}
```

**Count idiom:**

- Missing-key-তে map value হল **zero** (`0`) — তাই `counts[c]++` সরাসরি বাড়ায় (set/guard লাগে না)।
- billing → 3, shipping → 2, refund → 3, login → 1।

### Lines 17–21

```go
	fmt.Println("Ticket counts by category")

	for cat, n := range counts {
		fmt.Printf("  %-10s %d\n", cat, n)
	}
```

- `for cat, n := range counts` — key + value।
- `%-10s` — **left-aligned width-10** — কলাম-সারিবদ্ধ।
- *(নোট: map iteration order random — print-order প্রতি run-এ বদলাতে পারে।)*

---

## Expected Output

(ক্যাটাগরি-order random — সংখ্যা সঠিক)

```
Ticket counts by category
  refund     3
  login      1
  billing    3
  shipping   2
```

## মূল শিক্ষা / Key Takeaways

1. **Zero-value count** — `counts[c]++` বিনা-guard।
2. **`make(map[string]int)`** — initialized map।
3. **Key+value range** — `for cat, n := range counts`।
4. **`%-10s` alignment** — column layout।
5. **Random map order** — deterministic চাইলে keys sort।

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

`fmt` — for `Printf`, `Println`.

### Lines 5–9

```go
func main() {
	categories := []string{
		"billing", "shipping", "billing", "refund",
		"shipping", "billing", "login", "refund", "refund",
	}
```

9 tickets — categories: billing ×3, shipping ×2, refund ×3, login ×1.

### Line 11

```go
	counts := make(map[string]int)
```

**A frequency map** — `category → count`.

### Lines 13–15

```go
	for _, c := range categories {
		counts[c]++
	}
```

**The count idiom:**

- A missing key's map value is **zero** (`0`) — so `counts[c]++` increments directly (no guard/set needed).
- billing → 3, shipping → 2, refund → 3, login → 1.

### Lines 17–21

```go
	fmt.Println("Ticket counts by category")

	for cat, n := range counts {
		fmt.Printf("  %-10s %d\n", cat, n)
	}
```

- `for cat, n := range counts` — key + value.
- `%-10s` — **left-aligned width-10** — column alignment.
- *(Note: map iteration order is random — the print order may vary per run.)*

---

## Expected Output

(category order is random — the counts are exact)

```
Ticket counts by category
  refund     3
  login      1
  billing    3
  shipping   2
```

## Key Takeaways

1. **Zero-value counting** — `counts[c]++` with no guard.
2. **`make(map[string]int)`** — an initialized map.
3. **Key+value range** — `for cat, n := range counts`.
4. **`%-10s` alignment** — column layout.
5. **Random map order** — sort the keys for determinism.