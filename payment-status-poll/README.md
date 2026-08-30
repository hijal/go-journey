# payment-status-poll

Go-তে **goroutine**, **channel**, আর **`select`** দিয়ে payment-status polling-এ timeout সহ অপেক্ষা করা শেখার ছোট example।

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
    "fmt"
    "time"
)
```

দুটো package import করা হয়:

- `fmt` — console-এ output print করার জন্য।
- `time` — `time.Duration`, `time.Sleep`, আর `time.After` ব্যবহারের জন্য (sleep/delay আর timeout manage করতে)।

### Lines 8–11

```go
func checkPaymentStatus(statusCh chan<- string, delay time.Duration) {
    time.Sleep(delay)
    statusCh <- "confirmed"
}
```

`checkPaymentStatus` নামক একটা helper:

- domain-টা `chan<- string` — একটা **send-only channel**। Function-টা channel-এ শুধু data পাঠাতে পারে, read করতে পারে না।
- `delay time.Duration` — কতক্ষণ ঘুমিয়ে থাকবে (payment gateway-র মতো delayed response simulate করা)।
- `time.Sleep(delay)` — `delay` সময় অপেক্ষা করে, তারপর `statusCh <- "confirmed"` দিয়ে channel-এ `"confirmed"` পাঠায়।

মানে: payment "confirmed" হওয়াটা আসতে `delay` সময় লাগে।

### Lines 13–24

```go
func pollWithTimeout(delay, timeout time.Duration) string {
    statusCh := make(chan string)

    go checkPaymentStatus(statusCh, delay)

    select {
    case status := <-statusCh:
        return status
    case <-time.After(timeout):
        return "time-out"
    }
}
```

`pollWithTimeout` — মূল function:

- `make(chan string)` — একটা **unbuffered channel** (capacity 0)।
- `go checkPaymentStatus(statusCh, delay)` — goroutine হিসেবে status checker চালায়; এটা `delay` পরে answer দেবে।
- **`select`** — Go-র শক্তিশালী টুল: একসাথে **একাধিক channel operation-এর মধ্যে** অপেক্ষা করে, **যেটা আগে ready হয় সেটা চলে**:
    - `case status := <-statusCh:` — status channel-এ থেকে value আসলেই `status` আনে, `return status` (`"confirmed"`)।
    - `case <-time.After(timeout):` — `time.After(timeout)` একটা channel return করে যেটা `timeout` পর একবার value দেয়। যদি সেটা আগে ready হয় (মানে payment-টা `delay` সময়ের মধ্যে আসেনি), তাহলে `"time-out"` return করে।

অর্থাৎ **রেস**: payment confirm আর timeout — দুটোর মধ্যে যেটা আগে আসে, select সেটা নেয়। এটা হল **timeout-সহ poll** করার idiomatic Go উপায় — চিরকাল ঝুলে থাকা থেকে বাঁচায়।

> **নোট:** select-এর unwinnable case-এ (যেমন error-এ) ফিরে আসা যায় না; এখানে simple design-এ দুটো case-ই সবসময় একটার একটা জেতে, তাই কোনো blocking leak নেই। (Timeout-এর পরও goroutine যদি পরে send করে থাকে, সেটা অব্যবহৃত থেকে যায় — unbuffered channel-এ এটা block করতে পারে। ছোট example-এ এটা fine।)

### Line 26

```go
func main() {
```

Program-এর entry point।

### Lines 27–28

```go
fast := pollWithTimeout(50*time.Millisecond, 200*time.Millisecond)
fmt.Println("fast payment result:", fast)
```

`delay = 50ms`, `timeout = 200ms` — payment-টা **দ্রুত**: 50ms-এ `"confirmed"` আসে, যা 200ms timeout-এর **আগেই**। তাই `fast`-এর value `"confirmed"`, output: `fast payment result: confirmed`।

### Lines 30–31

```go
slow := pollWithTimeout(500*time.Millisecond, 200*time.Millisecond)
fmt.Println("slow payment result:", slow)
```

`delay = 500ms`, `timeout = 200ms` — payment-টা **ধীর**: 500ms লাগবে, কিন্তু timeout মাত্র 200ms। 200ms-এ timeout-এর case-টা আগে ready হয় → `slow`-এর value `"time-out"`, output: `slow payment result: time-out`।

### Line 32

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
fast payment result: confirmed
slow payment result: time-out
```

## মূল শিক্ষা / Key Takeaways

1. **`select`** — একাধিক channel-এর মধ্যে অপেক্ষা; যেটা আগে ready হয় সেটা চলে।
2. **`time.After(timeout)`** — timeout-এর জন্য timer channel; এটা fixed time-পর value দেয়।
3. **Timeout pattern** — দেরি-নিল external call-কে `select` দিয়ে সীমিত করা যায় — চিরকাল ঝুলে না থেকে।
4. **Goroutine + channel** — background-এ কাজ চালিয়ে result channel-এ পাঠানো।
5. **Send-only channel (`chan<-`)** — helper-এ write-only contract।

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
    "fmt"
    "time"
)
```

Two packages are imported:

- `fmt` — for printing output to the console.
- `time` — for `time.Duration`, `time.Sleep`, and `time.After` (managing delays and timeouts).

### Lines 8–11

```go
func checkPaymentStatus(statusCh chan<- string, delay time.Duration) {
    time.Sleep(delay)
    statusCh <- "confirmed"
}
```

A helper named `checkPaymentStatus`:

- The parameter `chan<- string` is a **send-only channel** — the function can only write to it, never read from it.
- `delay time.Duration` — how long to sleep (simulating a delayed response from a payment gateway).
- `time.Sleep(delay)` — waits for the delay, then `statusCh <- "confirmed"` sends `"confirmed"` into the channel.

In other words, the payment takes `delay` to get "confirmed".

### Lines 13–24

```go
func pollWithTimeout(delay, timeout time.Duration) string {
    statusCh := make(chan string)

    go checkPaymentStatus(statusCh, delay)

    select {
    case status := <-statusCh:
        return status
    case <-time.After(timeout):
        return "time-out"
    }
}
```

`pollWithTimeout` — the core function:

- `make(chan string)` — an **unbuffered channel** (capacity 0).
- `go checkPaymentStatus(statusCh, delay)` — launches the status checker as a goroutine; it will answer after `delay`.
- **`select`** — Go's powerful tool: it waits on **multiple channel operations** at once and **runs whichever is ready first**:
    - `case status := <-statusCh:` — if a value arrives on the status channel, bring it into `status` and `return status` (`"confirmed"`).
    - `case <-time.After(timeout):` — `time.After(timeout)` returns a channel that fires once after `timeout`. If that is ready first (meaning the payment didn't arrive within `timeout`), return `"time-out"`.

So it's a **race** between the payment confirmation and the timeout — whichever comes first, `select` takes it. This is the idiomatic Go way to **poll with a timeout** — it avoids hanging forever.

> **Note:** There's no way for `select` to fail in an unwinnable case here; with this simple design each call always resolves one of the two cases, so there's no blocking leak. (If the goroutine sends after a timeout, that send stays unused — on an unbuffered channel it could block; fine for this small example.)

### Line 26

```go
func main() {
```

Program entry point.

### Lines 27–28

```go
fast := pollWithTimeout(50*time.Millisecond, 200*time.Millisecond)
fmt.Println("fast payment result:", fast)
```

`delay = 50ms`, `timeout = 200ms` — the payment is **fast**: confirmation arrives in 50ms, **before** the 200ms timeout. So `fast` becomes `"confirmed"`: output `fast payment result: confirmed`.

### Lines 30–31

```go
slow := pollWithTimeout(500*time.Millisecond, 200*time.Millisecond)
fmt.Println("slow payment result:", slow)
```

`delay = 500ms`, `timeout = 200ms` — the payment is **slow**: it would take 500ms, but the timeout is only 200ms. At 200ms the timeout case is ready first → `slow` becomes `"time-out"`: output `slow payment result: time-out`.

### Line 32

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
fast payment result: confirmed
slow payment result: time-out
```

## Key Takeaways

1. **`select`** — waits on multiple channels; the one ready first runs.
2. **`time.After(timeout)`** — a timer channel that fires once after a fixed duration.
3. **Timeout pattern** — `select` caps how long you'll wait on a slow external call — no hanging forever.
4. **Goroutine + channel** — run work in the background and send results through a channel.
5. **Send-only channel (`chan<-`)** — a write-only contract for helpers.
