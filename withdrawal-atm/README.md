# withdrawal-atm

Go-তে **`if / else`** শর্ত আর `fmt.Printf` formatting দিয়ে ATM-withdrawal (টাকা তোলা) ভ্যালিডেশন শেখার ছোট example।

**📖 ভাষা নির্বাচন করুন / Choose language:**

[🇧🇩 বাংলা](#bangla) • [🇬🇧 English](#english)

---

<a name="bangla"></a>

## 🇧🇩 বাংলা সংস্করণ

### Lines 1–3

```go
package main

import "fmt"
```

- `package main` — একটা executable program।
- `fmt` — output print করার জন্য।

### Line 5

```go
func main() {
```

Program-এর entry point।

### Lines 6–7

```go
var balance float64 = 4500.0
var requestAmount float64 = 6000.0
```

- `balance` — account-এ বর্তমান টাকা, `4500.0` (float64)।
- `requestAmount` — user-এর তোলার অনুরোধ, `6000.0` (float64)।

দুটোই `var` keyword + explicit type দিয়ে declare করা।

### Lines 9–14

```go
if requestAmount <= balance {
    balance -= requestAmount
    fmt.Printf("Withdrawal successful. New balance: %.2f\n", balance)
} else {
    fmt.Printf("Insufficient funds: requested %.2f, available %.2f\n", requestAmount, balance)
}
```

Withdrawal ভ্যালিডেশন:

- `if requestAmount <= balance` — অনুরোধ-করা পরিমাণ ব্যালেন্সের মধ্যে আছে কিনা।
  - আছে (`balance -= requestAmount` দিয়ে balance থেকে কেটে নেয়): `Withdrawal successful. New balance: <X>`।
  - নেই: `Insufficient funds: requested <X>, available <Y>`।
- এখানে requestAmount (6000) > balance (4500), তাই `else` branch চলে।

`fmt.Printf`-এ `%.2f` float-কে দশমিকের ২ ঘরে দেখায়।

---

## Expected Output

```
Insufficient funds: requested 6000.00, available 4500.00
```

## মূল শিক্ষা / Key Takeaways

1. **`if / else`** — শর্ত true হলে `if` branch, না হলে `else` branch।
2. **`balance -= requestAmount`** — compound assignment (balance থেকে বিয়োগ)।
3. **`var ... float64`** — explicit typed declaration।
4. **`fmt.Printf` / `%.2f`** — float format (২ দশমিক)।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Lines 1–3

```go
package main

import "fmt"
```

- `package main` — an executable program.
- `fmt` — for console output.

### Line 5

```go
func main() {
```

Program entry point.

### Lines 6–7

```go
var balance float64 = 4500.0
var requestAmount float64 = 6000.0
```

- `balance` — the current money in the account, `4500.0` (float64).
- `requestAmount` — the user's withdrawal request, `6000.0` (float64).

Both declared with `var` + an explicit type.

### Lines 9–14

```go
if requestAmount <= balance {
    balance -= requestAmount
    fmt.Printf("Withdrawal successful. New balance: %.2f\n", balance)
} else {
    fmt.Printf("Insufficient funds: requested %.2f, available %.2f\n", requestAmount, balance)
}
```

Withdrawal validation:

- `if requestAmount <= balance` — whether the requested amount fits within the balance.
  - If yes (subtracts it with `balance -= requestAmount`): `Withdrawal successful. New balance: <X>`.
  - If not: `Insufficient funds: requested <X>, available <Y>`.
- Here requestAmount (6000) > balance (4500), so the `else` branch runs.

In `fmt.Printf`, `%.2f` formats a float with 2 decimal places.

---

## Expected Output

```
Insufficient funds: requested 6000.00, available 4500.00
```

## Key Takeaways

1. **`if / else`** — the `if` branch runs when the condition is true, otherwise the `else` branch.
2. **`balance -= requestAmount`** — compound assignment (subtract from balance).
3. **`var ... float64`** — explicit typed declaration.
4. **`fmt.Printf` / `%.2f`** — float formatting (2 decimals).
