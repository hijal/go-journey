# payment-risk-flagging

Go-তে **switch statement** (`switch` without condition), **`fallthrough`**, **`append`** আর `[]string` slice নিয়ে risk flagging শেখার ছোট example।

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

### Lines 5–21

```go
func riskFlags(amount float64) []string {
    var flags []string

    switch {
    case amount >= 100000:
        flags = append(flags, "manual-review")
        fallthrough
    case amount >= 10000:
        flags = append(flags, "extra-verification")
        fallthrough
    case amount >= 1000:
        flags = append(flags, "log-audit-trail")
    default:
        // small amount, no flags needed
    }
    return flags
}
```

একটা function `riskFlags` যা একটা `float64` (`amount`) নেয় এবং `[]string` (string-এর slice) return করে — transaction-এ কী কী risk flag প্রযোজ্য তা ধারণ করে।

- `var flags []string` — একটা খালি string slice বানায়।
- `switch { ... }` — **switch without condition**। এখানে প্রতিটা `case`-এ একটা শর্ত (boolean expression) দেওয়া হয়; প্রথমটাই match হয় যে চলে।
- **`fallthrough`** — এটা Go-র একটি বিশেষ keyword। সাধারণত Go-তে switch-এ একটি case match হলে বাকিগুলো আর চেক হয় না। কিন্তু `fallthrough` বললে **পরের case-টাও match-ish হিসেবে execute হয়** (তার শর্ত যাচাই না করেই)।

**এই function-টা কীভাবে কাজ করে (amount = 150000):**

1. `case amount >= 100000` → true (150000 ≥ 100000)। `"manual-review"` add হয়। তারপর `fallthrough` → পরের case-এ যায়।
2. `case amount >= 10000` → `fallthrough`-এর কারণে শর্ত না দেখেই execute হয়, `"extra-verification"` add হয়, আবার `fallthrough`।
3. `case amount >= 1000` → execute হয়, `"log-audit-trail"` add হয়। এখানে `fallthrough` নেই, তাই এখানেই শেষ।
4. `default` চালে না।

ফলে দাঁড়ায়: `["manual-review", "extra-verification", "log-audit-trail"]`।

> **সতর্কতা:** `fallthrough` ব্যবহার করলে মনে রাখতে হবে — পরের case-এর শর্ত আর যাচাই হয় না। তাই amount যত বড়, তত বেশি case-এর block চলে (যেহেতু 10万-এ থামে না)। এখানে ইচ্ছাকৃতভাবে cascade behavior দেখানো হয়েছে। সাধারণত চাইলে এই লজিকটা `if...else if` দিয়েও লেখা যায়, কিন্তু `fallthrough` দিয়ে slice-এ progressively flag accumulate করা হয়েছে।

- ছোট amount-এর জন্য (যেমন 500) `default` চলে — কোনো comment ছাড়া কিছুই হয় না (ছোট পরিমাণে risk flag লাগে না)।

### Line 23

```go
func main() {
```

Program-এর entry point।

### Line 24

```go
amount := 150000.0
```

`150000.0` একটা `float64`, তাই `:=` দিয়ে `amount`-এর type inference হয় `float64`।

### Lines 25–28

```go
if amount <= 0 {
    fmt.Println("Invalid amount")
    return
}
```

**validity check** — amount-টা 0 বা negative হলে `"Invalid amount"` print করে `return` দিয়ে function শেষ। (এখানে amount 150000, তাই এটা চলে না।)

### Line 30

```go
fmt.Println("Amount:", amount)
```

`Amount: 150000` print করে।

### Line 31

```go
fmt.Println("Flags:", riskFlags(amount))
```

`riskFlags(amount)` call করে return করা slice-টাকে print করে। Output: `Flags: [manual-review extra-verification log-audit-trail]`।

---

## Expected Output

```
Amount: 150000
Flags: [manual-review extra-verification log-audit-trail]
```

## মূল শিক্ষা / Key Takeaways

1. **Switch without condition** — `switch { }`-এ প্রতিটা `case`-এ boolean শর্ত; প্রথম match-টা চলে।
2. **`fallthrough`** — শর্ত না দেখে পরের case-ও execute করায় — cascade/accumulate logic-এর জন্য।
3. **`append`** — slice-এর শেষে নতুন element যোগ করে।
4. **`[]string`** — string-এর slice, যা dynamic বাড়ে।
5. **`default`** — কোনো case match না হলে চলে।
6. **Validity guard** — function শুরুতে অবৈধ input check করে early return।

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

### Lines 5–21

```go
func riskFlags(amount float64) []string {
    var flags []string

    switch {
    case amount >= 100000:
        flags = append(flags, "manual-review")
        fallthrough
    case amount >= 10000:
        flags = append(flags, "extra-verification")
        fallthrough
    case amount >= 1000:
        flags = append(flags, "log-audit-trail")
    default:
        // small amount, no flags needed
    }
    return flags
}
```

A function `riskFlags` that takes a `float64` (`amount`) and returns a `[]string` (slice of strings) holding which risk flags apply to the transaction.

- `var flags []string` — creates an empty string slice.
- `switch { ... }` — a **switch without a condition**. Each `case` here holds a condition (boolean expression); the first one that matches runs.
- **`fallthrough`** — a special Go keyword. Normally in Go, once one case matches, the others are skipped. But `fallthrough` forces the **next case to execute as well** (without even checking its condition).

**How the function works (amount = 150000):**

1. `case amount >= 100000` → true (150000 ≥ 100000). Adds `"manual-review"`. Then `fallthrough` → goes to the next case.
2. `case amount >= 10000` → executes because of `fallthrough` (condition not re-checked), adds `"extra-verification"`, then `fallthrough` again.
3. `case amount >= 1000` → executes, adds `"log-audit-trail"`. No `fallthrough` here, so it stops.
4. `default` does not run.

Result: `["manual-review", "extra-verification", "log-audit-trail"]`.

> **Caveat:** With `fallthrough`, remember the next case's condition is not re-evaluated. So the larger the amount, the more case blocks run (it doesn't stop at 100000). This intentionally shows a cascade behavior. Usually you'd write such logic with `if...else if`, but here flags are progressively accumulated into the slice using `fallthrough`.

- For a small amount (e.g. 500) the `default` runs — nothing happens (small amounts need no risk flags).

### Line 23

```go
func main() {
```

Program entry point.

### Line 24

```go
amount := 150000.0
```

`150000.0` is a `float64`, so `:=` infers `amount`'s type as `float64`.

### Lines 25–28

```go
if amount <= 0 {
    fmt.Println("Invalid amount")
    return
}
```

A **validity check** — if the amount is 0 or negative, print `"Invalid amount"` and `return` to end the function. (Here the amount is 150000, so this doesn't run.)

### Line 30

```go
fmt.Println("Amount:", amount)
```

Prints `Amount: 150000`.

### Line 31

```go
fmt.Println("Flags:", riskFlags(amount))
```

Calls `riskFlags(amount)` and prints the returned slice. Output: `Flags: [manual-review extra-verification log-audit-trail]`.

---

## Expected Output

```
Amount: 150000
Flags: [manual-review extra-verification log-audit-trail]
```

## Key Takeaways

1. **Switch without condition** — in `switch { }`, each `case` holds a boolean condition; the first match runs.
2. **`fallthrough`** — executes the next case too without checking its condition — for cascade/accumulate logic.
3. **`append`** — adds a new element to the end of a slice.
4. **`[]string`** — a slice of strings that grows dynamically.
5. **`default`** — runs when no case matches.
6. **Validity guard** — an early-return check for invalid input at the start of a function.
