# higher-order-retry-helper

Go-তে **higher-order function** (`withRetry` operation function নেয়), **`errors.Is`** unwrap-check আর **`%w` error-wrapping** দিয়ে retry helper শেখার ছোট example।

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
	"time"
)
```

- `errors` — sentinel error + `errors.Is`।
- `fmt` — `Errorf`, `Println`।
- `time` — `time.Sleep`।

### Line 9

```go
var errGatewayDown = errors.New("Payment gateway down")
```

**Sentinel error** — একটা named package-level error, যা wrap/compare করা যায়।

### Lines 11–28

```go
func withRetry(attempts int, operation func() error) error {
	var lastErr error

	for i := range attempts {
		err := operation()

		if err == nil {
			fmt.Printf("attempt %d: success\n", i+1)
			return nil
		}

		lastErr = err

		fmt.Printf("attempt %d failed: %v\n", i+1, err)
		time.Sleep(100 * time.Millisecond)
	}
	return fmt.Errorf("withRetry: all %d attempts failed: %w", attempts, lastErr)
}
```

`withRetry` — একটা **higher-order function**: একটা `operation func() error` নেয় এবং retry logic সরবরাহ করে (প্রদান করে):

- `for i := range attempts` — attempts numbers-round loop।
- প্রতিটা round-এ `operation()` call; `nil` হলে success print + `nil` return (early success)।
- fail হলে `lastErr` update, fail message print, `100ms` অপেক্ষা (backoff-লাইট)।
- সব attempt শেষে `fmt.Errorf("...: %w", lastErr)` — **`%w` wrapping**: শেষ error-টাকে wrap করা হয় যাতে caller `errors.Is`-এ unwrap করতে পারে।
- return value: `"withRetry: all N attempts failed: <last err>"`।

**এই উদাহরণে** — higher-order pattern: `operation`-টা যেকোনো fail-able callback (ব্যর্থ হতে পারে এমন ফাংশন); retry generic-reusable।

### Lines 30–45

```go
func main() {
	calls := 0
	flakyCharge := func() error {
		calls++
		if calls <= 2 {
			return fmt.Errorf("charge 4999 cents: %w", errGatewayDown)
		}
		return nil
	}

	if err := withRetry(5, flakyCharge); err != nil {
		fmt.Println("gave up:", err)
	} else {
		fmt.Println("payment captured")
	}
```

`flakyCharge` — একটা closure যা প্রথম ২ বার fail করে, পরে succeed:

- attempt 1 fail (`calls=1`), attempt 2 fail (`calls=2`), attempt 3 success (`calls=3`)।
- `withRetry(5, flakyCharge)` — attempt 3-এ `nil` return → `"payment captured"`।

**`%w` signal:** প্রতিটা fail-এ `fmt.Errorf("charge 4999 cents: %w", errGatewayDown)` — sentinel-এ wrap করা, যাতে `errors.Is` later unwrap করতে পারে।

### Lines 47–58

```go
	fmt.Println("---")

	deadCharge := func() error {
		return fmt.Errorf("charge 1299 cents: %w", errGatewayDown)
	}

	if err := withRetry(3, deadCharge); err != nil {
		fmt.Println("final error:", err)
		if errors.Is(err, errGatewayDown) {
			fmt.Println("gateway down - enqueue job for later")
		}
	}
}
```

`deadCharge` — **সবসময় fail** (গেটে always error):

- 3 attempts fail; `withRetry` শেষ attempt-এর error-টাকে wrap করে return।
- `fmt.Println("final error:", ...)` — `withRetry: all 3 attempts failed: charge 1299 cents: Payment gateway down`।
- **`errors.Is(err, errGatewayDown)`** — nil-error independent: wrapping chain-এর মধ্যে sentinel মিললো কি না। এখানে `%w` এর মাধ্যমে match হয় (গেটে always-down), তাই print: `gateway down - enqueue job for later`।

**Failure-handling decision:** caller-টি error-র root cause জানতে পারে — retry exhausted হলেও।

---

## Expected Output

```
attempt 1 failed: charge 4999 cents: Payment gateway down
attempt 2 failed: charge 4999 cents: Payment gateway down
attempt 3: success
payment captured
---
attempt 1 failed: charge 1299 cents: Payment gateway down
attempt 2 failed: charge 1299 cents: Payment gateway down
attempt 3 failed: charge 1299 cents: Payment gateway down
final error: withRetry: all 3 attempts failed: charge 1299 cents: Payment gateway down
gateway down - enqueue job for later
```

## মূল শিক্ষা / Key Takeaways

1. **Higher-order function** — `withRetry(attempts, op func() error)` — generic retry wrapper।
2. **Sentinel error** — `errGatewayDown` — named error-টা `errors.New`-এ।
3. **`%w` wrapping** — fail message-এ Original error চেইন-এ রাখা।
4. **`errors.Is`** — wrapping-এর মধ্য দিয়ে root cause match।
5. **Closure-as-operation** — `operation`-টা state capture করে (প্রথম ২ fail)।