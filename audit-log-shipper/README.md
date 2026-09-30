# audit-log-shipper

Go-তে **time/size-based batching** — `time.Ticker` + select, দুইভাবে flush (batch ভরলে, বা সময় পড়লে), শেষে `close` মানে final flush।

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
	"fmt"
	"strings"
	"time"
)
```

`strings` (Join), `time` (Ticker, Sleep)।

### Lines 9–12

```go
const (
	maxBatch      = 3
	flushInterval = 50 * time.Millisecond
)
```

**দুইটি flush trigger** — batch-এ ৩টা event হলে, অথবা প্রতি ৫০ms পরপর।

### Lines 14–28

```go
func shipper(events <-chan string, done chan<- struct{}) {
	defer close(done)

	ticker := time.NewTicker(flushInterval)

	defer ticker.Stop()
	batch := make([]string, 0, maxBatch)

	flush := func(reason string) {
		if len(batch) == 0 {
			return
		}
		fmt.Printf("flush %-6s -> [%s]\n", reason, strings.Join(batch, " "))
		batch = batch[:0]
	}
```

- **`defer close(done)`** — যেকোনো ভাবে শেষ হোক, main আটকে থাকবে না।
- **`defer ticker.Stop()`** — goroutine ছাড়লে timer ফাঁকা পড়ে না (resource leak-এর একটি সাধারণ উৎস)।
- **`make([]string, 0, maxBatch)`** — pre-allocate, তাই append-এ realloc হয় না।
- **flush কেন্দ্রিক `reason`** — "size" / "timer" / "closed" কোনটা হিসেবে ফ্লাশ হলো লগে দেখায়।
- **`if len(batch) == 0 { return }`** — খালি batch ছাপে না।
- **`batch = batch[:0]`** — slice-কে **শূন্য করে পুনর্ব্যবহার** (নতুন slice allocate না করে)।

### Lines 30–45

```go
	for {
		select {
		case e, ok := <-events:
			if !ok {
				flush("closed")
				return
			}
			batch = append(batch, e)
			if len(batch) == maxBatch {
				flush("size")
			}
		case <-ticker.C:
			flush("timer")
		}
	}
}
```

- **`select`** — event আসলে নাও, আর না আসলে সময় পড়লে ফ্লাশ করো।
- **channel বন্ধ (`!ok`)** → **final flush** — বাকি সব event হারিয়ে যাবে না।
- **`batch[:0]` পুনর্ব্যবহার** — `append` করলে আবার একই array-তে লেখে, তাই কোনো append/copy pitfall নেই।

### Lines 47–62

```go
func main() {
	events := make(chan string)
	done := make(chan struct{})

	go shipper(events, done)

	for i := range 5 {
		events <- fmt.Sprintf("login#%d", i+1)
	}

	time.Sleep(80 * time.Millisecond)
	events <- "login#6"
	events <- "login#7"
	close(events)
	<-done
}
```

- প্রথম ৫টা event পাঠানো হয় → ৩টা "size" flush-এ যায়, বাকি ২টা পড়ে থাকে।
- **৮০ms ঘুমানোর কারণে** ৫০ms ticker অবশ্যই একবার চলে → বাকি ২টা "timer" flush-এ যায়।
- **`close(events)`** → "closed" flush → **`close(done)`** → `<-done` main-কে ছাড়ে।

---

## Expected Output

```
flush size   -> [login#1 login#2 login#3]
flush timer  -> [login#4 login#5]
flush closed -> [login#6 login#7]
```

## মূল শিক্ষা / Key Takeaways

1. **দুই trigger, এক জায়গায়** — size আর timer দুটোই `select`-এ।
2. **`close` মানেই final flush** — বাকি data নষ্ট হয় না।
3. **`ticker.Stop()`** — ছাড়লে timer leak।
4. **`batch[:0]`** — slice reuse, নতুন allocation নয়।
5. **`done` channel** — শেষ হওয়ার নিশ্চয়তা (WaitGroup-এর চেয়ে হালকা, শুধু একবারের signal-এর জন্য)।

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
	"fmt"
	"strings"
	"time"
)
```

`strings` (Join), `time` (Ticker, Sleep).

### Lines 9–12

```go
const (
	maxBatch      = 3
	flushInterval = 50 * time.Millisecond
)
```

**Two flush triggers** — once the batch holds 3 events, or every 50ms, whichever comes first.

### Lines 14–28

```go
func shipper(events <-chan string, done chan<- struct{}) {
	defer close(done)

	ticker := time.NewTicker(flushInterval)

	defer ticker.Stop()
	batch := make([]string, 0, maxBatch)

	flush := func(reason string) {
		if len(batch) == 0 {
			return
		}
		fmt.Printf("flush %-6s -> [%s]\n", reason, strings.Join(batch, " "))
		batch = batch[:0]
	}
```

- **`defer close(done)`** — however it ends, main is never left hanging.
- **`defer ticker.Stop()`** — without it the timer keeps running after the goroutine exits (a common resource leak).
- **`make([]string, 0, maxBatch)`** — pre-allocated, so append never reallocates.
- **A flush closure taking `reason`** — the log shows whether a batch went out because of "size", "timer", or "closed".
- **`if len(batch) == 0 { return }`** — never print an empty batch.
- **`batch = batch[:0]`** — reuse the same backing array instead of allocating a new slice.

### Lines 30–45

```go
	for {
		select {
		case e, ok := <-events:
			if !ok {
				flush("closed")
				return
			}
			batch = append(batch, e)
			if len(batch) == maxBatch {
				flush("size")
			}
		case <-ticker.C:
			flush("timer")
		}
	}
}
```

- The **`select`** — take an event if one arrives, otherwise flush when the timer fires.
- **Channel closed (`!ok`)** → **final flush**, so nothing pending is lost.
- **`batch[:0]` reuse** — a later `append` writes into the same array, so there is no append/copy pitfall here.

### Lines 47–62

```go
func main() {
	events := make(chan string)
	done := make(chan struct{})

	go shipper(events, done)

	for i := range 5 {
		events <- fmt.Sprintf("login#%d", i+1)
	}

	time.Sleep(80 * time.Millisecond)
	events <- "login#6"
	events <- "login#7"
	close(events)
	<-done
}
```

- The first 5 events go out in a "size" flush of 3, leaving 2 pending.
- **The 80ms sleep guarantees** the 50ms ticker fires, so those 2 leave in a "timer" flush.
- **`close(events)`** → "closed" flush → **`close(done)`** → `<-done` releases main.

---

## Expected Output

```
flush size   -> [login#1 login#2 login#3]
flush timer  -> [login#4 login#5]
flush closed -> [login#6 login#7]
```

## Key Takeaways

1. **Two triggers in one place** — both size and timer live in the same `select`.
2. **`close` means flush the rest** — pending data is never dropped.
3. **`ticker.Stop()`** — skipping it leaks the timer.
4. **`batch[:0]`** — slice reuse with no new allocation.
5. **A `done` channel** — lighter than a `sync.WaitGroup` when all you need is a one-shot completion signal.