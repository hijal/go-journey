# test-log-router

Go-তে **multi-value `case`** দিয়ে log destination routing আর **table-driven unit test** শেখার ছোট example। এতে `main.go` আর `main_test.go` আছে।

**চালানো:** `go run .` • `go test -v`

**📖 ভাষা নির্বাচন করুন / Choose language:**

[🇧🇩 বাংলা](#bangla) • [🇬🇧 English](#english)

---

<a name="bangla"></a>

## 🇧🇩 বাংলা সংস্করণ

### `main.go`

#### Lines 1–3

```go
package main

import "fmt"
```

- `package main` — একটা executable program। test file-টাও একই package-এ।
- `fmt` — output print করার জন্য।

#### Lines 5–18

```go
func logDestination(level string) string {
    switch level {
    case "DEBUG", "TRACE":
        return "local file (rotated daily)"
    case "INFO":
        return "stdout"
    case "WARN", "ERROR":
        return "stdout + alerting pipeline"
    case "FATAL":
        return "stdout + alerting pipeline + pager"
    default:
        return "unknown level: drop and log a config warning"
    }
}
```

`logDestination` function — একটা log severity (`level`) নেয় এবং destination (string) return করে:

- **Multi-value case:** `case "DEBUG", "TRACE":` — এক case-এ একাধিক value; level এইগুলোর যেকোনোটা হলে match।
  - `"DEBUG"`, `"TRACE"` → `local file (rotated daily)`
  - `"INFO"` → `stdout`
  - `"WARN"`, `"ERROR"` → `stdout + alerting pipeline`
  - `"FATAL"` → `stdout + alerting pipeline + pager`
- `default` → `unknown level: drop and log a config warning`।

#### Lines 20–27

```go
func main() {
    fmt.Println(logDestination("INFO"))
    fmt.Println(logDestination("DEBUG"))
    fmt.Println(logDestination("WARN"))
    fmt.Println(logDestination("ERROR"))
    fmt.Println(logDestination("FATAL"))
    fmt.Println(logDestination("VERBOSE"))
}
```

ছয়টা log level-এর destination print করে (প্রতিটা mapping দেখানোর জন্য)।

---

### `main_test.go`

#### Lines 1–3

```go
package main

import "testing"
```

- `package main` — একই package।
- `testing` — Go-র testing framework।

#### Line 5

```go
func TestLogDestination(t *testing.T) {
```

একটা **test function** — `Test` prefix, `*testing.T` argument। `go test` এটা চালায়।

#### Lines 6–17

```go
tests := []struct {
    name  string
    level string
    want  string
}{
    {name: "debug goes to local file", level: "DEBUG", want: "local file (rotated daily)"},
    ...
}
```

**Table-driven test** — একটা anonymous struct slice, প্রতিটা row-এ: `name` (case label), `level` (input), `want` (expected output)। main.go-র প্রতিটা case-এর mapping-টা এখানে যাচাই করা হয়।

#### Line 19

```go
for _, tc := range tests {
```

প্রতিটা test case-এ loop।

#### Line 20

```go
t.Run(tc.name, func(t *testing.T) {
```

`t.Run` — প্রতিটা case আলাদা **subtest** হিসেবে (named)। fail হলে কোনটা ভেঙেছে স্পষ্ট।

#### Lines 21–24

```go
got := logDestination(tc.level)
if got != tc.want {
    t.Errorf("logDestination(%q) = %q, want %q", tc.level, got, tc.want)
}
```

- `got := logDestination(tc.level)` — actual result।
- `got != tc.want` হলে `t.Errorf` দিয়ে failure report (input, actual, expected `%q` linked)।

---

## Expected Output

`go run .`:

```
stdout
local file (rotated daily)
stdout + alerting pipeline
stdout + alerting pipeline
stdout + alerting pipeline + pager
unknown level: drop and log a config warning
```

`go test -v`: সব subtests PASS।

## মূল শিক্ষা / Key Takeaways

1. **Multi-value `case`** — `case "DEBUG", "TRACE":` — multiple values, যেকোনোটা match হলেই চলে।
2. **`return` inside switch** — প্রতি case-এ immediate return, fallthrough ছাড়া।
3. **`default`** — unmatched হলে।
4. **Table-driven test** — struct slice দিয়ে অনেক case যাচাই।
5. **`t.Run` subtests + `t.Errorf`** — যাচাই ও diagnostic।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### `main.go`

#### Lines 1–3

```go
package main

import "fmt"
```

- `package main` — an executable program. The test file is in the same package.
- `fmt` — for console output.

#### Lines 5–18

```go
func logDestination(level string) string {
    switch level {
    case "DEBUG", "TRACE":
        return "local file (rotated daily)"
    case "INFO":
        return "stdout"
    case "WARN", "ERROR":
        return "stdout + alerting pipeline"
    case "FATAL":
        return "stdout + alerting pipeline + pager"
    default:
        return "unknown level: drop and log a config warning"
    }
}
```

`logDestination` function — takes a log severity (`level`) and returns a destination (string):

- **Multi-value case:** `case "DEBUG", "TRACE":` — one case lists multiple values; it matches if `level` is any of them.
  - `"DEBUG"`, `"TRACE"` → `local file (rotated daily)`
  - `"INFO"` → `stdout`
  - `"WARN"`, `"ERROR"` → `stdout + alerting pipeline`
  - `"FATAL"` → `stdout + alerting pipeline + pager`
- `default` → `unknown level: drop and log a config warning`.

#### Lines 20–27

```go
func main() {
    fmt.Println(logDestination("INFO"))
    fmt.Println(logDestination("DEBUG"))
    fmt.Println(logDestination("WARN"))
    fmt.Println(logDestination("ERROR"))
    fmt.Println(logDestination("FATAL"))
    fmt.Println(logDestination("VERBOSE"))
}
```

Prints the destination for six log levels (to show each mapping).

---

### `main_test.go`

#### Lines 1–3

```go
package main

import "testing"
```

- `package main` — the same package.
- `testing` — Go's testing framework.

#### Line 5

```go
func TestLogDestination(t *testing.T) {
```

A **test function** — `Test` prefix, `*testing.T` argument. `go test` runs it.

#### Lines 6–17

```go
tests := []struct {
    name  string
    level string
    want  string
}{
    {name: "debug goes to local file", level: "DEBUG", want: "local file (rotated daily)"},
    ...
}
```

**Table-driven test** — an anonymous struct slice; each row has `name` (case label), `level` (input), and `want` (expected output). The mappings from main.go are verified here.

#### Line 19

```go
for _, tc := range tests {
```

Loops over each test case.

#### Line 20

```go
t.Run(tc.name, func(t *testing.T) {
```

`t.Run` — each case as its own named **subtest**. If one fails, it's clear which one broke.

#### Lines 21–24

```go
got := logDestination(tc.level)
if got != tc.want {
    t.Errorf("logDestination(%q) = %q, want %q", tc.level, got, tc.want)
}
```

- `got := logDestination(tc.level)` — the actual result.
- If `got != tc.want`, report a failure with `t.Errorf` (input, actual, expected — `%q` quoted).

---

## Expected Output

`go run .`:

```
stdout
local file (rotated daily)
stdout + alerting pipeline
stdout + alerting pipeline
stdout + alerting pipeline + pager
unknown level: drop and log a config warning
```

`go test -v`: all subtests PASS.

## Key Takeaways

1. **Multi-value `case`** — `case "DEBUG", "TRACE":` — runs if any value matches.
2. **`return` inside switch** — immediate return per case, no fallthrough.
3. **`default`** — for unmatched levels.
4. **Table-driven test** — verifying many cases with a struct slice.
5. **`t.Run` subtests + `t.Errorf`** — verification and diagnostics.
