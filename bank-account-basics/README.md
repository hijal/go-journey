# bank-account-basics

Go-তে **struct zero value + composite literal + field mutation** শেখার ছোট example — `Account` struct-এর default value, নাম দিয়ে initialize, balance update।

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

`fmt` — `Printf`।

### Lines 5–10

```go
type Account struct {
	Owner   string
	Number  string
	Balance float64
	Active  bool
}
```

`Account` struct — ব্যাংক account-এর model:

- `Owner string` — মালিকের নাম।
- `Number string` — অ্যাকাউন্ট নম্বর (string — leading zeros/phone-স্টাইল)।
- `Balance float64` — টাকার পরিমাণ।
- `Active bool` — সক্রিয় কিনা।

### Lines 12–14

```go
func main() {
	var acc Account
	fmt.Printf("Zero value: %+v\n", acc)
```

**Zero value:**

- `var acc Account` — কোনো value ছাড়া declare → প্রতিটা field-এর type-এর **zero value** পায়:
  - string → `""`
  - float64 → `0`
  - bool → `false`
- `%+v` — field name সহ print: `{Owner: Number: Balance:0 Active:false}`।

**মূল ধারণা:** Go-তে struct-কে initialize না করেই ব্যবহার করা যায় — zero value সবসময় valid।

### Lines 16–21

```go
	acc = Account{
		Owner:   "John Doe",
		Number:  "123456789",
		Balance: 5000.0,
		Active:  true,
	}
```

**Composite literal** — field-name-ভিত্তিক initialize:

- `Account{...}` — struct literal, `Field: value` syntax।
- Field-order-এর উপর নির্ভরতা নেই — নাম দিয়ে set (positional-এ order mismatch-এর risk ছিল, এখানে নেই)।

### Lines 23–24

```go
	acc.Balance -= 1000
	fmt.Printf("After withdrawal: %+v\n", acc)
```

**Field mutation:**

- `acc.Balance -= 1000` — `5000.0 − 1000` → `4000`। ডট নোটেশন দিয়ে struct-এর field modify।

---

## Expected Output

```
Zero value: {Owner: Number: Balance:0 Active:false}
After withdrawal: {Owner:John Doe Number:123456789 Balance:4000 Active:true}
```

## মূল শিক্ষা / Key Takeaways

1. **Zero value** — `var acc Account` → field-type-ভিত্তিক defaults।
2. **`%+v`** — field name সহ print।
3. **Composite literal** — `Account{Owner: ...}` field-name-ভিত্তিক init।
4. **Field access/mutation** — `acc.Balance -= 1000`।
5. **Struct as data model** — সম্পর্কিত data এক unit-এ।

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

`fmt` — for `Printf`.

### Lines 5–10

```go
type Account struct {
	Owner   string
	Number  string
	Balance float64
	Active  bool
}
```

`Account` struct — a bank account model:

- `Owner string` — the owner's name.
- `Number string` — the account number (a string — keeps leading zeros, phone-style).
- `Balance float64` — the amount of money.
- `Active bool` — whether the account is active.

### Lines 12–14

```go
func main() {
	var acc Account
	fmt.Printf("Zero value: %+v\n", acc)
```

**Zero value:**

- `var acc Account` — declared without a value — each field gets its type's **zero value**:
  - string → `""`
  - float64 → `0`
  - bool → `false`
- `%+v` — prints with field names: `{Owner: Number: Balance:0 Active:false}`.

**Core idea:** in Go, a struct can be used without explicit initialization — the zero value is always valid.

### Lines 16–21

```go
	acc = Account{
		Owner:   "John Doe",
		Number:  "123456789",
		Balance: 5000.0,
		Active:  true,
	}
```

**Composite literal** — field-name-based initialization:

- `Account{...}` — struct literal, `Field: value` syntax.
- No dependence on field order — set by name (positional order would risk mismatches).

### Lines 23–24

```go
	acc.Balance -= 1000
	fmt.Printf("After withdrawal: %+v\n", acc)
```

**Field mutation:**

- `acc.Balance -= 1000` — `5000.0 − 1000` → `4000`. Dots notation for modifying struct fields.

---

## Expected Output

```
Zero value: {Owner: Number: Balance:0 Active:false}
After withdrawal: {Owner:John Doe Number:123456789 Balance:4000 Active:true}
```

## Key Takeaways

1. **Zero value** — `var acc Account` → defaults from field types.
2. **`%+v`** — print with field names.
3. **Composite literal** — field-name-based init `Account{Owner: ...}`.
4. **Field access/mutation** — `acc.Balance -= 1000`.
5. **Struct as a data model** — related data in one unit.