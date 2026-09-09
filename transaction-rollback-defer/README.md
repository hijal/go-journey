# transaction-rollback-defer

Go-তে **`defer`-based rollback pattern** শেখার ছোট example — transfer-এ debit করার পর error হলে `defer` closure-টা balance restore করে।

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

- `errors` — sentinel error।
- `fmt` — printing/formatting।

### Lines 8–9

```go
var errReviewRequired = errors.New("credit needs manual review")
var balances = map[string]int{"acc-1": 5000, "acc-2": 1000}
```

- `errReviewRequired` — sentinel error।
- `balances` — account-to-balance map (package-level state)।

### Lines 11–36

```go
func transfer(from, to string, amount int) error {
	if amount <= 0 {
		return fmt.Errorf("transfer: invalid amount %d", amount)
	}
	if balances[from] < amount {
		return fmt.Errorf("transfer: insufficient funds in %s: have %d", from, balances[from])
	}

	balances[from] -= amount

	failed := false

	defer func() {
		if failed {
			balances[from] += amount
			fmt.Printf("role back %d from %s\n", amount, from)
		}
	}()

	if amount > 2000 {
		failed = true
		return fmt.Errorf("transfer: %w", errReviewRequired)
	}
	balances[to] += amount
	return nil
}
```

`transfer` — debit-then-credit transfer:

- **Guard clauses**:
  - `amount <= 0` → invalid amount error।
  - `balances[from] < amount` → insufficient funds error।
- **Debit:** `balances[from] -= amount` — sender-থেকে টাকা কাটা (state mutated)।
- **Rollback flag:** `failed := false` — পরে set হবে।
- **`defer` rollback closure:**
  - `defer func() {...}()` — function শেষের সময় (return-এর পর) চালায়, `failed`-টার value read-করা হয় তখন (closure — deferred-এ যাওয়ার সময় variable-এর source check এখানে)।
  - `if failed` → `balances[from] += amount` (প্রাথমিক debit restore) + message।
- **Business rule:** `amount > 2000` → transaction-টা manual review লাগে, **failed** — `failed = true`, error return করে credit-এ পৌঁছায় না।
- success path: `balances[to] += amount` + `nil`।

**Pattern:** debit সবসময় হয়; শুধুমাত্র credit fail-এই rollback (partial updates undo)। Debit-এ framework-লাইক guarantee: fail-এ restore।

**Note** `role back` in the message is a typo in the example source (it prints as-is).

### Lines 38–47

```go
func main() {
	if err := transfer("acc-1", "acc-2", 1500); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("after ok transfer:", balances)

	if err := transfer("acc-1", "acc-2", 2500); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("after failed transfer:", balances)
}
```

**Transfer 1 — `transfer("acc-1", "acc-2", 1500)`:**

- 1500 ≤ 2000 → success path। acc-1: 5000−1500 = 3500; acc-2: 1000+1500 = 2500।
- `failed` false → deferred rollback চলে না।
- Output: `after ok transfer: map[acc-1:3500 acc-2:2500]`

**Transfer 2 — `transfer("acc-1", "acc-2", 2500)`:**

- 2500 > 2000 → `failed = true`, error return।
- Deferred closure: `failed` true → `balances["acc-1"] += 2500` — debit restore → acc-1 আবার 3500। acc-2-তে কিছু যায় না (2500)।
- Output: `role back 2500 from acc-1`, `error: transfer: credit needs manual review`, তারপর `after failed transfer: map[acc-1:3500 acc-2:2500]` (**state restore হয়েছে** debit-এর আগে মতো)।

---

## Expected Output

```
after ok transfer: map[acc-1:3500 acc-2:2500]
role back 2500 from acc-1
error: transfer: credit needs manual review
after failed transfer: map[acc-1:3500 acc-2:2500]
```

## মূল শিক্ষা / Key Takeaways

1. **`defer` + rollback flag** — partial update fail-এ undo করা।
2. **Debit-first, credit-later** — transfer-এর সিকোয়েন্স।
3. **Sentinel + `%w`** — `fmt.Errorf("...: %w", errReviewRequired)`।
4. **Guard clauses** — amount/balance check আগে।
5. **State consistency** — rollback পর balances আগের মতো।

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

- `errors` — for the sentinel error.
- `fmt` — for printing/formatting.

### Lines 8–9

```go
var errReviewRequired = errors.New("credit needs manual review")
var balances = map[string]int{"acc-1": 5000, "acc-2": 1000}
```

- `errReviewRequired` — a sentinel error.
- `balances` — the account-to-balance map (package-level state).

### Lines 11–36

```go
func transfer(from, to string, amount int) error {
	if amount <= 0 {
		return fmt.Errorf("transfer: invalid amount %d", amount)
	}
	if balances[from] < amount {
		return fmt.Errorf("transfer: insufficient funds in %s: have %d", from, balances[from])
	}

	balances[from] -= amount

	failed := false

	defer func() {
		if failed {
			balances[from] += amount
			fmt.Printf("role back %d from %s\n", amount, from)
		}
	}()

	if amount > 2000 {
		failed = true
		return fmt.Errorf("transfer: %w", errReviewRequired)
	}
	balances[to] += amount
	return nil
}
```

`transfer` — a debit-then-credit transfer:

- **Guard clauses**:
  - `amount <= 0` → an invalid amount error.
  - `balances[from] < amount` → an insufficient funds error.
- **Debit:** `balances[from] -= amount` — deducts from the sender (state mutated).
- **Rollback flag:** `failed := false` — set later if needed.
- **`defer` rollback closure:**
  - `defer func() {...}()` — runs when the function ends (after return); reads `failed` at that time (a closure). 
  - `if failed` → `balances[from] += amount` (restores the original debit) + a message.
- **Business rule:** `amount > 2000` → the transaction needs manual review — `failed = true`, returns an error so the credit never happens.
- Success path: `balances[to] += amount` + `nil`.

**Pattern:** the debit always happens; only the credit can fail, and the rollback undoes partial updates. A framework-like guarantee: restore on failure.

**Note:** `role back` in the message is a typo in the example source (it's printed as-is).

### Lines 38–47

```go
func main() {
	if err := transfer("acc-1", "acc-2", 1500); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("after ok transfer:", balances)

	if err := transfer("acc-1", "acc-2", 2500); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("after failed transfer:", balances)
}
```

**Transfer 1 — `transfer("acc-1", "acc-2", 1500)`:**

- 1500 ≤ 2000 → success path. acc-1: 5000−1500 = 3500; acc-2: 1000+1500 = 2500.
- `failed` false → the deferred rollback doesn't run.
- Output: `after ok transfer: map[acc-1:3500 acc-2:2500]`

**Transfer 2 — `transfer("acc-1", "acc-2", 2500)`:**

- 2500 > 2000 → `failed = true`, error returned.
- The deferred closure: `failed` is true → `balances["acc-1"] += 2500` — the debit is restored → acc-1 back to 3500. Nothing goes to acc-2.
- Output: `role back 2500 from acc-1`, `error: transfer: credit needs manual review`, then `after failed transfer: map[acc-1:3500 acc-2:2500]` (**state restored** to what it was before the debit).

---

## Expected Output

```
after ok transfer: map[acc-1:3500 acc-2:2500]
role back 2500 from acc-1
error: transfer: credit needs manual review
after failed transfer: map[acc-1:3500 acc-2:2500]
```

## Key Takeaways

1. **`defer` + rollback flag** — undoing a partial update on failure.
2. **Debit-first, credit-later** — the transfer sequence.
3. **Sentinel + `%w`** — `fmt.Errorf("...: %w", errReviewRequired)`.
4. **Guard clauses** — amount/balance checks up front.
5. **State consistency** — balances are back to the original after rollback.