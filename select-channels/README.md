# select-channels

Go-তে **`select` + `time.After` timeout + async fetch pattern** — দুটো channel-এর মধ্যে ready-first race।

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
	"time"
)
```

`fmt` (print) + `time` (delay + timeout)। No sync import হওয়া সত্ত্বেও — concurrency ঠিক আছে (channel-ই synchronization)।

### Lines 8–16

```go
func fetchExchangeRate(delay time.Duration) <-chan float64 {
	ch := make(chan float64, 1)
	go func() {
		time.Sleep(delay)
		ch <- 122.40
	}()

	return ch
}
```

**Async fetch factory**:

- `make(chan float64, 1)` — **buffer 1** — send block করে না (goroutine-কে delay পরে send)।
- goroutine `delay` অপেক্ষা করে, তারপর rate পাঠায়।
- **`<-chan float64`** — **receive-only** channel type (sender-safe API — consumer send করতে পারে না)।

### Lines 18–24

```go
func main() {
	select {
	case rate := <-fetchExchangeRate(50 * time.Millisecond):
		fmt.Println("USD->BDT:", rate)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("rate service timed out!")
	}
}
```

**`select`** — যেটা-ই আগে ready, সেটি-ই চলে:

- `fetchExchangeRate(50ms)` — 50ms পরে value।
- `time.After(200ms)` — ২০০ms পরে timeout case।

50 < 200 → rate case wins। Buffer 1 থাকায় rate `<-` সাথে-সাথে আসে।

---

## Expected Output

```
USD->BDT: 122.4
```

## মূল শিক্ষা / Key Takeaways

1. **`select`** — multiple channel-op-এর readiness-এ first wins।
2. **`time.After(d)`** — এক-শট timeout case।
3. **Buffered `chan T, n`** — send-ব্লক-এ delay ছাড়া।
4. **`<-chan T`** — receive-only directional channel।
5. **Async fetch + select** — rate/timeout-এর race composition।

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
	"time"
)
```

`fmt` (printing) + `time` (delay + timeout). No `sync` import is needed — the channel is the synchronization.

### Lines 8–16

```go
func fetchExchangeRate(delay time.Duration) <-chan float64 {
	ch := make(chan float64, 1)
	go func() {
		time.Sleep(delay)
		ch <- 122.40
	}()

	return ch
}
```

**The async fetch factory**:

- `make(chan float64, 1)` — a **buffer of 1** — the send won't block (the goroutine sends after the delay).
- The goroutine waits `delay`, then sends the rate.
- **`<-chan float64`** — a **receive-only** channel type (a safe API for consumers).

### Lines 18–24

```go
func main() {
	select {
	case rate := <-fetchExchangeRate(50 * time.Millisecond):
		fmt.Println("USD->BDT:", rate)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("rate service timed out!")
	}
}
```

**`select`** — whichever case becomes ready first runs:

- `fetchExchangeRate(50ms)` — a value after 50ms.
- `time.After(200ms)` — a timeout case after 200ms.

50 < 200 → the rate case wins. With buffer 1, the rate is ready for `<-` immediately.

---

## Expected Output

```
USD->BDT: 122.4
```

## Key Takeaways

1. **`select`** — first ready wins across multiple channel ops.
2. **`time.After(d)`** — a one-shot timeout case.
3. **Buffered `chan T, n`** — no send-blocking delay.
4. **`<-chan T`** — a receive-only directional channel.
5. **Async fetch + select** — a composed rate/timeout.