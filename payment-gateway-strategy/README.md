# payment-gateway-strategy

Go-তে **Strategy pattern + function type** শেখার ছোট example — `PaymentProcessor func(amount, currency) error` দিয়ে checkout-এ যেকোনো payment gateway (Stripe/PayPal) swap করা।

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

### Line 3

```go
import "fmt"
```

`fmt` — `Printf`, `Errorf`।

### Line 5

```go
type PaymentProcessor func(amount float64, currency string) error
```

`PaymentProcessor` — একটা **function type**:

- Signature: `(amount float64, currency string) error`।
- এটা একটা **strategy interface**: যেকোনো function এই signature-এ মিললেই `PaymentProcessor`-হিসেবে ব্যবহার করা যায় (explicit interface-এর দরকার নেই)।

### Lines 7–10

```go
func processStripe(amount float64, currency string) error {
	fmt.Printf("Processing $%.2f %s via Stripe API...\n", amount, currency)
	return nil
}
```

`processStripe` — একটা concrete strategy:

- Stripe-তে payment প্রসেস করে, message print করে।
- `$%.2f` — দু-দশমিকে float format।
- Success-এ `nil` return।

### Lines 12–15

```go
func ProcessPayPal(amount float64, currency string) error {
	fmt.Printf("Processing $%.2f %s via PayPal API...\n", amount, currency)
	return nil
}
```

`ProcessPayPal` — আরেকটা strategy:

- PayPal-এ payment, একই signature — তাই দুটো-ই `PaymentProcessor`-এ pass করা যায়।
- `ProcessPayPal` **exported** (ক্যাপিটাল P), `processStripe` unexported — কিন্তু কোনো behavioral পার্থক্য নেই, naming convention-এর বিষয় (stripe-টা এই package-এর ভেতরেই private)।

### Lines 17–27

```go
func Checkout(processor PaymentProcessor, amount float64, currency string) error {
	fmt.Println("validating cart...")

	if err := processor(amount, currency); err != nil {
		return fmt.Errorf("payment failed: %w", err)
	}

	fmt.Println("sending recipient email...")

	return nil
}
```

`Checkout` — **client / context**:

- `processor PaymentProcessor` — কোন gateway ব্যবহার হবে সেটা caller ঠিক করে (strategy injection)।
- Cart validate → `processor(amount, currency)` call।
- Error-এ: `fmt.Errorf("payment failed: %w", err)` — **`%w`** দিয়ে original error wrap — caller `errors.Is`/`Unwrap` করতে পারে।
- Success-এ: recipient email message + `nil`।

**Decoupling:** Checkout সরাসরি Stripe/PayPal জানে না — শুধু signature-টা। নতুন gateway যোগ করতে `Checkout` বদলাতে হয় না।

### Lines 29–35

```go
func main() {
	_ = Checkout(processStripe, 99.99, "usd")

	fmt.Println("----")

	_ = Checkout(ProcessPayPal, 10.11, "BDT")
}
```

- `_ =` — error check না করে ignore (এই demo-য় intentionally)।
- প্রথমে Stripe-এ `$99.99 usd`, তারপর PayPal-এ `$10.11 BDT` — একই `Checkout`, ভিন্ন strategy।

---

## Expected Output

```
validating cart...
Processing $99.99 usd via Stripe API...
sending recipient email...
----
validating cart...
Processing $10.11 BDT via PayPal API...
sending recipient email...
```

## মূল শিক্ষা / Key Takeaways

1. **Function type** — `type PaymentProcessor func(...) error` — interface ছাড়াই strategy।
2. **Strategy pattern** — `Checkout`-এ strategy inject, রানটাইমে swap।
3. **Dependency injection** — Checkout আলগা, gateway details বিচ্ছিন্ন।
4. **`%w` wrapping** — processor-এর error-কে context যোগ করে wrap।
5. **Exported/unexported** — `ProcessPayPal` vs `processStripe`, উভয়ই একই type-এ fit।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Line 3

```go
import "fmt"
```

`fmt` — for `Printf`, `Errorf`.

### Line 5

```go
type PaymentProcessor func(amount float64, currency string) error
```

`PaymentProcessor` — a **function type**:

- Signature: `(amount float64, currency string) error`.
- It's a **strategy interface**: any function matching this signature can be used as a `PaymentProcessor` — no explicit interface needed.

### Lines 7–10

```go
func processStripe(amount float64, currency string) error {
	fmt.Printf("Processing $%.2f %s via Stripe API...\n", amount, currency)
	return nil
}
```

`processStripe` — one concrete strategy:

- Processes a payment through Stripe, printing a message.
- `$%.2f` — float formatted to two decimals.
- Returns `nil` on success.

### Lines 12–15

```go
func ProcessPayPal(amount float64, currency string) error {
	fmt.Printf("Processing $%.2f %s via PayPal API...\n", amount, currency)
	return nil
}
```

`ProcessPayPal` — another strategy:

- Processes a payment through PayPal, same signature — both fit `PaymentProcessor`.
- `ProcessPayPal` is **exported** (capital P), `processStripe` unexported — but no behavioral difference, purely a naming convention (stripe stays private to this package).

### Lines 17–27

```go
func Checkout(processor PaymentProcessor, amount float64, currency string) error {
	fmt.Println("validating cart...")

	if err := processor(amount, currency); err != nil {
		return fmt.Errorf("payment failed: %w", err)
	}

	fmt.Println("sending recipient email...")

	return nil
}
```

`Checkout` — the **client / context**:

- `processor PaymentProcessor` — the caller decides which gateway to use (strategy injection).
- Validates the cart → calls `processor(amount, currency)`.
- On error: `fmt.Errorf("payment failed: %w", err)` — wraps the original error with **`%w`** so callers can use `errors.Is`/`Unwrap`.
- On success: messages about the recipient email, then `nil`.

**Decoupling:** Checkout doesn't know Stripe or PayPal directly — just the signature. Adding a new gateway requires no change to `Checkout`.

### Lines 29–35

```go
func main() {
	_ = Checkout(processStripe, 99.99, "usd")

	fmt.Println("----")

	_ = Checkout(ProcessPayPal, 10.11, "BDT")
}
```

- `_ =` — the error is deliberately ignored in this demo.
- First Stripe with `$99.99 usd`, then PayPal with `$10.11 BDT` — same `Checkout`, different strategy.

---

## Expected Output

```
validating cart...
Processing $99.99 usd via Stripe API...
sending recipient email...
----
validating cart...
Processing $10.11 BDT via PayPal API...
sending recipient email...
```

## Key Takeaways

1. **Function type** — `type PaymentProcessor func(...) error` — strategy without an interface.
2. **Strategy pattern** — inject the strategy into `Checkout`, swap at runtime.
3. **Dependency injection** — Checkout stays generic, gateway details stay separate.
4. **`%w` wrapping** — adds context while wrapping the processor's error.
5. **Exported/unexported** — `ProcessPayPal` vs `processStripe`, both fit the same type.