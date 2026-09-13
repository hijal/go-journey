# log-analysis-slicing

Go-তে **slice view/share backing array + slicing (`logs[len-3:]`) + `slices.IndexFunc`** শেখার ছোট example — লগ হ্যান্ডলিং।

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
	"slices"
)
```

- `fmt` — `Printf`, `Println`।
- `slices` — `IndexFunc` (Go 1.21+)।

### Lines 8–15

```go
func main() {
	logs := []string{
		"INFO service started",
		"INFO request received",
		"ERROR db connection refused",
		"INFO retrying",
		"ERROR timeout after 3s",
	}
```

5টা log line — INFO/ERROR mix। Full slice।

### Lines 17–18

```go
	tail := logs[len(logs)-3:] // last 3 logs
	fmt.Println("tail:", tail)
```

**Slicing** — `logs[len(logs)-3:]`:

- index `2` থেকে শেষ পর্যন্ত → শেষ ৩টা: `ERROR db connection refused`, `INFO retrying`, `ERROR timeout after 3s`।
- **Slice-view:** tail নতুন copy নয় — logs-এর পেছনের **একই backing array**-র window।

### Lines 20–26

```go
	i := slices.IndexFunc(logs, func(line string) bool {
		return len(line) > 5 && line[:5] == "ERROR"
	})

	if i >= 0 {
		fmt.Printf("first error at index %d: %q\n", i, logs[i])
	}
```

`slices.IndexFunc` — predicate মিললে **প্রথম index**:

- `line[:5] == "ERROR"` — prefix-চেক (প্রথম 5 character)।
- `len(line) > 5` — **guard**: ছোট line-এ `[:5]` slice panic করত (out-of-bounds)। Python-like `startswith` এখানে দুটো শর্ত।
- প্রথম "ERROR" line index `2` → `first error at index 2: "ERROR db connection refused"`।

### Lines 28–30

```go
	tail[0] = "INFO tail overwritten"

	fmt.Println("original[2:]", logs[2])
```

**মূল শিক্ষা — shared backing array:**

- `tail[0] = "INFO tail overwritten"` — tail-এর 0-তম element-টা বদলানো।
- কারণ tail আর logs একই memory pool-এর উপর — `logs[2]`-ও বদলে গেছে → `original[2:] INFO tail overwritten`।
- **Slice-গুলো copy নয়; reference-ভিত্তিক window** — একটা জায়গায় লেখা অন্যজায়গায় দৃশ্য।

---

## Expected Output

```
tail: [ERROR db connection refused INFO retrying ERROR timeout after 3s]
first error at index 2: "ERROR db connection refused"
original[2:] INFO tail overwritten
```

## মূল শিক্ষা / Key Takeaways

1. **Slicing** — `logs[a:]`, `logs[:b]`, `logs[a:b]` — subset window।
2. **Slice = view** — একই backing array; mutation দুই জায়গায় দৃশ্য।
3. **`slices.IndexFunc`** — predicate-based first-index lookup।
4. **Bounds guard** — `len(line) > 5` + `line[:5]`।
5. **Tail access** — `logs[len-3:]` latest-N।

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
	"slices"
)
```

- `fmt` — for `Printf`, `Println`.
- `slices` — for `IndexFunc` (Go 1.21+).

### Lines 8–15

```go
func main() {
	logs := []string{
		"INFO service started",
		"INFO request received",
		"ERROR db connection refused",
		"INFO retrying",
		"ERROR timeout after 3s",
	}
```

5 log lines — an INFO/ERROR mix. The full slice.

### Lines 17–18

```go
	tail := logs[len(logs)-3:] // last 3 logs
	fmt.Println("tail:", tail)
```

**Slicing** — `logs[len(logs)-3:]`:

- from index `2` to the end → the last 3: `ERROR db connection refused`, `INFO retrying`, `ERROR timeout after 3s`.
- **Slice-view:** tail isn't a new copy — it's a window over the **same backing array** as logs.

### Lines 20–26

```go
	i := slices.IndexFunc(logs, func(line string) bool {
		return len(line) > 5 && line[:5] == "ERROR"
	})

	if i >= 0 {
		fmt.Printf("first error at index %d: %q\n", i, logs[i])
	}
```

`slices.IndexFunc` — the **first index** matching the predicate:

- `line[:5] == "ERROR"` — a prefix check (first 5 characters).
- `len(line) > 5` — a **guard**: on a short line, `[:5]` would panic (out of bounds). A Python-like `startswith` takes both conditions here.
- The first "ERROR" line is at index `2` → `first error at index 2: "ERROR db connection refused"`.

### Lines 28–30

```go
	tail[0] = "INFO tail overwritten"

	fmt.Println("original[2:]", logs[2])
```

**The core lesson — shared backing array:**

- `tail[0] = "INFO tail overwritten"` — modifies tail's 0-th element.
- Because tail and logs sit on the same memory pool, `logs[2]` also changed → `original[2:] INFO tail overwritten`.
- **Slices are not copies; they're reference-based windows** — a write in one place is visible in the other.

---

## Expected Output

```
tail: [ERROR db connection refused INFO retrying ERROR timeout after 3s]
first error at index 2: "ERROR db connection refused"
original[2:] INFO tail overwritten
```

## Key Takeaways

1. **Slicing** — `logs[a:]`, `logs[:b]`, `logs[a:b]` — subset windows.
2. **Slice = view** — shared backing array; mutations are visible in both.
3. **`slices.IndexFunc`** — predicate-based first-index lookup.
4. **Bounds guard** — `len(line) > 5` + `line[:5]`.
5. **Tail access** — `logs[len-3:]` for the latest N.