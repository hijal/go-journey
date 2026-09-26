# graceful-shutdown-worker

Go-তে **context-based graceful shutdown** — producer/consumer/heartbeat-চালানো service, timeout-এ সব Stage পরিষ্কারভাবে থেমে যায়।

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
	"context"
	"fmt"
	"sync"
	"time"
)
```

`context` (deadline + cancel signal)।

### Lines 10–23

```go
func heartbeat(ctx context.Context, every time.Duration) {
	ticker := time.NewTicker(every)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("heartbeat: stopping")
			return
		case <-ticker.C:
			fmt.Println("heartbeat: service alive")
		}
	}
}
```

**Heartbeat loop** — `select`-এ ডুয়েল-চ্যানেল: ticker নাকি cancel — Done আসলেই **clean return**।

### Lines 25–39

```go
func consumer(ctx context.Context, name string, queue <-chan string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s: stopping\n", name)
			return
		case msg, ok := <-queue:
			if !ok {
				return
			}
			fmt.Printf("%s: handled %s\n", name, msg)
			time.Sleep(40 * time.Millisecond)
		}
	}
}
```

**Consumer** — message consume করে (40ms kill করে); shutdown চাইলে বর্তমান কাজ শেষ করে এক-জনমেই চলে আসে।

### Lines 41–63

```go
	ctx, stop := context.WithTimeout(context.Background(), 250*time.Millisecond)
	defer stop()

	...
	wg.Go(func() { heartbeat(ctx, 100*time.Millisecond) })
	wg.Go(func() { consumer(ctx, "worker-1", queue) })

	wg.Go(func() {
		for i := 1; ; i++ {
			select {
			case <-ctx.Done():
				return
			case queue <- fmt.Sprintf("msg-%d", i):
				time.Sleep(60 * time.Millisecond)
			}
		}
	})

	wg.Wait()
	fmt.Println("shutdown complete, reason:", context.Cause(ctx))
```

- **250ms timeout** — সময় শেষ হলেই সব কিছুর জন্য cancel।
- Producer — প্রতি 60ms message; Done-এ থামে (queue-র ওপর blocked-থাকা situation-এ-ও)।
- `wg.Wait()` — সব কো-রুটাইন **গ্রেসফুলি বন্ধ**।
- **`context.Cause(ctx)`** — কেন shutdown → `DeadlineExceeded`।

---

## Expected Output (order varies)

```
worker-1: handled msg-1
worker-1: handled msg-2
heartbeat: service alive
worker-1: handled msg-3
worker-1: handled msg-4
heartbeat: service alive
worker-1: handled msg-5
heartbeat: stopping
worker-1: stopping
shutdown complete, reason: context deadline exceeded
```

## মূল শিক্ষা / Key Takeaways

1. **Shutdown via context** — কো-রুটাইনগুলো নিজেরা মতোই পরিষ্কার exit করে।
2. **`select + Done`** — long-running loop-এর প্রয়োজনীয় guard।
3. **`wg.Wait`** — সব Stage বন্ধ-হওয়া নিশ্চিত।
4. **`context.Cause`** — shutdown-এর কারণ report।
5. **No force-kill** — বর্তমান কাজ শেষ/বাতিল, state consistent।

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
	"context"
	"fmt"
	"sync"
	"time"
)
```

`context` (deadline + cancel signal).

### Lines 10–23

```go
func heartbeat(ctx context.Context, every time.Duration) {
	ticker := time.NewTicker(every)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("heartbeat: stopping")
			return
		case <-ticker.C:
			fmt.Println("heartbeat: service alive")
		}
	}
}
```

**The heartbeat loop** — a dual-channel `select`: ticker vs cancel — on `Done`, a **clean return**.

### Lines 25–39

```go
func consumer(ctx context.Context, name string, queue <-chan string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s: stopping\n", name)
			return
		case msg, ok := <-queue:
			if !ok {
				return
			}
			fmt.Printf("%s: handled %s\n", name, msg)
			time.Sleep(40 * time.Millisecond)
		}
	}
}
```

**The consumer** — consumes messages (40ms each); on shutdown it stops at the next safe point.

### Lines 41–63

```go
	ctx, stop := context.WithTimeout(context.Background(), 250*time.Millisecond)
	defer stop()

	...
	wg.Go(func() { heartbeat(ctx, 100*time.Millisecond) })
	wg.Go(func() { consumer(ctx, "worker-1", queue) })

	wg.Go(func() {
		for i := 1; ; i++ {
			select {
			case <-ctx.Done():
				return
			case queue <- fmt.Sprintf("msg-%d", i):
				time.Sleep(60 * time.Millisecond)
			}
		}
	})

	wg.Wait()
	fmt.Println("shutdown complete, reason:", context.Cause(ctx))
```

- **250ms timeout** — cancels everything when time is up.
- Producer — a message every 60ms; stops on `Done` (even if blocked on the queue).
- `wg.Wait()` — every goroutine **stops gracefully**.
- **`context.Cause(ctx)`** — why shutdown happened → `DeadlineExceeded`.

---

## Expected Output (order varies)

```
worker-1: handled msg-1
worker-1: handled msg-2
heartbeat: service alive
worker-1: handled msg-3
worker-1: handled msg-4
heartbeat: service alive
worker-1: handled msg-5
heartbeat: stopping
worker-1: stopping
shutdown complete, reason: context deadline exceeded
```

## Key Takeaways

1. **Shutdown via context** — goroutines exit on their own, cleanly.
2. **`select + Done`** — the essential guard for long-running loops.
3. **`wg.Wait`** — guarantees every stage is stopped.
4. **`context.Cause`** — reports the shutdown reason.
5. **No force-kill** — current work is finished/cancelled, state stays consistent.