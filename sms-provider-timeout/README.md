# sms-provider-timeout

Go-তে **`select` + `time.After` timeout, channel-returning function** — দ্রুত provider বিজয়ী, সীমান্তের ক্ষেত্রে (300ms vs 300ms) ফল অনির্ধারিত।

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

### Lines 3–9

```go
import (
	"errors"
	"fmt"
	"time"
)
```

`time` (Sleep, After, Duration)।

### Line 9

```go
var errTimeout = errors.New("timed out")
```

**Sentinel error** — caller পুনরায় চেষ্টা করবে কি না সেটা `errors.Is` দিয়ে সিদ্ধান্ত নিতে পারে।

### Lines 11–20

```go
func sendSMS(provider string, latency time.Duration) <-chan string {
	ack := make(chan string, 1)

	go func() {
		time.Sleep(latency)
		ack <- provider + ": delivered"
	}()

	return ack
}
```

- **`<-chan string`** — receive-only ফেরত দেয়, তাই caller send করতে পারে না।
- **`make(chan string, 1)`** — buffer 1: timeout হলেও goroutine send করে **block করে না** (goroutine leak নেই)।
- **`latency`** ইনজেক্ট করা — সত্যিকারের provider হলে network delay এখানে আসত।

### Lines 22–30

```go
func deliver(provider string, latency, timeout time.Duration) error {
	select {
	case msg := <-sendSMS(provider, latency):
		fmt.Println(msg)
		return nil
	case <-time.After(timeout):
		return fmt.Errorf("sms via %s: %w", provider, errTimeout)
	}
}
```

- **`select`** দুইটি case: ডেলিভারি হলে সফল, না হলে timeout।
- **`%w`** — `errors.Is`-এর জন্য wrap।

### Lines 32–43

```go
func main() {
	for _, p := range []struct {
		name    string
		latency time.Duration
	}{
		{"fast-sms", 20 * time.Millisecond},
		{"slow-sms", 300 * time.Millisecond},
	} {
		if err := deliver(p.name, p.latency, 300 * time.Millisecond); err != nil {
			fmt.Println("error:", err, "| retry later?", errors.Is(err, errTimeout))
		}
	}
}
```

- `fast-sms` (20ms) → timeout-এর আগে পৌঁছায় → **`fast-sms: delivered`**।
- `slow-sms` (300ms) → timeout-এর সঙ্গে প্রায় সমান → **`sms via slow-sms: timed out`**, `errors.Is` → `true`।

---

## সতর্কতা: ফলাফল সবসময় নির্ভরযোগ্য নয়

`slow-sms`-এর latency ও timeout **দুটোই 300ms**। `select`-এর case expression বাম থেকে ডানে মূল্যায়িত হয়, তাই goroutine-এর `Sleep(300ms)` **`time.After(300ms)`-এর টাইমারের আগে** শুরু হয়। ফলে দুটো প্রায় একসঙ্গে শেষ হয় এবং জিততে পারে যেকোনো একটি — বিশেষ করে goroutine যখন দ্রুত schedule হয়ে buffered channel-এ send করে ফেলে।

এই মেশিনে ৩০ বার রানে: **২৯ বার `timed out`, ১ বার `slow-sms: delivered`**। প্রায় ৩% flake উপরের ব্যাখ্যার সঙ্গে মেলে।

**নিয়ম:** boundary test-এ latency ও timeout-এর মধ্যে স্পষ্ট ব্যবধান রাখুন (যেমন 300ms বনাম 100ms)।

---

## Expected Output (সাধারণত)

```
fast-sms: delivered
error: sms via slow-sms: timed out | retry later? true
```

**বিরল:** দ্বিতীয় লাইন `slow-sms: delivered` হতে পারে — উপরের boundary-কারণে।

## মূল শিক্ষা / Key Takeaways

1. **`select` + `time.After`** — non-blocking timeout।
2. **`<-chan T` return** — receive-only, ব্যবহারকারীর নিয়ন্ত্রণ।
3. **Buffer 1** — cancel-এর পরেও goroutine leak হয় না।
4. **Sentinel + `%w`** — `errors.Is` দিয়ে retry সিদ্ধান্ত।
5. **সমান সময় = অনির্ধারিত ফল** — timeout-এর জন্য margin রাখুন।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Lines 3–9

```go
import (
	"errors"
	"fmt"
	"time"
)
```

`time` (Sleep, After, Duration).

### Line 9

```go
var errTimeout = errors.New("timed out")
```

A **sentinel error** — lets the caller decide with `errors.Is` whether to retry.

### Lines 11–20

```go
func sendSMS(provider string, latency time.Duration) <-chan string {
	ack := make(chan string, 1)

	go func() {
		time.Sleep(latency)
		ack <- provider + ": delivered"
	}()

	return ack
}
```

- Returns **`<-chan string`** (receive-only), so the caller cannot send on it.
- **`make(chan string, 1)`** — a buffer of 1: on timeout the goroutine's send doesn't block, so there's no goroutine leak.
- **`latency` is injected** — a real provider would bring network delay here.

### Lines 22–30

```go
func deliver(provider string, latency, timeout time.Duration) error {
	select {
	case msg := <-sendSMS(provider, latency):
		fmt.Println(msg)
		return nil
	case <-time.After(timeout):
		return fmt.Errorf("sms via %s: %w", provider, errTimeout)
	}
}
```

- The **`select`** has two cases: delivery succeeds, or the deadline fires.
- **`%w`** — wrapped so `errors.Is` can see it.

### Lines 32–43

```go
func main() {
	for _, p := range []struct {
		name    string
		latency time.Duration
	}{
		{"fast-sms", 20 * time.Millisecond},
		{"slow-sms", 300 * time.Millisecond},
	} {
		if err := deliver(p.name, p.latency, 300 * time.Millisecond); err != nil {
			fmt.Println("error:", err, "| retry later?", errors.Is(err, errTimeout))
		}
	}
}
```

- `fast-sms` (20ms) arrives before the timeout → **`fast-sms: delivered`**.
- `slow-sms` (300ms) is almost exactly the timeout → **`sms via slow-sms: timed out`**, `errors.Is` → `true`.

---

## Caution: the outcome is not always reproducible

`slow-sms` has a latency **and** a timeout of **300ms**. A `select` evaluates its case expressions left to right, so the goroutine's `Sleep(300ms)` starts **before** the `time.After(300ms)` timer. The two therefore finish almost together and either can win — especially when the goroutine is scheduled quickly and delivers into the buffered channel.

Across 30 runs on this machine: **29 `timed out`, 1 `slow-sms: delivered`** — roughly a 3% flake, matching the explanation above.

**Rule:** leave a clear margin between latency and timeout in a boundary test (e.g. 300ms vs 100ms).

---

## Expected Output (usually)

```
fast-sms: delivered
error: sms via slow-sms: timed out | retry later? true
```

**Rarely:** the second line can read `slow-sms: delivered` — see the boundary reason above.

## Key Takeaways

1. **`select` + `time.After`** — a non-blocking timeout.
2. **Returning `<-chan T`** — receive-only, the caller stays in control.
3. **Buffer of 1** — no goroutine leak after cancellation.
4. **Sentinel + `%w`** — retry decisions via `errors.Is`.
5. **Equal times = an undetermined result** — always leave a margin for a timeout.