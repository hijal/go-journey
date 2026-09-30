# report-worker-graceful

Go-তে **graceful shutdown** — `signal.NotifyContext` + `context.WithTimeout` (চেইন করা), `select`-এ `ctx.Done()`, এবং যা শেষ হয়নি তার হিসাব রাখা।

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

### Lines 3–9

```go
import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"time"
)
```

`context` (WithTimeout/Err/Done), `log/slog` (structured logger), `os/signal` (NotifyContext), `time`।

### Lines 11–19

```go
func newLogger() *slog.Logger {
	opts := &slog.HandlerOptions{ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey {
			return slog.Attr{}
		}
		return a
	}}
	return slog.New(slog.NewTextHandler(os.Stdout, opts))
}
```

- **`log/slog`** — stdlib structured logger (`key, value` জোড়া)।
- **`ReplaceAttr`** — timestamp বাদ দিয়ে log পড়তে সহজ, output ছোট ও deterministic।

### Lines 21–23

```go
func processReport(id int) {
	time.Sleep(time.Millisecond * 40)
}
```

প্রতিটা report-এ ৪০ms লাগে (নকল দীর্ঘ কাজ)। তাই ১০টা job ১০০ms deadline-এর মধ্যে শেষ হবে না — ঠিক এটাই শেখার উদ্দেশ্য।

### Lines 25–46

```go
func consume(ctx context.Context, log *slog.Logger, jobs <-chan int) int {
	processed := 0

	for {
		if ctx.Err() != nil {
			return processed
		}

		select {
		case <-ctx.Done():
			return processed
		case id, ok := <-jobs:
			if !ok {
				return processed
			}
			log.Info("report started", "id", id)
			processReport(id)
			processed++
			log.Info("report finished", "id", id)
		}
	}
}
```

- **`<-ctx.Done()`** — shutdown signal; এটি `select`-এর case, তাই কাজ চলাকালীনও সাথে সাথে ধরা পড়ে।
- **`ctx.Err() != nil`** — loop-এর শুরুতে আরেকটা পরীক্ষা, যাতে select-এর আগেই বোঝা যায় context মরে গেছে কিনা।
- **`id, ok := <-jobs`** — `!ok` মানে queue শেষ (স্বাভাবিক শেষ)।
- **তিনটি exit পথ**: context timeout, signal, অথবা queue শেষ — তিনটাতেই `processed` ফেরত হয়, তাই হিসাব হারায় না।
- **`processReport` ctx পায় না** — কাজ চলাকালীন থামানো যায় না, কাজ শেষ হলেই পরের iteration-এ সিদ্ধান্ত হয় (graceful মানে এটাই)।

### Lines 48–55

```go
func main() {
	log := newLogger()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()
```

- **`signal.NotifyContext(..., os.Interrupt)`** — Ctrl+C পেলেই `ctx` cancel হবে (graceful shutdown-এর আসল ট্রিগার)।
- **`defer stop()`** — signal handler সরিয়ে দেয়, resource release।
- **`context.WithTimeout(ctx, ...)`** — **গুরুত্বপূর্ণ:** দ্বিতীয় context-টি **আগের `ctx`-এর উপরেই** তৈরি হয় (চেইন করা), তাই timeout হলেও Ctrl+C-ও কাজ করে।
  - ✗ যদি `context.WithTimeout(context.Background(), ...)` লেখা হতো, নতুন context আগেরটাকে ছিঁড়ে ফেলত — সিগনাল path নিষ্ক্রিয় হয়ে যেত, আর শুধু timeout-ই কাজ করত।
- **`defer cancel()`** — timeout timer ফ্রি করে, অপ্রয়োজনীয় context আটকে রাখে না।

### Lines 57–68

```go
	jobs := make(chan int, 10)

	for id := range 10 {
		jobs <- id + 1
	}
	close(jobs)

	done := make(chan int)

	go func() { done <- consume(ctx, log, jobs) }()
	processed := <-done
	log.Info("shutdown complete", "processed", processed, "left_in_queue", len(jobs), "reason", ctx.Err())
}
```

- **buffer 10** — সব ১০টা job পাঠানো যায় consumer চালু না হওয়া পর্যন্ত (deadlock-এড়ানো)।
- **`close(jobs)`** — queue শেষ হয়েছে বলে ইঙ্গিত, তবে deadline আগে আসতে পারে।
- **`done` channel** — worker-এর ফেরত মান গ্রহণ; `go func(){ done <- ... }()` বন্ধ না করলে goroutine leak হতো।
- **`left_in_queue`** — ১০-এর মধ্যে ৩টি শেষ, **৭টি বাকি**।
- **`reason=ctx.Err()`** — কেন থামল (`context deadline exceeded`)।

---

## Expected Output

```
level=INFO msg="report started" id=1
level=INFO msg="report finished" id=1
level=INFO msg="report started" id=2
level=INFO msg="report finished" id=2
level=INFO msg="report started" id=3
level=INFO msg="report finished" id=3
level=INFO msg="shutdown complete" processed=3 left_in_queue=7 reason="context deadline exceeded"
```

## পর্যবেক্ষণ

১০০ms deadline-এ ৪০ms-এর ৩টা কাজ (১২০ms > ১০০ms) বাকি ৭টা queue-তে থেকে যায়। `signal.NotifyContext` ও `WithTimeout`-এর চেইন দুটোই চালু থাকলেও এখানে timeout এগিয়ে আসে, তাই `reason` দেখায় `context deadline exceeded`; Ctrl+C চাপলে একই জায়গায় `context canceled` দেখাবে।

## মূল শিক্ষা / Key Takeaways

1. **Context চেইন করো** — `WithTimeout(ctx, ...)`; `context.Background()` পুনর্ব্যবহার করলে সিগনাল হারিয়ে যায়।
2. **`defer stop()` + `defer cancel()`** — দুটোই resource release, একটাও বাদ দেওয়া যাবে না।
3. **`select`-এ `<-ctx.Done()`** — চলমান কাজের মাঝেই বাতিলের অনুরোধ ধরা পড়ে।
4. **তিনটি exit পথ** — timeout, signal, queue শেষ; প্রতিটিতেই অবস্থান ফেরত দেওয়া হয়।
5. **`left_in_queue`** — graceful হলেও কাজ অসম্পূর্ণ থাকতে পারে, তাই হিসাব রাখা জরুরি।
6. **buffered `jobs` channel** — producer আটকে না যায়।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Lines 3–9

```go
import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"time"
)
```

`context` (WithTimeout/Err/Done), `log/slog` (structured logger), `os/signal` (NotifyContext), `time`.

### Lines 11–19

```go
func newLogger() *slog.Logger {
	opts := &slog.HandlerOptions{ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey {
			return slog.Attr{}
		}
		return a
	}}
	return slog.New(slog.NewTextHandler(os.Stdout, opts))
}
```

- **`log/slog`** is the stdlib structured logger (`key, value` pairs).
- **`ReplaceAttr`** drops the timestamp, keeping the log short and deterministic to read.

### Lines 21–23

```go
func processReport(id int) {
	time.Sleep(time.Millisecond * 40)
}
```

Each report takes 40ms (mocked long work). So 10 jobs cannot finish inside the 100ms deadline — that is exactly the point of the example.

### Lines 25–46

```go
func consume(ctx context.Context, log *slog.Logger, jobs <-chan int) int {
	processed := 0

	for {
		if ctx.Err() != nil {
			return processed
		}

		select {
		case <-ctx.Done():
			return processed
		case id, ok := <-jobs:
			if !ok {
				return processed
			}
			log.Info("report started", "id", id)
			processReport(id)
			processed++
			log.Info("report finished", "id", id)
		}
	}
}
```

- **`<-ctx.Done()`** is a `select` case, so a shutdown request is noticed even while work is in flight.
- **`ctx.Err() != nil`** is a second check at the top of the loop, catching a dead context before the `select`.
- **`id, ok := <-jobs`** — `!ok` means the queue ended (a normal finish).
- **Three exit paths** — context timeout, signal, or queue drained; all of them return `processed`, so no work is unaccounted for.
- **`processReport` never sees the context** — running work is not interrupted; the decision comes on the next iteration (that is what "graceful" means here).

### Lines 48–55

```go
func main() {
	log := newLogger()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()
```

- **`signal.NotifyContext(..., os.Interrupt)`** cancels `ctx` on Ctrl+C (the real graceful-shutdown trigger).
- **`defer stop()`** unregisters the signal handler, releasing resources.
- **`context.WithTimeout(ctx, ...)`** — **Key point:** the second context is derived **from the first `ctx`** (chained), so Ctrl+C still works even after the timeout is added.
  - ✗ Writing `context.WithTimeout(context.Background(), ...)` would build a brand-new context and silently discard the signal one, leaving only the timeout alive.
- **`defer cancel()`** frees the timer and stops the context leaking.

### Lines 57–68

```go
	jobs := make(chan int, 10)

	for id := range 10 {
		jobs <- id + 1
	}
	close(jobs)

	done := make(chan int)

	go func() { done <- consume(ctx, log, jobs) }()
	processed := <-done
	log.Info("shutdown complete", "processed", processed, "left_in_queue", len(jobs), "reason", ctx.Err())
}
```

- **A buffer of 10** lets all 10 jobs be sent before the consumer even starts (avoiding a deadlock).
- **`close(jobs)`** signals that the queue is complete, though the deadline may arrive first.
- **The `done` channel** receives the worker's return value; without the `go func(){ done <- ... }()` wrapper the send would block and leak the goroutine.
- **`left_in_queue`** is **7** (3 of 10 finished).
- **`reason=ctx.Err()`** explains why it stopped: `context deadline exceeded`.

---

## Expected Output

```
level=INFO msg="report started" id=1
level=INFO msg="report finished" id=1
level=INFO msg="report started" id=2
level=INFO msg="report finished" id=2
level=INFO msg="report started" id=3
level=INFO msg="report finished" id=3
level=INFO msg="shutdown complete" processed=3 left_in_queue=7 reason="context deadline exceeded"
```

## Observation

Within the 100ms deadline only 3 of the 40ms jobs fit (120ms > 100ms), leaving 7 in the queue. Both `signal.NotifyContext` and the `WithTimeout` chain stay live, but the timeout wins here, so `reason` reports `context deadline exceeded`; press Ctrl+C instead and the same line reports `context canceled`.

## Key Takeaways

1. **Chain your contexts** — `WithTimeout(ctx, ...)`; reusing `context.Background()` silently kills signal handling.
2. **`defer stop()` + `defer cancel()`** — both release resources; neither may be skipped.
3. **`<-ctx.Done()` inside `select`** — a cancellation is noticed mid-flight.
4. **Three exit paths** — timeout, signal, drained queue; each one reports its state.
5. **`left_in_queue`** — graceful shutdown still leaves work undone, so the count matters.
6. **A buffered `jobs` channel** — keeps the producer from blocking.