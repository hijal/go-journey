# search-query-matcher

Go-তে **`strings.Fields` normalization + `strings.Contains` + case-insensitive AND-match** শেখার ছোট example — keyword search matcher।

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

### Lines 3–6

```go
import (
	"fmt"
	"strings"
)
```

- `fmt` — `Println`।
- `strings` — `Fields`, `ToLower`, `Contains`।

### Lines 8–16

```go
func matches(query, product string) bool {
	product = strings.ToLower(product)
	for _, word := range strings.Fields(strings.ToLower(query)) {
		if !strings.Contains(product, word) {
			return false
		}
	}
	return true
}
```

**Matcher logic:**

- `product = strings.ToLower(product)` — product-কে lowercase (case-insensitive)।
- `strings.Fields(strings.ToLower(query))` — query lowercase + **whitespace-normalize**: `"  wireless   MOUSE "` → `["wireless", "mouse"]` (Fields একাধিক space-ও handle করে)।
- প্রতিটি `word`: `strings.Contains(product, word)` — substring মিলবে না-হলে early `false`।
- সব মিললে `true` — **AND semantics** (all-terms)।

*(Contains = substring; একটা শব্দ product-র মাঝে থাকলেই হয়।)*

### Lines 18–25

```go
func main() {
	products := []string{
		"Wireless Mouse Pro",
		"USB-C Charging Cable",
		"Mechanical Keyboard",
	}

	query := "  wireless   MOUSE "
```

৩টা product + messy query (leading-space + double-space + UPPER)।

### Lines 27–31

```go
	for _, p := range products {
		if matches(query, p) {
			fmt.Println("HIT:", p)
		}
	}
```

- "Wireless Mouse Pro" — wireless + mouse দুটোই → **HIT**।
- বাকিগুলো কোনোটা-ই নাই → না।

---

## Expected Output

```
HIT: Wireless Mouse Pro
```

## মূল শিক্ষা / Key Takeaways

1. **`strings.Fields`** — whitespace-normalize + tokenization।
2. **`strings.ToLower`** — case-insensitive।
3. **`strings.Contains`** — substring match।
4. **Early-return** — না-পাওয়া শব্দে `false`।
5. **AND semantics** — সব শব্দ দরকার।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Lines 3–6

```go
import (
	"fmt"
	"strings"
)
```

- `fmt` — for `Println`.
- `strings` — for `Fields`, `ToLower`, `Contains`.

### Lines 8–16

```go
func matches(query, product string) bool {
	product = strings.ToLower(product)
	for _, word := range strings.Fields(strings.ToLower(query)) {
		if !strings.Contains(product, word) {
			return false
		}
	}
	return true
}
```

**The matcher logic:**

- `product = strings.ToLower(product)` — lowercase the product (case-insensitivity).
- `strings.Fields(strings.ToLower(query))` — lowercase the query + **whitespace-normalize**: `"  wireless   MOUSE "` → `["wireless", "mouse"]` (Fields handles multiple spaces too).
- Each `word`: `strings.Contains(product, word)` — an unmatched substring returns `false` early.
- All matched → `true` — **AND semantics** (all-terms).

*(Contains does a substring check — a word just has to appear somewhere in the product.)*

### Lines 18–25

```go
func main() {
	products := []string{
		"Wireless Mouse Pro",
		"USB-C Charging Cable",
		"Mechanical Keyboard",
	}

	query := "  wireless   MOUSE "
```

3 products + a messy query (leading space, double space, UPPER).

### Lines 27–31

```go
	for _, p := range products {
		if matches(query, p) {
			fmt.Println("HIT:", p)
		}
	}
```

- "Wireless Mouse Pro" — both wireless + mouse present → **HIT**.
- The others contain none → skipped.

---

## Expected Output

```
HIT: Wireless Mouse Pro
```

## Key Takeaways

1. **`strings.Fields`** — whitespace-normalize + tokenize.
2. **`strings.ToLower`** — case-insensitivity.
3. **`strings.Contains`** — substring matching.
4. **Early-return** — a missing word returns `false`.
5. **AND semantics** — every word required.