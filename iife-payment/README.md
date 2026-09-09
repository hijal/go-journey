# iife-payment

Go-তে **IIFE (Immediately Invoked Function Expression)** শেখার ছোট example — একটা anonymous function declare + সাথে-সাথে call, result variable-এ।

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

### Lines 6–7

```go
amount := 1250.75
currency := "BDT"
```

Input: payment amount 1250.75, currency "BDT"।

### Lines 9–14

```go
summary := func(amt float64, cur string) string {
	if amt <= 0 {
		return "REJECT: invalid amount"
	}
	return fmt.Sprintf("%s %.2f accepted", cur, amt)
}(amount, currency)
```

**IIFE** — `func(...) ... {...}(args)`:

- একটা **anonymous function** define করা হয়েছে (`func(amt float64, cur string) string {...}`), ঠিক define করার পরেই `(amount, currency)` দিয়ে **immediately invoked**।
- Arguments `amount`, `currency` function-এর parameter-এ যায়।
- Return value সরাসরি `summary`-তে assign।
- লজিক:
  - `amt <= 0` → rejection string।
  - else → `fmt.Sprintf("%s %.2f accepted", cur, amt)` — currency + 2-decimal amount + "accepted"।

**টা এখানে:** এটা একটা isolated validation+build — scope-এ temporary logic-one-pass pipeline-এর মতো; নামের function declare না করে কোথাও logic রিইউজ না করলেও fine।

**Result value:** `summary` = `"BDT 1250.75 accepted"` (amount positive)।

### Line 16

```go
fmt.Println(summary)
```

Output: `BDT 1250.75 accepted`.

### Line 17

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
BDT 1250.75 accepted
```

## মূল শিক্ষা / Key Takeaways

1. **IIFE** — `func() T {...}()` — declare + immediately invoke in one expression।
2. **No polluting scope** — লজিকটা local, নাম্বার function-এ দরকার হয় না।
3. **Immediate result** — return value এক লাইনে assign-off।
4. **`%.2f`** — 2-decimal float formatting।
5. **Validation + build** — একটা pass-এ reject check + message build।

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

### Lines 6–7

```go
amount := 1250.75
currency := "BDT"
```

Input: payment amount 1250.75, currency "BDT".

### Lines 9–14

```go
summary := func(amt float64, cur string) string {
	if amt <= 0 {
		return "REJECT: invalid amount"
	}
	return fmt.Sprintf("%s %.2f accepted", cur, amt)
}(amount, currency)
```

**IIFE** — `func(...) ... {...}(args)`:

- An **anonymous function** is defined (`func(amt float64, cur string) string {...}`), and **immediately invoked** with `(amount, currency)` right after the definition.
- The arguments `amount`, `currency` flow into the function's parameters.
- The return value is assigned directly to `summary`.
- The logic:
  - `amt <= 0` → a rejection string.
  - else → `fmt.Sprintf("%s %.2f accepted", cur, amt)` — currency + 2-decimal amount + "accepted".

**What it's for here:** an isolated one-pass validation + build — no named function needed, no scope pollution.

**Result value:** `summary` = `"BDT 1250.75 accepted"` (amount is positive).

### Line 16

```go
fmt.Println(summary)
```

Output: `BDT 1250.75 accepted`.

### Line 17

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
BDT 1250.75 accepted
```

## Key Takeaways

1. **IIFE** — `func() T {...}()` — declare and immediately invoke in one expression.
2. **No scope pollution** — the logic stays local; no named function needed.
3. **Immediate result** — the return value is captured directly.
4. **`%.2f`** — 2-decimal float formatting.
5. **Validation + build** — reject-check and message building in one pass.