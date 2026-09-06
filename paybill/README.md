# paybill

Go-তে **business rule validation + custom error** শেখার ছোট example -- bill payment function যা amount ও balance যাচাই করে সিদ্ধান্ত দেয়.

**ভাষা নির্বাচন করুন / Choose language:**

[বাংলা](#bangla) * [English](#english)

---

<a name="bangla"></a>

## বাংলা সংস্করণ

### Lines 1-6

```go
package main

import (
    "errors"
    "fmt"
)
```

- `package main` -- executable program.
- `errors` -- custom error তৈরি করতে.
- `fmt` -- `Sprintf` ও print করতে.

### Lines 8-19

```go
func payBill(amount, balance float64) (string, error) {
    if amount <= 0 {
        return "", errors.New("amount must be greater than zero")
    }

    if amount > balance {
        return "", errors.New("insufficiant balance")
    }

    code := fmt.Sprintf("TXN-%.0f-OK", amount)
    return code, nil
}
```

`payBill` দুটো rule check করে:

- `amount <= 0` -- zero বা negative amount invalid -> error return.
- `amount > balance` -- balance-এর বেশি spend করা যাবে না -> error return.
- সব ঠিক থাকলে `TXN-<amount>-OK` format-এ transaction code return করে.

### Lines 21-34

```go
func main() {
    code, err := payBill(500, 2000)
    if err != nil {
        fmt.Println("payment failed:", err)
        return
    }

    fmt.Println("payment successfull:", code)

    _, err = payBill(5000, 2000)
    if err != nil {
        fmt.Println("Payment failed:", err)
        return
    }
}
```

দুটো call:

- `payBill(500, 2000)` -- 500 <= 2000, success -> "TXN-500-OK"
- `payBill(5000, 2000)` -- 5000 > 2000, error -> "insufficiant balance"

---

## Expected Output

```
payment successfull: TXN-500-OK
Payment failed: insufficiant balance
```

## মূল শিক্ষা / Key Takeaways

1. **Guard clauses** -- invalid input early return করা clean code pattern.
2. **`errors.New`** -- static string error তৈরি.
3. **`fmt.Sprintf` + format verb** -- `%.0f` zero-decimal float, transaction code generate করতে.
4. **Error as second return** -- Go-তে standard `(value, error)` signature.

---

---

<a name="english"></a>

## English Version

### Lines 1-6

```go
package main

import (
    "errors"
    "fmt"
)
```

- `package main` -- an executable program.
- `errors` -- to create custom errors.
- `fmt` -- for `Sprintf` and printing.

### Lines 8-19

```go
func payBill(amount, balance float64) (string, error) {
    if amount <= 0 {
        return "", errors.New("amount must be greater than zero")
    }

    if amount > balance {
        return "", errors.New("insufficiant balance")
    }

    code := fmt.Sprintf("TXN-%.0f-OK", amount)
    return code, nil
}
```

`payBill` enforces two business rules:

- `amount <= 0` -- zero or negative amount is invalid, returns an error.
- `amount > balance` -- cannot spend more than the balance, returns an error.
- otherwise returns a transaction code formatted as `TXN-<amount>-OK`.

### Lines 21-34

```go
func main() {
    code, err := payBill(500, 2000)
    if err != nil {
        fmt.Println("payment failed:", err)
        return
    }

    fmt.Println("payment successfull:", code)

    _, err = payBill(5000, 2000)
    if err != nil {
        fmt.Println("Payment failed:", err)
        return
    }
}
```

Two calls:

- `payBill(500, 2000)` -- 500 <= 2000, success -> "TXN-500-OK"
- `payBill(5000, 2000)` -- 5000 > 2000, error -> "insufficiant balance"

---

## Expected Output

```
payment successfull: TXN-500-OK
Payment failed: insufficiant balance
```

## Key Takeaways

1. **Guard clauses** -- early-return on invalid input keeps code clean.
2. **`errors.New`** -- creates a static string error.
3. **`fmt.Sprintf` + format verb** -- `%.0f` formats a float with no decimals, used here for the transaction code.
4. **Error as second return** -- Go's standard `(value, error)` signature.
