# batch-settlement-errgroup

Go-তে **`golang.org/x/sync/errgroup`** — bounded concurrent batch (limit 3), প্রথম error-এ early cancel, `errors.Is` sentinel-check।

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
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)
```

`errgroup` — external dependency (`go get golang.org/x/sync`), error-returning goroutine-গ্রুপ।

### Lines 12–17

```go
type task struct {
	account string
	delay   time.Duration
}

var errAccountFrozen = errors.New("account frozen")
```

Task + sentinel error।

### Lines 19–32

```go
func settle(ctx context.Context, account string, d time.Duration) error {
	select {
	case <-ctx.Done():
		return fmt.Errorf("settle %s: %w", account, ctx.Err())
	case <-time.After(d):
	}

	if account == "ACC-203" {
		return fmt.Errorf("settle %s: %w", account, errAccountFrozen)
	}

	fmt.Println("settled", account)
	return nil
}
```

**Cancellation-aware work** — cancel-এ immediate return; নির্দিষ্ট account-এ frozen-error।

### Lines 34–51

```go
	g, ctx := errgroup.WithContext(context.Background())

	g.SetLimit(3)

	for _, t := range tasks {
		g.Go(func() error {
			return settle(ctx, t.account, t.delay)
		})
	}
```

- **`WithContext`** — গ্রুপের নিজস্ব ctx; কোনো গোরুটিন error দিলে auto-cancel।
- **`SetLimit(3)`** — সর্বোচ্চ 3 concurrent (worker-pool-like)।
- সব task জমা → গ্রুপ চালায়।

### Lines 53–58

```go
	if err := g.Wait(); err != nil {
		fmt.Println("batch failed:", err)
		fmt.Println("frozen account involved?", errors.Is(err, errAccountFrozen))
		return
	}
	fmt.Println("all accounts settled")
```

**`g.Wait()`** — প্রথম error-টা ফেরায়। এখানে ACC-203-এর frozen error → `errors.Is` → true।

---

## Expected Output

```
settled ACC-201
settled ACC-202
batch failed: settle ACC-203: account frozen
frozen account involved? true
```

## মূল শিক্ষা / Key Takeaways

1. **`errgroup.WithContext`** — সম্মিলিত cancellation।
2. **`SetLimit(n)`** — concurrency cap built-in।
3. **First-error return** — whole-batch fail-fast।
4. **`ctx` propagation** — pending task-গুলোও cancel-এ সাড়া দেয়।
5. **`errors.Is`** — সুস্পষ্ট কারণ-report।

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
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)
```

`errgroup` — an external dependency (`go get golang.org/x/sync`) for an error-returning goroutine group.

### Lines 12–17

```go
type task struct {
	account string
	delay   time.Duration
}

var errAccountFrozen = errors.New("account frozen")
```

A task + a sentinel error.

### Lines 19–32

```go
func settle(ctx context.Context, account string, d time.Duration) error {
	select {
	case <-ctx.Done():
		return fmt.Errorf("settle %s: %w", account, ctx.Err())
	case <-time.After(d):
	}

	if account == "ACC-203" {
		return fmt.Errorf("settle %s: %w", account, errAccountFrozen)
	}

	fmt.Println("settled", account)
	return nil
}
```

**Cancellation-aware work** — returns immediately on cancel; a specific account fails with the frozen error.

### Lines 34–51

```go
	g, ctx := errgroup.WithContext(context.Background())

	g.SetLimit(3)

	for _, t := range tasks {
		g.Go(func() error {
			return settle(ctx, t.account, t.delay)
		})
	}
```

- **`WithContext`** — the group owns a context; any goroutine error auto-cancels it.
- **`SetLimit(3)`** — at most 3 concurrent (worker-pool-like).
- All tasks are submitted, then the group runs them.

### Lines 53–58

```go
	if err := g.Wait(); err != nil {
		fmt.Println("batch failed:", err)
		fmt.Println("frozen account involved?", errors.Is(err, errAccountFrozen))
		return
	}
	fmt.Println("all accounts settled")
```

**`g.Wait()`** — returns the first error. Here it's ACC-203's frozen error → `errors.Is` → true.

---

## Expected Output

```
settled ACC-201
settled ACC-202
batch failed: settle ACC-203: account frozen
frozen account involved? true
```

## Key Takeaways

1. **`errgroup.WithContext`** — collective cancellation.
2. **`SetLimit(n)`** — a built-in concurrency cap.
3. **First-error return** — whole-batch fail-fast.
4. **`ctx` propagation** — pending tasks also react to cancel.
5. **`errors.Is`** — a clear cause report.