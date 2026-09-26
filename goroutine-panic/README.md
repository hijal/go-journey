# goroutine-panic

Go-তে **goroutine-এর panic recover + slog logging** — `wg.Go`-চালানো task-এ panic হলে service crash না হয়ে log হয়।

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

### Lines 3–8

```go
import (
	"fmt"
	"log/slog"
	"sync"
)
```

`slog` (structured error log) + `sync` (`WaitGroup`)।

### Lines 9–18

```go
func runSafely(wg *sync.WaitGroup, task string, fn func()) {
	wg.Go(func() {
		defer func() {
			if r := recover(); r != nil {
				slog.Error("task panicked", "task", task, "panic", r)
			}
		}()
		fn()
	})
}
```

**Safe wrapper:**

- `wg.Go`-তে বেঁধে goroutine চালায় — Done automatically।
- **`defer + recover`** — `fn()` panic করলে caught; task নাম আর panic value `slog.Error`-এ log।

### Lines 23–29

```go
	runSafely(&wg, "sync-orders", func() { fmt.Println("orders synced") })
	runSafely(&wg, "sync-refunds", func() {
		var refunds map[string]int
		refunds["tx-1"] = 100
	})
	wg.Wait()
	fmt.Println("service still alive")
```

- `sync-orders` — normal।
- `sync-refunds` — **nil map-এ write** → panic (`assignment to entry in nil map`) — caught + logged।
- `wg.Wait` পরে — **service Crash-নয়, alive-ই থেকে যায়**।

---

## Expected Output

```
orders synced
service still alive
2026/... ERROR task panicked task=sync-refunds panic="assignment to entry in nil map"
```

## মূল শিক্ষা / Key Takeaways

1. **Goroutine panic প্রোগ্রাম panic করে** — recover ছাড়া whole process crash।
2. **`defer` + `recover`** — goroutine-এর মধ্যে safety net।
3. **`slog.Error`** — key-value structured log।
4. **`wg.Go` + panic** — LIFO defer-এ log হয়, তারপর Done → কোনো deadlock নেই।
5. **Recovered panic = service বেঁচে থাকে** ("still alive")।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Lines 3–8

```go
import (
	"fmt"
	"log/slog"
	"sync"
)
```

`slog` (structured error log) + `sync` (`WaitGroup`).

### Lines 9–18

```go
func runSafely(wg *sync.WaitGroup, task string, fn func()) {
	wg.Go(func() {
		defer func() {
			if r := recover(); r != nil {
				slog.Error("task panicked", "task", task, "panic", r)
			}
		}()
		fn()
	})
}
```

**A safe wrapper:**

- Runs `fn` in a goroutine bound to `wg.Go` — `Done` happens automatically.
- **`defer` + `recover`** — if `fn()` panics it's caught; the task name and panic value are logged via `slog.Error`.

### Lines 23–29

```go
	runSafely(&wg, "sync-orders", func() { fmt.Println("orders synced") })
	runSafely(&wg, "sync-refunds", func() {
		var refunds map[string]int
		refunds["tx-1"] = 100
	})
	wg.Wait()
	fmt.Println("service still alive")
```

- `sync-orders` — normal.
- `sync-refunds` — **writing to a nil map** → panics (`assignment to entry in nil map`) — caught + logged.
- After `wg.Wait` — the **service is still alive, not crashed**.

---

## Expected Output

```
orders synced
service still alive
2026/... ERROR task panicked task=sync-refunds panic="assignment to entry in nil map"
```

## Key Takeaways

1. **A goroutine panic panics the program** — the whole process crashes without recover.
2. **`defer` + `recover`** — a safety net inside the goroutine.
3. **`slog.Error`** — key-value structured logging.
4. **`wg.Go` + panic** — the LIFO defer logs first, then `Done` — no deadlock.
5. **A recovered panic = the service lives on** ("still alive").