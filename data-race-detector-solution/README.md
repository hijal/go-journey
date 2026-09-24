# data-race-detector-solution

`data-race-detector`-এর **fix** — `sync.Mutex` দিয়ে shared counter-কে safe করা।

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

`fmt` (print) + `sync` (`WaitGroup` + **`Mutex`**)।

### Lines 8–13

```go
func main() {
	var (
		wg        sync.WaitGroup
		mu        sync.Mutex
		pageViews int
	)
```

Grouped `var` — শেয়ারড state-এর পাশেই lock।

### Lines 15–23

```go
	for range 1000 {
		wg.Go(func() {
			mu.Lock()
			defer mu.Unlock()
			pageViews++
		})
	}
	wg.Wait()
	fmt.Println("page views:", pageViews)
```

**The fix:**

- **`mu.Lock()`** — section-এ ঢোকার আগে exclusive access।
- **`defer mu.Unlock()`** — exit-এ lock ছাড়ে।
- **`pageViews++`** — এখন protected critical section — একটার-পরে-একটা, update হারায় না।

### Before vs After

| version | plain run | `go run -race` |
|---|---|---|
| `data-race-detector` | 992/968/995 (ভুল) | `DATA RACE` warning |
| **this (solution)** | সবসময় **1000** | clean |

> Mutex নিয়ম: shared data + goroutine = lock/atomic নয়-হলে race। Lock scope যত ছোট তত ভালো।

---

## Expected Output

```
page views: 1000
```

## মূল শিক্ষা / Key Takeaways

1. **`mu.Lock()` / `defer mu.Unlock()`** — critical section guard।
2. **Grouped `var`** — lock + data পাশাপাশি declare।
3. **Race-মুক্ত** — `-race` clean।
4. **Always-deterministic** — 1000 ভুল হয় না।
5. **`wg.Go` + mutex** — concurrency-তে safe composition।

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

`fmt` (printing) + `sync` (`WaitGroup` + **`Mutex`**).

### Lines 8–13

```go
func main() {
	var (
		wg        sync.WaitGroup
		mu        sync.Mutex
		pageViews int
	)
```

A grouped `var` — the lock sits right next to the shared state it protects.

### Lines 15–23

```go
	for range 1000 {
		wg.Go(func() {
			mu.Lock()
			defer mu.Unlock()
			pageViews++
		})
	}
	wg.Wait()
	fmt.Println("page views:", pageViews)
```

**The fix:**

- **`mu.Lock()`** — takes exclusive access before entering the section.
- **`defer mu.Unlock()`** — releases the lock on exit.
- **`pageViews++`** — now a protected critical section — one at a time, no lost updates.

### Before vs After

| version | plain run | `go run -race` |
|---|---|---|
| `data-race-detector` | 992/968/995 (wrong) | `DATA RACE` warning |
| **this (solution)** | always **1000** | clean |

> Mutex rule: shared data + goroutines = lock/atomic or you will race. Keep the lock scope as small as possible.

---

## Expected Output

```
page views: 1000
```

## Key Takeaways

1. **`mu.Lock()` / `defer mu.Unlock()`** — guarding the critical section.
2. **Grouped `var`** — declaring lock + data side by side.
3. **Race-free** — `-race` clean.
4. **Always deterministic** — never wrong, always 1000.
5. **`wg.Go` + mutex** — safe composition under concurrency.