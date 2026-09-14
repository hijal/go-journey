# wrap-text-exercise

Go-তে **greedy word-wrap + `strings.Fields` + `strings.Builder` + line-width tracking** শেখার ছোট example — text wrapping।

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
- `strings` — `Fields`, `Builder`।

### Lines 8–28

```go
func wrapText(text string, width int) string {
	var b strings.Builder

	lineLen := 0

	for _, word := range strings.Fields(text) {
		switch {
		case lineLen == 0:
			b.WriteString(word)
			lineLen = len(word)
		case lineLen+1+len(word) <= width:
			b.WriteString(" " + word)
			lineLen += 1 + len(word)
		default:
			b.WriteString("\n" + word)
			lineLen = len(word)
		}
	}

	return b.String()
}
```

**Word-wrap logic (greedy first-fit):**

- `strings.Fields(text)` — text-কে whitespace-tokenize (multiple-space normalized)।
- `lineLen` — বর্তমান লাইনে কতটুকু জায়গা নেওয়া হয়েছে।
- তিনটি case:
  - `lineLen == 0` — লাইন খালি, word সরাসরি বসাও।
  - `lineLen+1+len(word) ≤ width` — এক space + word বসালে পুরো লাইন width-এ ঢুকবে → একসাথে।
  - `default` — space সহ ঢুকবে না → নতুন লাইন (`\n`) দিয়ে শুরু।

*(`strings.Builder` — string concat এ loop-এ efficient; `+` বারবার নতুন string allocate করে।)*

### Lines 30–32

```go
func main() {
	fmt.Println(wrapText("go is simple but strings are deep", 12))
}
```

Width 12-তে ৭টা শব্দ:

- "go is simple" (11) → "but" যোগ করলে 15 → newline
- "but strings" (11) → "are" 15 → newline
- "are deep" (8)

---

## Expected Output

```
go is simple
but strings
are deep
```

## মূল শিক্ষা / Key Takeaways

1. **Greedy first-fit wrap** — width-এ ঢুকবে না-হলেই break।
2. **`strings.Fields`** — multi-space normalize + tokenize।
3. **`strings.Builder`** — efficient string concat।
4. **`lineLen` tracking** — বর্তমান line width মনে রাখা।

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
- `strings` — for `Fields`, `Builder`.

### Lines 8–28

```go
func wrapText(text string, width int) string {
	var b strings.Builder

	lineLen := 0

	for _, word := range strings.Fields(text) {
		switch {
		case lineLen == 0:
			b.WriteString(word)
			lineLen = len(word)
		case lineLen+1+len(word) <= width:
			b.WriteString(" " + word)
			lineLen += 1 + len(word)
		default:
			b.WriteString("\n" + word)
			lineLen = len(word)
		}
	}

	return b.String()
}
```

**The word-wrap logic (greedy first-fit):**

- `strings.Fields(text)` — tokenize by whitespace (multi-space normalized).
- `lineLen` — how much of the current line is already filled.
- Three cases:
  - `lineLen == 0` — fresh line, write the word directly.
  - `lineLen+1+len(word) ≤ width` — a space + word fit within width → place on the same line.
  - `default` — won't fit → break to a new line (`\n`).

*(`strings.Builder` — efficient string concatenation in a loop; `+` allocates a new string each time.)*

### Lines 30–32

```go
func main() {
	fmt.Println(wrapText("go is simple but strings are deep", 12))
}
```

With width 12 and 7 words:

- "go is simple" (11) → adding "but" makes 15 → newline
- "but strings" (11) → "are" is 15 → newline
- "are deep" (8)

---

## Expected Output

```
go is simple
but strings
are deep
```

## Key Takeaways

1. **Greedy first-fit wrapping** — break when the next word won't fit.
2. **`strings.Fields`** — whitespace normalization + tokenization.
3. **`strings.Builder`** — efficient string concatenation.
4. **`lineLen` tracking** — remembering the current line width.