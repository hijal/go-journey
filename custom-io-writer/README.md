# custom-io-writer

Go-তে **`io.Writer` interface implementation + custom `Write(p []byte)` method** শেখার ছোট example — log collector।

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
	"io"
)
```

- `fmt` — `Println`, `Printf`।
- `io` — `Writer` interface।

### Lines 8–10

```go
type LogCollector struct {
	lines []string
}
```

Receiver type — accumulated log lines।

### Lines 12–15

```go
func (l *LogCollector) Write(p []byte) (int, error) {
	l.lines = append(l.lines, string(p))
	return len(p), nil
}
```

**`io.Writer` contract** — `Write([]byte) (int, error)`:

- `[]byte`-কে `string` করে append।
- `len(p)` return (সব বাইট গ্রহণ করেছি)।

*(Standard library: `fmt.Fprintln(w, ...)` auto-pass `[]byte` including trailing `\n`।)*

### Lines 17–26

```go
	func main() {
		collector := &LogCollector{}

		var w io.Writer = collector
		fmt.Fprintln(w, "server started")
		fmt.Fprintln(w, "connection accepted")

		for i, line := range collector.lines {
			fmt.Printf("log[%d]: %s", i, line)
		}
	}
```

- `var w io.Writer = collector` — static check: `LogCollector` satisfies `Writer`।
- `fmt.Fprintln` + `fmt.Fprintf` সব মানক `io.Writer` accept করে — custom collector plug-in।
- `%s` prints raw line (trailing `\n` — visual newline)।

---

## Expected Output

```
log[0]: server started
log[1]: connection accepted
```

## মূল শিক্ষা / Key Takeaways

1. **`io.Writer`** — একটা method interface।
2. **`Write([]byte) (int, error)`** — contract (byte-slice + return-কত।)।
3. **Static satisfaction** — `var w io.Writer = collector` check।
4. **Library interop** — `fmt.Fprintln(w, ...)` custom Writer-এ pass।

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
	"io"
)
```

- `fmt` — for `Println`, `Printf`.
- `io` — for the `Writer` interface.

### Lines 8–10

```go
type LogCollector struct {
	lines []string
}
```

The receiver type — accumulates log lines.

### Lines 12–15

```go
func (l *LogCollector) Write(p []byte) (int, error) {
	l.lines = append(l.lines, string(p))
	return len(p), nil
}
```

**The `io.Writer` contract** — `Write([]byte) (int, error)`:

- Converts `[]byte` to `string` and appends.
- Returns `len(p)` (all bytes accepted).

*(Standard library functions like `fmt.Fprintln` auto-pass `[]byte` including a trailing `\n`.)*

### Lines 17–26

```go
	func main() {
		collector := &LogCollector{}

		var w io.Writer = collector
		fmt.Fprintln(w, "server started")
		fmt.Fprintln(w, "connection accepted")

		for i, line := range collector.lines {
			fmt.Printf("log[%d]: %s", i, line)
		}
	}
```

- `var w io.Writer = collector` — a static check that `LogCollector` satisfies `Writer`.
- `fmt.Fprintln`, `fmt.Fprintf` — all accept a `io.Writer`; plug in the custom collector.
- `%s` prints the raw line (the trailing `\n` provides a visual newline).

---

## Expected Output

```
log[0]: server started
log[1]: connection accepted
```

## Key Takeaways

1. **`io.Writer`** — a one-method interface.
2. **`Write([]byte) (int, error)`** — the contract (byte-slice + return count).
3. **Static satisfaction** — `var w io.Writer = collector` as a compile-time check.
4. **Library interop** — `fmt.Fprintln(w, ...)` works with any `io.Writer`.