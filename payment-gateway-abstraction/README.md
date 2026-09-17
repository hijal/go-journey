# payment-gateway-abstraction

Go-তে **interface abstraction + multiple implementations + sentinel error (`errors.Is` ও `%w` wrap)** শেখার ছোট example — payment gateway (Stripe/Bkash) checkout।

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
	"errors"
	"fmt"
	"log"
)
```

- `errors` — `New`, `Is`।
- `log` — `Fatal`।
- `fmt` — `Sprintf`, `Errorf`, `Printf`।

### Lines 9–13

```go
type PaymentProcessor interface {
	Charge(customerID string, amountCents int64) (txId string, err error)
	Refund(txID string, amountCents int64) error
	Name() string
}
```

**Core abstraction** — ৩টা method: charge, refund, name। দুই gateway এই interface implement করবে।

### Line 15

```go
var ErrInvalidAmount = errors.New("amount must be positive")
```

**Sentinel error** — package-level shared value, `errors.Is` দিয়ে check করা যাবে।

### Lines 17–23

```go
type StripeProcessor struct {
	APIKey string
}

type BkashProcessor struct {
	MerchantNumber string
}
```

দুটো concrete provider struct (ভিন্ন config field)।

### Lines 25–43

```go
var _ PaymentProcessor = (*StripeProcessor)(nil)

func (s *StripeProcessor) Charge(customerID string, amountCents int64) (txId string, err error) {
	if amountCents <= 0 {
		return "", ErrInvalidAmount
	}
	return fmt.Sprintf("ch_%s_%d", customerID, amountCents), nil
}

func (s *StripeProcessor) Refund(txID string, amountCents int64) error {
	if amountCents <= 0 {
		return ErrInvalidAmount
	}
	return nil
}

func (s *StripeProcessor) Name() string {
	return "stripe"
}
```

Stripe implementation: validation + prefixed tx id `ch_...` + "stripe" নাম + compile-time check।

### Lines 45–63

```go
var _ PaymentProcessor = (*BkashProcessor)(nil)

func (b *BkashProcessor) Charge(customerID string, amountCents int64) (txId string, err error) {
	if amountCents <= 0 {
		return "", ErrInvalidAmount
	}
	return fmt.Sprintf("BK-%s-%d", customerID, amountCents), nil
}
```

Bkash implementation — একই interface, ভিন্ন tx format `BK-...`।

### Lines 65–72

```go
func Checkout(p PaymentProcessor, customerID string, amountCents int64) (string, error) {
	txID, err := p.Charge(customerID, amountCents)

	if err != nil {
		return "", fmt.Errorf("checkout via %s: %w", p.Name(), err)
	}
	return txID, nil
}
```

**Polymorphic helper** — interface param; `%w` দিয়ে error wrap (cause chain সংরক্ষণ)।

### Lines 74–87

```go
	providers := map[string]PaymentProcessor{
		"US": &StripeProcessor{APIKey: "sk_test_123"},
		"BD": &BkashProcessor{MerchantNumber: "01700000000"},
	}

	for _, country := range []string{"US", "BD"} {
		txID, err := Checkout(providers[country], "c_123", 2599)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: charged via %s, tx=%s\n", country, providers[country].Name(), txID)
	}
```

Country → provider map, loop-এ polymorphic dispatch (US → Stripe, BD → Bkash)।

### Lines 89–91

```go
	if _, err := Checkout(providers["US"], "cust42", 0); errors.Is(err, ErrInvalidAmount) {
		fmt.Println("rejected:", err)
	}
```

**`errors.Is`** — `%w` wrapper ভেদ করে original sentinel পায় → "rejected: ..." print।

---

## Expected Output

```
US: charged via stripe, tx=ch_c_123_2599
BD: charged via bkash, tx=BK-c_123-2599
rejected: checkout via stripe: amount must be positive
```

## মূল শিক্ষা / Key Takeaways

1. **Interface abstraction** — এক contract, একাধিক provider।
2. **Sentinel error + `errors.Is`** — shared value identity check।
3. **`%w` wrapping** — cause chain অটুট রাখে।
4. **`var _ I = (*T)(nil)`** — compile-time interface gate।
5. **Map + interface** — runtime provider selection।

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
	"errors"
	"fmt"
	"log"
)
```

- `errors` — `New`, `Is`.
- `log` — `Fatal`.
- `fmt` — `Sprintf`, `Errorf`, `Printf`.

### Lines 9–13

```go
type PaymentProcessor interface {
	Charge(customerID string, amountCents int64) (txId string, err error)
	Refund(txID string, amountCents int64) error
	Name() string
}
```

**The core abstraction** — 3 methods: charge, refund, name. Both gateways implement it.

### Line 15

```go
var ErrInvalidAmount = errors.New("amount must be positive")
```

A **sentinel error** — a package-level shared value checkable via `errors.Is`.

### Lines 17–23

```go
type StripeProcessor struct {
	APIKey string
}

type BkashProcessor struct {
	MerchantNumber string
}
```

Two concrete provider structs (with different config fields).

### Lines 25–43

```go
var _ PaymentProcessor = (*StripeProcessor)(nil)

func (s *StripeProcessor) Charge(customerID string, amountCents int64) (txId string, err error) {
	if amountCents <= 0 {
		return "", ErrInvalidAmount
	}
	return fmt.Sprintf("ch_%s_%d", customerID, amountCents), nil
}

func (s *StripeProcessor) Refund(txID string, amountCents int64) error {
	if amountCents <= 0 {
		return ErrInvalidAmount
	}
	return nil
}

func (s *StripeProcessor) Name() string {
	return "stripe"
}
```

The Stripe implementation: validation + prefixed tx id `ch_...` + the name "stripe" + a compile-time check.

### Lines 45–63

```go
var _ PaymentProcessor = (*BkashProcessor)(nil)

func (b *BkashProcessor) Charge(customerID string, amountCents int64) (txId string, err error) {
	if amountCents <= 0 {
		return "", ErrInvalidAmount
	}
	return fmt.Sprintf("BK-%s-%d", customerID, amountCents), nil
}
```

The Bkash implementation — same interface, different tx format `BK-...`.

### Lines 65–72

```go
func Checkout(p PaymentProcessor, customerID string, amountCents int64) (string, error) {
	txID, err := p.Charge(customerID, amountCents)

	if err != nil {
		return "", fmt.Errorf("checkout via %s: %w", p.Name(), err)
	}
	return txID, nil
}
```

**A polymorphic helper** — an interface param; `%w` wraps the error (preserving the cause chain).

### Lines 74–87

```go
	providers := map[string]PaymentProcessor{
		"US": &StripeProcessor{APIKey: "sk_test_123"},
		"BD": &BkashProcessor{MerchantNumber: "01700000000"},
	}

	for _, country := range []string{"US", "BD"} {
		txID, err := Checkout(providers[country], "c_123", 2599)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: charged via %s, tx=%s\n", country, providers[country].Name(), txID)
	}
```

A country → provider map, with polymorphic dispatch in the loop (US → Stripe, BD → Bkash).

### Lines 89–91

```go
	if _, err := Checkout(providers["US"], "cust42", 0); errors.Is(err, ErrInvalidAmount) {
		fmt.Println("rejected:", err)
	}
```

**`errors.Is`** — reaches the original sentinel through the `%w` wrapper → prints "rejected: ...".

---

## Expected Output

```
US: charged via stripe, tx=ch_c_123_2599
BD: charged via bkash, tx=BK-c_123-2599
rejected: checkout via stripe: amount must be positive
```

## Key Takeaways

1. **Interface abstraction** — one contract, multiple providers.
2. **Sentinel error + `errors.Is`** — identity check on a shared value.
3. **`%w` wrapping** — keeps the cause chain intact.
4. **`var _ I = (*T)(nil)`** — a compile-time interface gate.
5. **Map + interface** — runtime provider selection.