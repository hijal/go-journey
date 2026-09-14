# sku-validation

Go-তে **staged guard-clause validation + `strings.Split` + `strconv.Atoi` + `fmt.Errorf`** শেখার ছোট example — SKU format validator (fail-fast)।

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
	"fmt"
	"strconv"
	"strings"
)
```

- `fmt` — `Println`, `Errorf`।
- `strconv` — `Atoi`।
- `strings` — `Split`।

### Lines 9–28

```go
func validateSKU(sku string) error {
	parts := strings.Split(sku, "-")
	if len(parts) != 2 {
		return fmt.Errorf("SKU must look like CAT-NUMBER, got %q", sku)
	}

	category, idText := parts[0], parts[1]
	if len(category) != 3 {
		return fmt.Errorf("category must be 3 letters, got %q", category)
	}

	id, err := strconv.Atoi(idText)
	if err != nil {
		return fmt.Errorf("product number %q is not a valid integer", idText)
	}
	if id < 1 {
		return fmt.Errorf("product number must be positive, got %d", id)
	}
	return nil
}
```

**Staged guard-clause (fail-fast) — সব পরীক্ষা early-return:**

- `strings.Split(sku, "-")` → ঠিক দুটো টুকরো না-হলে (`BAD`) reject।
- `len(category) != 3` — ঠিক 3 অক্ষর (`BOOK` 4-হওয়ায় reject)।
- `strconv.Atoi(idText)` — numeric নয় (`abc`) reject।
- `id < 1` — positive-guard।
- প্রতিটি error `fmt.Errorf` + `%q`/`%d` details।

*(এক-একটা শর্তের বদলে single-layer check — error-এ ঠিক কোন স্তরে fail হলো তা স্পষ্ট।)*

### Lines 30–37

```go
func main() {
	for _, sku := range []string{"BOOK-1024", "TOY-7", "BAD", "PEN-abc"} {
		if err := validateSKU(sku); err != nil {
			fmt.Println(sku, "REJECTED:", err)
		} else {
			fmt.Println(sku, "ACCEPTED")
		}
	}
}
```

৪টা নমুনা — ১টা good, ৩টা ভিন্ন-ভাবে bad (structure, length, numeric)।

---

## Expected Output

```
BOOK-1024 REJECTED: category must be 3 letters, got "BOOK"
TOY-7 ACCEPTED
BAD REJECTED: SKU must look like CAT-NUMBER, got "BAD"
PEN-abc REJECTED: product number "abc" is not a valid integer
```

## মূল শিক্ষা / Key Takeaways

1. **Guard-clause early-return** — staged fails।
2. **`strings.Split`** — structural parts।
3. **`strconv.Atoi`** — numeric-string parse + err।
4. **`fmt.Errorf` + `%q`** — informative diagnostic।
5. **Positive-guard** — domain rule (SKU number ≥ 1)।

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
	"fmt"
	"strconv"
	"strings"
)
```

- `fmt` — for `Println`, `Errorf`.
- `strconv` — for `Atoi`.
- `strings` — for `Split`.

### Lines 9–28

```go
func validateSKU(sku string) error {
	parts := strings.Split(sku, "-")
	if len(parts) != 2 {
		return fmt.Errorf("SKU must look like CAT-NUMBER, got %q", sku)
	}

	category, idText := parts[0], parts[1]
	if len(category) != 3 {
		return fmt.Errorf("category must be 3 letters, got %q", category)
	}

	id, err := strconv.Atoi(idText)
	if err != nil {
		return fmt.Errorf("product number %q is not a valid integer", idText)
	}
	if id < 1 {
		return fmt.Errorf("product number must be positive, got %d", id)
	}
	return nil
}
```

**Staged guard clauses (fail-fast) — every check early-returns:**

- `strings.Split(sku, "-")` — an exact 2-part split or reject (`BAD`).
- `len(category) != 3` — exactly 3 letters (`BOOK` has 4 → reject).
- `strconv.Atoi(idText)` — non-numeric (`abc`) rejects.
- `id < 1` — a positive guard.
- Each error via `fmt.Errorf` + `%q`/`%d` details.

*(One-level checks let the caller see exactly which stage failed.)*

### Lines 30–37

```go
func main() {
	for _, sku := range []string{"BOOK-1024", "TOY-7", "BAD", "PEN-abc"} {
		if err := validateSKU(sku); err != nil {
			fmt.Println(sku, "REJECTED:", err)
		} else {
			fmt.Println(sku, "ACCEPTED")
		}
	}
}
```

4 samples — 1 good, 3 bad in different ways (structure, length, numeric).

---

## Expected Output

```
BOOK-1024 REJECTED: category must be 3 letters, got "BOOK"
TOY-7 ACCEPTED
BAD REJECTED: SKU must look like CAT-NUMBER, got "BAD"
PEN-abc REJECTED: product number "abc" is not a valid integer
```

## Key Takeaways

1. **Guard-clause early-return** — staged fails.
2. **`strings.Split`** — structural parts.
3. **`strconv.Atoi`** — numeric-string parsing + err.
4. **`fmt.Errorf` + `%q`** — an informative diagnostic.
5. **Positive-guard** — a domain rule (SKU number ≥ 1).