# transaction-fraud-screening

Go-তে **struct slice**, **`for range` loop**, **`continue`**, আর nested **`if`** দিয়ে transaction fraud screening (জালিয়াতি যাচাই) শেখার ছোট example।

**📖 ভাষা নির্বাচন করুন / Choose language:**

[🇧🇩 বাংলা](#bangla) • [🇬🇧 English](#english)

---

<a name="bangla"></a>

## 🇧🇩 বাংলা সংস্করণ

### Lines 1–3

```go
package main

import "fmt"
```

- `package main` — একটা executable program।
- `fmt` — output print করার জন্য।

### Lines 5–9

```go
type Transaction struct {
    ID     string
    Amount float64
    Online bool
}
```

একটা struct `Transaction`: `ID` (ট্রানজেকশন ID), `Amount` (পরিমাণ, float64), `Online` (অনলাইন ট্রানজেকশন কিনা, bool)।

### Line 11

```go
func main() {
```

Program-এর entry point।

### Lines 12–17

```go
transactions := []Transaction{
    {ID: "TX-001", Amount: 800, Online: true},
    {ID: "TX-002", Amount: 250000, Online: true},
    {ID: "TX-003", Amount: 1200, Online: false},
    {ID: "TX-004", Amount: 95000, Online: false},
}
```

`[]Transaction` slice — ৪টা transaction (positional field initialization সহ struct literal)।

### Line 19

```go
flagged := 0
```

`flagged` counter — কতটা ট্রানজেকশন attention-এ দরকার, শুরুতে 0।

### Line 20

```go
for _, tx := range transactions {
```

প্রতিটা transaction-এ loop। `_` index discard, `tx` = প্রতিটা Transaction।

### Lines 21–24

```go
if tx.Amount < 50000 {
    fmt.Printf("%s approved\n", tx.ID)
    continue
}
```

- `if tx.Amount < 50000` — 50000-এর কম হলে automatic **approved** print করে **`continue`** (বাকি logic skip, পরের transaction-এ যায়)।
- TX-001 (800) ও TX-003 (1200) approved।

### Line 26

```go
flagged++
```

এখানে পৌঁছালে বুঝি amount ≥ 50000 — তাই `flagged` counter বাড়ায়।

### Lines 28–32

```go
if tx.Amount > 200000 && tx.Online {
    fmt.Printf("%s: BLOCKED(large online transaction)\n", tx.ID)
} else {
    fmt.Printf("%s: flagged for manual review\n", tx.ID)
}
```

নested logic:

- `tx.Amount > 200000 && tx.Online` — 200000-এর বেশি **এবং** অনলাইন হলে → `BLOCKED(large online transaction)`।
- না হলে → `flagged for manual review` (manual review-এ পাঠানো)।
  - TX-002 (250000 online) → BLOCKED।
  - TX-004 (95000 offline) → manual review।

### Lines 35–39

```go
if flagged == 0 {
    fmt.Println("All transactions passed automatic checks")
} else {
    fmt.Printf("%d transaction(s) need attention\n", flagged)
}
```

সবশেষে:

- `flagged == 0` হলে → "All transactions passed automatic checks"।
- না হলে → কতটা attention-এ দরকার দেখায়। এখানে flagged=2 (TX-002, TX-004) ⇒ `2 transaction(s) need attention`।

---

## Expected Output

```
TX-001 approved
TX-002: BLOCKED(large online transaction)
TX-003 approved
TX-004: flagged for manual review
2 transaction(s) need attention
```

## মূল শিক্ষা / Key Takeaways

1. **Struct slice** — `[]Transaction` — অনেকগুলো struct ধরে।
2. **`for range` + positional struct literal** — `{ID: "...", Amount: ..., Online: ...}`।
3. **`continue`** — শর্তে স্কিপ করে পরের iteration।
4. **Nested `&&` condition** — amount + online-এর সমন্বিত check।
5. **Counter variable** — `flagged++` দিয়ে tracking।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Lines 1–3

```go
package main

import "fmt"
```

- `package main` — an executable program.
- `fmt` — for console output.

### Lines 5–9

```go
type Transaction struct {
    ID     string
    Amount float64
    Online bool
}
```

A struct `Transaction`: `ID` (transaction ID), `Amount` (the amount, float64), `Online` (whether it's an online transaction, bool).

### Line 11

```go
func main() {
```

Program entry point.

### Lines 12–17

```go
transactions := []Transaction{
    {ID: "TX-001", Amount: 800, Online: true},
    {ID: "TX-002", Amount: 250000, Online: true},
    {ID: "TX-003", Amount: 1200, Online: false},
    {ID: "TX-004", Amount: 95000, Online: false},
}
```

A `[]Transaction` slice — 4 transactions (using positional struct literals).

### Line 19

```go
flagged := 0
```

The `flagged` counter — how many transactions need attention, starting at 0.

### Line 20

```go
for _, tx := range transactions {
```

Loops over each transaction. `_` discards the index, `tx` = each Transaction.

### Lines 21–24

```go
if tx.Amount < 50000 {
    fmt.Printf("%s approved\n", tx.ID)
    continue
}
```

- `if tx.Amount < 50000` — for amounts under 50000, print automatic **approved** and **`continue`** (skip the rest, move to the next transaction).
- TX-001 (800) and TX-003 (1200) are approved.

### Line 26

```go
flagged++
```

Reaching here means the amount is ≥ 50000 — so increment the `flagged` counter.

### Lines 28–32

```go
if tx.Amount > 200000 && tx.Online {
    fmt.Printf("%s: BLOCKED(large online transaction)\n", tx.ID)
} else {
    fmt.Printf("%s: flagged for manual review\n", tx.ID)
}
```

Nested logic:

- `tx.Amount > 200000 && tx.Online` — greater than 200000 **and** online → `BLOCKED(large online transaction)`.
- Otherwise → `flagged for manual review` (send for manual review).
  - TX-002 (250000 online) → BLOCKED.
  - TX-004 (95000 offline) → manual review.

### Lines 35–39

```go
if flagged == 0 {
    fmt.Println("All transactions passed automatic checks")
} else {
    fmt.Printf("%d transaction(s) need attention\n", flagged)
}
```

At the end:

- If `flagged == 0` → "All transactions passed automatic checks".
- Otherwise, show how many need attention. Here flagged=2 (TX-002, TX-004) ⇒ `2 transaction(s) need attention`.

---

## Expected Output

```
TX-001 approved
TX-002: BLOCKED(large online transaction)
TX-003 approved
TX-004: flagged for manual review
2 transaction(s) need attention
```

## Key Takeaways

1. **Struct slice** — `[]Transaction` — holds many structs.
2. **`for range` + positional struct literal** — `{ID: "...", Amount: ..., Online: ...}`.
3. **`continue`** — skip and move to the next iteration.
4. **Nested `&&` condition** — combined amount + online check.
5. **Counter variable** — tracking with `flagged++`.
