# concurrent-webhook

Go-তে **goroutine**, **channel**, আর **`sync.WaitGroup`** দিয়ে একসাথে (concurrently) webhook পাঠানো শেখার ছোট example।

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
    "sync"
)
```

দুটো package import করা হয়:

- `fmt` — console-এ output print করার জন্য।
- `sync` — **`sync.WaitGroup`** ব্যবহারের জন্য। WaitGroup দিয়ে আমরা main-কে goroutine-গুলো শেষ হওয়া পর্যন্ত অপেক্ষা করাতে পারি।

### Lines 8–11

```go
func sendWebhook(url string, results chan<- string, wg *sync.WaitGroup) {
    defer wg.Done()
    results <- fmt.Sprintf("delivered to %s", url)
}
```

`sendWebhook` নামক function যা একটা merchant URL নেয় এবং সেই URL-এ webhook "deliver" করার simulation করে:

- `url string` — যেখানে webhook পাঠানো হবে।
- `results chan<- string` — একটা **send-only channel** (`chan<-`)। মজার বিষয়: function parameter-এ `chan<- string` লিখলে channel থেকে accept-ও করা যায়, কিন্তু এখানে বোঝানো হচ্ছে এই function **শুধু পাঠাবে** (write-only)। এতে type system নিজেই নিশ্চিত করে যে function channel-এ শুধু data দেয়, read করে না।
- `wg *sync.WaitGroup` — WaitGroup-এর pointer, যাতে `Done()` call করা যায়। Pointer নিতে হয় কারণ `Done()` মূল WaitGroup-এর internal counter-টা পরিবর্তন করে; value দিলে copy-তে পরিবর্তন হতো, যা কাজ করত না।

- `defer wg.Done()` — function শেষ হলে (যেকোনোভাবে `return` করলেও) WaitGroup-এর counter 1 কমানো হবে।
- `results <- fmt.Sprintf(...)` — channel-এ একটা formatted message পাঠায়: `"delivered to <url>"`।

### Line 13

```go
func main() {
```

Program-এর entry point।

### Lines 14–18

```go
endpoints := []string{
    "https://merchant-a.example.com/webhook",
    "https://merchant-b.example.com/webhook",
    "https://merchant-c.example.com/webhook",
}
```

`endpoints` — একটা `[]string` slice যাতে তিনটা merchant-এর webhook URL আছে, যেগুলোতে আমরা webhook পাঠাতে চাই।

### Line 20

```go
results := make(chan string, len(endpoints))
```

`make(chan string, len(endpoints))` — একটা **buffered channel** বানায়। `len(endpoints)` মানে capacity `3`। Buffered channel-এর সুবিধা: **receiver না থাকলেও** ৩টা value পর্যন্ত জমা রাখতে পারে — এতে goroutine-গুলো block করে বসে থাকবে না। যদি unbuffered হতো, তাহলে প্রতিটা send-এর জন্য receiver লাগত।

### Line 21

```go
var wg sync.WaitGroup
```

`wg` নামে একটা `sync.WaitGroup` declare করা হয়। ওটা একটা counter রাখে; আমরা বলি "৩টা goroutine চলছে", ওটা `Wait()` দিয়ে counter ০ হওয়া পর্যন্ত অপেক্ষা করে।

### Lines 23–26

```go
for _, endpoint := range endpoints {
    wg.Add(1)
    go sendWebhook(endpoint, results, &wg)
}
```

`endpoints`-এর প্রতিটা URL-এর উপর loop:

- `wg.Add(1)` — WaitGroup counter-এ 1 যোগ করে (আরেকটা কাজ শুরু হচ্ছে)।
- `go sendWebhook(endpoint, results, &wg)` — **`go` keyword** দিয়ে `sendWebhook` টা **goroutine** হিসেবে চালায়। অর্থাৎ এটা আলাদা concurrently চলমান task-এ চলে — main এখানে থেমে থাকে না। তিনটা URL → তিনটা goroutine একসাথে চলে। `&wg` pointer হিসেবে pass করা হয়।

### Lines 28–29

```go
wg.Wait()
close(results)
```

- `wg.Wait()` — সব goroutine শেষ (০ counter) হওয়া পর্যন্ত main **block** করে থাকে। এটা না করলে main আগেই শেষ হয়ে যেত এবং goroutine-গুলো মাঝপথে অদৃশ্য হয়ে যেত।
- `close(results)` — channel-এ আর কোনো ডেটা আসবে না ঘোষণা করে। এটা দরকারি কারণ পরে আমরা `range` দিয়ে channel-টা iterate করব — channel বন্ধ না করলে `range` চিরকাল থাকত।

### Lines 31–33

```go
for message := range results {
    fmt.Println(message)
}
```

`for ... range` দিয়ে channel-এর সব পাঠানো value-টা পড়ে (buffered-এ জমা থাকা ৩টা message) এবং print করে। Channel close হয়ে যাওয়ায় loop এমনিতেই শেষ হয়।

> **নোট:** goroutine-গুলো যেহেতু একসাথে চলে, message-গুলোর **order নিশ্চিত নয়** — কোন merchant আগে `"delivered"` হবে তা নির্ভর করে scheduling-এর উপর। সাধারণত উপরে-থেকে-নিচে order আসে, কিন্তু guarantee নয়।

### Line 34

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
delivered to https://merchant-a.example.com/webhook
delivered to https://merchant-b.example.com/webhook
delivered to https://merchant-c.example.com/webhook
```

## মূল শিক্ষা / Key Takeaways

1. **Goroutine (`go fn()`)** — function-কে একসাথে চালানো; main block করে না।
2. **Channel (`chan`)** — goroutine-এর মধ্যে data পাঠানোর পথ; `<-` দিয়ে write, `range` দিয়ে read+close।
3. **`sync.WaitGroup`** — goroutine-গুলো শেষ হওয়া পর্যন্ত `Wait()`; প্রতিটায় `Add(1)` আর শেষে `Done()`।
4. **Send-only channel (`chan<-`)** — function-কে শুধু write করতে বাধ্য করার compile-time contract।
5. **Buffered channel** — capacity-র মধ্যে receiver ছাড়াই value জমা রাখা যায়।
6. **`close()`** — channel-এ আর data নেই বলে ঘোষণা; `range` শেষ করা।

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
    "sync"
)
```

Two packages are imported:

- `fmt` — for printing output to the console.
- `sync` — for **`sync.WaitGroup`**. A WaitGroup lets `main` wait until all goroutines finish.

### Lines 8–11

```go
func sendWebhook(url string, results chan<- string, wg *sync.WaitGroup) {
    defer wg.Done()
    results <- fmt.Sprintf("delivered to %s", url)
}
```

A function `sendWebhook` that takes a merchant URL and simulates "delivering" a webhook to it:

- `url string` — where the webhook should be sent.
- `results chan<- string` — a **send-only channel** (`chan<-`). Fun fact: a function can still accept a bidirectional channel, but declaring `chan<- string` signals that this function will **only write** to the channel. The type system guarantees it never reads.
- `wg *sync.WaitGroup` — a pointer to a WaitGroup, so we can call `Done()`. It must be a pointer because `Done()` mutates the WaitGroup's internal counter; passing by value would modify a copy and do nothing.

- `defer wg.Done()` — when the function finishes (no matter how), decrement the WaitGroup counter by 1.
- `results <- fmt.Sprintf(...)` — sends a formatted message into the channel: `"delivered to <url>"`.

### Line 13

```go
func main() {
```

Program entry point.

### Lines 14–18

```go
endpoints := []string{
    "https://merchant-a.example.com/webhook",
    "https://merchant-b.example.com/webhook",
    "https://merchant-c.example.com/webhook",
}
```

`endpoints` — a `[]string` slice holding three merchant webhook URLs we want to deliver to.

### Line 20

```go
results := make(chan string, len(endpoints))
```

`make(chan string, len(endpoints))` creates a **buffered channel**. The capacity (`len(endpoints)`) is `3`. The benefit of a buffered channel: it can hold up to 3 values even **without a receiver**, so the goroutines never block waiting to send. With an unbuffered channel, every send would need a ready receiver.

### Line 21

```go
var wg sync.WaitGroup
```

Declares a `sync.WaitGroup` named `wg`. It keeps a counter; we say "3 goroutines are running", and `Wait()` blocks until the counter reaches 0.

### Lines 23–26

```go
for _, endpoint := range endpoints {
    wg.Add(1)
    go sendWebhook(endpoint, results, &wg)
}
```

Loop over each URL in `endpoints`:

- `wg.Add(1)` — increments the WaitGroup counter (one more task starting).
- `go sendWebhook(endpoint, results, &wg)` — the **`go` keyword** launches `sendWebhook` as a **goroutine**, i.e. it runs concurrently — `main` doesn't stop here. Three URLs → three goroutines running in parallel. `&wg` is passed as a pointer.

### Lines 28–29

```go
wg.Wait()
close(results)
```

- `wg.Wait()` — blocks `main` until all goroutines have finished (counter reaches 0). Without this, `main` would exit early and the goroutines would vanish mid-work.
- `close(results)` — announces that no more data will be sent on the channel. This matters because we'll iterate the channel with `range` next — without closing, `range` would wait forever.

### Lines 31–33

```go
for message := range results {
    fmt.Println(message)
}
```

`for ... range` drains all buffered values from the channel (the 3 messages) and prints each. Because the channel is closed, the loop ends naturally.

> **Note:** Since the goroutines run concurrently, the **ordering of messages is not guaranteed** — which merchant gets `"delivered"` first depends on the scheduler. It typically appears top-to-bottom, but it's not guaranteed.

### Line 34

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
delivered to https://merchant-a.example.com/webhook
delivered to https://merchant-b.example.com/webhook
delivered to https://merchant-c.example.com/webhook
```

## Key Takeaways

1. **Goroutine (`go fn()`)** — run a function concurrently; `main` doesn't block.
2. **Channel (`chan`)** — a path for data between goroutines; write with `<-`, read+iterate with `range`, end with `close`.
3. **`sync.WaitGroup`** — `Wait()` until goroutines finish; `Add(1)` per task and `Done()` at the end.
4. **Send-only channel (`chan<-`)** — compile-time contract that forces a function to only write.
5. **Buffered channel** — can hold values without a ready receiver, up to its capacity.
6. **`close()`** — declares no more data is coming; lets `range` terminate.
