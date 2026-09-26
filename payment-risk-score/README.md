# payment-risk-score

Go-তে **goroutine + unbuffered channel দিয়ে overlapping work** — risk model parallel-এ চলে, এদিকে billing address validate হয়।

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

`time` (simulated latency)।

### Lines 8–11

```go
func riskScore(cardBIN string, out chan<- int) {
	time.Sleep(80 * time.Millisecond)
	out <- len(cardBIN) * 7
}
```

**Producer** — 80ms-এর risk-model চালিয়ে স্কোর (BIN দৈর্ঘ্য × 7) পাঠায়। `chan<-` — send-only parameter type।

### Lines 13–20

```go
func main() {
	scores := make(chan int)

	go riskScore("457173", scores)
	fmt.Println("validating billing address while the risk model runs")
	time.Sleep(30 * time.Millisecond)

	score := <-scores
	fmt.Println("risk score:", score)
```

**Overlap** — main-টা risk goroutine-এর সাথে সাথে billing validation করতে থাকে; 30ms পরে receive-এ wait করে। **ফলাফল: 80ms মোট, sequential হলে 110ms।**

### Lines 22–25

```go
	if score > 30 {
		fmt.Println("transaction flagged for manual review")
	}
```

`len("457173") * 7` = 6 × 7 = **42** > 30 → flagged।

---

## Expected Output

```
validating billing address while the risk model runs
risk score: 42
transaction flagged for manual review
```

## মূল শিক্ষা / Key Takeaways

1. **Unbuffered channel** — send-কে receiver জুটে না-পাওয়া অবধি block।
2. **Overlap** — independent কাজ (risk + billing) সময় ভাগ করতে পারে।
3. **`chan<-` send-only** — API-তে misdirection কমালো।
4. **Simple coordination** — `go` + channel, কোনো mutex নয়।
5. **Real similarity** — 80ms risk | 30ms validate ≈ মোট 80ms।

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

`time` (simulated latency).

### Lines 8–11

```go
func riskScore(cardBIN string, out chan<- int) {
	time.Sleep(80 * time.Millisecond)
	out <- len(cardBIN) * 7
}
```

**Producer** — runs a 80ms risk model and sends the score (BIN length × 7). `chan<-` — a send-only parameter type.

### Lines 13–20

```go
func main() {
	scores := make(chan int)

	go riskScore("457173", scores)
	fmt.Println("validating billing address while the risk model runs")
	time.Sleep(30 * time.Millisecond)

	score := <-scores
	fmt.Println("risk score:", score)
```

**Overlap** — main keeps validating the billing address while the risk goroutine works; it only blocks on the receive afterwards (another 50ms). **Net: 80ms total, ~110ms if sequential.**

### Lines 22–25

```go
	if score > 30 {
		fmt.Println("transaction flagged for manual review")
	}
```

`len("457173") * 7` = 6 × 7 = **42** > 30 → flagged.

---

## Expected Output

```
validating billing address while the risk model runs
risk score: 42
transaction flagged for manual review
```

## Key Takeaways

1. **Unbuffered channel** — a send blocks until a receiver pairs up.
2. **Overlap** — independent work (risk + billing) can share time.
3. **`chan<-` send-only** — less misdirection in the API.
4. **Simple coordination** — `go` + channel, no mutex.
5. **Real difference** — 80ms risk | 30ms validate ≈ 80ms total.