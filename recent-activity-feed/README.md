# recent-activity-feed

Go-তে **builtin `min` tail-trim + `slices.Clone`+`slices.Reverse` in-place-safety** শেখার ছোট example — latest-N events newest-first।

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
- `slices` — `Clone`, `Reverse` (Go 1.21+)।

### Lines 8–9

```go
func main() {
	events := []int{1700000100, 1700000250, 1700000400, 1700000550, 1700000700, 1700000850, 1700001000}
```

7টা ascending timestamp (Unix epoch)। (Feed-এর "পুরাতন→নতুন" source-order।)

### Line 11

```go
	recent := events[len(events)-min(5, len(events)):]
```

**Tail-trim — latest N:**

- `min(5, len(events))` — **builtin `min` (Go 1.21+)**: 5 new 7 — `min` returns `5`। (যদি 5-এর কম event থাকত, `len`-ই limit — out-of-bounds guard।)
- `events[len-5:]` — শেষ 5টা (সবচেয়ে নতুন, ascending-order-এর tail)।

### Lines 12–13

```go
	newestFirst := slices.Clone(recent)
	slices.Reverse(newestFirst)
```

**Reverse — copy-তে:**

- `slices.Reverse` **in-place** mutate করে।
- তাই আগে `slices.Clone(recent)` — নিজস্ব copy-তে reverse; `events` untouched।

### Lines 15–16

```go
	fmt.Println("newest first:", newestFirst)
	fmt.Println("original untouched:", events[:2], "...")
```

- `newestFirst` — reverse-chronological।
- Last line-টা **Clone-এর প্রমাণ**: original এখনো ascending।

---

## Expected Output

```
newest first: [1700001000 1700000850 1700000700 1700000550 1700000400]
original untouched: [1700000100 1700000250] ...
```

## মূল শিক্ষা / Key Takeaways

1. **Builtin `min`** — Go 1.21+ clamp/list-length guard।
2. **Tail trimming** — `events[len-N:]`।
3. **Clone before in-place-op** — Reverse mutate-র আগে copy।
4. **Newest-first** — asc array-র tail-এর reverse।

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
- `slices` — for `Clone`, `Reverse` (Go 1.21+).

### Lines 8–9

```go
func main() {
	events := []int{1700000100, 1700000250, 1700000400, 1700000550, 1700000700, 1700000850, 1700001000}
```

7 ascending timestamps (Unix epoch). (The feed's "old→new" source order.)

### Line 11

```go
	recent := events[len(events)-min(5, len(events)):]
```

**Tail-trim — the latest N:**

- `min(5, len(events))` — **builtin `min` (Go 1.21+)**: 5 vs 7 — `min` returns `5`. (With fewer than 5 events, `len` becomes the limit — an out-of-bounds guard.)
- `events[len-5:]` — the last 5 (the newest, the ascending tail).

### Lines 12–13

```go
	newestFirst := slices.Clone(recent)
	slices.Reverse(newestFirst)
```

**Reverse — on a copy:**

- `slices.Reverse` mutates **in place**.
- So first `slices.Clone(recent)` — the reverse happens on a private copy; `events` stays untouched.

### Lines 15–16

```go
	fmt.Println("newest first:", newestFirst)
	fmt.Println("original untouched:", events[:2], "...")
```

- `newestFirst` — reverse-chronological.
- The last line is the **proof of Clone**: the original is still ascending.

---

## Expected Output

```
newest first: [1700001000 1700000850 1700000700 1700000550 1700000400]
original untouched: [1700000100 1700000250] ...
```

## Key Takeaways

1. **Builtin `min`** — Go 1.21+ clamp/length guard.
2. **Tail trimming** — `events[len-N:]`.
3. **Clone before an in-place op** — a copy before `Reverse` mutates.
4. **Newest-first** — reverse of the ascending tail.