# key-value-config-parser

Go-তে **`strings.Split` + `TrimSpace` + `HasPrefix` + `Cut` mini-INI parser** শেখার ছোট example।

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

- `fmt` — `Printf`।
- `strings` — `Split`, `TrimSpace`, `HasPrefix`, `Cut`।

### Lines 8–13

```go
func main() {
	rawConfig := `
# app settings
server_port = 8080
  env = production 
`
```

**Raw multiline config** (backtick string):

- খালি লাইন, comment (`# ...`), এবং ইন্ডেন্ট (leading `  `) ও trailing-space (production ` `) সহ config — realist text।

### Lines 14–19

```go
	for _, line := range strings.Split(rawConfig, "\n") {
		line = strings.TrimSpace(line)

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
```

**Line-scan + skip-filters:**

- `strings.Split(rawConfig, "\n")` — line-wise।
- `line = strings.TrimSpace(line)` — leading/trailing whitespace (ইন্ডেন্ট, trailing space) ক্লিন।
- `line == ""` খালি → `continue`; `strings.HasPrefix(line, "#")` comment → `continue`।

### Lines 21–27

```go
		key, value, found := strings.Cut(line, "=")

		if !found {
			fmt.Println("invalid line:", line)
			return
		}
		fmt.Printf("%s => %q\n", strings.TrimSpace(key), strings.TrimSpace(value))
	}
```

**Cut — key/value split:**

- `strings.Cut(line, "=")` → প্রথম `=`-এর আগে/পরে + found-flag।
- `!found` — `=`-বিহীন → `invalid line` + `return` (fail-fast)।
- আবার `TrimSpace` key/value-তে (যেমন `  env` → `env`)।
- `%q` — quoted-value display (`"8080"`, `"production"`)।

---

## Expected Output

```
server_port => "8080"
env => "production"
```

## মূল শিক্ষা / Key Takeaways

1. **Multiline raw-string** — `` `...` `` config block।
2. **`TrimSpace`+skip** — blank/comment filtering।
3. **`strings.Cut`** — one-call key/value split + found flag।
4. **Fail-fast** — `=`-নাই → error + `return`।
5. **`%q` display** — quoted output।

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

- `fmt` — for `Printf`.
- `strings` — for `Split`, `TrimSpace`, `HasPrefix`, `Cut`.

### Lines 8–13

```go
func main() {
	rawConfig := `
# app settings
server_port = 8080
  env = production 
`
```

A **raw multiline config** (backtick string):

- An empty line, a comment (`# ...`), plus lines with indentation (leading `  `) and trailing space (production ` `) — realistic text.

### Lines 14–19

```go
	for _, line := range strings.Split(rawConfig, "\n") {
		line = strings.TrimSpace(line)

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
```

**Line-scan + skip filters:**

- `strings.Split(rawConfig, "\n")` — line-wise.
- `line = strings.TrimSpace(line)` — cleans leading/trailing whitespace (indent, trailing space).
- Empty (`line == ""`) → `continue`; comment (`strings.HasPrefix(line, "#")`) → `continue`.

### Lines 21–27

```go
		key, value, found := strings.Cut(line, "=")

		if !found {
			fmt.Println("invalid line:", line)
			return
		}
		fmt.Printf("%s => %q\n", strings.TrimSpace(key), strings.TrimSpace(value))
	}
```

**Cut — the key/value split:**

- `strings.Cut(line, "=")` → everything before/after the first `=` plus a found flag.
- `!found` — no `=` → `invalid line` + `return` (fail-fast).
- `TrimSpace` again on key/value (e.g. `  env` → `env`).
- `%q` — a quoted-value display (`"8080"`, `"production"`).

---

## Expected Output

```
server_port => "8080"
env => "production"
```

## Key Takeaways

1. **Multiline raw-string** — a `` `...` `` config block.
2. **`TrimSpace`+skip** — blank/comment filtering.
3. **`strings.Cut`** — a one-call key/value split + found flag.
4. **Fail-fast** — no `=` → error + `return`.
5. **`%q` display** — quoted output.