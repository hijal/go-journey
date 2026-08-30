# devops-tooling

Go-তে **multiple assignment (swap/failover)**, **`strconv.Atoi`** আর **error handling** শেখার ছোট example — দুটো value-কে এক লাইনে অদলবদল করা এবং string-কে number-এ convert করা।

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
	"fmt"
	"strconv"
)
```

দুটো package:

- `fmt` — output print।
- `strconv` — string-কে number (`Atoi`)।

### Line 8

```go
func main() {
```

Program-এর entry point।

### Line 9

```go
primary, backup := "server-a.internal", "server-b.internal"
```

**Multiple declaration** — এক লাইনে দুটো variable `:=` দিয়ে declare:

- `primary` = `"server-a.internal"`
- `backup` = `"server-b.internal"`

Failover scenario-র setup — primary অথবা backup server।

### Line 10

```go
fmt.Println("Before failover — primary:", primary, "backup:", backup)
```

Failover-এর আগের state print: `Before failover — primary: server-a.internal backup: server-b.internal`।

### Line 12

```go
primary, backup = backup, primary
```

**Multiple assignment** — এক লাইনে দুটো value অদলবদল (swap)। Go-তে temporary variable লাগে না; ডান পাশের value-গুলো আগে evaluate হয়, তারপর assign। Failover observable : primary এখন আগের backup।

### Line 13

```go
fmt.Println("After failover  — primary:", primary, "backup:", backup)
```

Failover-এর পর: `After failover  — primary: server-b.internal backup: server-a.internal`।

### Line 15

```go
portStr := "8080"
```

`portStr` — একটা numeric string (`"8080"`), parse করার জন্য।

### Lines 16–17

```go
port, err := strconv.Atoi(portStr)
```

`strconv.Atoi` — string-কে `int`-এ convert করে, **দুটো value** return করে:

- `port` — convert-করা number (8080)।
- `err` — সফল হলে `nil`; fail হলে error।

### Lines 18–21

```go
if err != nil {
	fmt.Println("invalid port:", err)
	return
}
```

Error check: `err != nil` হলে message print করে `return` দিয়ে function শেষ হয়। (এখানে parse সফল, তাই এটা চলে না।) ফেইল-ফাস্ট pattern — port parse fail হলে বেশি এগোনো হয় না।

### Line 23

```go
fmt.Println("Parsed port:", port)
```

সফল parse-এর result: `Parsed port: 8080`।

### Lines 25–27

```go
invalidPort := "not-a-number"
_, err = strconv.Atoi(invalidPort)
fmt.Println("Expected parse error:", err)
```

এখন ইচ্ছাকৃত **invalid** string parse করা হয়:

- `invalidPort := "not-a-number"` — number নয়।
- `_, err = strconv.Atoi(invalidPort)` — `_` দিয়ে first return (number) discard করা হয়, শুধু `err` নেওয়া হয়। `=` (assignment) — `err` আগে declare হয়েছে।
- Print: `Expected parse error: strconv.Atoi: parsing "not-a-number": invalid syntax` — error-টা পরে semantic.

### Line 28

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
Before failover — primary: server-a.internal backup: server-b.internal
After failover  — primary: server-b.internal backup: server-a.internal
Parsed port: 8080
Expected parse error: strconv.Atoi: parsing "not-a-number": invalid syntax
```

## মূল শিক্ষা / Key Takeaways

1. **Multiple assignment** — `a, b = b, a` দিয়ে value swap — temp variable লাগে না।
2. **`strconv.Atoi`** — string→int, সাথে error-দুটো value return।
3. **Error-check pattern** — `if err != nil { ...; return }` fail-fast।
4. **`_` discard** — return value-গুলোর যেটা দরকারি নয় discard করা।
5. **`:=` vs `=`** — নতুন declare-তে `:=`, existing variable-এ `=`।

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
	"fmt"
	"strconv"
)
```

Two packages:

- `fmt` — for output.
- `strconv` — for string→number conversion (`Atoi`).

### Line 8

```go
func main() {
```

Program entry point.

### Line 9

```go
primary, backup := "server-a.internal", "server-b.internal"
```

**Multiple declaration** — two variables declared in one line with `:=`:

- `primary` = `"server-a.internal"`
- `backup` = `"server-b.internal"`

Setup for a failover scenario — primary and backup servers.

### Line 10

```go
fmt.Println("Before failover — primary:", primary, "backup:", backup)
```

Prints the state before failover: `Before failover — primary: server-a.internal backup: server-b.internal`.

### Line 12

```go
primary, backup = backup, primary
```

**Multiple assignment** — swaps two values in one line. Go needs no temporary variable; the right-hand values are evaluated first, then assigned. The observable failover: primary is now the old backup.

### Line 13

```go
fmt.Println("After failover  — primary:", primary, "backup:", backup)
```

After failover: `After failover  — primary: server-b.internal backup: server-a.internal`.

### Line 15

```go
portStr := "8080"
```

`portStr` — a numeric string (`"8080"`), ready to parse.

### Lines 16–17

```go
port, err := strconv.Atoi(portStr)
```

`strconv.Atoi` — converts a string to `int`, returning **two values**:

- `port` — the converted number (8080).
- `err` — `nil` on success, the error on failure.

### Lines 18–21

```go
if err != nil {
	fmt.Println("invalid port:", err)
	return
}
```

The error check: if `err != nil`, print the message and `return` to end the function. (Parsing succeeded here, so this doesn't run.) It's a fail-fast pattern — you don't proceed past a failed port parse.

### Line 23

```go
fmt.Println("Parsed port:", port)
```

The successful parse's result: `Parsed port: 8080`.

### Lines 25–27

```go
invalidPort := "not-a-number"
_, err = strconv.Atoi(invalidPort)
fmt.Println("Expected parse error:", err)
```

Now an intentionally **invalid** string is parsed:

- `invalidPort := "not-a-number"` — not a number.
- `_, err = strconv.Atoi(invalidPort)` — the first return (the number) is discarded with `_`, keeping only `err`. Uses `=` (assignment) — `err` was declared earlier.
- Prints: `Expected parse error: strconv.Atoi: parsing "not-a-number": invalid syntax` — the error after the fact.

### Line 28

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
Before failover — primary: server-a.internal backup: server-b.internal
After failover  — primary: server-b.internal backup: server-a.internal
Parsed port: 8080
Expected parse error: strconv.Atoi: parsing "not-a-number": invalid syntax
```

## Key Takeaways

1. **Multiple assignment** — `a, b = b, a` swaps values — no temp variable needed.
2. **`strconv.Atoi`** — string→int, returning the value along with an error.
3. **Error-check pattern** — `if err != nil { ...; return }` is fail-fast.
4. **`_` discard** — drop return values you don't need.
5. **`:=` vs `=`** — `:=` declares new, `=` assigns existing.