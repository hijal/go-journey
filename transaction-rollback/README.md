# transaction-rollback

Go-তে **`defer`** ব্যবহার করে transaction-pattern-এ **fail-safe rollback** শেখার ছোট example — `defer` কীভাবে function-এর প্রতিটি exit path-এ cleanup নিশ্চিত করে।

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

দুটো package import করা হয়:

- `errors` — `errors.New` দিয়ে plain error তৈরি করার জন্য।
- `fmt` — console-এ output print করার জন্য।

### Lines 8–10

```go
type Transaction struct {
    committed bool
}
```

`Transaction` নামক struct type declare করে, যার একটা **unexported** field `committed` (bool)। এটা track করে transaction commit হয়েছে কিনা। (Bare minimum representation — real app-এ এখানে database/state থাকত।)

### Lines 12–15

```go
func (t *Transaction) commit() {
    t.committed = true
    fmt.Println("transaction committed")
}
```

`commit` method — **pointer receiver** `(t *Transaction)`। `committed = true` সেট করে এবং "transaction committed" print করে। Pointer receiver দরকারি কারণ আমরা মূল struct-টাই modify করতে চাই।

### Lines 17–21

```go
func (t *Transaction) rollback() {
    if !t.committed {
        fmt.Println("transaction rolled back")
    }
}
```

`rollback` method — যদি transaction **committed না হয়** (`!t.committed`), তাহলে "transaction rolled back" print করে। Committed হয়ে গেলে কিছুই করা হয় না — অর্থাৎ commit সফল হলে rollback "কিছু না করাই right"।

### Lines 23–35

```go
func transferFunds(amount float64) error {
    tx := &Transaction{}

    defer tx.rollback()

    if amount <= 0 {
        return errors.New("invalid transfer amount")
    }

    fmt.Println("transferring", amount)
    tx.commit()
    return nil
}
```

`transferFunds` — মূল function:

- `tx := &Transaction{}` — নতুন transaction (pointer) বানায়; `committed=false` দিয়ে শুরু।
- **`defer tx.rollback()`** — সবচেয়ে গুরুত্বপূর্ণ লাইন। `defer` এটা schedule করে যাতে **function-টা return হওয়ার ঠিক আগে** (যদিও যেখানেই `return` হোক — success অথবা error —) `tx.rollback()` চালানো হয়।
  - Success path: `tx.commit()` আগে `committed=true` সেট করে, `return nil`। তারপর deferred rollback চলে — কিন্তু `committed` true, তাই কিছুই print হয় না।
  - Error path: `amount <= 0` হলে error return করে — `committed` এখনও false। Deferred rollback চলে → "transaction rolled back" print হয়।

**এটাই `defer`-এর শক্তি:** rollback cleanup-টা function-এর শুরুতে লেখা হয়, কিন্তু execution-টা function-এর শেষে নিশ্চিত হয় — প্রতিটা exit path-এ (সফল হোক বা ব্যর্থ) একবার। `defer` মনে রাখে আর একবারই চালায়, comment-কে জানতেও হয় না কার কোন return-এ।

### Line 37

```go
func main() {
```

Program-এর entry point।

### Lines 39–41

```go
if err := transferFunds(100); err != nil {
    fmt.Println("error", err)
}
```

`transferFunds(100)` call করে — বৈধ amount। ভেতরে "transferring 100" print হয়, `commit()` চলে ("transaction committed" print), `return nil` — কোনো error নেই, তাই main-এ কিছুই print হয় না।

### Lines 43–46

```go
fmt.Println("--------------")
if err := transferFunds(-10); err != nil {
    fmt.Println("error", err)
}
```

`transferFunds(-10)` call করে — **invalid** amount।

- ভেতরে `amount <= 0` → `errors.New(...)` দিয়ে error return করে।
- **Return-এর আগে** deferred `tx.rollback()` চলে → "transaction rolled back" print হয় (committed false)।
- তারপর main-এ error-টা পৌঁছে, `if err != nil` true হয় → "error invalid transfer amount" print হয়।

### Line 47

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
transferring 100
transaction committed
--------------
transaction rolled back
error invalid transfer amount
```

## মূল শিক্ষা / Key Takeaways

1. **`defer`** — function শেষে (যেকোনো return path-এ) cleanup নিশ্চিত করার mechanism।
2. **Rollback guard** — committed না হলে rollback হয়; committed হলে কিছুই না।
3. **Pointer receiver** — method-কে মূল struct-টাই modify করতে হয়।
4. **Fail-safe** — error-এও `defer`-এর cleanup চলে — resource leak/partial-state বাঁচায়।
5. **Error handling** — `if err := ...; err != nil` দিয়ে error check।

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

Two packages are imported:

- `errors` — for creating plain errors with `errors.New`.
- `fmt` — for printing output to the console.

### Lines 8–10

```go
type Transaction struct {
    committed bool
}
```

Declares a struct type `Transaction` with a single **unexported** field `committed` (bool). It tracks whether the transaction has been committed. (A bare-minimum representation — a real app would hold database/state here.)

### Lines 12–15

```go
func (t *Transaction) commit() {
    t.committed = true
    fmt.Println("transaction committed")
}
```

The `commit` method — **pointer receiver** `(t *Transaction)`. It sets `committed = true` and prints "transaction committed". A pointer receiver is required because we want to mutate the original struct.

### Lines 17–21

```go
func (t *Transaction) rollback() {
    if !t.committed {
        fmt.Println("transaction rolled back")
    }
}
```

The `rollback` method — if the transaction is **not committed** (`!t.committed`), it prints "transaction rolled back". If already committed, it does nothing — i.e. once a commit succeeded, rolling back means "doing nothing is correct".

### Lines 23–35

```go
func transferFunds(amount float64) error {
    tx := &Transaction{}

    defer tx.rollback()

    if amount <= 0 {
        return errors.New("invalid transfer amount")
    }

    fmt.Println("transferring", amount)
    tx.commit()
    return nil
}
```

`transferFunds` — the core function:

- `tx := &Transaction{}` — creates a new transaction (pointer), starting with `committed=false`.
- **`defer tx.rollback()`** — the most important line. `defer` schedules `tx.rollback()` to run **right before the function returns** — from whichever `return` (success or error):
  - Success path: `tx.commit()` first sets `committed=true`, then `return nil`. The deferred rollback then runs — but `committed` is true, so nothing prints.
  - Error path: `amount <= 0` causes an early `return` of the error — `committed` is still false. The deferred rollback runs → prints "transaction rolled back".

**This is the power of `defer`:** the rollback cleanup is written at the top of the function, but guaranteed to execute at the end — on every exit path (success or failure). `defer` remembers it and runs it exactly once; nobody has to remember to call rollback before each `return`.

### Line 37

```go
func main() {
```

Program entry point.

### Lines 39–41

```go
if err := transferFunds(100); err != nil {
    fmt.Println("error", err)
}
```

Calls `transferFunds(100)` — a valid amount. Inside, "transferring 100" prints, `commit()` runs (prints "transaction committed"), `return nil` — no error, so nothing prints back in `main`.

### Lines 43–46

```go
fmt.Println("--------------")
if err := transferFunds(-10); err != nil {
    fmt.Println("error", err)
}
```

Calls `transferFunds(-10)` — an **invalid** amount.

- Inside, `amount <= 0` → returns `errors.New(...)`.
- **Before returning**, the deferred `tx.rollback()` runs → prints "transaction rolled back" (committed is false).
- Then the error reaches `main`, `if err != nil` is true → prints "error invalid transfer amount".

### Line 47

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
transferring 100
transaction committed
--------------
transaction rolled back
error invalid transfer amount
```

## Key Takeaways

1. **`defer`** — guarantees cleanup runs at the end of a function, from any return path.
2. **Rollback guard** — rollback happens if not committed; does nothing if committed.
3. **Pointer receiver** — methods must be able to mutate the original struct.
4. **Fail-safe** — deferred cleanup also runs on error — prevents resource leaks / partial state.
5. **Error handling** — checked with the `if err := ...; err != nil` pattern.