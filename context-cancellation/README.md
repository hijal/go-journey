# context-cancellation

Go-তে **`context.WithTimeout` + `ctx.Done()` cancellation + `errors.Is(DeadlineExceeded)`** শেখার ছোট example — long-running report job।

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
	"errors"
	"fmt"
	"time"
)
```

`context` (deadline/cancel), `errors` (`Is`), `time` (delay)।

### Lines 10–20

```go
func generateReport(ctx context.Context) error {
	for step := range 5 {
		select {
		case <-ctx.Done():
			return fmt.Errorf("report stopped at step %d: %w", step, ctx.Err())
		case <-time.After(100 * time.Millisecond):
			fmt.Println("finished step", step)
		}
	}
	return nil
}
```

**Cancellation-aware loop** — প্রতি step-এ `select`:

- **`<-ctx.Done()`** — cancel/deadline হলে channel close হয় → early return, `ctx.Err()` wrapped।
- **`time.After(100ms)`** — normal step finish।

### Line 23

```go
	ctx, cancel := context.WithTimeout(context.Background(), 250*time.Millisecond)
	defer cancel()
```

**`WithTimeout`** — 250ms-এর মধ্যে deadline; `defer cancel()` — শেষে resource cleanup।

### Lines 26–31

```go
	if err := generateReport(ctx); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Println("timeout:", err)
			return
		}
		fmt.Println("unexpected error:", err)
	}
```

**`errors.Is(err, context.DeadlineExceeded)`** — `%w` wrap-এর ভেতরেও sentinel মেলে; timeout-বনাম-unexpected আলাদা করা যায়।

### Timing

| step | window | result |
|---|---|---|
| 0 | 0–100ms | finished |
| 1 | 100–200ms | finished |
| 2 | 200–300ms | **250ms-এ `ctx.Done()` wins** |

---

## Expected Output

```
finished step 0
finished step 1
timeout: report stopped at step 2: context deadline exceeded
```

## মূল শিক্ষা / Key Takeaways

1. **`context.WithTimeout`** — deadline-carrying context।
2. **`<-ctx.Done()`** — cancellation signal (channel close)।
3. **`ctx.Err()`** — `DeadlineExceeded`/`Canceled` কারণ।
4. **`errors.Is(DeadlineExceeded)`** — wrapped-এ-ও matching।
5. **`defer cancel()`** — always release resources।

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
	"errors"
	"fmt"
	"time"
)
```

`context` (deadline/cancel), `errors` (`Is`), `time` (delay).

### Lines 10–20

```go
func generateReport(ctx context.Context) error {
	for step := range 5 {
		select {
		case <-ctx.Done():
			return fmt.Errorf("report stopped at step %d: %w", step, ctx.Err())
		case <-time.After(100 * time.Millisecond):
			fmt.Println("finished step", step)
		}
	}
	return nil
}
```

**A cancellation-aware loop** — `select` at every step:

- **`<-ctx.Done()`** — on cancel/deadline the channel closes → early return with `ctx.Err()` wrapped.
- **`time.After(100ms)`** — the normal step finish.

### Line 23

```go
	ctx, cancel := context.WithTimeout(context.Background(), 250*time.Millisecond)
	defer cancel()
```

**`WithTimeout`** — a deadline at 250ms; `defer cancel()` cleans up at the end.

### Lines 26–31

```go
	if err := generateReport(ctx); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Println("timeout:", err)
			return
		}
		fmt.Println("unexpected error:", err)
	}
```

**`errors.Is(err, context.DeadlineExceeded)`** — still matches the sentinel inside the `%w` wrap; lets you tell a timeout from an unexpected error.

### Timing

| step | window | result |
|---|---|---|
| 0 | 0–100ms | finished |
| 1 | 100–200ms | finished |
| 2 | 200–300ms | **`ctx.Done()` wins at 250ms** |

---

## Expected Output

```
finished step 0
finished step 1
timeout: report stopped at step 2: context deadline exceeded
```

## Key Takeaways

1. **`context.WithTimeout`** — a deadline-carrying context.
2. **`<-ctx.Done()`** — the cancellation signal (channel close).
3. **`ctx.Err()`** — `DeadlineExceeded`/`Canceled` cause.
4. **`errors.Is(DeadlineExceeded)`** — matching even when wrapped.
5. **`defer cancel()`** — always release resources.