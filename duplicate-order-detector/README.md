# duplicate-order-detector

Go-তে **map-based seen-set + `slices.Contains` dedupe + sort** শেখার ছোট example — duplicate order ID শনাক্ত (deterministic, একবার-ই লিস্টে)।

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

- `fmt` — `Println`।
- `slices` — `Contains`, `Sort` (Go 1.21+)।

### Lines 8–11

```go
func main() {
	orders := []string{
		"ORD-101", "ORD-102", "ORD-101", "ORD-103", "ORD-102", "ORD-101",
	}
```

6টা order — ORD-101 **৩বার**, ORD-102 **২বার**, ORD-103 একবার।

### Lines 13–14

```go
	seen := make(map[string]bool)
	duplicates := []string{}
```

- `seen` — "পূর্বে দেখা হয়েছে" flag-set।
- `duplicates` — শনাক্ত ডুপ-হোল্ডার।

### Lines 16–21

```go
	for _, id := range orders {
		if seen[id] && !slices.Contains(duplicates, id) {
			duplicates = append(duplicates, id)
		}
		seen[id] = true
	}
```

**Detect loop:**

- `if seen[id]` — আগে **দ্বিতীয়বার** আসছে মানে duplicate।
- `!slices.Contains(duplicates, id)` — ডুপ-লিস্টে আগে ঢোকেনি → `append`। (ORD-101 ৩য়বার আসলে এই guard-টা appending আটকায় — প্রতিটা duplicate **একবার-ই** লিস্টে।)
- `seen[id] = true` — প্রতিবার mark (আগেরটা hold করে)।

### Lines 23–24

```go
	slices.Sort(duplicates)
	fmt.Println("duplicate order ids:", duplicates)
```

- `slices.Sort` — deterministic-report order (যদি map-order hop না হয়)।
- Print।

**Walkthrough:**

1. ORD-101 — seen? no → mark ✓
2. ORD-102 — seen? no → mark ✓
3. **ORD-101** — seen ✓ → duplicates+=101
4. ORD-103 — mark ✓
5. **ORD-102** — seen ✓ → duplicates+=102
6. ORD-101 — ৩য়বার, duplicate-এ আছে → skip

---

## Expected Output

```
duplicate order ids: [ORD-101 ORD-102]
```

## মূল শিক্ষা / Key Takeaways

1. **Set-scan** — `map[string]bool` presence-check।
2. **`slices.Contains` dedupe** — list-এ একবার-ই।
3. **`slices.Sort`** — deterministic output।
4. **First-pass O(n)** — single loop detection।

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

- `fmt` — for `Println`.
- `slices` — for `Contains`, `Sort` (Go 1.21+).

### Lines 8–11

```go
func main() {
	orders := []string{
		"ORD-101", "ORD-102", "ORD-101", "ORD-103", "ORD-102", "ORD-101",
	}
```

6 orders — ORD-101 **3 times**, ORD-102 **2 times**, ORD-103 once.

### Lines 13–14

```go
	seen := make(map[string]bool)
	duplicates := []string{}
```

- `seen` — a "already seen" flag set.
- `duplicates` — the detected-dup holder.

### Lines 16–21

```go
	for _, id := range orders {
		if seen[id] && !slices.Contains(duplicates, id) {
			duplicates = append(duplicates, id)
		}
		seen[id] = true
	}
```

**Detect loop:**

- `if seen[id]` — appearing **a second time** means a duplicate.
- `!slices.Contains(duplicates, id)` — not yet in the list → `append`. (On the 3rd ORD-101 this guard stops the append — each duplicate hits the list **exactly once**.)
- `seen[id] = true` — marks it each pass (keeps earlier entries).

### Lines 23–24

```go
	slices.Sort(duplicates)
	fmt.Println("duplicate order ids:", duplicates)
```

- `slices.Sort` — a deterministic report order.
- Prints.

**Walkthrough:**

1. ORD-101 — seen? no → mark ✓
2. ORD-102 — seen? no → mark ✓
3. **ORD-101** — seen ✓ → duplicates+=101
4. ORD-103 — mark ✓
5. **ORD-102** — seen ✓ → duplicates+=102
6. ORD-101 — 3rd time, already listed → skip

---

## Expected Output

```
duplicate order ids: [ORD-101 ORD-102]
```

## Key Takeaways

1. **Set-scan** — `map[string]bool` presence checks.
2. **`slices.Contains` dedupe** — listed once.
3. **`slices.Sort`** — deterministic output.
4. **First-pass O(n)** — a single-loop detection.