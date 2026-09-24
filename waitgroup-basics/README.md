# waitgroup-basics

Go-তে **`sync.WaitGroup` — classic (`Add`/`Done`/`Wait`) + modern (`wg.Go`)** শেখার ছোট example — goroutine-দের জন্য সঠিক waiting mechanism।

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
	"sync"
	"time"
)
```

- `sync` — `WaitGroup` (thread-safe counter)।
- `time` — `Sleep` (এখানে শুধু কাজ-সিমুলেশন, sync না)।

### Lines 9–12

```go
func generateInvoice(orderID int) {
	time.Sleep(50 * time.Millisecond)
	fmt.Println("invoice ready for order", orderID)
}
```

কাজ-সিমুলেটিং func — 50ms "work", তারপর print।

### Lines 14–15

```go
func main() {
	var wg sync.WaitGroup
```

**`WaitGroup`** — zero-value usable; goroutine count track করে।

### Lines 17–27

```go
	// classic pattern
	for _, id := range []int{101, 102, 103} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			generateInvoice(id)
		}()
	}
	wg.Wait()
	fmt.Println("all invoices done with classic pattern")
	fmt.Println()
```

**Classic pattern:**

- **`wg.Add(1)`** — প্রতিটি goroutine spawn-এর আগে counter +1।
- `go func(){...}()` — goroutine; `defer wg.Done()` — শেষে counter −1।
- Go 1.22+ loop var — `id` প্রতি iteration-এ ভিন্ন (`101,102,103`)।
- **`wg.Wait()`** — counter 0 না হওয়া পর্যন্ত block; কাজ শেষ হলেই পরের line।

### Lines 29–37

```go
	// modern pattern
	for _, id := range []int{105, 106, 107} {
		wg.Go(func() {
			generateInvoice(id)
		})
	}
	wg.Wait()

	fmt.Println("all invoices done with modern pattern")
```

**Modern pattern (Go 1.24+ `WaitGroup.Go`):**

- **`wg.Go(func)`** — `Add(1)` + `go func()` + `Done()` **এক call-এ**; counter নিজেই manage হয় (এখনো finish-এর আগে `Wait` করতে হবে)।
- কম boilerplate, rookie mistake (Done ভুলে যাওয়া) কমবার সম্ভাবনা।

> **time.Sleep ব্যবহার করো না অপেক্ষার জন্য** — তুমি জানো না কত সময় লাগবে: কম হলে অসম্পূর্ণ, বেশি হলে নষ্ট। `WaitGroup` ঠিক সময়েই ছাড়বে।

---

## Expected Output

```
invoice ready for order 102
invoice ready for order 103
invoice ready for order 101
all invoices done with classic pattern

invoice ready for order 106
invoice ready for order 105
invoice ready for order 107
all invoices done with modern pattern
```

(প্রথম দুই block-এর ভেতরে line-গুলোর order goroutine-এ nondeterministic; প্রতিটা block-এর শেষ line সবসময় last)

## মূল শিক্ষা / Key Takeaways

1. **`Add(n)`** — counter বাড়াও (spawn-এর আগে)।
2. **`Done()`** — counter এক কমাও (প্রায়ই `defer`)।
3. **`Wait()`** — 0 পর্যন্ত block।
4. **`wg.Go()`** — Go 1.24+: Add+go+Done একসাথে।
5. **Sleep অ-নির্ভরযোগ্য** — WaitGroup-ই সঠিক।
6. **Thread-safe counter** — mutation goroutine-safe।

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
	"sync"
	"time"
)
```

- `sync` — the `WaitGroup` (a thread-safe counter).
- `time` — `Sleep` (here only simulating work, not syncing).

### Lines 9–12

```go
func generateInvoice(orderID int) {
	time.Sleep(50 * time.Millisecond)
	fmt.Println("invoice ready for order", orderID)
}
```

A work-simulating func — 50ms of "work", then a print.

### Lines 14–15

```go
func main() {
	var wg sync.WaitGroup
```

**`WaitGroup`** — usable at zero value; tracks the goroutine count.

### Lines 17–27

```go
	// classic pattern
	for _, id := range []int{101, 102, 103} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			generateInvoice(id)
		}()
	}
	wg.Wait()
	fmt.Println("all invoices done with classic pattern")
	fmt.Println()
```

**The classic pattern:**

- **`wg.Add(1)`** — counter +1 before each goroutine spawn.
- `go func(){...}()` — the goroutine; `defer wg.Done()` — counter −1 on exit.
- Go 1.22+ loop vars — `id` is per-iteration (`101,102,103`).
- **`wg.Wait()`** — blocks until the counter hits 0; execution continues only when the work is done.

### Lines 29–37

```go
	// modern pattern
	for _, id := range []int{105, 106, 107} {
		wg.Go(func() {
			generateInvoice(id)
		})
	}
	wg.Wait()

	fmt.Println("all invoices done with modern pattern")
```

**The modern pattern (Go 1.24+ `WaitGroup.Go`):**

- **`wg.Go(func)`** — `Add(1)` + `go func()` + `Done()` **in one call**; the counter is managed for you (you still `Wait` before finishing).
- Less boilerplate, fewer chances for the rookie mistake of forgetting `Done`.

> **Don't use `time.Sleep` to wait** — you don't know how long things take: too little and the work is incomplete, too much and time is wasted. `WaitGroup` releases at exactly the right moment.

---

## Expected Output

```
invoice ready for order 102
invoice ready for order 103
invoice ready for order 101
all invoices done with classic pattern

invoice ready for order 106
invoice ready for order 105
invoice ready for order 107
all invoices done with modern pattern
```

(within each block the line order is nondeterministic per goroutine; each block's last line always comes last)

## Key Takeaways

1. **`Add(n)`** — bump the counter (before spawning).
2. **`Done()`** — decrement (often via `defer`).
3. **`Wait()`** — block until 0.
4. **`wg.Go()`** — Go 1.24+: Add+go+Done in one.
5. **Sleep is unreliable** — `WaitGroup` is the correct tool.
6. **Thread-safe counter** — goroutine-safe mutation.