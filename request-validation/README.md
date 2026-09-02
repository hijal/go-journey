# request-validation

Go-তে **function with multiple return/logic**, **`strings` package** (`TrimSpace`, `Contains`) আর **input validation** শেখার ছোট example।

**📖 ভাষা নির্বাচন করুন / Choose language:**

[🇧🇩 বাংলা](#bangla) • [🇬🇧 English](#english)

---

<a name="bangla"></a>

## 🇧🇩 বাংলা সংস্করণ

### Lines 1–6

```go
package main

import (
    "fmt"
    "strings"
)
```

- `package main` — একটা executable program।
- `fmt` — output print করার জন্য।
- `strings` — string-এর সাথে কাজ করা (`TrimSpace`, `Contains`) — ইমেইল validation-এর জন্য।

### Lines 8–27

```go
func registerUser(email, password string) string {
    email = strings.TrimSpace(email)

    if email == "" {
        return "400: email is required"
    }

    if password == "" {
        return "400: password is required"
    }

    if len(password) < 8 {
        return "400: password must be at least 8 characters long"
    }

    if !strings.Contains(email, "@") {
        return "400: email must be valid"
    }
    return "Registration successful"
}
```

`registerUser` function — `email` আর `password` (দুটো string) নেয় এবং একটা string return করে (success message বা error message):

- `email = strings.TrimSpace(email)` — আগে-পেছের ফাঁকা জায়গা (space) সরিয়ে নেয়। যেমন `"  rat@x.com "` → `"rat@x.com"`।
- `if email == ""` — খালি হলে `"400: email is required"`।
- `if password == ""` — খালি password হলে `"400: password is required"`।
- `if len(password) < 8` — ৮ অক্ষরের কম হলে `"400: password must be at least 8 characters long"`। (`len` দিয়ে length)।
- `if !strings.Contains(email, "@")` — ইমেইলে `@` না থাকলে `"400: email must be valid"`। (`!` মানে not)।
- সব ঠিক থাকলে `"Registration successful"`।

**নোট:** এখানে একাধিক `return` — যেকোনো validation fail হলে সাথে-সাথে return। এতে একটা guard-clause (early-return) pattern তৈরি হয়।

### Line 29

```go
func main() {
```

Program-এর entry point।

### Lines 30–34

```go
fmt.Println(registerUser("  ", "secret123"))
fmt.Println(registerUser("rata@example.com", "short"))
fmt.Println(registerUser("rataexample.com", "secret123"))
fmt.Println(registerUser("rata@example.com", "secret123"))
```

চারটা ভিন্ন case registerUser-কে test করে:

1. `("  ", "secret123")` — email «খালি» (spaces) → TrimSpace-এ খালি → `"400: email is required"`।
2. `("rata@example.com", "short")` — password ৮-এর কম → `"400: password must be at least 8 characters long"`।
3. `("rataexample.com", "secret123")` — email-এ `@` নাই → `"400: email must be valid"`।
4. `("rata@example.com", "secret123")` — সব ঠিক → `"Registration successful"`।

---

## Expected Output

```
400: email is required
400: password must be at least 8 characters long
400: email must be valid
Registration successful
```

## মূল শিক্ষা / Key Takeaways

1. **Guard clauses (early return)** — প্রতিটা invalid case-এ সাথে-সাথে return, বাকিটা পড়ে না।
2. **`strings.TrimSpace`** — string-এর আগে-পরে whitespace সরায়।
3. **`strings.Contains`** — substring আছে কিনা check করে।
4. **`len`** — string-এর দৈর্ঘ্য (length)।
5. **`!` (logical NOT)** — শর্ত উল্টায়।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Lines 1–6

```go
package main

import (
    "fmt"
    "strings"
)
```

- `package main` — an executable program.
- `fmt` — for console output.
- `strings` — for string utilities (`TrimSpace`, `Contains`) used in email validation.

### Lines 8–27

```go
func registerUser(email, password string) string {
    email = strings.TrimSpace(email)

    if email == "" {
        return "400: email is required"
    }

    if password == "" {
        return "400: password is required"
    }

    if len(password) < 8 {
        return "400: password must be at least 8 characters long"
    }

    if !strings.Contains(email, "@") {
        return "400: email must be valid"
    }
    return "Registration successful"
}
```

`registerUser` function — takes `email` and `password` (both strings) and returns a string (a success message or an error message):

- `email = strings.TrimSpace(email)` — removes leading/trailing whitespace. E.g. `"  rat@x.com "` → `"rat@x.com"`.
- `if email == ""` — if empty, return `"400: email is required"`.
- `if password == ""` — if empty, return `"400: password is required"`.
- `if len(password) < 8` — if fewer than 8 characters, return `"400: password must be at least 8 characters long"`. (`len` gives the length.)
- `if !strings.Contains(email, "@")` — if there's no `@`, return `"400: email must be valid"`. (`!` means NOT).
- If all good, return `"Registration successful"`.

**Note:** Multiple `return` statements — any failed validation returns immediately. This creates a guard-clause (early-return) pattern.

### Line 29

```go
func main() {
```

Program entry point.

### Lines 30–34

```go
fmt.Println(registerUser("  ", "secret123"))
fmt.Println(registerUser("rata@example.com", "short"))
fmt.Println(registerUser("rataexample.com", "secret123"))
fmt.Println(registerUser("rata@example.com", "secret123"))
```

Tests four different cases to `registerUser`:

1. `("  ", "secret123")` — email is whitespace → TrimSpace makes it empty → `"400: email is required"`.
2. `("rata@example.com", "short")` — password shorter than 8 → `"400: password must be at least 8 characters long"`.
3. `("rataexample.com", "secret123")` — email has no `@` → `"400: email must be valid"`.
4. `("rata@example.com", "secret123")` — all good → `"Registration successful"`.

---

## Expected Output

```
400: email is required
400: password must be at least 8 characters long
400: email must be valid
Registration successful
```

## Key Takeaways

1. **Guard clauses (early return)** — each invalid case returns immediately without reading the rest.
2. **`strings.TrimSpace`** — strips leading/trailing whitespace.
3. **`strings.Contains`** — checks whether a substring exists.
4. **`len`** — gives the length of a string.
5. **`!` (logical NOT)** — inverts a condition.
