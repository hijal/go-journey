# network-retry-mechanism

Go-তে **retry mechanism** শেখার ছোট example — exponential backoff-এর মতো linear backoff + jitter সহ flaky API call-টা ম্যানেজ করে।

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

### Lines 3–8

```go
import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)
```

- `errors` — `errors.New` দিয়ে error তৈরি।
- `fmt` — `Errorf`, `Printf`।
- `math/rand` — jitter + flaky simulate-র জন্য random।
- `time` — `Duration`, `Sleep`।

### Lines 10–29

```go
func Retry(maxAttempts int, baseDelay time.Duration, operation func() error) error {
	var err error

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		err = operation()

		if err == nil {
			return nil
		}

		if attempt < maxAttempts {
			delay := baseDelay * time.Duration(attempt)
			jitter := time.Duration(rand.Intn(100) * int(time.Millisecond))
			time.Sleep(delay + jitter)
			fmt.Printf("Attempt %d failed: %v. Retrying...\n", attempt, err)
		}
	}

	return fmt.Errorf("operation failed after %d attempts: %w", maxAttempts, err)
}
```

`Retry` — একটা operation-কে `maxAttempts` বার চেষ্টা করে:

- `var err error` — শেষ failure টা রাখার জন্য (outside loop-এ declare, কারণ loop শেষে এটা দরকার)।
- Loop `attempt := 1; attempt <= maxAttempts` — এক-ভিত্তিক attempt counter।
- `err = operation()` — কাজটা call; fail হলে `err != nil`।
- **Success:** `if err == nil { return nil }` — সফল হলেই সাথে সাথে return (আর retry লাগবে না)।
- **Retry block** (`attempt < maxAttempts`):
  - `delay := baseDelay * time.Duration(attempt)` — **linear backoff**: attempt বাড়লে delay বাড়ে (100ms, 200ms, ...)
  - `jitter := rand.Intn(100) * time.Millisecond` — 0–99ms random jitter, যেন সব client একসাথে আবার hit না করে।
  - `time.Sleep(delay + jitter)` — retry-র আগে অপেক্ষা।
  - `Printf` — কোন attempt fail-লো message।
- Loop শেষে: `fmt.Errorf("operation failed after %d attempts: %w", maxAttempts, err)` — শেষ error-টা **`%w`** দিয়ে wrap, সবার শেষে report। `%w` error-টাকে chain-এ রাখে, caller `errors.Is`/`Unwrap` করতে পারে।

**Important:** attempt counter top-level-এ `err != nil`-এ return করে, call করে না — success হলেই exit, fail-এ retry।

### Lines 31–45

```go
func main() {
	flakyAPICall := func() error {
		if rand.Float32() < 0.7 {
			return errors.New("connection timeout")
		}
		return nil
	}

	err := Retry(3, 100*time.Millisecond, flakyAPICall)

	if err != nil {
		fmt.Println("final failure:", err)
	} else {
		fmt.Println("API call succeed")
	}
}
```

- `flakyAPICall` — একটি closure: **70% chance** `"connection timeout"` error, 30% chance success।
- `Retry(3, 100*time.Millisecond, flakyAPICall)` — max 3 attempt, base delay 100ms।
- Result: fail-হলে `final failure: ...`, success-এ `API call succeed`।

---

## Expected Output

(প্রতি run-এ random-এর জন্য ভিন্ন; ৩টা representative case)

**Case A — সব fail (≈34% chance):**
```
Attempt 1 failed: connection timeout. Retrying...
Attempt 2 failed: connection timeout. Retrying...
final failure: operation failed after 3 attempts: connection timeout
```

**Case B — ৩য় attempt-এ সফল:**
```
Attempt 1 failed: connection timeout. Retrying...
Attempt 2 failed: connection timeout. Retrying...
API call succeed
```

**Case C — প্রথম attempt-এ সফল:**
```
API call succeed
```

## মূল শিক্ষা / Key Takeaways

1. **Retry loop** — `for attempt := 1; attempt <= maxAttempts; attempt++`।
2. **Fail-fast on success** — `err == nil` এ সাথে সাথে return।
3. **Linear backoff + jitter** — `baseDelay * attempt` + random jitter, thundering herd এড়ানো।
4. **`%w` wrapping** — শেষ error-টা wrap, caller-কে actual cause।
5. **Closure as operation** — fail/succeed ফাংশন-কল-এর ভেতরে।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Lines 3–8

```go
import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)
```

- `errors` — for creating errors with `errors.New`.
- `fmt` — for `Errorf`, `Printf`.
- `math/rand` — for jitter and simulating flakiness.
- `time` — for `Duration`, `Sleep`.

### Lines 10–29

```go
func Retry(maxAttempts int, baseDelay time.Duration, operation func() error) error {
	var err error

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		err = operation()

		if err == nil {
			return nil
		}

		if attempt < maxAttempts {
			delay := baseDelay * time.Duration(attempt)
			jitter := time.Duration(rand.Intn(100) * int(time.Millisecond))
			time.Sleep(delay + jitter)
			fmt.Printf("Attempt %d failed: %v. Retrying...\n", attempt, err)
		}
	}

	return fmt.Errorf("operation failed after %d attempts: %w", maxAttempts, err)
}
```

`Retry` — tries an operation up to `maxAttempts` times:

- `var err error` — holds the last failure (declared outside the loop, needed after it).
- Loop `attempt := 1; attempt <= maxAttempts` — 1-based attempt counter.
- `err = operation()` — runs the task; `err != nil` on failure.
- **Success:** `if err == nil { return nil }` — return immediately on success (no more retries).
- **Retry block** (`attempt < maxAttempts`):
  - `delay := baseDelay * time.Duration(attempt)` — **linear backoff**: delay grows with each attempt (100ms, 200ms, ...).
  - `jitter := rand.Intn(100) * time.Millisecond` — random 0–99ms jitter so all clients don't hit again at once.
  - `time.Sleep(delay + jitter)` — waits before the next try.
  - `Printf` — reports which attempt failed.
- After the loop: `fmt.Errorf("operation failed after %d attempts: %w", maxAttempts, err)` — wraps the last error with **`%w`** and reports it. `%w` keeps the error in the chain so callers can use `errors.Is`/`Unwrap`.

**Note:** the success check returns on `err == nil` — exit on success, retry on failure.

### Lines 31–45

```go
func main() {
	flakyAPICall := func() error {
		if rand.Float32() < 0.7 {
			return errors.New("connection timeout")
		}
		return nil
	}

	err := Retry(3, 100*time.Millisecond, flakyAPICall)

	if err != nil {
		fmt.Println("final failure:", err)
	} else {
		fmt.Println("API call succeed")
	}
}
```

- `flakyAPICall` — a closure: **70% chance** of a `"connection timeout"` error, 30% success.
- `Retry(3, 100*time.Millisecond, flakyAPICall)` — at most 3 attempts, 100ms base delay.
- On failure: `final failure: ...`; on success: `API call succeed`.

---

## Expected Output

(Random — varies per run; 3 representative cases)

**Case A — all attempts fail (≈34% chance):**
```
Attempt 1 failed: connection timeout. Retrying...
Attempt 2 failed: connection timeout. Retrying...
final failure: operation failed after 3 attempts: connection timeout
```

**Case B — success on attempt 3:**
```
Attempt 1 failed: connection timeout. Retrying...
Attempt 2 failed: connection timeout. Retrying...
API call succeed
```

**Case C — success on first attempt:**
```
API call succeed
```

## Key Takeaways

1. **Retry loop** — `for attempt := 1; attempt <= maxAttempts; attempt++`.
2. **Fail-fast on success** — return immediately when `err == nil`.
3. **Linear backoff + jitter** — `baseDelay * attempt` plus random jitter to avoid a thundering herd.
4. **`%w` wrapping** — wrap the final error to preserve the underlying cause.
5. **Closure as operation** — fail/succeed happens inside the function call.