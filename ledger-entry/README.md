# ledger-entry

Go-তে **custom type (`type Money float64`)** আর struct slice-এর উপর loop চালিয়ে double-entry ledger-এ debit/credit-এর ধরন শেখার ছোট example।

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
type Money float64
```

একটা **custom (named) type** `Money` define করা হয় যার underlying type হলো `float64`। বরাবর প্রতিটা জায়গায় খালি `float64` ব্যবহার না করে টাকা-সংক্রান্ত data-কে `Money` type দিয়ে represent করা হয়। এতে কোড আরও self-documenting হয় — অর্থাৎ "এটা শুধু কোনো decimal না, এটা টাকার পরিমাণ"।

> **দুর্বলতা:** এখানে `Debit: 5000` ও `Credit: 5000` দুটোই `Money` হিসেবে type-নিরাপদ — দুটো `Money` যোগ/তুলনা করা যায়। তবে `Money` আর `float64`-এর মাঝে implicit conversion হয় না।

### Lines 7–11

```go
type LedgerEntry struct {
    AccountID string
    Debit     Money
    Credit    Money
}
```

`LedgerEntry` নামক একটা struct type declare করে। এটা double-entry ledger-এর **একটা entry/line-item** represent করে:

- `AccountID` — কোন account-এর এটা (string)।
- `Debit` — এই account-এ debit-এর পরিমাণ (`Money`)।
- `Credit` — এই account-এ credit-এর পরিমাণ (`Money`)।

### Line 13

```go
func main() {
```

Program-এর entry point।

### Lines 14–25

```go
entries := []LedgerEntry{
    {
        AccountID: "cach",
        Debit:     5000,
        Credit:    0,
    },
    {
        AccountID: "revenue",
        Debit:     0,
        Credit:    5000,
    },
}
```

`entries` — একটা `[]LedgerEntry` (struct-এর slice)। দুটো entry আছে:

1. `AccountID: "cach"` — `Debit: 5000`, `Credit: 0`। (লক্ষ্য করো, "cash"-এর typo ইচ্ছাকৃত/উদাহরণ-মাত্র।)
2. `AccountID: "revenue"` — `Debit: 0`, `Credit: 5000`।

এটা double-entry accounting-এর সৌন্দর্য দেখায়: একটা transaction (e.g. নগদে আয়) আয়-হয় debit হিসেবে (cash-এ) আর credit হিসেবে (revenue-এ) — দুটো দিকেই `5000` লেখা হয়। => মোট debit = মোট credit।

`Debit`/`Credit`-এর value টা literals `5000`... সাধারণ সংখ্যা। যেহেতু field type `Money` (underlying `float64`), unttyped constant `5000` automatically `Money`-এ convert হয়।

### Line 27

```go
var totalDebit, totalCredit Money
```

দুটো accumulator variable declare করে — `totalDebit` আর `totalCredit`, দুটোই type `Money`। `var` দিয়ে declare করায় দুটোই **zero value** (`0`) দিয়ে শুরু হয়।

### Lines 29–33

```go
for _, entry := range entries {
    totalCredit += entry.Credit
    totalDebit += entry.Debit
    fmt.Println(entry.AccountID, "- debit:", entry.Debit, "credit:", entry.Credit)
}
```

`for ... range` loop `entries` slice-এর প্রতিটা `LedgerEntry`-তে যায়:

- `totalCredit += entry.Credit` — entry-র credit accumulative যোগ করে।
- `totalDebit += entry.Debit` — entry-র debit যোগ করে।
- `fmt.Println` — প্রতিটা entry-র account id, debit, credit print করে।

### Line 35

```go
fmt.Println("Books balanced:", totalDebit == totalCredit)
```

সব entry process করার পর মোট debit আর মোট credit তুলনা করা হয়। এখানে `totalDebit = 5000` আর `totalCredit = 5000` সমান, তাই `Books balanced: true` print হয়। **Double-entry accounting-র core rule:** ledger তখনই balanced, যখন debit-এর যোগফল = credit-এর যোগফল।

### Line 36

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
cach - debit: 5000 credit: 0
revenue - debit: 0 credit: 5000
Books balanced: true
```

## মূল শিক্ষা / Key Takeaways

1. **Custom type** — `type Money float64` দিয়ে ডোমেইন-নির্দিষ্ট type; কোড self-documenting হয়।
2. **Struct slice** — `[]LedgerEntry` দিয়ে রেকর্ডের collection।
3. **`for ... range` + accumulator** — সব entry-র debit/credit যোগ করা।
4. **Double-entry balance check** — মোট debit == মোট credit মানেই ledger balanced।
5. **Zero value** — `var` দিয়ে declare করলে `Money` (`float64`) `0` দিয়ে শুরু হয়।

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
type Money float64
```

Defines a **custom (named) type** `Money` whose underlying type is `float64`. Instead of using a bare `float64` everywhere, money-related data is represented as `Money`. This makes the code more self-documenting — "this isn't just any decimal, it's an amount of money".

> **Caveat:** Here `Debit: 5000` and `Credit: 5000` are both type-safe — two `Money` values can be added/compared. But there is no implicit conversion between `Money` and plain `float64`.

### Lines 7–11

```go
type LedgerEntry struct {
    AccountID string
    Debit     Money
    Credit    Money
}
```

Declares a struct type `LedgerEntry` representing **one entry/line-item** of a double-entry ledger:

- `AccountID` — which account this entry belongs to (string).
- `Debit` — the debit amount on this account (`Money`).
- `Credit` — the credit amount on this account (`Money`).

### Line 13

```go
func main() {
```

Program entry point.

### Lines 14–25

```go
entries := []LedgerEntry{
    {
        AccountID: "cach",
        Debit:     5000,
        Credit:    0,
    },
    {
        AccountID: "revenue",
        Debit:     0,
        Credit:    5000,
    },
}
```

`entries` — a `[]LedgerEntry` (slice of structs). There are two entries:

1. `AccountID: "cach"` — `Debit: 5000`, `Credit: 0`. (The typo for "cash" is intentional/illustrative only.)
2. `AccountID: "revenue"` — `Debit: 0`, `Credit: 5000`.

This shows the beauty of double-entry accounting: one transaction (e.g. income received in cash) is recorded as a debit (to cash) and a credit (to revenue) — `5000` on both sides. Total debit = total credit.

The `Debit`/`Credit` values are the literals `5000`... plain numbers. Because the field type is `Money` (underlying `float64`), the untyped constant `5000` is automatically converted to `Money`.

### Line 27

```go
var totalDebit, totalCredit Money
```

Declares two accumulator variables — `totalDebit` and `totalCredit`, both of type `Money`. Declared with `var`, both start at the **zero value** (`0`).

### Lines 29–33

```go
for _, entry := range entries {
    totalCredit += entry.Credit
    totalDebit += entry.Debit
    fmt.Println(entry.AccountID, "- debit:", entry.Debit, "credit:", entry.Credit)
}
```

The `for ... range` loop goes over each `LedgerEntry` in the `entries` slice:

- `totalCredit += entry.Credit` — accumulate the credits.
- `totalDebit += entry.Debit` — accumulate the debits.
- `fmt.Println` — prints each entry's account id, debit, and credit.

### Line 35

```go
fmt.Println("Books balanced:", totalDebit == totalCredit)
```

After processing all entries, the total debit is compared to the total credit. Here `totalDebit = 5000` and `totalCredit = 5000` are equal, so `Books balanced: true` is printed. **Core rule of double-entry accounting:** the ledger is balanced when the sum of debits equals the sum of credits.

### Line 36

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
cach - debit: 5000 credit: 0
revenue - debit: 0 credit: 5000
Books balanced: true
```

## Key Takeaways

1. **Custom type** — `type Money float64` gives a domain-specific type; code becomes self-documenting.
2. **Struct slice** — a `[]LedgerEntry` collection of records.
3. **`for ... range` + accumulator** — summing debits/credits across all entries.
4. **Double-entry balance check** — total debit == total credit means the ledger is balanced.
5. **Zero value** — declared with `var`, `Money` (`float64`) starts at `0`.