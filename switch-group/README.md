# switch-group

Go-তে **multi-value `case`** (এক case-এ একাধিক value) দিয়ে log destination routing শেখার ছোট example।

**📖 ভাষা নির্বাচন করুন / Choose language:**

[🇧🇩 বাংলা](#bangla) • [🇬🇧 English](#english)

---

<a name="bangla"></a>

## 🇧🇩 বাংলা সংস্করণ

### Lines 1–3

```go
package main

import "fmt"
```

- `package main` — একটা executable program।
- `fmt` — output print করার জন্য।

### Lines 5–18

```go
func logDestination(level string) string {
    switch level {
    case "DEBUG", "Trace":
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

`logDestination` function — একটা `level` (log severity, string) নেয় এবং একটা destination (string) return করে।

- **Multi-value case:** `case "DEBUG", "Trace"` — এক case-এ একাধিক value কমা দিয়ে। `level` যদি এই value-গুলোর **যেকোনোটা** হয়, তাহলে case match হয়।
  - `"DEBUG"`, `"Trace"` → `local file (rotated daily)`
  - `"INFO"` → `stdout`
  - `"WARN"`, `"ERROR"` → `stdout + alerting pipeline`
  - `"FATAL"` → `stdout + alerting pipeline + pager`
- `default` — কোনো case match না হলে → `unknown level: drop and log a config warning`।

**নোট:** switch-এ `case "DEBUG", "Trace"`-এ ডবল quote-এর ভেতরে `Trace` বড় হাতের `T` — তাই `"Trace"` আর `"trace"` আলাদা (case-sensitive)।

### Line 20

```go
func main() {
```

Program-এর entry point।

### Lines 21–26

```go
fmt.Println("DEBUG   ->", logDestination("DEBUG"))
fmt.Println("INFO    ->", logDestination("INFO"))
fmt.Println("WARN    ->", logDestination("WARN"))
fmt.Println("FATAL   ->", logDestination("FATAL"))
fmt.Println("VERBOSE ->", logDestination("VERBOSE"))
```

পাঁচটা log level-এর destination দেখায়:

1. `"DEBUG"` → `local file (rotated daily)`
2. `"INFO"` → `stdout`
3. `"WARN"` → `stdout + alerting pipeline`
4. `"FATAL"` → `stdout + alerting pipeline + pager`
5. `"VERBOSE"` → কোনো case match নেই → `unknown level: drop and log a config warning`

---

## Expected Output

```
DEBUG   -> local file (rotated daily)
INFO    -> stdout
WARN    -> stdout + alerting pipeline
FATAL   -> stdout + alerting pipeline + pager
VERBOSE -> unknown level: drop and log a config warning
```

## মূল শিক্ষা / Key Takeaways

1. **Multi-value `case`** — `case "A", "B":` — value-গুলোর যেকোনোটা match হলেই চলে।
2. **`return` inside switch** — প্রতিটা case-এ সাথে-সাথে return, fallthrough ছাড়া।
3. **`default`** — কোনো case match না হলে।
4. **Case-sensitivity** — string case-গুলো exact match হয়।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Lines 1–3

```go
package main

import "fmt"
```

- `package main` — an executable program.
- `fmt` — for console output.

### Lines 5–18

```go
func logDestination(level string) string {
    switch level {
    case "DEBUG", "Trace":
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

`logDestination` function — takes a `level` (log severity, string) and returns a destination (string).

- **Multi-value case:** `case "DEBUG", "Trace"` — one case lists multiple values separated by commas. The case matches if `level` equals **any** of those values.
  - `"DEBUG"`, `"Trace"` → `local file (rotated daily)`
  - `"INFO"` → `stdout`
  - `"WARN"`, `"ERROR"` → `stdout + alerting pipeline`
  - `"FATAL"` → `stdout + alerting pipeline + pager`
- `default` — if no case matches → `unknown level: drop and log a config warning`.

**Note:** In `case "DEBUG", "Trace"`, `Trace` has an uppercase `T`, so `"Trace"` and `"trace"` are different (case-sensitive).

### Line 20

```go
func main() {
```

Program entry point.

### Lines 21–26

```go
fmt.Println("DEBUG   ->", logDestination("DEBUG"))
fmt.Println("INFO    ->", logDestination("INFO"))
fmt.Println("WARN    ->", logDestination("WARN"))
fmt.Println("FATAL   ->", logDestination("FATAL"))
fmt.Println("VERBOSE ->", logDestination("VERBOSE"))
```

Shows the destination for five log levels:

1. `"DEBUG"` → `local file (rotated daily)`
2. `"INFO"` → `stdout`
3. `"WARN"` → `stdout + alerting pipeline`
4. `"FATAL"` → `stdout + alerting pipeline + pager`
5. `"VERBOSE"` → no case matches → `unknown level: drop and log a config warning`

---

## Expected Output

```
DEBUG   -> local file (rotated daily)
INFO    -> stdout
WARN    -> stdout + alerting pipeline
FATAL   -> stdout + alerting pipeline + pager
VERBOSE -> unknown level: drop and log a config warning
```

## Key Takeaways

1. **Multi-value `case`** — `case "A", "B":` — runs if any of the values match.
2. **`return` inside switch** — returns immediately per case, no fallthrough.
3. **`default`** — runs when no case matches.
4. **Case-sensitivity** — string cases match exactly.
