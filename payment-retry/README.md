# payment-retry

Go-তে **sentinel error + `errors.Is` + `%w` wrap**, **exponential backoff (`time.Sleep`)** আর **countdown retry loop** শেখার ছোট example.

**ভাষা নির্বাচন করুন / Choose language:**

[বাংলা](#bangla) * [English](#english)

---

<a name="bangla"></a>

## বাংলা সংস্করণ

### Lines 1-7

```go
package main

import (
	"errors"
	"fmt"
	"time"
)
```

- `package main` -- একটা executable program.
- `errors` -- sentinel error বানানো (`errors.New`) আর চেনার (`errors.Is`) জন্য.
- `fmt` -- output + error wrap (`%w`) এর জন্য.
- `time` -- delay (`time.Sleep`) আর duration (`time.Millisecond`) এর জন্য.

### Lines 9-10

```go
var errPaymentDeclined = errors.New("payment declined")
var gatewayCalls int
```

দুটো **package-level variable**:

- `errPaymentDeclined` -- **sentinel error**: একটা fixed value যা দিয়ে "declined" failure চেনা যায়. `var` + `errors.New` দিয়ে একবার বানিয়ে বারবার compare করা হয়.
- `gatewayCalls` -- fake gateway কতবার call হয়েছে তার counter (zero value 0). প্রথম 2 বার fail, 3য় বার success simulate করতে ব্যবহার হয়.

### Lines 12-18

```go
func charge(amount int) error {
	gatewayCalls++
	if gatewayCalls < 3 {
		return fmt.Errorf("charge %d BDT: %w", amount, errPaymentDeclined)
	}
	return nil
}
```

`charge` -- payment gateway simulate করে:

- `gatewayCalls++` -- প্রতিটা call-এ counter বাড়ে.
- প্রথম 2 বার (`gatewayCalls` 1, 2): `fmt.Errorf("charge %d BDT: %w", ...)` দিয়ে wrapped error return করে. `%w` দিয়ে `errPaymentDeclined`-কে ভেতরে wrap করা হয় -- উপরের message (`charge 1500 BDT: ...`) যোগ হয়, কিন্তু `errors.Is` দিয়ে ভেতরের sentinel-টা এখনো চেনা যায়.
- 3য় বার থেকে (`gatewayCalls >= 3`): `nil` return -- success.

### Line 20

```go
func main() {
```

Program-এর entry point.

### Lines 21-22

```go
const amount = 1500
backoff := 100 * time.Millisecond
```

- `const amount = 1500` -- 1500 BDT charge (constant, বদলায় না).
- `backoff := 100 * time.Millisecond` -- প্রথম retry-এর আগে 100ms অপেক্ষা. প্রতিবার double হবে (exponential).

### Lines 24-26

```go
remainingAttempts := 5
for remainingAttempts > 0 {
	remainingAttempts--
```

**Countdown retry loop** -- `remainingAttempts` 5 থেকে শুরু, প্রতিটা attempt-এ 1 কমে. `> 0` থাকা পর্যন্ত চলে (সর্বোচ্চ 5 বার চেষ্টা).

### Lines 28-32

```go
	err := charge(amount)
	if err == nil {
		fmt.Printf("payment of %d BDT succeeded\n", amount)
		break
	}
```

- `charge(amount)` call করে.
- `err == nil` মানে success -- message print করে `break` দিয়ে loop থেকে বেরিয়ে যাওয়া.

### Lines 34-37

```go
	if !errors.Is(err, errPaymentDeclined) {
		fmt.Println("unexpected error - aborting:", err)
		return
	}
```

**Error classification:** `errors.Is(err, errPaymentDeclined)` চেক করে error-টা wrapped হলেও ভেতরে sentinel-টা আছে কি না:

- Declined error হলে -> নিচে গিয়ে retry করে.
- অন্য কোনো unexpected error হলে -> retry না করে সাথে-সাথে `return` (abort). অচেনা error-এ বারবার retry করা বিপজ্জনক, তাই এই guard.

### Lines 39-41

```go
	fmt.Printf("charge failed (%v); retrying in %v\n", err, backoff)
	time.Sleep(backoff)
	backoff *= 2
```

Retry-এর আগে:

- `%v` দিয়ে error আর wait time print করে.
- `time.Sleep(backoff)` -- পরের attempt-এর আগে অপেক্ষা করে.
- `backoff *= 2` -- অপেক্ষা double করে (100ms -> 200ms -> 400ms...). এটাই **exponential backoff** -- fail বারবার হলে সার্ভারের উপর চাপ কমাতে wait বাড়ানো হয়.

**Flow:**

1. attempt 1 -> `gatewayCalls`=1 -> fail -> prints "retrying in 100ms" -> sleep 100ms -> backoff=200ms
2. attempt 2 -> `gatewayCalls`=2 -> fail -> prints "retrying in 200ms" -> sleep 200ms -> backoff=400ms
3. attempt 3 -> `gatewayCalls`=3 -> `nil` -> prints "payment of 1500 BDT succeeded" -> `break`

### Lines 44-46

```go
if remainingAttempts == 0 {
	fmt.Println("gave up: payment could not be completed")
}
```

Loop শেষে `remainingAttempts == 0` মানে সব 5 বার শেষ, একবারও success হয়নি -- তখন give-up message. এখানে 3য় attempt-এ success হয়, তাই এটা print হয় না.

### Line 47

```go
}
```

Closing brace -- `main` function শেষ হয়.

---

## Expected Output

```
charge failed (charge 1500 BDT: payment declined); retrying in 100ms
charge failed (charge 1500 BDT: payment declined); retrying in 200ms
payment of 1500 BDT succeeded
```

## মূল শিক্ষা / Key Takeaways

1. **Sentinel error** -- `errors.New` দিয়ে fixed error value; `==` এর বদলে `errors.Is` দিয়ে চেনা.
2. **`%w` wrap** -- context (`charge 1500 BDT`) যোগ করে, ভেতরের error চেনা যায়.
3. **Error classification** -- চেনা (declined) হলে retry, অচেনা হলে abort.
4. **Exponential backoff** -- `backoff *= 2` দিয়ে প্রতিবার wait double.
5. **Countdown loop** -- `remainingAttempts--` + `> 0`; শেষে `== 0` মানে give up.
6. **`err == nil`** -- success মানে `nil` error.

---

---

<a name="english"></a>

##  English Version

### Lines 1-7

```go
package main

import (
	"errors"
	"fmt"
	"time"
)
```

- `package main` -- an executable program.
- `errors` -- for creating a sentinel error (`errors.New`) and matching it (`errors.Is`).
- `fmt` -- for output + error wrapping (`%w`).
- `time` -- for delays (`time.Sleep`) and durations (`time.Millisecond`).

### Lines 9-10

```go
var errPaymentDeclined = errors.New("payment declined")
var gatewayCalls int
```

Two **package-level variables**:

- `errPaymentDeclined` -- a **sentinel error**: a fixed value used to recognize a "declined" failure. Created once with `var` + `errors.New` and compared repeatedly.
- `gatewayCalls` -- counts how many times the fake gateway was called (zero value 0). Used to fail the first 2 times and succeed on the 3rd.

### Lines 12-18

```go
func charge(amount int) error {
	gatewayCalls++
	if gatewayCalls < 3 {
		return fmt.Errorf("charge %d BDT: %w", amount, errPaymentDeclined)
	}
	return nil
}
```

`charge` -- simulates a payment gateway:

- `gatewayCalls++` -- increments on every call.
- First 2 calls (`gatewayCalls` 1, 2): returns a wrapped error via `fmt.Errorf("charge %d BDT: %w", ...)`. `%w` wraps `errPaymentDeclined` inside -- the outer message (`charge 1500 BDT: ...`) is added, but the inner sentinel is still detectable with `errors.Is`.
- From the 3rd call (`gatewayCalls >= 3`): returns `nil` -- success.

### Line 20

```go
func main() {
```

Program entry point.

### Lines 21-22

```go
const amount = 1500
backoff := 100 * time.Millisecond
```

- `const amount = 1500` -- a 1500 BDT charge (constant, never changes).
- `backoff := 100 * time.Millisecond` -- waits 100ms before the first retry. Doubles each time (exponential).

### Lines 24-26

```go
remainingAttempts := 5
for remainingAttempts > 0 {
	remainingAttempts--
```

A **countdown retry loop** -- `remainingAttempts` starts at 5 and decreases each attempt. Runs while `> 0` (at most 5 tries).

### Lines 28-32

```go
	err := charge(amount)
	if err == nil {
		fmt.Printf("payment of %d BDT succeeded\n", amount)
		break
	}
```

- Calls `charge(amount)`.
- `err == nil` means success -- prints a message and `break`s out of the loop.

### Lines 34-37

```go
	if !errors.Is(err, errPaymentDeclined) {
		fmt.Println("unexpected error - aborting:", err)
		return
	}
```

**Error classification:** `errors.Is(err, errPaymentDeclined)` checks whether the sentinel is inside, even when wrapped:

- Declined error -> falls through and retries below.
- Any other unexpected error -> `return` immediately (abort). Retrying an unknown error blindly is dangerous, hence this guard.

### Lines 39-41

```go
	fmt.Printf("charge failed (%v); retrying in %v\n", err, backoff)
	time.Sleep(backoff)
	backoff *= 2
```

Before retrying:

- Prints the error and wait time with `%v`.
- `time.Sleep(backoff)` -- waits before the next attempt.
- `backoff *= 2` -- doubles the wait (100ms -> 200ms -> 400ms...). This is **exponential backoff** -- the wait grows so repeated failures don't hammer the server.

**Flow:**

1. attempt 1 -> `gatewayCalls`=1 -> fail -> prints "retrying in 100ms" -> sleep 100ms -> backoff=200ms
2. attempt 2 -> `gatewayCalls`=2 -> fail -> prints "retrying in 200ms" -> sleep 200ms -> backoff=400ms
3. attempt 3 -> `gatewayCalls`=3 -> `nil` -> prints "payment of 1500 BDT succeeded" -> `break`

### Lines 44-46

```go
if remainingAttempts == 0 {
	fmt.Println("gave up: payment could not be completed")
}
```

After the loop, `remainingAttempts == 0` means all 5 tries were used with no success -- print a give-up message. Here the 3rd attempt succeeds, so this never prints.

### Line 47

```go
}
```

Closing brace -- ends the `main` function.

---

## Expected Output

```
charge failed (charge 1500 BDT: payment declined); retrying in 100ms
charge failed (charge 1500 BDT: payment declined); retrying in 200ms
payment of 1500 BDT succeeded
```

## Key Takeaways

1. **Sentinel error** -- a fixed error value via `errors.New`; match with `errors.Is`, not `==`.
2. **`%w` wrap** -- adds context (`charge 1500 BDT`) while keeping the inner error detectable.
3. **Error classification** -- retry on known (declined), abort on unknown.
4. **Exponential backoff** -- `backoff *= 2` doubles the wait each time.
5. **Countdown loop** -- `remainingAttempts--` + `> 0`; `== 0` at the end means give up.
6. **`err == nil`** -- a `nil` error means success.
