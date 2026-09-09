# shared-wallet-closure

Go-তে **multiple closures sharing one captured state** শেখার ছোট example — `openWallet` ৩টা function return করে (deposit, withdraw, current), সব একই `balance` variable-এ read/write করে।

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

Console-এ output print করার জন্য `fmt` package import করা হয়।

### Lines 5–23

```go
func openWallet(owner string) (deposit func(int), withdraw func(int) error, current func() int) {
	balance := 0

	deposit = func(amount int) {
		balance += amount
		fmt.Printf("%s deposit %d, balance %d\n", owner, amount, balance)
	}

	withdraw = func(amount int) error {
		if amount > balance {
			return fmt.Errorf("%s withdraw %d: insufficient funds(balance %d)", owner, amount, balance)
		}

		balance -= amount
		fmt.Printf("%s withdraw %d, balance %d\n", owner, amount, balance)
		return nil
	}
	return deposit, withdraw, func() int { return balance }
}
```

`openWallet` — একটা "wallet maker":

- **Multiple return** — ৩টা named function type: `deposit func(int)`, `withdraw func(int) error`, `current func() int`.
- `balance := 0` — **captured shared state**।
- `deposit` — closure: `balance += amount`, print।
- `withdraw` — closure: 
  - `amount > balance` → error return (insufficient funds)। কিছু reduce হয় না।
  - else `balance -= amount`, print, `nil`।
- শেষ return: `deposit`, `withdraw`, আর একটা **inline anonymous** `current` (শুধু `balance` return)।

**Shared state:** ৩টা function-ই **একই `balance`**-কে capture করে — deposit বাড়ায়, withdraw কমায়, current পড়ে। Function-এর ভেতরে `owner`-ও capture (print-এ ব্যবহার)। State বাইরে visible নয় — private closure state (encapsulation)।

### Lines 25–39

```go
func main() {
	deposit, withdraw, current := openWallet("john")

	deposit(500)
	deposit(300)

	if err := withdraw(200); err != nil {
		fmt.Println("error:", err)
	}

	if err := withdraw(1000); err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("final balance:", current())
}
```

`openWallet("john")` — ৩টা function assign:

- `deposit(500)` → balance 500
- `deposit(300)` → balance 800
- `withdraw(200)` → 800−200 = 600 সফল
- `withdraw(1000)` → 1000 > 600 → error print (`insufficient funds(balance 600)`)
- `current()` → খালি balance 600-টা পড়ে → `final balance: 600`

Function-গুলো আলাদা variable-এ — কিন্তু state শেয়ার করা (তে হুইয়ার ওয়ালেট ৮০০→৬০০ sequence-এ)। 

### Line 40

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
john deposit 500, balance 500
john deposit 300, balance 800
john withdraw 200, balance 600
error: john withdraw 1000: insufficient funds(balance 600)
final balance: 600
```

## মূল শিক্ষা / Key Takeaways

1. **Shared captured state** — ৩টা closure একই `balance`-এ access করে।
2. **Multiple return functions** — deposit/withdraw/current trio।
3. **Encapsulation** — state বাইরে private, শুধু closures-এর মাধ্যমে।
4. **Guard clause** — withdraw-এ insufficient-funds check।
5. **`current` read** — state-কে read-only দেখায় (function-এ wrap)।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Lines 5–23

```go
func openWallet(owner string) (deposit func(int), withdraw func(int) error, current func() int) {
	balance := 0

	deposit = func(amount int) {
		balance += amount
		fmt.Printf("%s deposit %d, balance %d\n", owner, amount, balance)
	}

	withdraw = func(amount int) error {
		if amount > balance {
			return fmt.Errorf("%s withdraw %d: insufficient funds(balance %d)", owner, amount, balance)
		}

		balance -= amount
		fmt.Printf("%s withdraw %d, balance %d\n", owner, amount, balance)
		return nil
	}
	return deposit, withdraw, func() int { return balance }
}
```

`openWallet` — a "wallet maker":

- **Multiple return** — 3 named function types: `deposit func(int)`, `withdraw func(int) error`, `current func() int`.
- `balance := 0` — **captured shared state**.
- `deposit` — closure: `balance += amount`, prints.
- `withdraw` — closure:
  - `amount > balance` → returns an error (insufficient funds). Nothing is deducted.
  - else `balance -= amount`, prints, returns `nil`.
- Final return: `deposit`, `withdraw`, plus an **inline anonymous** `current` (just returns `balance`).

**Shared state:** all 3 functions capture the **same `balance`** — deposit adds, withdraw subtracts, current reads. The `owner` is also captured (used in prints). The state isn't visible outside — private closure state (encapsulation).

### Lines 25–39

```go
func main() {
	deposit, withdraw, current := openWallet("john")

	deposit(500)
	deposit(300)

	if err := withdraw(200); err != nil {
		fmt.Println("error:", err)
	}

	if err := withdraw(1000); err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("final balance:", current())
}
```

`openWallet("john")` — assigns the 3 functions:

- `deposit(500)` → balance 500
- `deposit(300)` → balance 800
- `withdraw(200)` → 800−200 = 600 success
- `withdraw(1000)` → 1000 > 600 → error printed (`insufficient funds(balance 600)`)
- `current()` → reads the current balance 600 → `final balance: 600`

The functions are separate variables — but the state is shared (one wallet going 800 → 600).

### Line 40

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
john deposit 500, balance 500
john deposit 300, balance 800
john withdraw 200, balance 600
error: john withdraw 1000: insufficient funds(balance 600)
final balance: 600
```

## Key Takeaways

1. **Shared captured state** — 3 closures access the same `balance`.
2. **Multiple return functions** — deposit/withdraw/current trio.
3. **Encapsulation** — the state is private outside; only reachable via the closures.
4. **Guard clause** — insufficient-funds check on withdraw.
5. **`current` read** — exposes state read-only (wrapped in a function).