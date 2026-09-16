# multi-payment-checkout

Go-তে **interface polymorphism + cents-to-currency formatting** শেখার ছোট example — multi-payment checkout।

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

`fmt` — `Println`, `Sprintf`।

### Lines 5–7

```go
type PaymentMethod interface {
	Pay(amountCents int64) string
}
```

**Contract** — `Pay(amountCents)` → string। যেকোনো type-এর এটাই body হলে interface-টা satisfies করে।

### Lines 9–15

```go
type Card struct {
	last4 string
}

func (c Card) Pay(amountCents int64) string {
	return fmt.Sprintf("card ****%s charged ৳%.2f", c.last4, float64(amountCents)/100)
}
```

**`Card` implementation** — last-4 ধরে masked card; `amountCents/100` → টাকা (`%.2f` 2-decimal)।

### Lines 17–23

```go
type MobileWallet struct {
	Phone string
}

func (w MobileWallet) Pay(amountCents int64) string {
	return fmt.Sprintf("wallet %s paid ৳%.2f", w.Phone, float64(amountCents)/100)
}
```

**`MobileWallet` implementation** — phone-নির্ভর; একই interface-এর আরেক concrete type।

### Lines 25–27

```go
func checkout(method PaymentMethod, amountCents int64) {
	fmt.Println(method.Pay(amountCents))
}
```

Interface param — concrete type জানা দরকার নেই; `Pay` call-ই dynamic dispatch।

### Lines 29–40

```go
	checkout(Card{last4: "4242"}, 125000)
	checkout(MobileWallet{Phone: "+8801700-000000"}, 59900)

	methods := []PaymentMethod{
		Card{last4: "1111"},
		MobileWallet{Phone: "+8801811-111111"},
	}

	for _, m := range methods {
		checkout(m, 25000)
	}
```

- দুটি আলাদা concrete type interfere-নেই।
- `[]PaymentMethod` — mixed slice + range-এ polymorphic call।

*(cent: 125000 → ৳1250.00, 59900 → ৳599.00, 25000 → ৳250.00।)*

---

## Expected Output

```
card ****4242 charged ৳1250.00
wallet +8801700-000000 paid ৳599.00
card ****1111 charged ৳250.00
wallet +8801811-111111 paid ৳250.00
```

## মূল শিক্ষা / Key Takeaways

1. **Interface polymorphism** — এক `PaymentMethod`, নানা implementations।
2. **Dynamic dispatch** — `method.Pay()` concrete-নির্ভর।
3. **Interface slice** — `[]PaymentMethod` mixed।
4. **Cents to currency** — `int64/100` + `%.2f`।
5. **`int64`** — বড় cent-amount overflow-safe।

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

`fmt` — for `Println`, `Sprintf`.

### Lines 5–7

```go
type PaymentMethod interface {
	Pay(amountCents int64) string
}
```

**The contract** — `Pay(amountCents)` → string. Any type with that method satisfies the interface.

### Lines 9–15

```go
type Card struct {
	last4 string
}

func (c Card) Pay(amountCents int64) string {
	return fmt.Sprintf("card ****%s charged ৳%.2f", c.last4, float64(amountCents)/100)
}
```

**The `Card` implementation** — masks with the last-4 digits; `amountCents/100` converts to the main unit (`%.2f` gives 2 decimals).

### Lines 17–23

```go
type MobileWallet struct {
	Phone string
}

func (w MobileWallet) Pay(amountCents int64) string {
	return fmt.Sprintf("wallet %s paid ৳%.2f", w.Phone, float64(amountCents)/100)
}
```

**The `MobileWallet` implementation** — phone-based; another concrete type behind the same interface.

### Lines 25–27

```go
func checkout(method PaymentMethod, amountCents int64) {
	fmt.Println(method.Pay(amountCents))
}
```

An interface param — the concrete type isn't needed; `Pay` dispatches dynamically.

### Lines 29–40

```go
	checkout(Card{last4: "4242"}, 125000)
	checkout(MobileWallet{Phone: "+8801700-000000"}, 59900)

	methods := []PaymentMethod{
		Card{last4: "1111"},
		MobileWallet{Phone: "+8801811-111111"},
	}

	for _, m := range methods {
		checkout(m, 25000)
	}
```

- Two different concrete types work with no interference.
- `[]PaymentMethod` — a mixed slice; polymorphic calls in the range loop.

*(cents: 125000 → ৳1250.00, 59900 → ৳599.00, 25000 → ৳250.00.)*

---

## Expected Output

```
card ****4242 charged ৳1250.00
wallet +8801700-000000 paid ৳599.00
card ****1111 charged ৳250.00
wallet +8801811-111111 paid ৳250.00
```

## Key Takeaways

1. **Interface polymorphism** — one `PaymentMethod`, many implementations.
2. **Dynamic dispatch** — `method.Pay()` depends on the concrete type.
3. **Interface slices** — `[]PaymentMethod` mixed.
4. **Cents to currency** — `int64/100` + `%.2f`.
5. **`int64`** — overflow-safe for big cent amounts.