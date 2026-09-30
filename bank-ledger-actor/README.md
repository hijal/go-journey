# bank-ledger-actor

Go-তে **actor model** — একটাই goroutine (actor) balance map-এর মালিক, বাকি সবাই message পাঠায় ও reply channel-এ উত্তর পায়। কোনো mutex নেই।

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

### Lines 3–7

```go
import (
	"errors"
	"fmt"
	"sync"
)
```

`errors` (New/Is), `fmt`, `sync` (WaitGroup)।

### Line 9

```go
var errInsufficientFunds = errors.New("insufficient funds")
```

**Sentinel error** — `errors.Is` দিয়ে চিনতে পারা যায়।

### Lines 11–17

```go
type opKind int

const (
	deposit opKind = iota
	withdraw
	balance
)
```

- `opKind` — typed enum; `iota` দিয়ে 0, 1, 2।
- `balance` শুধু একটা **প্রশ্ন**, কোনো mutation নয়।

### Lines 19–29

```go
type response struct {
	balance int
	err     error
}

type request struct {
	kind    opKind
	account string
	amount  int
	reply   chan response
}
```

- **request** = message: কী করতে হবে (`kind`), কোন account, কত টাকা, আর উত্তর কোথায় পাঠাতে হবে (`reply` channel)।
- **response** = reply: balance অথবা error।
- **`reply chan response`** — প্রতিটি request-এর নিজস্ব reply channel, তাই অন্য কাজের উত্তরের সাথে গুলিয়ে যায় না।

### Lines 31–52

```go
func ledger(requests <-chan request) {
	balances := map[string]int{}

	for req := range requests {
		switch req.kind {
		case deposit:
			balances[req.account] += req.amount
		case withdraw:
			if balances[req.account] < req.amount {
				req.reply <- response{
					balance: balances[req.account],
					err:     fmt.Errorf("withdraw %d from %s: %w", req.amount, req.account, errInsufficientFunds),
				}
				continue
			}
			balances[req.account] -= req.amount
		}
		req.reply <- response{
			balance: balances[req.account],
		}
	}
}
```

- **`balances := map[string]int{}`** — এই map **শুধু এই goroutine-ই** ছুঁতে পারে (actor-ই এর একমাত্র owner)। তাই lock দরকার হয় না, আর data race-ও সম্ভব না।
- **`for req := range requests`** — channel বন্ধ হলে actor শেষ।
- **`balance` কোনো case নেই** — শুধু শেষের reply-তে বর্তমান balance পঠানো হয় (শুধু পড়া)।
- **অপর্যাপ্ত তহবিলে** — `%w`-wrapped error পাঠিয়ে `continue`, তাই balance অক্ষত থাকে (rollback লাগে না)।

### Lines 54–60

```go
func call(requests chan<- request, kind opKind, account string, amount int) (int, error) {
	reply := make(chan response, 1)
	requests <- request{kind: kind, account: account, amount: amount, reply: reply}

	res := <-reply
	return res.balance, res.err
}
```

- **`make(chan response, 1)`** — buffer 1, তাই reply পাঠাতে actor কখনো block করে না; অপেক্ষা করে শুধু caller।
- **`<-reply`** — এই call-ই নিজের উত্তরের জন্য অপেক্ষা করে।

### Lines 62–96

```go
func main() {
	requests := make(chan request)
	go ledger(requests)

	var wg sync.WaitGroup
	for range 100 {
		wg.Add(1)

		go func() {
			defer wg.Done()

			if _, err := call(requests, deposit, "ACC-77", 10); err != nil {
				fmt.Println("deposit failed:", err)
			}
		}()
	}
	wg.Wait()

	bal, _ := call(requests, balance, "ACC-77", 0)
	fmt.Println("balance after deposits:", bal)

	if _, err := call(requests, withdraw, "ACC-77", 5000); err != nil {
		fmt.Println("error:", err, "| insufficient?", errors.Is(err, errInsufficientFunds))
	}

	bal, err := call(requests, withdraw, "ACC-77", 400)

	if err != nil {
		fmt.Println("unexpected error:", err)
		return
	}

	fmt.Println("balance after withdrawing 400:", bal)
	close(requests)
}
```

- ১০০টি deposit goroutine → ১০০×১০ = **১০০০** (সবাই এক actor-এ সিরিয়ালাইজ, তাই যোগ exactly ১০০০)।
- **`errors.Is(err, errInsufficientFunds)`** → `true`।
- **শেষে `close(requests)`** — actor-এর loop শেষ, goroutine বেরিয়ে যায়।

---

## Expected Output

```
balance after deposits: 1000
error: withdraw 5000 from ACC-77: insufficient funds | insufficient? true
balance after withdrawing 400: 600
```

## মূল শিক্ষা / Key Takeaways

1. **Actor = একা মালিক** — `balances` map-এ একটি goroutine-ই লেখে, তাই mutex অপ্রয়োজনীয়।
2. **Request/reply channel** — সরাসরি function call নয়, message পাঠানো হয়।
3. **`reply chan response` প্রতি request-এ আলাদা** — উত্তর গুলায় না।
4. **buffer 1 reply channel** — actor-কে block করে না।
5. **`%w` + `errors.Is`** — typed error propagate ও পরীক্ষা।
6. **`close(requests)`** — actor-কে ধ্বংস করার শেষ পথ।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Lines 3–7

```go
import (
	"errors"
	"fmt"
	"sync"
)
```

`errors` (New/Is), `fmt`, `sync` (WaitGroup).

### Line 9

```go
var errInsufficientFunds = errors.New("insufficient funds")
```

A **sentinel error** — identifiable through `errors.Is`.

### Lines 11–17

```go
type opKind int

const (
	deposit opKind = iota
	withdraw
	balance
)
```

- `opKind` is a typed enum; `iota` gives 0, 1, 2.
- `balance` is only a **query**, never a mutation.

### Lines 19–29

```go
type response struct {
	balance int
	err     error
}

type request struct {
	kind    opKind
	account string
	amount  int
	reply   chan response
}
```

- A **request** is the message: what to do (`kind`), which account, how much, and where the answer should go (`reply` channel).
- A **response** is the reply: a balance or an error.
- **`reply chan response`** — every request owns its own reply channel, so answers never get mixed up between callers.

### Lines 31–52

```go
func ledger(requests <-chan request) {
	balances := map[string]int{}

	for req := range requests {
		switch req.kind {
		case deposit:
			balances[req.account] += req.amount
		case withdraw:
			if balances[req.account] < req.amount {
				req.reply <- response{
					balance: balances[req.account],
					err:     fmt.Errorf("withdraw %d from %s: %w", req.amount, req.account, errInsufficientFunds),
				}
				continue
			}
			balances[req.account] -= req.amount
		}
		req.reply <- response{
			balance: balances[req.account],
		}
	}
}
```

- **`balances := map[string]int{}`** — only this goroutine may touch that map (the actor is its sole owner). Hence no lock is needed and a data race is impossible.
- **`for req := range requests`** — the actor finishes when the channel closes.
- **No `case` for `balance`** — it simply reports the current balance in the final reply (a read-only op).
- **On insufficient funds** it sends a `%w`-wrapped error and `continue`s, so the balance stays untouched (no rollback needed).

### Lines 54–60

```go
func call(requests chan<- request, kind opKind, account string, amount int) (int, error) {
	reply := make(chan response, 1)
	requests <- request{kind: kind, account: account, amount: amount, reply: reply}

	res := <-reply
	return res.balance, res.err
}
```

- **`make(chan response, 1)`** — a buffer of 1, so the actor never blocks on a reply send.
- **`<-reply`** — this call waits for its own answer and nothing else.

### Lines 62–96

```go
func main() {
	requests := make(chan request)
	go ledger(requests)

	var wg sync.WaitGroup
	for range 100 {
		wg.Add(1)

		go func() {
			defer wg.Done()

			if _, err := call(requests, deposit, "ACC-77", 10); err != nil {
				fmt.Println("deposit failed:", err)
			}
		}()
	}
	wg.Wait()

	bal, _ := call(requests, balance, "ACC-77", 0)
	fmt.Println("balance after deposits:", bal)

	if _, err := call(requests, withdraw, "ACC-77", 5000); err != nil {
		fmt.Println("error:", err, "| insufficient?", errors.Is(err, errInsufficientFunds))
	}

	bal, err := call(requests, withdraw, "ACC-77", 400)

	if err != nil {
		fmt.Println("unexpected error:", err)
		return
	}

	fmt.Println("balance after withdrawing 400:", bal)
	close(requests)
}
```

- 100 deposit goroutines → 100×10 = **1000**; they serialise inside the single actor, so the total is exactly 1000.
- **`errors.Is(err, errInsufficientFunds)`** → `true`.
- **`close(requests)`** at the end — the actor's loop ends and the goroutine exits.

---

## Expected Output

```
balance after deposits: 1000
error: withdraw 5000 from ACC-77: insufficient funds | insufficient? true
balance after withdrawing 400: 600
```

## Key Takeaways

1. **An actor has exclusive ownership** — one goroutine writes `balances`, so no mutex is required.
2. **Request/reply channels** — communication happens by message, not a direct function call.
3. **A separate `reply` channel per request** — answers never get crossed.
4. **A buffer of 1 on the reply channel** — keeps the actor from ever blocking.
5. **`%w` + `errors.Is`** — propagate and identify typed errors.
6. **`close(requests)`** — the standard way to shut an actor down.