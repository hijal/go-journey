# context-done-err

Go-তে **`ctx.Done()` দিয়ে timeout** — `select`-এ result বনাম context, `ctx.Err()` `%w` wrap।

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
	"context"
	"fmt"
	"time"
)
```

`context` (WithTimeout), `time` (Sleep/Millisecond)।

### Lines 9–22

```go
func fetchRate(ctx context.Context) (float64, error) {
	result := make(chan float64, 1)
	go func() {
		time.Sleep(time.Millisecond * 300)
		result <- 119.55
	}()

	select {
	case r := <-result:
		return r, nil

	case <-ctx.Done():
		return 0, fmt.Errorf("fetch rate: %w", ctx.Err())
	}
}
```

- **`make(chan float64, 1)`** — buffer 1, তাই timeout হলেও goroutine-টা send করে **block করে না** (goroutine leak নেই)।
- **`select`** — দুই পথ:
  - `case r := <-result` — 300ms পরে value এলে সফল।
  - `case <-ctx.Done()` — deadline (100ms) আগে শেষ হলে error।
- **`ctx.Err()`** = `context.DeadlineExceeded`; `%w` দিয়ে wrap।

### Lines 25–37

```go
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)

	defer cancel()

	rate, err := fetchRate(ctx)

	if err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Println("rate:", rate)
}
```

`defer cancel()` — context-এর resource সবসময় মুক্ত করতে।

---

## Expected Output

```
error fetch rate: context deadline exceeded
```

## মূল শিক্ষা / Key Takeaways

1. **`ctx.Done()`** — cancel/deadline-এর signal channel।
2. **`select`** — কাজ হলে result, না হলে ctx।
3. **Buffered chan 1** — leak-free cancellation।
4. **`%w` + `ctx.Err()`** — unwrappable error (`errors.Is`)।

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
	"context"
	"fmt"
	"time"
)
```

`context` (WithTimeout), `time` (Sleep/Millisecond).

### Lines 9–22

```go
func fetchRate(ctx context.Context) (float64, error) {
	result := make(chan float64, 1)
	go func() {
		time.Sleep(time.Millisecond * 300)
		result <- 119.55
	}()

	select {
	case r := <-result:
		return r, nil

	case <-ctx.Done():
		return 0, fmt.Errorf("fetch rate: %w", ctx.Err())
	}
}
```

- **`make(chan float64, 1)`** — a buffer of 1, so even on timeout the goroutine's send doesn't block (no goroutine leak).
- **`select`** — two paths:
  - `case r := <-result` — success if the value arrives after 300ms.
  - `case <-ctx.Done()` — error when the 100ms deadline fires first.
- **`ctx.Err()`** is `context.DeadlineExceeded`, wrapped with `%w`.

### Lines 25–37

```go
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)

	defer cancel()

	rate, err := fetchRate(ctx)

	if err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Println("rate:", rate)
}
```

`defer cancel()` — always releases the context's resources.

---

## Expected Output

```
error fetch rate: context deadline exceeded
```

## Key Takeaways

1. **`ctx.Done()`** — the channel signalled on cancel/deadline.
2. **`select`** — result if it lands, ctx if it doesn't.
3. **Buffered channel of 1** — leak-free cancellation.
4. **`%w` + `ctx.Err()`** — an unwrappable error (`errors.Is`).