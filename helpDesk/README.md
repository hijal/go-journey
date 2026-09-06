# helpDesk

Go-তে **function pipeline** pattern শেখার ছোট example -- slice of functions বানিয়ে string processing steps chain করা.

**ভাষা নির্বাচন করুন / Choose language:**

[বাংলা](#bangla) * [English](#english)

---

<a name="bangla"></a>

## বাংলা সংস্করণ

### Lines 1-6

```go
package main

import (
    "fmt"
    "strings"
)
```

- `package main` -- একটা executable program.
- `fmt` -- output print করতে.
- `strings` -- `TrimSpace`, `ReplaceAll`, `Fields`, `Join` এর জন্য.

### Lines 8-22

```go
func buildPipeline() []func(string) string {
    trim := strings.TrimSpace

    removePlaceholder := func(s string) string {
        return strings.ReplaceAll(s, "[TICKET]", "")
    }

    collapseSpaces := func(s string) string {
        return strings.Join(strings.Fields(s), " ")
    }

    return []func(string) string{
        trim, removePlaceholder, collapseSpaces,
    }
}
```

Pipeline তৈরি:

- `trim` -- `strings.TrimSpace` function-কে value হিসেবে assign (function reference).
- `removePlaceholder` -- anonymous function যা `[TICKET]` placeholder মুছে দেয়.
- `collapseSpaces` -- multiple spaces collapse করে single space-এ.
- return type `[]func(string) string` -- string-processor function-এর slice.

### Lines 24-30

```go
func processTicket(raw string, steps []func(string) string) string {
    text := raw
    for _, step := range steps {
        text = step(text)
    }
    return text
}
```

প্রতিটা step sequentially apply হয়; আগের step-এর output পরের step-এ যায়.

### Lines 32-38

```go
func main() {
    raw := "  [TICKET]   My   payment   failed   twice  "

    steps := buildPipeline()
    fmt.Printf("raw: %q\n", raw)
    fmt.Printf("clean: %q\n", processTicket(raw, steps))
}
```

Input `raw`:
1. `trim` -> leading/trailing spaces মুছবে
2. `removePlaceholder` -> `[TICKET]` মুছবে
3. `collapseSpaces` -> multiple spaces single করবে

---

## Expected Output

```
raw: "  [TICKET]   My   payment   failed   twice  "
clean: "My payment failed twice"
```

## মূল শিক্ষা / Key Takeaways

1. **Function as value** -- `strings.TrimSpace` কে variable-এ assign করা যায়.
2. **Function type slice** -- `[]func(string) string` -- pipeline of string-transformers.
3. **Sequential processing** -- একটা step-এর output পরের step-এ input হয়.
4. **Composability** -- pipeline steps independently build/test করা যায়.

---

---

<a name="english"></a>

## English Version

### Lines 1-6

```go
package main

import (
    "fmt"
    "strings"
)
```

- `package main` -- an executable program.
- `fmt` -- for console output.
- `strings` -- for `TrimSpace`, `ReplaceAll`, `Fields`, `Join`.

### Lines 8-22

```go
func buildPipeline() []func(string) string {
    trim := strings.TrimSpace

    removePlaceholder := func(s string) string {
        return strings.ReplaceAll(s, "[TICKET]", "")
    }

    collapseSpaces := func(s string) string {
        return strings.Join(strings.Fields(s), " ")
    }

    return []func(string) string{
        trim, removePlaceholder, collapseSpaces,
    }
}
```

Pipeline builder:

- `trim` -- assigns `strings.TrimSpace` as a function value.
- `removePlaceholder` -- anonymous function that strips the `[TICKET]` placeholder.
- `collapseSpaces` -- collapses runs of whitespace into a single space.
- return type `[]func(string) string` -- a slice of string-processor functions.

### Lines 24-30

```go
func processTicket(raw string, steps []func(string) string) string {
    text := raw
    for _, step := range steps {
        text = step(text)
    }
    return text
}
```

Each step is applied in order; the previous step's output becomes the next step's input.

### Lines 32-38

```go
func main() {
    raw := "  [TICKET]   My   payment   failed   twice  "

    steps := buildPipeline()
    fmt.Printf("raw: %q\n", raw)
    fmt.Printf("clean: %q\n", processTicket(raw, steps))
}
```

Input `raw` is processed by:
1. `trim` -- strips leading/trailing spaces
2. `removePlaceholder` -- removes `[TICKET]`
3. `collapseSpaces` -- collapses repeated spaces

---

## Expected Output

```
raw: "  [TICKET]   My   payment   failed   twice  "
clean: "My payment failed twice"
```

## Key Takeaways

1. **Function as value** -- `strings.TrimSpace` can be assigned to a variable.
2. **Function type slice** -- `[]func(string) string` is a pipeline of string transformers.
3. **Sequential processing** -- each step's output feeds the next step's input.
4. **Composability** -- pipeline steps can be built and tested independently.
