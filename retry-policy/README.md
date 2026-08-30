# retry-policy

Go-তে **retry loop**, **`time.Sleep` / backoff**, **`const`**, আর `fmt.Printf` দিয়ে retry policy শেখার ছোট example।

**📖 ভাষা নির্বাচন করুন / Choose language:**

[🇧🇩 বাংলা](#bangla) • [🇬🇧 English](#english)

---

<a name="bangla"></a>

## 🇧🇩 বাংলা সংস্করণ

### Lines 1–6

```go
package main

import (
    "fmt"
    "time"
)
```

- `package main` — একটা executable program।
- `fmt` — output print করার জন্য।
- `time` — delay (`time.Sleep`) আর duration type (`time.Duration`, `time.Millisecond`) এর জন্য।

### Lines 8–9

```go
const maxRetries = 3
const baseBackoff = 500 * time.Millisecond
```

দুটো **constant**:

- `maxRetries = 3` — সর্বোচ্চ কতবার retry করা যাবে।
- `baseBackoff = 500 * time.Millisecond` — retry-এর মধ্যে অপেক্ষার **base** সময় 500 millisecond। অনুবর্তী retry-তে এটা গুণ হয়ে বাড়বে (exponential backoff-এর ধারণা)।

### Line 11

```go
func main() {
```

Program-এর entry point।

### Line 12

```go
attempt := 0
```

`attempt` counter, শুরুতে 0। (এটা এখানে হয়েছে কতবার চেষ্টা করেছে সে count।)

### Line 14

```go
for attempt <= maxRetries {
```

`for` loop চলে যতক্ষণ `attempt <= maxRetries` (0,1,2,3 — মোট ৪টা iteration)। Loop condition-টা `<=` হওয়ায় প্রথমে-end attempt-টাও (0-indexed) include হয়।

### Line 15

```go
ok := doRequest(attempt)
```

`doRequest(attempt)` call করে result-টা `ok` (bool)-এ রাখে। ফাংশনটা simulate করে আমাদের request সফল হয়েছে কি না।

### Lines 17–20

```go
if ok {
    fmt.Printf("success after %d attempt(s)\n", attempt+1)
    return
}
```

যদি `ok` true হয়:

- `fmt.Printf("success after %d attempt(s)\n", attempt+1)` — `%d` placeholder-এ `attempt+1` বসে (1-আধারিত attempt সংখ্যা print করতে)। 
- `return` — মিলেই function থেকে বেরিয়ে যায়।

### Lines 21–22

```go
attempt++
time.Sleep(baseBackoff * time.Duration(attempt))
```

সফল না হলে:

- `attempt++` — counter বাড়ায়।
- `time.Sleep(...)` — কিছু সময় অপেক্ষা করে পরের attempt-এর আগে। `baseBackoff * time.Duration(attempt)` — প্রথম retry-তে 500ms, দ্বিতীয়তে 1000ms, ইত্যাদি — মানে প্রতিবার অপেক্ষা বাড়ে (backoff)। `time.Duration(attempt)` দিয়ে int-কে Duration-এ convert করা হয় (কারণ multiplication-এর জন্য দুই operand-এর type match হোক)।

### Line 23

```go
}
```

Closing brace — `for` loop শেষ। (সব attempt শেষেও সফল না হলে এখানে আসে।)

### Line 25

```go
fmt.Println("giving up after", maxRetries, "retries")
```

সব retry ব্যর্থ হলে এই message print হয়।

### Lines 28–31

```go
func doRequest(attempt int) bool {
    fmt.Println("attempt", attempt+1)
    return attempt >= 2
}
```

`doRequest` function যা একটা attempt নেয় এবং bool return করে:

- `fmt.Println("attempt", attempt+1)` — attempt number print করে (1-আধারিত)।
- `return attempt >= 2` — 3য় attempt (index 2) থেকে `true` return করে। অর্থাৎ আপাতত: attempt index 0,1 (attempt #1, #2) fail; attempt index 2 (attempt #3) থেকে success।

**Flow (main-এ):**

1. attempt=0 → doRequest(0) → prints "attempt 1", fails (0 >= 2 false) → attempt=1, sleep 500ms
2. attempt=1 → doRequest(1) → prints "attempt 2", fails (1 >= 2 false) → attempt=2, sleep 1000ms
3. attempt=2 → doRequest(2) → prints "attempt 3", succeeds (2 >= 2 true) → prints "success after 3 attempt(s)", returns

---

## Expected Output

```
attempt 1
attempt 2
attempt 3
success after 3 attempt(s)
```

## মূল শিক্ষা / Key Takeaways

1. **Retry loop** — `for` দিয়ে নির্দিষ্ট limit পর্যন্ত পুনরায় চেষ্টা।
2. **`time.Sleep`** — retry-এর মধ্যে delay।
3. **Backoff** — প্রতিবার অপেক্ষা বাড়ানো (`baseBackoff * attempt`), যেন repeatedly fail-এ সার্ভার চাপ না পড়ে।
4. **`time.Duration`** — সময়-ব্যবধান type; int-কে Duration-এ convert করতে `time.Duration(x)`।
5. **`fmt.Printf` / `%d`** — পূর্ণসংখ্যা placeholder-সহ format output।
6. **Early `return`** — সফল হলে সাথে-সাথে বেরিয়ে যাওয়া।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Lines 1–6

```go
package main

import (
    "fmt"
    "time"
)
```

- `package main` — an executable program.
- `fmt` — for console output.
- `time` — for delays (`time.Sleep`) and duration types (`time.Duration`, `time.Millisecond`).

### Lines 8–9

```go
const maxRetries = 3
const baseBackoff = 500 * time.Millisecond
```

Two **constants**:

- `maxRetries = 3` — the maximum number of retries allowed.
- `baseBackoff = 500 * time.Millisecond` — the **base** wait time of 500 ms between retries. In successive retries it's multiplied to grow (the idea of exponential backoff).

### Line 11

```go
func main() {
```

Program entry point.

### Line 12

```go
attempt := 0
```

An `attempt` counter, starting at 0 (tracks how many times it has tried).

### Line 14

```go
for attempt <= maxRetries {
```

The `for` loop runs while `attempt <= maxRetries` (0,1,2,3 — 4 iterations total). Because the condition is `<=`, the first attempt (0-indexed) is included.

### Line 15

```go
ok := doRequest(attempt)
```

Calls `doRequest(attempt)` and stores the result in `ok` (bool). This function simulates whether our request succeeded.

### Lines 17–20

```go
if ok {
    fmt.Printf("success after %d attempt(s)\n", attempt+1)
    return
}
```

If `ok` is true:

- `fmt.Printf("success after %d attempt(s)\n", attempt+1)` — the `%d` placeholder is replaced by `attempt+1` (to print a 1-based attempt count).
- `return` — immediately exits the function.

### Lines 21–22

```go
attempt++
time.Sleep(baseBackoff * time.Duration(attempt))
```

If not successful:

- `attempt++` — increments the counter.
- `time.Sleep(...)` — waits before the next attempt. `baseBackoff * time.Duration(attempt)` — 500ms on the first retry, 1000ms on the second, etc. — i.e. the wait grows each time (backoff). `time.Duration(attempt)` converts the int to a Duration so the multiplication operands match types.

### Line 23

```go
}
```

Closing brace — end of the `for` loop. (We reach here if all attempts failed.)

### Line 25

```go
fmt.Println("giving up after", maxRetries, "retries")
```

Prints this message if all retries failed.

### Lines 28–31

```go
func doRequest(attempt int) bool {
    fmt.Println("attempt", attempt+1)
    return attempt >= 2
}
```

`doRequest` takes an attempt and returns a bool:

- `fmt.Println("attempt", attempt+1)` — prints the attempt number (1-based).
- `return attempt >= 2` — returns `true` from the 3rd attempt (index 2) onward. So attempts 0,1 (#1, #2) fail; index 2 (#3) succeeds.

**Flow (in main):**

1. attempt=0 → doRequest(0) → prints "attempt 1", fails (0 ≥ 2 false) → attempt=1, sleep 500ms
2. attempt=1 → doRequest(1) → prints "attempt 2", fails (1 ≥ 2 false) → attempt=2, sleep 1000ms
3. attempt=2 → doRequest(2) → prints "attempt 3", succeeds (2 ≥ 2 true) → prints "success after 3 attempt(s)", returns

---

## Expected Output

```
attempt 1
attempt 2
attempt 3
success after 3 attempt(s)
```

## Key Takeaways

1. **Retry loop** — retrying with `for` up to a limit.
2. **`time.Sleep`** — a delay between retries.
3. **Backoff** — growing the wait each time (`baseBackoff * attempt`) so the server isn't hammered on repeated failures.
4. **`time.Duration`** — a time-interval type; use `time.Duration(x)` to convert an int.
5. **`fmt.Printf` / `%d`** — formatted output with an integer placeholder.
6. **Early `return`** — exiting immediately on success.
