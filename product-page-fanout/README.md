# product-page-fanout

Go-তে **`wg.Go` parallel fetch fan-out + `errors.Join` degraded-Report** — product page-ের stock/price/reviews একসাথে আর `~120ms`-এ সংগৃহীত।

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
	"errors"
	"fmt"
	"sync"
	"time"
)
```

`sync` (WaitGroup) + `errors` (Join/Is)।

### Lines 10–14 / 16

```go
type productPage struct {
	stock   int
	price   int
	reviews float64
}

var errReviewsDown = errors.New("reviews service unavailable")
```

Page struct + একটি sentinel error।

### Lines 18–31

```go
func fetchStock(sku string) (int, error) {
	time.Sleep(120 * time.Millisecond)
	return 34, nil
}

func fetchPrice(sku string) (int, error) {
	time.Sleep(120 * time.Millisecond)
	return 2450, nil
}

func fetchReviews(sku string) (float64, error) {
	time.Sleep(60 * time.Millisecond)
	return 0, errReviewsDown
}
```

তিনটা সিমুলেটেড service — stock/price ঠিক আছে, reviews down।

### Lines 38–72

```go
	var (
		page productPage
		wg   sync.WaitGroup
		errs [3]error
	)

	wg.Go(func() { /* fetchStock → page.stock */ })
	wg.Go(func() { /* fetchPrice → page.price */ })
	wg.Go(func() { /* fetchReviews → page.reviews */ })
```

**Fan-out** — তিনটা fetch ঠিক সমান্তরাল। প্রতিটি goroutine নিজের **ভিন্ন field/error-slot**-এ লিখে → কোনো data race নেই।

### Lines 74–79

```go
	wg.Wait()

	fmt.Printf("page %s: stock=%d price=%d reviews=%.1f\n",
		sku, page.stock, page.price, page.reviews)
	fmt.Println("concurrent fetch took",
		time.Since(start).Round(10*time.Millisecond), "(sequential would be ~270ms)")
```

সবচেয়ে ধীর (120ms) fetch-ই মোট সময় — sequential `120+120+60=300ms`-এর বদলে **~120ms**।

### Lines 81–84

```go
	if err := errors.Join(errs[:]...); err != nil {
		fmt.Println("degraded:", err)
		fmt.Println("reviews down?", errors.Is(err, errReviewsDown))
	}
```

**Partial failure graceful** — `errors.Join` সব error জোড়া; `errors.Is` sentinel-এ scan।

---

## Expected Output

```
page SKU-77120: stock=34 price=2450 reviews=0.0
concurrent fetch took 120ms (sequential would be ~270ms)
degraded: reviews: reviews service unavailable
reviews down? true
```

## মূল শিক্ষা / Key Takeaways

1. **Fan-out** — স্বতন্ত্র data-fetch goroutine-টো ভাগে ভাগ।
2. **Elapsed-time win** — 300ms → overall ~120ms-এ।
3. **Race-free writes** — প্রতিটি goroutine নিজস্ব struct-field / errs-slot লিখে।
4. **Graceful degradation** — আংশিক fail-এও page render হয়।
5. **`errors.Join` + `Is`** — সব error-এর কাছে পৌঁছানো + cause-check।

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
	"errors"
	"fmt"
	"sync"
	"time"
)
```

`sync` (WaitGroup) + `errors` (Join/Is).

### Lines 10–14 / 16

```go
type productPage struct {
	stock   int
	price   int
	reviews float64
}

var errReviewsDown = errors.New("reviews service unavailable")
```

A page struct + a sentinel error.

### Lines 18–31

```go
func fetchStock(sku string) (int, error) {
	time.Sleep(120 * time.Millisecond)
	return 34, nil
}

func fetchPrice(sku string) (int, error) {
	time.Sleep(120 * time.Millisecond)
	return 2450, nil
}

func fetchReviews(sku string) (float64, error) {
	time.Sleep(60 * time.Millisecond)
	return 0, errReviewsDown
}
```

Three simulated services — stock/price fine, reviews down.

### Lines 38–72

```go
	var (
		page productPage
		wg   sync.WaitGroup
		errs [3]error
	)

	wg.Go(func() { /* fetchStock → page.stock */ })
	wg.Go(func() { /* fetchPrice → page.price */ })
	wg.Go(func() { /* fetchReviews → page.reviews */ })
```

**Fan-out** — all three fetches run in parallel. Each goroutine writes to its own **distinct field/error slot** → no data race.

### Lines 74–79

```go
	wg.Wait()

	fmt.Printf("page %s: stock=%d price=%d reviews=%.1f\n",
		sku, page.stock, page.price, page.reviews)
	fmt.Println("concurrent fetch took",
		time.Since(start).Round(10*time.Millisecond), "(sequential would be ~270ms)")
```

The slowest fetch (120ms) sets the total — **~120ms** instead of `120+120+60=300ms`.

### Lines 81–84

```go
	if err := errors.Join(errs[:]...); err != nil {
		fmt.Println("degraded:", err)
		fmt.Println("reviews down?", errors.Is(err, errReviewsDown))
	}
```

**Graceful partial failure** — `errors.Join` combines errors; `errors.Is` scans for the sentinel.

---

## Expected Output

```
page SKU-77120: stock=34 price=2450 reviews=0.0
concurrent fetch took 120ms (sequential would be ~270ms)
degraded: reviews: reviews service unavailable
reviews down? true
```

## Key Takeaways

1. **Fan-out** — split independent data fetches across goroutines.
2. **Elapsed-time win** — one ~120ms round instead of ~300ms in sequence.
3. **Race-free writes** — each goroutine writes its own struct field / err slot.
4. **Graceful degradation** — the page still renders on partial failure.
5. **`errors.Join` + `Is`** — reach every error and still cause-check.