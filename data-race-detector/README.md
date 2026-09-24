# data-race-detector

Go-তে **data race hard-নিয়ে experiment** — ইচ্ছাকৃত-ভাবে unsafe concurrent increment (`-race` detector দেখাতে)।

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
	"sync"
)
```

`fmt` (print) + `sync` (`WaitGroup`)।

### Lines 8–10

```go
func main() {
	var wg sync.WaitGroup
	pageViews := 0
```

একটা shared counter — goroutine-দের মিলে increment করবে।

### Lines 11–17

```go
	for range 1000 {
		wg.Go(func() {
			pageViews++
		})
	}
	wg.Wait()
	fmt.Println("page views:", pageViews)
```

**The bug:** 1000টা goroutine `pageViews++` — `++` atomic নয় (read-modify-write), তাই একসাথে access করলে update হারায়।

### Why it's "broken on purpose"

- **Plain run** — count প্রায়ই 1000-এর কম (গোরুটীন race-এ losing update)।
- **`go run -race`** — `WARNING: DATA RACE` + exact line।

> **Fix দেখো:** `data-race-detector-solution/` — mutex-guarded version।

---

## Expected Output

Plain run (ভিন্ন count — race-এর ফল):

```
page views: 992
page views: 968
page views: 995
```

`-race` run:

```
==================
WARNING: DATA RACE
Read at 0x... by goroutine 15:
  main.main.func1()
      .../data-race-detector/main.go:13
==================
```

## মূল শিক্ষা / Key Takeaways

1. **`++` atomic নয়** — ৩-step অপ (read, inc, write)।
2. **Losing update** — overlay-write ড্রপ count।
3. **`-race` detector** — bug synchronous-ভাবে ধরা পড়ে।
4. **`wg.Go`** — পরিবেশন count ট্র্যাক।
5. **Medicine:** mutex/atomic (সার্ব case)।

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
	"sync"
)
```

`fmt` (printing) + `sync` (`WaitGroup`).

### Lines 8–10

```go
func main() {
	var wg sync.WaitGroup
	pageViews := 0
```

A shared counter that all the goroutines are going to increment.

### Lines 11–17

```go
	for range 1000 {
		wg.Go(func() {
			pageViews++
		})
	}
	wg.Wait()
	fmt.Println("page views:", pageViews)
```

**The bug:** 1000 goroutines do `pageViews++` — `++` is not atomic (read-modify-write), so concurrent access loses updates.

### Why it's "broken on purpose"

- **Plain run** — the count is usually below 1000 (racing goroutines lose updates).
- **`go run -race`** — `WARNING: DATA RACE` with the exact line.

> **See the fix in** `data-race-detector-solution/` — the mutex-guarded version.

---

## Expected Output

Plain run (varying counts — the result of the race):

```
page views: 992
page views: 968
page views: 995
```

`-race` run:

```
==================
WARNING: DATA RACE
Read at 0x... by goroutine 15:
  main.main.func1()
      .../data-race-detector/main.go:13
==================
```

## Key Takeaways

1. **`++` is not atomic** — it's a 3-step op (read, inc, write).
2. **Losing updates** — overwrites drop the count.
3. **The `-race` detector** — catches the bug synchronously.
4. **`wg.Go`** — tracks the worker count.
5. **The medicine** — a mutex/atomic (any shared case).