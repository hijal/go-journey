# sql-clause-builder

Go-তে **variadic `...any`**, `strings.Repeat`-type placeholder-building দিয়ে SQL `IN` clause safe-ভাবে build শেখার ছোট example — placeholders (`?`) আর arguments-কে একসাথে return করা।

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

- `fmt` — `Sprintf`, `Println`।
- `strings` — `Join`।

### Lines 8–20

```go
func BuilderClause(column string, values ...any) (string, []any) {
	if len(values) == 0 {
		return "", nil
	}

	placeHolders := make([]string, len(values))
	for i := range values {
		placeHolders[i] = "?"
	}
	query := fmt.Sprintf("%s IN (%s)", column, strings.Join(placeHolders, ", "))
	return query, values
}
```

`BuilderClause` — একটা SQL `IN` clause build করে:

- `column string` — column-এর নাম।
- `values ...any` — **variadic**: IN-এর জন্য মানগুলো (`101, 102, 103`)। `any`-টাইপ — int, string, ইত্যাদি কোনো-টাইপ।
- **Multiple return** `(string, []any)` — clause string + args slice — prepared statement ব্যবহারের জন্য।
- `if len(values) == 0` — **guard**: value না থাকলে খালি string + nil args return (invalid `() IN ()` এড়ানো)।
- `placeHolders := make([]string, len(values))` — `?` placeholder-এর slice (per-value)।
- Loop-এ প্রতিটা `placeHolders[i] = "?"` — একটা `?` per parameter। **এরপর** `strings.Join(placeHolders, ", ")` = `?, ?, ?` — values count-এর সাথে match।
- `fmt.Sprintf("%s IN (%s)", column, placeholders)` — `user_id IN (?, ?, ?)`।
- `return query, values` — prepared statement-এ query + args-দুটোই return।

**নিরাপত্তা point:** মানগুলো clause-এ সরাসরি embed না — শুধু `?` placeholder; values আলাদা `[]any`-তে। এটা **SQL injection** থেকে বাঁচায় (prepared statement: মান কখনো query-string-এ string-interpolation হয় না)।

### Line 22

```go
func main() {
```

Program-এর entry point।

### Lines 23–25

```go
query, args := BuilderClause("user_id", 101, 102, 103)
fmt.Println("Query:", query)
fmt.Println("Args:", args)
```

- `BuilderClause("user_id", 101, 102, 103)` — column `user_id`, ৩টা value।
- `query` = `user_id IN (?, ?, ?)`, `args` = `[101 102 103]`।
- Output:
  - `Query: user_id IN (?, ?, ?)`
  - `Args: [101 102 103]`

### Line 26

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
Query: user_id IN (?, ?, ?)
Args: [101 102 103]
```

## মূল শিক্ষা / Key Takeaways

1. **Variadic `...any`** — IN-এর জন্য যেকোনো-সংখ্যক, যেকোনো-টাইপ মান।
2. **Multiple return** — clause + args-দুইটাই return — prepared statement-এর জন্য।
3. **Placeholder per value** — প্রতিটা মানের জন্য একটা `?` (count-match)।
4. **SQL-injection safety** — মান embed না করে, placeholder + args আলাদা।
5. **Guard clause** — খালি input → empty clause (invalid SQL এড়ানো)।

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

- `fmt` — for `Sprintf`, `Println`.
- `strings` — for `Join`.

### Lines 8–20

```go
func BuilderClause(column string, values ...any) (string, []any) {
	if len(values) == 0 {
		return "", nil
	}

	placeHolders := make([]string, len(values))
	for i := range values {
		placeHolders[i] = "?"
	}
	query := fmt.Sprintf("%s IN (%s)", column, strings.Join(placeHolders, ", "))
	return query, values
}
```

`BuilderClause` — builds a SQL `IN` clause:

- `column string` — the column name.
- `values ...any` — **variadic**: the values for IN (`101, 102, 103`). Type `any` — int, string, etc., any type.
- **Multiple return** `(string, []any)` — the clause string + the args slice — for use with prepared statements.
- `if len(values) == 0` — **guard**: no values → return an empty string + nil args (avoids invalid `() IN ()`).
- `placeHolders := make([]string, len(values))` — a slice for the `?` placeholders (one per value).
- The loop sets each `placeHolders[i] = "?"` — one `?` per parameter. **Then** `strings.Join(placeHolders, ", ")` = `?, ?, ?` — matching the values count.
- `fmt.Sprintf("%s IN (%s)", column, placeholders)` — `user_id IN (?, ?, ?)`.
- `return query, values` — returns both the query and args for the prepared statement.

**Security point:** the values aren't embedded directly in the clause — only `?` placeholders; the values go in a separate `[]any`. This protects against **SQL injection** (with prepared statements, values are never string-interpolated into the query).

### Line 22

```go
func main() {
```

Program entry point.

### Lines 23–25

```go
query, args := BuilderClause("user_id", 101, 102, 103)
fmt.Println("Query:", query)
fmt.Println("Args:", args)
```

- `BuilderClause("user_id", 101, 102, 103)` — column `user_id`, 3 values.
- `query` = `user_id IN (?, ?, ?)`, `args` = `[101 102 103]`.
- Output:
  - `Query: user_id IN (?, ?, ?)`
  - `Args: [101 102 103]`

### Line 26

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
Query: user_id IN (?, ?, ?)
Args: [101 102 103]
```

## Key Takeaways

1. **Variadic `...any`** — any number and type of values for IN.
2. **Multiple return** — both clause and args returned — for prepared statements.
3. **Placeholder per value** — one `?` per value (matching counts).
4. **SQL-injection safety** — values aren't embedded; placeholders + separate args.
5. **Guard clause** — empty input → empty clause (avoids invalid SQL).