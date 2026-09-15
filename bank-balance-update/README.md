# bank-balance-update

Go-তে **pointer-receiver method + validation guard + balance mutation** শেখার ছোট example — bank account deposit/withdraw।

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

`fmt` — `Println`, `Errorf`, `Printf`।

### Lines 5–8

```go
type Account struct {
	Owner   string
	Balance float64
}
```

`Account` struct — মালিকের নাম + ব্যালেন্স (float — টাকা fractional হতে পারে)।

### Lines 10–17

```go
func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("deposit amount must be positive, got %.2f", amount)
	}

	a.Balance += amount
	return nil
}
```

**Deposit:**

- `*Account` receiver — মূল struct-কেই mutate করে (value receiver হলে copy-তেই ঢুকত)।
- Guard: `amount <= 0` → error (concise `.2f` message)।
- `a.Balance += amount` — success-এ ব্যালেন্স বাড়ে।

### Lines 19–31

```go
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("withdrawal amount must be positive, got %.2f", amount)
	}

	if amount > a.Balance {
		return fmt.Errorf("insufficient funds: balance %.2f, requested %.2f", a.Balance, amount)
	}

	a.Balance -= amount

	return nil
}
```

**Withdraw — দুই-স্তর guard:**

1. `amount <= 0` — positive-check।
2. `amount > a.Balance` — insufficient-funds-check (balance + requested দুটোই error-এ)।

Reject না-হলে `a.Balance -= amount`।

### Lines 33–44

```go
func main() {
	acc := &Account{Owner: "Alice", Balance: 100.0}

	if err := acc.Deposit(50); err != nil {
		fmt.Println("deposit error:", err)
	}

	if err := acc.Withdraw(30); err != nil {
		fmt.Println("withdraw error:", err)
	}

	fmt.Printf("%s's balance: %.2f\n", acc.Owner, acc.Balance)
}
```

- প্রাথমিক ব্যালেন্স 100।
- `Deposit(50)` → 150, `Withdraw(30)` → 120।
- error হলে print, না-হলে ব্যালেন্স-continue।

---

## Expected Output

```
Alice's balance: 120.00
```

## মূল শিক্ষা / Key Takeaways

1. **Pointer receiver** — struct mutation।
2. **Validation guards** — positive + insufficient-funds।
3. **`fmt.Errorf`** — descriptive error (`.2f`)।
4. **Fail-safe ops** — ভুল হলে balance অপরিবর্তিত।

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

`fmt` — for `Println`, `Errorf`, `Printf`.

### Lines 5–8

```go
type Account struct {
	Owner   string
	Balance float64
}
```

The `Account` struct — owner's name + balance (float, money may be fractional).

### Lines 10–17

```go
func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("deposit amount must be positive, got %.2f", amount)
	}

	a.Balance += amount
	return nil
}
```

**Deposit:**

- `*Account` receiver — mutates the real struct (a value receiver would only touch a copy).
- Guard: `amount <= 0` → error (concise `.2f` message).
- `a.Balance += amount` — balance grows on success.

### Lines 19–31

```go
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("withdrawal amount must be positive, got %.2f", amount)
	}

	if amount > a.Balance {
		return fmt.Errorf("insufficient funds: balance %.2f, requested %.2f", a.Balance, amount)
	}

	a.Balance -= amount

	return nil
}
```

**Withdraw — a two-tier guard:**

1. `amount <= 0` — a positivity check.
2. `amount > a.Balance` — an insufficient-funds check (both values appear in the error).

Unless rejected, `a.Balance -= amount`.

### Lines 33–44

```go
func main() {
	acc := &Account{Owner: "Alice", Balance: 100.0}

	if err := acc.Deposit(50); err != nil {
		fmt.Println("deposit error:", err)
	}

	if err := acc.Withdraw(30); err != nil {
		fmt.Println("withdraw error:", err)
	}

	fmt.Printf("%s's balance: %.2f\n", acc.Owner, acc.Balance)
}
```

- Starts with a 100 balance.
- `Deposit(50)` → 150, `Withdraw(30)` → 120.
- Print the error if it happens; otherwise the balance flow continues.

---

## Expected Output

```
Alice's balance: 120.00
```

## Key Takeaways

1. **Pointer receiver** — struct mutation.
2. **Validation guards** — positive + insufficient funds.
3. **`fmt.Errorf`** — descriptive errors (`.2f`).
4. **Fail-safe ops** — a failed operation leaves the balance untouched.