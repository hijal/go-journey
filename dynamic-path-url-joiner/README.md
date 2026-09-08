# dynamic-path-url-joiner

Go-তে **variadic parameter** (`...string`), `strings.Trim`, `strings.Join` দিয়ে path-গুলোর clean join শেখার ছোট example — URL path-এর ঝামেলা (leading/trailing slash) ঠিক করা।

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
- `strings` — `Trim`, `Join`।

### Lines 8–17

```go
func JoinPath(parts ...string) string {
	var cleanParts []string

	for _, p := range parts {
		if p != "" {
			cleanParts = append(cleanParts, strings.Trim(p, "/"))
		}
	}
	return "/" + strings.Join(cleanParts, "/")
}
```

`JoinPath` — path parts-কে একটা clean slash-separated path-এ join:

- `parts ...string` — **variadic**: একাধিক string argument।
- `var cleanParts []string` — clean-করা parts-এর slice।
- Loop-এ প্রতিটা `p`:
  - `if p != ""` — **empty-skip**: খালি part বাদ দেওয়া হয় (double slash এড়াতে)।
  - `cleanParts = append(cleanParts, strings.Trim(p, "/"))` — **`strings.Trim(p, "/")`** p-এর starting ও trailing slash remove করে। এটা `TrimLeft`/`TrimRight` উভয়ই — দুদিকের `/` মুছে দেয়।
- `strings.Join(cleanParts, "/")` — clean parts-কে single `/` দিয়ে join।
- `"/" + ...` — leading slash prefix যোগ করে।

**Result** — input-এ slash থাকলেও বা না থাকলেও একটা clean absolute path।

### Line 19

```go
func main() {
```

Program-এর entry point।

### Line 20

```go
fmt.Println(JoinPath("/api/", "v1/", "/users", "101/"))
```

`JoinPath("/api/", "v1/", "/users", "101/")`:

- `/api/` → Trim → `api`
- `v1/` → Trim → `v1`
- `/users` → Trim → `users`
- `101/` → Trim → `101`
- Join: `"api/v1/users/101"`, prefix `"/"` → `/api/v1/users/101`

**See:** বাড়তি slashes (`//`, `/v1/`, `101/`) সব clean হয়ে গেছে — একটা সুন্দর স্বাভাবিক path।

### Line 21

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
/api/v1/users/101
```

## মূল শিক্ষা / Key Takeaways

1. **Variadic `...string`** — একাধিক path part argument।
2. **`strings.Trim(p, "/")`** — উভয় দিকের separator remove।
3. **Empty-skip** — `p != ""` খালি part বাদ দেয় (double slash আটকায়)।
4. **`strings.Join`** — part-গুলো single separator দিয়ে।
5. **Path normalization** — user input-এর slash ঝামেলা clean করা।

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
- `strings` — for `Trim`, `Join`.

### Lines 8–17

```go
func JoinPath(parts ...string) string {
	var cleanParts []string

	for _, p := range parts {
		if p != "" {
			cleanParts = append(cleanParts, strings.Trim(p, "/"))
		}
	}
	return "/" + strings.Join(cleanParts, "/")
}
```

`JoinPath` — joins path parts into one clean slash-separated path:

- `parts ...string` — **variadic**: multiple string arguments.
- `var cleanParts []string` — a slice for the cleaned parts.
- The loop over each `p`:
  - `if p != ""` — **empty-skip**: empty parts are dropped (avoids double slashes).
  - `cleanParts = append(cleanParts, strings.Trim(p, "/"))` — **`strings.Trim(p, "/")`** removes leading and trailing slashes from `p`. It covers both sides (like TrimLeft + TrimRight).
- `strings.Join(cleanParts, "/")` — joins the clean parts with a single `/`.
- `"/" + ...` — prefixes a leading slash.

**Result** — a clean absolute path regardless of stray slashes in the input.

### Line 19

```go
func main() {
```

Program entry point.

### Line 20

```go
fmt.Println(JoinPath("/api/", "v1/", "/users", "101/"))
```

`JoinPath("/api/", "v1/", "/users", "101/")`:

- `/api/` → Trim → `api`
- `v1/` → Trim → `v1`
- `/users` → Trim → `users`
- `101/` → Trim → `101`
- Join: `"api/v1/users/101"`, prefixed `"/"` → `/api/v1/users/101`

**See:** all the stray slashes (`//`, `/v1/`, `101/`) are cleaned into a single well-formed path.

### Line 21

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
/api/v1/users/101
```

## Key Takeaways

1. **Variadic `...string`** — multiple path-part arguments.
2. **`strings.Trim(p, "/")`** — removes the separator from both sides.
3. **Empty-skip** — `p != ""` drops empty parts (prevents double slashes).
4. **`strings.Join`** — parts joined with a single separator.
5. **Path normalization** — cleaning messy user-supplied slash input.