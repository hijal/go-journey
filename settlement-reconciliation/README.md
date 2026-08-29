# settlement-reconciliation

Go-তে **`goto` label**, **`errors.New`**, multi-dimensional slice (`[][]float64`), আর **`break scan`** (labeled break) দিয়ে retry + reconciliation শেখার ছোট example।

**📖 ভাষা নির্বাচন করুন / Choose language:**

[🇧🇩 বাংলা](#bangla) • [🇬🇧 English](#english)

---

<a name="bangla"></a>

## 🇧🇩 বাংলা সংস্করণ

### Lines 1–6

```go
package main

import (
    "errors"
    "fmt"
)
```

- `package main` — একটা executable program।
- `errors` — custom error তৈরি (`errors.New`) করার জন্য।
- `fmt` — output print করার জন্য।

### Lines 8–13

```go
func connectToBank(attempt int) error {
    if attempt < 3 {
        return errors.New("failed to connect to bank")
    }
    return nil
}
```

একটা function `connectToBank` যা `attempt` (int) নেয় এবং `error` return করে। Bank-এ connect করার simulation:

- `attempt < 3` হলে (1, 2 attempt) `errors.New("failed to connect to bank")` দিয়ে error return করে।
- `attempt >= 3` হলে `nil` (সফল) return করে।

অর্থাৎ প্রথম ২টা attempt fail হয়, ৩য় attempt থেকে success।

### Line 15

```go
func main() {
```

Program-এর entry point।

### Lines 16–17

```go
var err error
attempt := 1
```

`var err error` — একটা error variable (nil-এ initialize)। `attempt := 1` — retry counter, শুরুতে 1।

### Line 18

```go
retry:
```

একটা **label** `retry:` — এটা একটা mark, যেখানে `goto` দিয়ে ফিরে আসা যায়। (Not a statement, just a position marker।)

### Lines 19–25

```go
err = connectToBank(attempt)

if err != nil && attempt < 5 {
    fmt.Println("Attempt", attempt, "failed:", err)
    attempt++
    goto retry
}
```

- `connectToBank(attempt)` call করে error-টা `err`-এ রাখে।
- যদি error থাকে **আর** `attempt < 5` হয়, তাহলে failure message print করে, `attempt++` দিয়ে counter বাড়ায়, এবং **`goto retry`** দিয়ে আবার label-এ ফিরে যায় (retry করে)।
- 3য় attempt থেকে সফল হয়, তাই লুপটা 1, 2 fail করে, 3-এ connect হয়।

### Lines 27–30

```go
if err != nil {
    fmt.Println("giving up after", attempt, "attempts")
    return
}
```

Loop শেষের পরেও যদি `err != nil` থেকে যায় (অর্থাৎ ৫বারও fail), তাহলে "giving up" print করে `return`।

### Line 32

```go
fmt.Println("Connected on attempt", attempt)
```

Connect সফল হলে এই message print হয়।

### Lines 34–38

```go
batches := [][]float64{
    {500, 300, 200},
    {700, -1, 100},
    {900, 400},
}
```

`[][]float64` — একটা multi-dimensional (2D) slice, অর্থাৎ float64-এর slice-এর slice। তিনটা batch (sorting/group) আছে, প্রতিটায় কয়েকটা entry। দ্বিতীয় batch-এ `-1` একটা mismatched (অবৈধ) entry।

### Line 39

```go
scan:
```

আরেকটা **label** `scan:` — বাইরের `for` loop-টার আগে।

### Lines 40–48

```go
for batchIndex, batch := range batches {
    for _, entry := range batch {
        if entry < 0 {
            fmt.Println("mismatched entry in batch", batchIndex, ":", entry)
            break scan
        }
        fmt.Println("verified entry:", entry)
    }
}
```

দুই স্তরের nesting:

- বাইরের loop প্রতিটা `batch`-এ যায় (`batchIndex` = batch এর index, `batch` = সেই batch-এর slice)।
- ভেতরের loop সেই batch-এর প্রতিটা `entry`-তে যায়।

- যদি `entry < 0` হয় (অবৈধ/mismatch), message print করে **`break scan`** — `scan` label-এ চিহ্নিত **বাইরের** loop-কেও ছেড়ে বেরিয়ে যায় (inner + outer দুটোই stop)।
- নাহলে "verified entry" print করে।

> **Labeled break:** সাধারণ `break` শুধু ভেতরের loop-টা ভাঙে। কিন্তু `break scan` label-এর সাথে outer loop-ও ভাঙে — nested loop থেকে সম্পূর্ণ বেরিয়ে আসতে চাইলে এটা দরকারি।

---

## Expected Output

```
Attempt 1 failed: failed to connect to bank
Attempt 2 failed: failed to connect to bank
Connected on attempt 3
verified entry: 500
verified entry: 300
verified entry: 200
verified entry: 700
mismatched entry in batch 1 : -1
```

## মূল শিক্ষা / Key Takeaways

1. **`errors.New`** — plain error message তৈরি করে।
2. **`goto` + label** — নির্দিষ্ট জায়গায় ফিরে গিয়ে retry logic।
3. **`[][]float64`** — 2D/multi-dimensional slice।
4. **`break scan` (labeled break)** — nested loop থেকে outer loop-ও বেরিয়ে আসা।
5. **Retry pattern** — fail হলে counter বাড়িয়ে আবার চেষ্টা, নির্দিষ্ট limit-এ give up।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Lines 1–6

```go
package main

import (
    "errors"
    "fmt"
)
```

- `package main` — an executable program.
- `errors` — for creating custom errors with `errors.New`.
- `fmt` — for console output.

### Lines 8–13

```go
func connectToBank(attempt int) error {
    if attempt < 3 {
        return errors.New("failed to connect to bank")
    }
    return nil
}
```

A function `connectToBank` that takes an `attempt` (int) and returns an `error`. Simulates connecting to a bank:

- If `attempt < 3` (attempts 1, 2), returns an error with `errors.New("failed to connect to bank")`.
- If `attempt >= 3`, returns `nil` (success).

So the first two attempts fail, and success begins from the 3rd attempt.

### Line 15

```go
func main() {
```

Program entry point.

### Lines 16–17

```go
var err error
attempt := 1
```

`var err error` — an error variable (initialized to nil). `attempt := 1` — a retry counter, starting at 1.

### Line 18

```go
retry:
```

A **label** `retry:` — a position marker that `goto` can jump back to. (Not a statement, just a marker.)

### Lines 19–25

```go
err = connectToBank(attempt)

if err != nil && attempt < 5 {
    fmt.Println("Attempt", attempt, "failed:", err)
    attempt++
    goto retry
}
```

- Calls `connectToBank(attempt)` and stores the error in `err`.
- If there's an error **and** `attempt < 5`, print the failure message, increment the counter with `attempt++`, then **`goto retry`** to go back to the label and retry.
- Success happens from attempt 3, so the loop fails at 1 and 2, then connects at 3.

### Lines 27–30

```go
if err != nil {
    fmt.Println("giving up after", attempt, "attempts")
    return
}
```

If `err` is still non-nil after the loop (i.e. failed 5 times), print "giving up" and `return`.

### Line 32

```go
fmt.Println("Connected on attempt", attempt)
```

Prints this message when the connection succeeds.

### Lines 34–38

```go
batches := [][]float64{
    {500, 300, 200},
    {700, -1, 100},
    {900, 400},
}
```

`[][]float64` — a multi-dimensional (2D) slice — a slice of slices of float64. There are three batches (groups), each with a few entries. The second batch contains `-1`, a mismatched (invalid) entry.

### Line 39

```go
scan:
```

Another **label** `scan:` — placed before the outer `for` loop.

### Lines 40–48

```go
for batchIndex, batch := range batches {
    for _, entry := range batch {
        if entry < 0 {
            fmt.Println("mismatched entry in batch", batchIndex, ":", entry)
            break scan
        }
        fmt.Println("verified entry:", entry)
    }
}
```

Two levels of nesting:

- The outer loop goes through each `batch` (`batchIndex` = the batch's index, `batch` = that batch's slice).
- The inner loop goes through each `entry` of that batch.

- If `entry < 0` (invalid/mismatch), print a message and **`break scan`** — exits the **outer** loop marked by the `scan` label (stops both inner and outer loops).
- Otherwise, print "verified entry".

> **Labeled break:** A plain `break` only breaks the inner loop. But `break scan` breaks the outer loop as well — useful when you need to fully exit from a nested loop.

---

## Expected Output

```
Attempt 1 failed: failed to connect to bank
Attempt 2 failed: failed to connect to bank
Connected on attempt 3
verified entry: 500
verified entry: 300
verified entry: 200
verified entry: 700
mismatched entry in batch 1 : -1
```

## Key Takeaways

1. **`errors.New`** — creates a plain error message.
2. **`goto` + label** — jumps back to a specific point to retry logic.
3. **`[][]float64`** — a 2D / multi-dimensional slice.
4. **`break scan` (labeled break)** — exits the outer loop too from within a nested loop.
5. **Retry pattern** — on failure, bump the counter and try again, give up after a limit.
