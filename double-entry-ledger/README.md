# double-entry-ledger

Go-তে **struct**, **receiver methods**, **pointer receiver**, আর **error-handling** দিয়ে double-entry accounting-র basic debit/credit logic বুঝতে সহায়ক example।

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

Console-এ output print আর error format করার জন্য `fmt` package import করা হয়।

### Lines 5–8

```go
type Account struct {
    Code    string
    balance int64
}
```

`Account` নামক একটা struct type declare করে এর দুটো field:

- `Code` (বড় হাতের → **exported**) — account-কে identify করা একটা code (যেমন `"1000-CASH"`)।
- `balance` (ছোট হাতের → **unexported**) — বর্তমান balance **cents**-এ `int64` হিসেবে সংরক্ষিত। আমরা এটা ছোট হাতের রাখি, যাতে বাইরে থেকে সরাসরি পরিবর্তন না করে শুধু method দিয়ে পরিবর্তন করা যায় (encapsulation)।

### Lines 10–12

```go
func (a *Account) Credit(amountCents int64) {
    a.balance += amountCents
}
```

একটা **method** `Credit` যার **pointer receiver** `(a *Account)`। এটা `a.balance += amountCents` দিয়ে balance-এর সাথে `amountCents` যোগ করে। Pointer receiver ব্যবহার করি যাতে পরিবর্তনটা মূল account-এ প্রভাব ফেলে।

Double-entry accounting-তে এখানে **credit** balance বাড়ায় (যেমন cash account-এ আয় আসা বোঝায়)।

### Lines 14–20

```go
func (a *Account) Debit(amountCents int64) error {
    if a.balance < amountCents {
        return fmt.Errorf("insufficient balance in account %s: have %d, need %d", a.Code, a.balance, amountCents)
    }
    a.balance -= amountCents
    return nil
}
```

`Debit` নামক একটা method যা balance থেকে `amountCents` বিয়োগ করে। এটা একটা `error` return করে, যাতে caller-রা failure manage করতে পারে:

- আগে এটা check করে balance-টা চাওয়া amount-এর চেয়ে **কম** কিনা। যদি হয় (যথেষ্ট টাকা নেই), তাহলে `fmt.Errorf` দিয়ে (placeholders `%s`, `%d` — `a.Code`, `a.balance`, `amountCents`-এর জন্য) একটা error বানিয়ে balance **পরিবর্তন না করেই** return করে।
- নাহলে `a.balance -= amountCents` করে `nil` (no error) return করে।

### Lines 22–24

```go
func (a *Account) Balance() int64 {
    return a.balance
}
```

একটা **getter** method `Balance` যা বর্তমান balance return করে। যেহেতু `balance` unexported, বাইরের কোডের জন্য এটাই নিরাপদ উপায় value পড়ার।

### Line 26

```go
func main() {
```

Program-এর entry point।

### Lines 27–34

```go
cash := &Account{
    Code:    "1000-CASH",
    balance: 0,
}
revenue := &Account{
    Code:    "2000-REV",
    balance: 0,
}
```

`&Account{...}` (pointer) দিয়ে দুটো account তৈরি করা হয়। `cash` হলো কোম্পানির নগদ/cash account আর `revenue` হলো আয়/income account। দুটোর শুরুতে balance `0`।

### Line 36

```go
cash.Credit(50000)
```

`cash`-এর উপর `Credit` method call করে `50000` cents দিয়ে (₹500 বা $500)। এতে cash balance `50000` হয়ে যায়।

### Lines 37–39

```go
if err := cash.Debit(0); err != nil {
    fmt.Println(err)
}
```

`cash.Debit(0)` call করে — `0` debit করা, যা সবসময় success হয়। এটা Go-র **`if` with an initializer** ব্যবহার করে: আগে `Debit` call চলে (এর error `err`-এ যায়), তারপর `err != nil` চেক হয়। যেহেতু `0` debit করা succeeds, `err` হলো `nil` এবং কিছুই print হয় না।

### Lines 41–42

```go
fmt.Println("Cash:", cash.Balance())
fmt.Println("Revenue:", revenue.Balance())
```

`Balance()` getter দিয়ে balance-গুলো print করে। Output:

```
Cash: 50000
Revenue: 0
```

### Lines 44–46

```go
if err := cash.Debit(60000); err != nil {
    fmt.Println("debit failed", err)
}
```

`cash` থেকে `60000` cents debit করার চেষ্টা করে, যেখানে শুধু `50000` আছে। যেহেতু `60000 > 50000`, method একটা error return করে। `if` initializer সেটা capture করে, `err != nil` true হয়, তাই error message print হয়:

```
debit failed insufficient balance in account 1000-CASH: have 50000, need 60000
```

এটা দেখায় কিভাবে insufficient-balance অবস্থা program crash না করেই মসৃণভাবে handle করা যায়।

### Line 47

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
Cash: 50000
Revenue: 0
debit failed insufficient balance in account 1000-CASH: have 50000, need 60000
```

## মূল শিক্ষা / Key Takeaways

1. **Encapsulation** — unexported `balance` field-কে method দিয়েই পরিবর্তন করা হয়।
2. **Pointer receiver** — method নিজের struct-এর মূল value-কে modify করতে `*Account` receiver।
3. **Methods as behavior** — `Credit`, `Debit`, `Balance` — type-এর সাথে attached behavior।
4. **Error handling** — `Debit` error return করে; `if err := ...; err != nil` pattern দিয়ে check।
5. **`fmt.Errorf`** — format করা error message তৈরি করে।
6. **Double-entry concept** — credit বাড়ায়, debit কমানোর আগে balance check করে।

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

Imports the `fmt` package for console output and formatting errors.

### Lines 5–8

```go
type Account struct {
    Code    string
    balance int64
}
```

Declares a struct type `Account` with two fields:

- `Code` (uppercase → **exported**) — a code identifying the account (e.g. `"1000-CASH"`).
- `balance` (lowercase → **unexported**) — the current balance in **cents** stored as an `int64`. We keep it lowercase so it can only be changed through methods (encapsulation), not directly from outside.

### Lines 10–12

```go
func (a *Account) Credit(amountCents int64) {
    a.balance += amountCents
}
```

A **method** `Credit` with a **pointer receiver** `(a *Account)`. It adds `amountCents` to the balance using `a.balance += amountCents`. We use a pointer receiver so the change affects the original account.

In double-entry accounting, a **credit** increases the account balance here (representing income into, say, the cash account).

### Lines 14–20

```go
func (a *Account) Debit(amountCents int64) error {
    if a.balance < amountCents {
        return fmt.Errorf("insufficient balance in account %s: have %d, need %d", a.Code, a.balance, amountCents)
    }
    a.balance -= amountCents
    return nil
}
```

A method `Debit` that subtracts `amountCents` from the balance. It returns an `error` so callers can handle failures:

- First it checks if the balance is **less than** the requested amount. If so (not enough money), it builds an error with `fmt.Errorf` (using `%s`, `%d` placeholders for `a.Code`, `a.balance`, `amountCents`) and returns it **without** changing the balance.
- Otherwise it does `a.balance -= amountCents` and returns `nil` (no error).

### Lines 22–24

```go
func (a *Account) Balance() int64 {
    return a.balance
}
```

A **getter** method `Balance` that returns the current balance. Since `balance` is unexported, this is the safe way for outside code to read it.

### Line 26

```go
func main() {
```

Program entry point.

### Lines 27–34

```go
cash := &Account{
    Code:    "1000-CASH",
    balance: 0,
}
revenue := &Account{
    Code:    "2000-REV",
    balance: 0,
}
```

Creates two accounts using `&Account{...}` (pointers). `cash` is the company's cash account and `revenue` is the income/revenue account. Both start with a `0` balance.

### Line 36

```go
cash.Credit(50000)
```

Calls the `Credit` method on `cash` with `50000` cents (₹500.00 / $500). This increases the cash balance to `50000`.

### Lines 37–39

```go
if err := cash.Debit(0); err != nil {
    fmt.Println(err)
}
```

Calls `cash.Debit(0)` — debiting `0`, which always succeeds. It uses Go's **`if` with an initializer**: the `Debit` call runs first (its error goes into `err`), then `err != nil` is checked. Since debiting `0` succeeds, `err` is `nil` and nothing prints.

### Lines 41–42

```go
fmt.Println("Cash:", cash.Balance())
fmt.Println("Revenue:", revenue.Balance())
```

Prints the balances using the `Balance()` getter. Output:

```
Cash: 50000
Revenue: 0
```

### Lines 44–46

```go
if err := cash.Debit(60000); err != nil {
    fmt.Println("debit failed", err)
}
```

Tries to debit `60000` cents from `cash`, which only has `50000`. Since `60000 > 50000`, the method returns an error. The `if` initializer captures it, `err != nil` is true, so it prints the error message:

```
debit failed insufficient balance in account 1000-CASH: have 50000, need 60000
```

This demonstrates graceful handling of an insufficient-balance condition without the program crashing.

### Line 47

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
Cash: 50000
Revenue: 0
debit failed insufficient balance in account 1000-CASH: have 50000, need 60000
```

## Key Takeaways

1. **Encapsulation** — The unexported `balance` field is only changed through methods (encapsulation).
2. **Pointer receiver** — Use `*Account` receiver so methods can modify the original struct.
3. **Methods as behavior** — `Credit`, `Debit`, `Balance` — behavior attached to a type.
4. **Error handling** — `Debit` returns an error checked with the `if err := ...; err != nil` pattern.
5. **`fmt.Errorf`** — Creates formatted error messages.
6. **Double-entry concept** — Credit increases the balance; debit checks the balance before subtracting.
