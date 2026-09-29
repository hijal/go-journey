# payment-charge-channel

Go-তে **struct-carrying channel + sentinel error** — `chan<- chargeResult{txnID, err}`, `%w` + `errors.Is`।

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
	"errors"
	"fmt"
)
```

`errors` (New, Is), `fmt` (Sprintf, Errorf)।

### Line 8

```go
var errInvalidAmount = errors.New("invalid amount")
```

**Sentinel error** — package-level `var`, তাই শেষে `errors.Is` দিয়ে শনাক্ত করা যায়।

### Lines 10–13

```go
type chargeResult struct {
	txnID string
	err   error
}
```

**Result struct** — channel-এ যা যাই হোক, value + error একসাথে বহন করে (দুটো আলাদা value-return-এর বদলে)।

### Lines 15–23

```go
func charge(card string, amount int, out chan<- chargeResult) {
	if amount <= 0 {
		out <- chargeResult{
			err: fmt.Errorf("charge card %s: %w", card, errInvalidAmount),
		}
		return
	}
	out <- chargeResult{txnID: fmt.Sprintf("TXN-%s-%d", card, amount)}
}
```

- **`chan<- chargeResult`** — শুধু send-only; error হলেও **একই channel**-এ result পাঠায় (ফলে main-এর loop সবসময় একটা value পায়)।
- `%w`-wrap — `errors.Is(res.err, errInvalidAmount)`-এর জন্য।
- `amount <= 0` → `charge card 4242: invalid amount`।

### Lines 25–39

```go
func main() {
	results := make(chan chargeResult)

	for _, amount := range []int{1500, 0} {
		go charge("4242", amount, results)

		res := <-results

		if res.err != nil {
			fmt.Println("failed:", res.err, "| invalid amount?", errors.Is(res.err, errInvalidAmount))
			return
		}

		fmt.Println("charged:", res.txnID)
	}
}
```

- `1500` → সফল: `charged: TXN-4242-1500`।
- `0` → ব্যর্থ: `errors.Is` → `true`, তারপর `return` (পরের amount আর চলে না)।

---

## Expected Output

```
charged: TXN-4242-1500
failed: charge card 4242: invalid amount | invalid amount? true
```

## মূল শিক্ষা / Key Takeaways

1. **Struct on channel** — value ও error একসাথে, একটা path-এ।
2. **`chan<-` direction** — producer হিসেবে send-only।
3. **Sentinel + `%w`** — `errors.Is` দিয়ে typed check।
4. **`return` on error** — fail-fast, বাকি amount বাদ।

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
	"errors"
	"fmt"
)
```

`errors` (New, Is), `fmt` (Sprintf, Errorf).

### Line 8

```go
var errInvalidAmount = errors.New("invalid amount")
```

A **sentinel error** — a package-level `var`, so callers can recognise it later with `errors.Is`.

### Lines 10–13

```go
type chargeResult struct {
	txnID string
	err   error
}
```

A **result struct** — whatever travels over the channel, the value and the error move together (instead of two separate return values).

### Lines 15–23

```go
func charge(card string, amount int, out chan<- chargeResult) {
	if amount <= 0 {
		out <- chargeResult{
			err: fmt.Errorf("charge card %s: %w", card, errInvalidAmount),
		}
		return
	}
	out <- chargeResult{txnID: fmt.Sprintf("TXN-%s-%d", card, amount)}
}
```

- **`chan<- chargeResult`** — send-only, as the producer.
- On failure it still sends on the **same channel**, so main's loop always receives something.
- `%w` wrap — required for `errors.Is(res.err, errInvalidAmount)`.
- `amount <= 0` → `charge card 4242: invalid amount`.

### Lines 25–39

```go
func main() {
	results := make(chan chargeResult)

	for _, amount := range []int{1500, 0} {
		go charge("4242", amount, results)

		res := <-results

		if res.err != nil {
			fmt.Println("failed:", res.err, "| invalid amount?", errors.Is(res.err, errInvalidAmount))
			return
		}

		fmt.Println("charged:", res.txnID)
	}
}
```

- `1500` → success: `charged: TXN-4242-1500`.
- `0` → failure: `errors.Is` → `true`, then `return` (the remaining amounts never run).

---

## Expected Output

```
charged: TXN-4242-1500
failed: charge card 4242: invalid amount | invalid amount? true
```

## Key Takeaways

1. **Struct on a channel** — value and error travel on one path.
2. **The `chan<-` direction** — send-only for the producer.
3. **Sentinel + `%w`** — a typed check via `errors.Is`.
4. **`return` on error** — fail-fast, skipping the rest.