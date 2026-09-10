# structured-logging-variadic

Go-তে **variadic `...any`** দিয়ে key-value pairs-এ structured logging শেখার ছোট example — odd-argument warning, RFC3339 timestamp, context-data build।

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
	"strings"
	"time"
)
```

- `fmt` — `Sprintf`, `Printf`।
- `strings` — `Join`, `ToUpper`।
- `time` — `time.Now()`, formatting।

### Lines 9–28

```go
func Log(level string, message string, keysAndValues ...any) {
	if len(keysAndValues)%2 != 0 {
		fmt.Printf("warning: odd number of key-value arguments passed to logger")
	}

	timestamp := time.Now().Format(time.RFC3339)
	var contextData []string

	for i := 0; i < len(keysAndValues)-1; i += 2 {
		contextData = append(contextData, fmt.Sprintf("%v=%v", keysAndValues[i], keysAndValues[i+1]))
	}

	contextStr := ""

	if len(contextData) > 0 {
		contextStr = " | " + strings.Join(contextData, ", ")
	}

	fmt.Printf("[%s] [%s] %s%s\n", timestamp, strings.ToUpper(level), message, contextStr)
}
```

`Log` — structured logger:

- Parameters: `level`, `message`, `keysAndValues ...any` — **variadic** key-value pairs (alternating key, value)।
- **Odd-count guard:** `len(keysAndValues)%2 != 0` — জোড় না হলে warning print (একটা orphan key-value জোড়া ছাড়া)।
- `time.Now().Format(time.RFC3339)` — current time-কে RFC3339 format (`2026-09-10T10:20:48+06:00`)-এ।
- Loop `i := 0; i < len(...)-1; i += 2` — **pairs**: `(0,1), (2,3), ...`। প্রতিটা জোড়ায় `Sprintf("%v=%v", key, value)` → `host=localhost`।
- `strings.Join(contextData, ", ")` — সব pairs এক string-এ comma-দিয়ে।
- Empty হলে `""` (no context); else prepend `" | "`।
- Print: `[timestamp] [LEVEL] message | k=v, k2=v2`। `ToUpper(level)` level uppercase।

**Structured-data value:** মান-গুলো key-value-তে — grep/parse সহজ। Variadic-ভিত্তিতে caller-কে struct/map বানাতে হয় না, ফ্ল্যাট pairs pass করতে হয়।

### Lines 30–32

```go
func main() {
	Log("info", "server started")
	Log("error", "database connection failed", "host", "localhost", "port", "5432", "retry", "true")
}
```

দুটো log:

- `Log("info", "server started")` — no context → `[timestamp] [INFO] server started`
- `Log("error", "database connection failed", "host", "localhost", "port", "5432", "retry", "true")` — ৩টা key-value pair → `host=localhost, port=5432, retry=true`

---

## Expected Output

```
[2026-09-10T10:20:48+06:00] [INFO] server started
[2026-09-10T10:20:48+06:00] [ERROR] database connection failed | host=localhost, port=5432, retry=true
```

> Timestamp runtime অনুযায়ী বদলাবে — format একই।

## মূল শিক্ষা / Key Takeaways

1. **Variadic `...any`** — key-value pairs flat-ভাবে pass।
2. **Odd-count guard** — জোড়-সংখ্যক argument নিশ্চিত।
3. **`time.RFC3339`** — standardized timestamp format।
4. **Pair iteration** — `i += 2` stepping।
5. **`%v=%v` build** — key-value transformer।

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
	"strings"
	"time"
)
```

- `fmt` — for `Sprintf`, `Printf`.
- `strings` — for `Join`, `ToUpper`.
- `time` — for `time.Now()`, formatting.

### Lines 9–28

```go
func Log(level string, message string, keysAndValues ...any) {
	if len(keysAndValues)%2 != 0 {
		fmt.Printf("warning: odd number of key-value arguments passed to logger")
	}

	timestamp := time.Now().Format(time.RFC3339)
	var contextData []string

	for i := 0; i < len(keysAndValues)-1; i += 2 {
		contextData = append(contextData, fmt.Sprintf("%v=%v", keysAndValues[i], keysAndValues[i+1]))
	}

	contextStr := ""

	if len(contextData) > 0 {
		contextStr = " | " + strings.Join(contextData, ", ")
	}

	fmt.Printf("[%s] [%s] %s%s\n", timestamp, strings.ToUpper(level), message, contextStr)
}
```

`Log` — a structured logger:

- Parameters: `level`, `message`, `keysAndValues ...any` — **variadic** key-value pairs (alternating key, value).
- **Odd-count guard:** if `len(keysAndValues)%2 != 0`, print a warning (one orphan pair without its counterpart).
- `time.Now().Format(time.RFC3339)` — the current time in RFC3339 (`2026-09-10T10:20:48+06:00`).
- Loop `i := 0; i < len(...)-1; i += 2` — **pairs**: `(0,1), (2,3), ...`. Each pair becomes `Sprintf("%v=%v", key, value)` → `host=localhost`.
- `strings.Join(contextData, ", ")` — all pairs into one comma-separated string.
- Empty → `""` (no context); otherwise prepend `" | "`.
- Print: `[timestamp] [LEVEL] message | k=v, k2=v2`. `ToUpper(level)` uppercases the level.

**Value of structured data:** the context is key-value, easy to grep/parse. With variadic args, callers don't need a struct/map — they just pass flat pairs.

### Lines 30–32

```go
func main() {
	Log("info", "server started")
	Log("error", "database connection failed", "host", "localhost", "port", "5432", "retry", "true")
}
```

Two logs:

- `Log("info", "server started")` — no context → `[timestamp] [INFO] server started`
- `Log("error", "database connection failed", "host", "localhost", "port", "5432", "retry", "true")` — 3 key-value pairs → `host=localhost, port=5432, retry=true`

---

## Expected Output

```
[2026-09-10T10:20:48+06:00] [INFO] server started
[2026-09-10T10:20:48+06:00] [ERROR] database connection failed | host=localhost, port=5432, retry=true
```

> The timestamp will vary by run time — the format stays the same.

## Key Takeaways

1. **Variadic `...any`** — pass key-value pairs flat.
2. **Odd-count guard** — ensure an even number of arguments.
3. **`time.RFC3339`** — a standardized timestamp format.
4. **Pair iteration** — stepping by 2 (`i += 2`).
5. **`%v=%v` build** — key-value transformer.