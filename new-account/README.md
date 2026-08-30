# new-account

Go-তে **zero value** শেখার ছোট example — `var` দিয়ে declare করা variable-গুলো value না দিয়ে কী নিয়ে শুরু হয়, আর তারপর assign করা জানা।

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

### Line 5

```go
func main() {
```

Program-এর entry point।

### Lines 6–8

```go
var accountHolder string
var balanceInCents int64
var isActive bool
```

তিনটা variable **শুধু type দিয়ে** declare করা হয় — কোনো initial value দেওয়া হয়নি। মানে প্রতিটা তার type-এর **zero value** ধরে:

- `accountHolder` (`string`) → `""` (empty string) — print-এ blank দেখায়।
- `balanceInCents` (`int64`) → `0`।
- `isActive` (`bool`) → `false`।

**Zero value:** Go-তে declared-but-not-assigned variable-গুলো কখনোই "undefined" থাকে না — প্রতিটা type-এর নিজস্ব শূন্য-মান বসে। এটা C/C++-এর uninitialized memory-র বিপরীতে নিরাপদ।

### Lines 10–13

```go
fmt.Println("---before setup---")
fmt.Println("Account holder name:", accountHolder)
fmt.Println("Account balance(cents):", balanceInCents)
fmt.Println("Account active?:", isActive)
```

Setup-এর আগে state print করা হয়:

- `Account holder name: ` (খালি — string zero value)
- `Account balance(cents): 0`
- `Account active?: false`

`---before setup---` banner-এর সাথে — account-টা সম্পূর্ণ uninitialized (zero)।

### Lines 15–17

```go
accountHolder = "John Doe"
balanceInCents = 5000000
isActive = true
```

এখন **assign** করা হয় (`=` — variable-গুলো আগেই declare, তাই `:=` নয়):

- `accountHolder = "John Doe"` — নাম set।
- `balanceInCents = 5000000` — balance paisa-তে (5,000,000 paisa)।
- `isActive = true` — account active।

### Lines 19–22

```go
fmt.Println("---after setup---")
fmt.Println("Account holder name:", accountHolder)
fmt.Println("Account balance(cents):", balanceInCents)
fmt.Println("Account active?:", isActive)
```

Setup-এর পরে state print:

- `Account holder name: John Doe`
- `Account balance(cents): 5000000`
- `Account active?: true`

Zero value-এর পরিবর্তে real value।

### Line 23

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
---before setup---
Account holder name: 
Account balance(cents): 0
Account active?: false
---after setup---
Account holder name: John Doe
Account balance(cents): 5000000
Account active?: true
```

## মূল শিক্ষা / Key Takeaways

1. **Zero value** — `var x string/int64/bool` initialize-না করলেও নিজের type-এর শূন্য-মান পায়।
2. **String → `""`** — print-এ blank; int → `0`; bool → `false`।
3. **No "undefined" state** — Go-তে uninitialized variable থেকে পড়া নিরাপদ।
4. **`=` vs `:=`** — আগে declare করা variable-এ `=` দিয়ে value দেওয়া হয়।
5. **Setup pattern** — before/after print দেখায় declare-এ zero, assign-এ real value।

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

Imports the `fmt` package for console output.

### Line 5

```go
func main() {
```

Program entry point.

### Lines 6–8

```go
var accountHolder string
var balanceInCents int64
var isActive bool
```

Three variables declared **with type only** — no initial value. Each takes its type's **zero value**:

- `accountHolder` (`string`) → `""` (empty string) — prints as blank.
- `balanceInCents` (`int64`) → `0`.
- `isActive` (`bool`) → `false`.

**Zero value:** in Go, declared-but-unassigned variables are never "undefined" — each type starts at its own zero value. This is safe, unlike uninitialized memory in C/C++.

### Lines 10–13

```go
fmt.Println("---before setup---")
fmt.Println("Account holder name:", accountHolder)
fmt.Println("Account balance(cents):", balanceInCents)
fmt.Println("Account active?:", isActive)
```

Prints the state before setup:

- `Account holder name: ` (blank — the string zero value)
- `Account balance(cents): 0`
- `Account active?: false`

Together with the `---before setup---` banner — the account is fully uninitialized (zero).

### Lines 15–17

```go
accountHolder = "John Doe"
balanceInCents = 5000000
isActive = true
```

Now we **assign** (`=` — the variables were already declared, so no `:=`):

- `accountHolder = "John Doe"` — sets the name.
- `balanceInCents = 5000000` — balance in paisa (5,000,000 paisa).
- `isActive = true` — account active.

### Lines 19–22

```go
fmt.Println("---after setup---")
fmt.Println("Account holder name:", accountHolder)
fmt.Println("Account balance(cents):", balanceInCents)
fmt.Println("Account active?:", isActive)
```

Prints the state after setup:

- `Account holder name: John Doe`
- `Account balance(cents): 5000000`
- `Account active?: true`

Real values instead of zero values.

### Line 23

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
---before setup---
Account holder name: 
Account balance(cents): 0
Account active?: false
---after setup---
Account holder name: John Doe
Account balance(cents): 5000000
Account active?: true
```

## Key Takeaways

1. **Zero value** — `var x string/int64/bool` gets its type's zero value even without initialization.
2. **String → `""`** — prints blank; int → `0`; bool → `false`.
3. **No "undefined" state** — reading an uninitialized variable in Go is safe.
4. **`=` vs `:=`** — previously declared variables get values with `=`.
5. **Setup pattern** — before/after prints show zero at declare, real value at assign.