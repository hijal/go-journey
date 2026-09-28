# signal-channel-broadcast

Go-তে **close দিয়ে broadcast** — `chan struct{}` বন্ধ করলে সব blocked worker একসাথে চালু।

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

`sync` (WaitGroup), `time` (Sleep)।

### Lines 9–23

```go
func main() {
	shutdown := make(chan struct{})

	var wg sync.WaitGroup

	for id := range 10 {
		wg.Go(func() {
			<-shutdown

			fmt.Println("worker", id, "stopping")
		})
	}
	time.Sleep(time.Millisecond * 50)
	close(shutdown)
	wg.Wait()
}
```

- **`chan struct{}`** — শুধু signal বোঝাতে zero-size type; কোনো value বহন করে না।
- **`wg.Go(func(){})`** — Go 1.25+ helper, `Add(1)`+`go`+`Done` একসাথে (loop var per-iteration, তাই `id` ঠিক থাকে)।
- ১০টা worker-ই `<-shutdown`-এ **blocked**।
- **`close(shutdown)`** — একটা `close`-ই সবাইকে ছেড়ে দেয় (broadcast), যেন আলাদা আলাদা send না করতে হয়।
- **`wg.Wait()`** — সব শেষ হওয়ার অপেক্ষা।
- Print-এর **ক্রম nondeterministic**।

---

## Expected Output (order varies)

```
worker 4 stopping
worker 8 stopping
worker 7 stopping
...
worker 2 stopping
```

## মূল শিক্ষা / Key Takeaways

1. **`close` = broadcast** — সব receiver একসাথে মুক্ত।
2. **`chan struct{}`** — pure signal channel।
3. **`wg.Go`** — modern WaitGroup helper।
4. **`wg.Wait()`** — completion-এর synchronization।

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

`sync` (WaitGroup), `time` (Sleep).

### Lines 9–23

```go
func main() {
	shutdown := make(chan struct{})

	var wg sync.WaitGroup

	for id := range 10 {
		wg.Go(func() {
			<-shutdown

			fmt.Println("worker", id, "stopping")
		})
	}
	time.Sleep(time.Millisecond * 50)
	close(shutdown)
	wg.Wait()
}
```

- **`chan struct{}`** — a zero-size type used purely as a signal; it carries no value.
- **`wg.Go(func(){})`** — the Go 1.25+ helper that does `Add(1)` + `go` + `Done` in one (loop variables are per-iteration, so `id` is correct).
- All 10 workers are **blocked** on `<-shutdown`.
- **`close(shutdown)`** — a single `close` releases all of them at once (a broadcast), instead of one send per worker.
- **`wg.Wait()`** — waits for everyone to finish.
- The print **order is nondeterministic**.

---

## Expected Output (order varies)

```
worker 4 stopping
worker 8 stopping
worker 7 stopping
...
worker 2 stopping
```

## Key Takeaways

1. **`close` is a broadcast** — every receiver is released together.
2. **`chan struct{}`** — a pure signal channel.
3. **`wg.Go`** — the modern WaitGroup helper.
4. **`wg.Wait()`** — synchronizes on completion.