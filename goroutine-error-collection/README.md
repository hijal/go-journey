# goroutine-error-collection

Go-তে **goroutine-এর error collection (`errors.Join` + index-based slice write)** — parallel charge sender-মাঝে partial failure।

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
)
```

`errors` (`New`, `Join`, `Is`) + `sync` (`WaitGroup`)।

### Line 9

```go
var errCardExpired = errors.New("card expired")
```

Sentinel error।

### Lines 11–16

```go
func chargeCard(cardID string) error {
	if cardID == "card-expired" {
		return errCardExpired
	}
	return nil
}
```

Simulated charge — নির্দিষ্ট card-এর জন্য fail।

### Lines 19–31

```go
	cards := []string{"card-ok-1", "card-expired", "card-ok-2"}

	errs := make([]error, len(cards))

	var wg sync.WaitGroup

	for i, card := range cards {
		wg.Go(func() {
			if err := chargeCard(card); err != nil {
				errs[i] = fmt.Errorf("charge %s: %w", card, err)
			}
		})
	}
```

**Concurrent error collection:**

- `errs` — index-এ aligned error slot (প্রতি goroutine শুধু নিজের slot-ই লিখে)।
- **Different slot → no race** (distinct element = distinct memory)।
- Go 1.22+ loop var — `i`, `card` per-iteration safe।
- fail হলে `fmt.Errorf("... %w")` — context + cause।

### Lines 33–38

```go
	wg.Wait()

	if err := errors.Join(errs...); err != nil {
		fmt.Println("some charged failed:")
		fmt.Println(err)
		fmt.Println("any expired card?", errors.Is(err, errCardExpired))
	}
```

**`errors.Join(errs...)`** — সব non-nil error-কে একটা multi-error-এ join; কোনো error না থাকলে `nil`। `errors.Is` multi-error-টার ভেতরে scan করে।

---

## Expected Output

```
some charged failed:
charge card-expired: card expired
any expired card? true
```

## মূল শিক্ষা / Key Takeaways

1. **`errors.Join`** — সব goroutine-এর error-কে একত্রিত error।
2. **Index-aligned slice** — concurrency-এ কাঠামো-safe collection।
3. **Go 1.22 loop vars** — closure-এ per-iteration capture।
4. **`%w` + context** — কোন card fail হলো।
5. **`errors.Is` on Join** — individual sentinel-এ-ও reach।

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
)
```

`errors` (`New`, `Join`, `Is`) + `sync` (`WaitGroup`).

### Line 9

```go
var errCardExpired = errors.New("card expired")
```

A sentinel error.

### Lines 11–16

```go
func chargeCard(cardID string) error {
	if cardID == "card-expired" {
		return errCardExpired
	}
	return nil
}
```

A simulated charge — fails for a specific card.

### Lines 19–31

```go
	cards := []string{"card-ok-1", "card-expired", "card-ok-2"}

	errs := make([]error, len(cards))

	var wg sync.WaitGroup

	for i, card := range cards {
		wg.Go(func() {
			if err := chargeCard(card); err != nil {
				errs[i] = fmt.Errorf("charge %s: %w", card, err)
			}
		})
	}
```

**Concurrent error collection:**

- `errs` — an error slot aligned by index (goroutine i writes only its own slot).
- **Different slots → no race** (distinct elements are distinct memory).
- Go 1.22+ loop vars — `i`, `card` are per-iteration safe.
- On failure `fmt.Errorf("... %w")` — context + cause.

### Lines 33–38

```go
	wg.Wait()

	if err := errors.Join(errs...); err != nil {
		fmt.Println("some charged failed:")
		fmt.Println(err)
		fmt.Println("any expired card?", errors.Is(err, errCardExpired))
	}
```

**`errors.Join(errs...)`** — joins all non-nil errors into one multi-error; returns `nil` when there are none. `errors.Is` scans inside the multi-error.

---

## Expected Output

```
some charged failed:
charge card-expired: card expired
any expired card? true
```

## Key Takeaways

1. **`errors.Join`** — aggregates every goroutine's error into one.
2. **Index-aligned slice** — structured-safe collection under concurrency.
3. **Go 1.22 loop vars** — per-iteration capture in closures.
4. **`%w` + context** — which card failed.
5. **`errors.Is` on Join** — still reaches individual sentinels.