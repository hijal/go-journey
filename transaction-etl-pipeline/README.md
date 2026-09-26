# transaction-etl-pipeline

Go-তে **channel-based ETL pipeline** — source → parse → filter, malformed data আলাদা error-channel-এ।

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
	"strconv"
	"strings"
)
```

`strings` (Cut) + `strconv` (Atoi)।

### Lines 9–12

```go
type tx struct {
	id     string
	amount int
}
```

কাজের unit।

### Lines 14–25

```go
func source(lines []string) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		for _, line := range lines {
			out <- line
		}
	}()

	return out
}
```

**Stage 1 — Source**: raw lines-কে channel-এ streams; শেষে `close` → downstream-এ end-of-data।

### Lines 27–54

```go
func parseStage(in <-chan string) (<-chan tx, <-chan error) {
	out := make(chan tx)
	errc := make(chan error, 16)

	go func() {
		defer close(out)
		defer close(errc)

		for line := range in {
			id, raw, ok := strings.Cut(line, ",")
			if !ok {
				errc <- fmt.Errorf("malformed line %q", line)
				continue
			}
			amount, err := strconv.Atoi(raw)
			if err != nil {
				errc <- fmt.Errorf("parse amount of %s: %w", id, err)
				continue
			}
			out <- tx{id: id, amount: amount}
		}
	}()
	return out, errc
}
```

**Stage 2 — Parse**: valid → `out`, invalid → `errc` (buffered error-channel)। Error-ও **data-flow-এর অংশ** — skipped record-ও নষ্ট হয় না।

### Lines 56–69

```go
func filter(in <-chan tx, threshold int) <-chan tx {
	out := make(chan tx)

	go func() {
		defer close(out)
		for t := range in {
			if t.amount >= threshold {
				out <- t
			}
		}
	}()
	return out
}
```

**Stage 3 — Filter**: threshold-এর উপরে-টাই পাঠায় (`>= 500`)।

### Lines 81–98

```go
	txs, errc := parseStage(source(lines))
	highValue := filter(txs, 500)

	total, count := 0, 0
	for t := range highValue {
		count++
		total += t.amount
		fmt.Println("high-value:", t.id, t.amount)
	}

	for err := range errc {
		fmt.Println("skipped:", err)
	}

	fmt.Printf("%d high-value transactions, total %d BDT\n", count, total)
```

Pipeline-টা জুড়ে, error-গুলো পরে read। Input: 2500 ✓ / 180 ✗(below) / notanumber ✗(parse) / 9900 ✓ / (no comma) ✗ → **2 high-value, total 12400**।

---

## Expected Output

```
high-value: TX-1001 2500
high-value: TX-1004 9900
skipped: parse amount of TX-1003: strconv.Atoi: parsing "notanumber": invalid syntax
skipped: malformed line "brokenline"
2 high-value transactions, total 12400 BDT
```

## মূল শিক্ষা / Key Takeaways

1. **Pipeline of channels** — প্রত্যেক Stage = goroutine + channel।
2. **`close` cascade** — upstream close → downstream range শেষ।
3. **Error channel** — bad data কে মূল flow থেকে আলাদা।
4. **Buffered errc** — error-গুলো না-থেমে push করা যায়।
5. **Composability** — Stage যোগ/বদল করা সহজ।

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
	"strconv"
	"strings"
)
```

`strings` (Cut) + `strconv` (Atoi).

### Lines 9–12

```go
type tx struct {
	id     string
	amount int
}
```

The unit of work.

### Lines 14–25

```go
func source(lines []string) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		for _, line := range lines {
			out <- line
		}
	}()

	return out
}
```

**Stage 1 — Source**: streams raw lines into a channel; `close` at the end → end-of-data for the downstream.

### Lines 27–54

```go
func parseStage(in <-chan string) (<-chan tx, <-chan error) {
	out := make(chan tx)
	errc := make(chan error, 16)

	go func() {
		defer close(out)
		defer close(errc)

		for line := range in {
			id, raw, ok := strings.Cut(line, ",")
			if !ok {
				errc <- fmt.Errorf("malformed line %q", line)
				continue
			}
			amount, err := strconv.Atoi(raw)
			if err != nil {
				errc <- fmt.Errorf("parse amount of %s: %w", id, err)
				continue
			}
			out <- tx{id: id, amount: amount}
		}
	}()
	return out, errc
}
```

**Stage 2 — Parse**: valid → `out`, invalid → `errc` (buffered error-channel). Errors are part of the data flow — nothing is silently dropped.

### Lines 56–69

```go
func filter(in <-chan tx, threshold int) <-chan tx {
	out := make(chan tx)

	go func() {
		defer close(out)
		for t := range in {
			if t.amount >= threshold {
				out <- t
			}
		}
	}()
	return out
}
```

**Stage 3 — Filter**: passes only values at/above the threshold (`>= 500`).

### Lines 81–98

```go
	txs, errc := parseStage(source(lines))
	highValue := filter(txs, 500)

	total, count := 0, 0
	for t := range highValue {
		count++
		total += t.amount
		fmt.Println("high-value:", t.id, t.amount)
	}

	for err := range errc {
		fmt.Println("skipped:", err)
	}

	fmt.Printf("%d high-value transactions, total %d BDT\n", count, total)
```

The pipeline drains; errors are read afterwards. Input: 2500 ✓ / 180 ✗(below) / notanumber ✗(parse) / 9900 ✓ / (no comma) ✗ → **2 high-value, total 12400**.

---

## Expected Output

```
high-value: TX-1001 2500
high-value: TX-1004 9900
skipped: parse amount of TX-1003: strconv.Atoi: parsing "notanumber": invalid syntax
skipped: malformed line "brokenline"
2 high-value transactions, total 12400 BDT
```

## Key Takeaways

1. **Pipeline of channels** — each stage is a goroutine + channel.
2. **`close` cascade** — upstream close ends downstream ranges.
3. **Error channel** — bad data kept out of the main flow.
4. **Buffered errc** — errors can be pushed without blocking.
5. **Composability** — stages are easy to add or replace.