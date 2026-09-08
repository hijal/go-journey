# custom-logger

Go-তে **variadic parameter** (`...any`), `fmt.Sprint`, `strings.Join` আর `strings.ToUpper` দিয়ে custom logger function শেখার ছোট example।

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

- `fmt` — `Sprint`, `Printf`।
- `strings` — `Join`, `ToUpper`।

### Lines 8–15

```go
func Log(level string, message ...any) {
	strMsgs := make([]string, len(message))
	for i, msg := range message {
		strMsgs[i] = fmt.Sprint(msg)
	}
	finalMsg := strings.Join(strMsgs, " ")
	fmt.Printf("[%s] %s\n", strings.ToUpper(level), finalMsg)
}
```

`Log` — একটা logger function:

- `level string` — log-এর level (e.g. "info", "error", "debug")।
- `message ...any` — **variadic** parameter: যেকোনো সংখ্যক argument, প্রতিটা `any` (যেকোনো type)।
- `strMsgs := make([]string, len(message))` — same-length string slice।
- loop-এ প্রতিটা `msg`-কে `fmt.Sprint(msg)` দিয়ে **string-এ convert** করা হয় (Sprint string-এর একটা representation return করে)।
- `strings.Join(strMsgs, " ")` — সব string-কে space দিয়ে join।
- `strings.ToUpper(level)` — level-কে **uppercase** করে; `fmt.Printf` দিয়ে `[LEVEL] message` format।

**জেনেরিক input:** `Log("info", "User", 101, "from", ip)`-এর মতো — int, string, map — সব handle হয়, কারণ প্রতিটা `fmt.Sprint`-এ যায়।

### Line 17

```go
func main() {
```

Program-এর entry point।

### Lines 18–20

```go
Log("info", "User", 101, "logged in from", "192.168.1.1")
Log("error", "Failed to connect to database")
Log("debug", "Processing request", "with parameters:", map[string]string{"id": "123", "action": "update"})
```

তিনটা log call — mixing types:

- info: `"User", 101, "logged in from", "192.168.1.1"` → `[INFO] User 101 logged in from 192.168.1.1`
- error: একটা string → `[ERROR] Failed to connect to database`
- debug: string + map → `[DEBUG] Processing request with parameters: map[...]`

### Line 21

```go
fmt.Println("Logging completed.")
```

`Logging completed.` — সব log-এর পর।

### Line 22

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
[INFO] User 101 logged in from 192.168.1.1
[ERROR] Failed to connect to database
[DEBUG] Processing request with parameters: map[action:update id:123]
Logging completed.
```

## মূল শিক্ষা / Key Takeaways

1. **Variadic `...any`** — যেকোনো সংখ্যক, যেকোনো-টাইপ argument।
2. **`fmt.Sprint`** — যে-কোনো value-কে string representation-এ।
3. **`strings.Join`** — slice-কে separator দিয়ে এক string।
4. **`strings.ToUpper`** — formatting-এর জন্য।
5. **Exported `Log`** — Uppercase নাম, অন্য package-এ ব্যবহার-যোগ্য।

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

- `fmt` — for `Sprint`, `Printf`.
- `strings` — for `Join`, `ToUpper`.

### Lines 8–15

```go
func Log(level string, message ...any) {
	strMsgs := make([]string, len(message))
	for i, msg := range message {
		strMsgs[i] = fmt.Sprint(msg)
	}
	finalMsg := strings.Join(strMsgs, " ")
	fmt.Printf("[%s] %s\n", strings.ToUpper(level), finalMsg)
}
```

`Log` — a logger function:

- `level string` — the log level (e.g. "info", "error", "debug").
- `message ...any` — a **variadic** parameter: any number of arguments, each `any` (any type).
- `strMsgs := make([]string, len(message))` — a string slice of the same length.
- The loop converts each `msg` to a **string** with `fmt.Sprint(msg)` (Sprint returns a string representation).
- `strings.Join(strMsgs, " ")` — joins all strings with spaces.
- `strings.ToUpper(level)` — uppercases the level; `fmt.Printf` prints `[LEVEL] message`.

**Generic input:** calls like `Log("info", "User", 101, "from", ip)` — int, string, map — all handled, since each goes through `fmt.Sprint`.

### Line 17

```go
func main() {
```

Program entry point.

### Lines 18–20

```go
Log("info", "User", 101, "logged in from", "192.168.1.1")
Log("error", "Failed to connect to database")
Log("debug", "Processing request", "with parameters:", map[string]string{"id": "123", "action": "update"})
```

Three log calls mixing types:

- info: `"User", 101, "logged in from", "192.168.1.1"` → `[INFO] User 101 logged in from 192.168.1.1`
- error: one string → `[ERROR] Failed to connect to database`
- debug: string + map → `[DEBUG] Processing request with parameters: map[...]`

### Line 21

```go
fmt.Println("Logging completed.")
```

`Logging completed.` — after all logs.

### Line 22

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
[INFO] User 101 logged in from 192.168.1.1
[ERROR] Failed to connect to database
[DEBUG] Processing request with parameters: map[action:update id:123]
Logging completed.
```

## Key Takeaways

1. **Variadic `...any`** — any number of arguments of any type.
2. **`fmt.Sprint`** — any value to its string representation.
3. **`strings.Join`** — a slice into one string with a separator.
4. **`strings.ToUpper`** — for formatting.
5. **Exported `Log`** — uppercase name, usable from other packages.