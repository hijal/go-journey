# package-local-variable

Go-তে **package-level (global) variable** আর **local variable**-এর পার্থক্য বোঝার ছোট example — scope, lifetime, আর function-এর মাঝে state share করা।

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

### Lines 5–6

```go
// Package-level variable
var totalProcessedCents int64
```

`totalProcessedCents` — একটা **package-level (global) variable**:

- Function-গুলোর **বাইরে**, package-এ declare করা হয়।
- Scope: **সব function** থেকে access করা যায় (যেমন `processTransaction` আর `main`)।
- Lifetime: program-জুড়ে থাকে — function শেষ হলেও value মুছে যায় না।
- Zero value: `int64` → `0` দিয়ে শুরু।
- এটাই শেয়ার করা "ledger" — function call-গুলোর মাঝে running total ধরে রাখে।

### Lines 8–15

```go
func processTransaction(amountCents int64) {
	// local variable
	fee := amountCents / 100
	totalProcessedCents += amountCents - fee
	fmt.Printf("processed %d cents (fee %d), running total now %d\n",
		amountCents, fee, totalProcessedCents)

}
```

`processTransaction` — একটা transaction process করে:

- `fee := amountCents / 100` — একটা **local variable**। Integer-এ `fee = amount`-এর 1% (100 দিয়ে ভাগ)। Scope শুধু **এই function-টার ভেতরে**; function-এর বাইরে এটার অস্তিত্ব নাই।
- `totalProcessedCents += amountCents - fee` — package-level variable-টা **update** করা হয়: আগের total-এর সাথে net (amount minus fee) যোগ হয়। এটা function থেকে package-level variable mutate করার example।
- `fmt.Printf(...)` — `Println`-এর বদলে `Printf` — `%d` placeholder দিয়ে formatted।

### Line 16

```go
func main() {
```

Program-এর entry point।

### Lines 17–19

```go
processTransaction(10000)
processTransaction(25000)
processTransaction(5000)
```

তিনটা transaction:

- `10000` → fee 100, net 9900 → total 9900
- `25000` → fee 250, net 24750 → total 34650
- `5000` → fee 50, net 4950 → total 39600

লক্ষ্য করো: প্রতিটা call-এ `totalProcessedCents` তার আগের value-টা **মনে রেখে** যোগ করছে — package-level variable-ই কারণ। প্রতিটা call-এর শেষে local `fee`, `amountCents` আর `totalProcessedCents` survive না — local সব function শেষে মরে যায়, package-level থাকে।

### Line 21

```go
fmt.Println("Final ledger total:", totalProcessedCents)
```

সব transaction-এর পর final state read করা হয়: `Final ledger total: 39600`। `main`-এ package-level variable-টা এখনও accessibility আছে।

### Line 22

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
processed 10000 cents (fee 100), running total now 9900
processed 25000 cents (fee 250), running total now 34650
processed 5000 cents (fee 50), running total now 39600
Final ledger total: 39600
```

## মূল শিক্ষা / Key Takeaways

1. **Package-level variable** — function-এর বাইরে declare; সব function-এর scope; program lifetime।
2. **Local variable** — function-এর ভেতরে; শুধু সেই function-এ দৃশ্য; function শেষে যায়।
3. **Shared state** — package-level variable দিয়ে function call-গুলোর মাঝে data রাখা হয়।
4. **`Printf` + `%d`** — formatted output।
5. **Zero value** — `var totalProcessedCents int64` `0` দিয়ে শুরু।

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

### Lines 5–6

```go
// Package-level variable
var totalProcessedCents int64
```

`totalProcessedCents` — a **package-level (global) variable**:

- Declared **outside** any function, at package scope.
- Scope: reachable from **every function** (e.g. `processTransaction` and `main`).
- Lifetime: lasts for the whole program — its value survives a function ending.
- Zero value: starts at `0` for `int64`.
- This shared "ledger" keeps a running total across function calls.

### Lines 8–15

```go
func processTransaction(amountCents int64) {
	// local variable
	fee := amountCents / 100
	totalProcessedCents += amountCents - fee
	fmt.Printf("processed %d cents (fee %d), running total now %d\n",
		amountCents, fee, totalProcessedCents)

}
```

`processTransaction` — processes one transaction:

- `fee := amountCents / 100` — a **local variable**. In integers, `fee` is 1% of `amount` (divided by 100). Its scope is **only inside this function**; outside it doesn't exist.
- `totalProcessedCents += amountCents - fee` — **updates** the package-level variable: adds the net (amount minus fee) to the previous total. An example of mutating a package-level variable from inside a function.
- `fmt.Printf(...)` — `Printf` instead of `Println` — formatted output with `%d` placeholders.

### Line 16

```go
func main() {
```

Program entry point.

### Lines 17–19

```go
processTransaction(10000)
processTransaction(25000)
processTransaction(5000)
```

Three transactions:

- `10000` → fee 100, net 9900 → total 9900
- `25000` → fee 250, net 24750 → total 34650
- `5000` → fee 50, net 4950 → total 39600

Note how `totalProcessedCents` **remembers** its previous value between calls — that's the package-level variable at work. The locals `fee`/`amountCents` and the total don't survive; locals die at the end of each function, the package-level one persists.

### Line 21

```go
fmt.Println("Final ledger total:", totalProcessedCents)
```

Reads the final state after all transactions: `Final ledger total: 39600`. The package-level variable is still accessible in `main`.

### Line 22

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
processed 10000 cents (fee 100), running total now 9900
processed 25000 cents (fee 250), running total now 34650
processed 5000 cents (fee 50), running total now 39600
Final ledger total: 39600
```

## Key Takeaways

1. **Package-level variable** — declared outside functions; visible to all; program lifetime.
2. **Local variable** — declared inside a function; visible only there; gone when it ends.
3. **Shared state** — package-level variables carry data between function calls.
4. **`Printf` + `%d`** — formatted output.
5. **Zero value** — `var totalProcessedCents int64` starts at `0`.