# stock-feed-merge

Go-তে **দুইটি channel merge + nil-channel idiom** — `select`-এ দুই feed, close হলে channel-টি `nil` করে সেই case নিষ্ক্রিয় করা।

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

`time` (Sleep)।

### Lines 8–12

```go
type quote struct {
	exchange string
	symbol   string
	price    float64
}
```

একটি struct — প্রতিটি feed-এর একটা মূল্য (DSE/CSE, symbol `GP`, price)।

### Lines 14–27

```go
func feed(exchange string, start time.Duration, prices []float64) <-chan quote {
	out := make(chan quote)

	go func() {
		defer close(out)
		time.Sleep(start)

		for _, p := range prices {
			out <- quote{exchange: exchange, symbol: "GP", price: p}
			time.Sleep(40 * time.Millisecond)
		}
	}()
	return out
}
```

- **`<-chan quote`** — receive-only ফেরত দেয়।
- `start` ইনজেকশন — দুই exchange আলাদা সময়ে শুরু (DSE 0ms, CSE 20ms)।
- **`time.Sleep(40ms)`** — feed-এর মধ্যে 40ms বিরতি, তাই দুই feed **সমান্তরাল** (interleaved) আসে।
- **`defer close(out)`** — সব price পাঠানোর পর channel বন্ধ, consumer বুঝবে আর কিছু আসবে না।

### Lines 29–55

```go
func main() {
	dse := feed("DSE", 0, []float64{310.5, 311.0, 309.8})
	cse := feed("CSE", 20*time.Millisecond, []float64{310.2, 310.9})

	best := 0.0

	for dse != nil || cse != nil {
		select {
		case q, ok := <-dse:
			if !ok {
				dse = nil
				fmt.Println("DSE feed closed")
				continue
			}
			fmt.Printf("%s %s %.1f\n", q.exchange, q.symbol, q.price)
			best = max(best, q.price)
		case q, ok := <-cse:
			if !ok {
				cse = nil
				fmt.Println("CSE feed closed")
				continue
			}
			fmt.Printf("%s %s %.1f\n", q.exchange, q.symbol, q.price)
			best = max(best, q.price)
		}
	}
	fmt.Printf("best price seen: %.1f\n", best)
}
```

- **`for dse != nil || cse != nil`** — দুটো feed-ই বন্ধ না হলে চলবে।
- **`select`** — যে feed আগে ready, সেটাই নেবে; দুটো 40ms/20ms বিরতিতে বাজে ধরে এখানে interleave হয়।
- **`dse = nil`** — এটাই মূল কৌশল: **`select`-এ nil channel কখনো ready হয় না**, তাই ওই case নিজে থেকে নিষ্ক্রিয় হয়ে যায় (block-for-ever)। `continue` দিয়ে শুধু ওই কাজটা বাদ দেওয়া হয়।
- **`max`** builtin — চলমান সর্বোচ্চ দাম।
- শেষে `best price seen: 311.0` (DSE-এর 311.0 সবচেয়ে বেশি)।

---

## Expected Output

```
DSE GP 310.5
CSE GP 310.2
DSE GP 311.0
CSE GP 310.9
DSE GP 309.8
CSE feed closed
DSE feed closed
best price seen: 311.0
```

## মূল শিক্ষা / Key Takeaways

1. **দুই channel, এক `select`** — merge/fan-in-এর সরল রূপ।
2. **nil channel = disabled case** — `dse = nil` দিয়ে সেই case নিষ্ক্রিয়।
3. **comma-ok + `continue`** — বন্ধ হওয়া feed হ্যান্ডেল করা।
4. **Loop condition-এ nil check** — শেষ হওয়া টের পাওয়া।
5. **Interleaving** — 40ms/20ms বিরতির কারণেই output ধারাবাহিক।

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

`time` (Sleep).

### Lines 8–12

```go
type quote struct {
	exchange string
	symbol   string
	price    float64
}
```

A struct — one quote per feed (DSE/CSE, symbol `GP`, price).

### Lines 14–27

```go
func feed(exchange string, start time.Duration, prices []float64) <-chan quote {
	out := make(chan quote)

	go func() {
		defer close(out)
		time.Sleep(start)

		for _, p := range prices {
			out <- quote{exchange: exchange, symbol: "GP", price: p}
			time.Sleep(40 * time.Millisecond)
		}
	}()
	return out
}
```

- Returns **`<-chan quote`** (receive-only).
- `start` is injected so the two exchanges begin at different times (DSE 0ms, CSE 20ms).
- **`time.Sleep(40ms)`** — a gap between quotes, so the two feeds arrive **in parallel** (interleaved).
- **`defer close(out)`** — closes once all prices are sent, telling the consumer nothing more is coming.

### Lines 29–55

```go
func main() {
	dse := feed("DSE", 0, []float64{310.5, 311.0, 309.8})
	cse := feed("CSE", 20*time.Millisecond, []float64{310.2, 310.9})

	best := 0.0

	for dse != nil || cse != nil {
		select {
		case q, ok := <-dse:
			if !ok {
				dse = nil
				fmt.Println("DSE feed closed")
				continue
			}
			fmt.Printf("%s %s %.1f\n", q.exchange, q.symbol, q.price)
			best = max(best, q.price)
		case q, ok := <-cse:
			if !ok {
				cse = nil
				fmt.Println("CSE feed closed")
				continue
			}
			fmt.Printf("%s %s %.1f\n", q.exchange, q.symbol, q.price)
			best = max(best, q.price)
		}
	}
	fmt.Printf("best price seen: %.1f\n", best)
}
```

- **`for dse != nil || cse != nil`** — keep looping while either feed is open.
- The **`select`** takes whichever feed is ready first; the 40ms/20ms spacing makes the interleaving predictable here.
- **`dse = nil`** is the central trick: a **nil channel is never ready in a `select`**, so that case disables itself (blocks forever). The `continue` just skips the rest of that branch.
- **`max`** builtin — tracks the running best price.
- Finally `best price seen: 311.0` (DSE's 311.0 is the highest).

---

## Expected Output

```
DSE GP 310.5
CSE GP 310.2
DSE GP 311.0
CSE GP 310.9
DSE GP 309.8
CSE feed closed
DSE feed closed
best price seen: 311.0
```

## Key Takeaways

1. **Two channels, one `select`** — the simple form of merge/fan-in.
2. **A nil channel is a disabled case** — `dse = nil` switches that case off.
3. **comma-ok + `continue`** — handling a feed that has closed.
4. **nil in the loop condition** — detects that everything has finished.
5. **Interleaving** — the 40ms/20ms spacing is what makes the output readable.