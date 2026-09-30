# order-csv-pipeline

Go-তে **দুই-ধাপের channel pipeline** — `parse` → `addVAT` → consumer, error-as-data + `continue` (কোনো ধাপ থামে না)।

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

`strings` (Cut), `strconv` (Atoi)।

### Lines 9–14

```go
type order struct {
	id    string
	net   int
	total int
	err   error
}
```

**একটি struct-এ সব** — value ও error একসাথে বহন করে, তাই কোনো ধাপে error হলেও সেই item বাদ দেওয়া যায় (নিজস্ব `err` field আছে)।

### Lines 16–37

```go
func parse(lines []string) <-chan order {
	out := make(chan order)

	go func() {
		defer close(out)
		for _, line := range lines {
			id, amt, found := strings.Cut(line, ",")

			if !found {
				out <- order{id: line, err: fmt.Errorf("parse %q: missing comma", line)}
				continue
			}
			n, err := strconv.Atoi(amt)
			if err != nil {
				out <- order{id: id, err: fmt.Errorf("parse %q: %w", line, err)}
				continue
			}
			out <- order{id: id, net: n}
		}
	}()
	return out
}
```

- **ধাপ ১ (source)** — প্রতিটি লাইন parse করে channel-এ পাঠায়।
- `strings.Cut` — comma-এর দুই অংশ (`found` false মানে comma নেই)।
- **`%w`** — `Atoi`-এর original error ধরে রাখে (`strconv.Atoi: parsing "abc": invalid syntax`)।
- **`continue`** — ভুল হলেও পরের লাইনে যায়, pipeline থামে না।
- **`defer close(out)`** — শেষে channel বন্ধ → পরের ধাপের `range` শেষ হয়।
- Unbuffered channel → স্বাভাবিক backpressure।

### Lines 39–54

```go
func addVAT(in <-chan order) <-chan order {
	out := make(chan order)

	go func() {
		defer close(out)

		for o := range in {
			if o.err == nil {
				o.total = o.net + o.net*15/100
			}
			out <- o
		}
	}()

	return out
}
```

- **ধাপ ২ (transform)** — input range করে VAT যোগ করে আরেক channel-এ পাঠায়।
- **`if o.err == nil`** — error-যুক্ত item-এ VAT লাগে না, error-সহ সরাসরি পাঠানো হয় (forwarding)।
- **`<-chan order` → `<-chan order`** — direction দিয়ে API সীমিত।

### Lines 56–73

```go
func main() {
	lines := []string{
		"ORD-1,1000",
		"ORD-2,abc",
		"ORD-3,2400",
		"broken-line",
	}

	var revenue int
	for o := range addVAT(parse(lines)) {
		if o.err != nil {
			fmt.Println("skip:", o.err)
			continue
		}
		fmt.Printf("%s net=%d total=%d\n", o.id, o.net, o.total)
		revenue += o.total
	}
	fmt.Println("revenue:", revenue)
}
```

- **`addVAT(parse(lines))`** — দুই ধাপ এক বাক্যে জোড়া লাগানো (composition)।
- শুধু error-বিহীন item revenue-তে যোগ হয়: 1150 + 2760 = **3910**।

---

## Expected Output

```
ORD-1 net=1000 total=1150
skip: parse "ORD-2,abc": strconv.Atoi: parsing "abc": invalid syntax
ORD-3 net=2400 total=2760
skip: parse "broken-line": missing comma
revenue: 3910
```

## মূল শিক্ষা / Key Takeaways

1. **Channel = pipeline** — প্রতিটি ধাপ একটা ধাপের ইনপুট-আউটপুট।
2. **`<-chan` return** — উৎস শুধু পড়ে/লেখে, composition সহজ।
3. **`defer close` দিয়ে chain শেষ** — `range` শেষ হয়।
4. **Error as data** — struct-এ `err` field; ভুল item ধাপ থামায় না।
5. **`continue`** — bad record skip, বাকিরা প্রক্রিয়া হয়।

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

`strings` (Cut), `strconv` (Atoi).

### Lines 9–14

```go
type order struct {
	id    string
	net   int
	total int
	err   error
}
```

**Everything in one struct** — value and error travel together, so an item that fails in one stage can simply be skipped (it carries its own `err` field).

### Lines 16–37

```go
func parse(lines []string) <-chan order {
	out := make(chan order)

	go func() {
		defer close(out)
		for _, line := range lines {
			id, amt, found := strings.Cut(line, ",")

			if !found {
				out <- order{id: line, err: fmt.Errorf("parse %q: missing comma", line)}
				continue
			}
			n, err := strconv.Atoi(amt)
			if err != nil {
				out <- order{id: id, err: fmt.Errorf("parse %q: %w", line, err)}
				continue
			}
			out <- order{id: id, net: n}
		}
	}()
	return out
}
```

- **Stage 1 (source)** — parses each line and sends it on the channel.
- `strings.Cut` — the two halves around the comma (`found` false means no comma).
- **`%w`** — keeps the original `Atoi` error (`strconv.Atoi: parsing "abc": invalid syntax`).
- **`continue`** — moves to the next line; the pipeline never stops.
- **`defer close(out)`** — closes at the end, so the next stage's `range` finishes.
- An unbuffered channel gives natural backpressure.

### Lines 39–54

```go
func addVAT(in <-chan order) <-chan order {
	out := make(chan order)

	go func() {
		defer close(out)

		for o := range in {
			if o.err == nil {
				o.total = o.net + o.net*15/100
			}
			out <- o
		}
	}()

	return out
}
```

- **Stage 2 (transform)** — ranges the input, adds VAT, sends onward.
- **`if o.err == nil`** — no VAT on errored items; they're forwarded as-is.
- **`<-chan order` → `<-chan order`** — directions restrict the API.

### Lines 56–73

```go
func main() {
	lines := []string{
		"ORD-1,1000",
		"ORD-2,abc",
		"ORD-3,2400",
		"broken-line",
	}

	var revenue int
	for o := range addVAT(parse(lines)) {
		if o.err != nil {
			fmt.Println("skip:", o.err)
			continue
		}
		fmt.Printf("%s net=%d total=%d\n", o.id, o.net, o.total)
		revenue += o.total
	}
	fmt.Println("revenue:", revenue)
}
```

- **`addVAT(parse(lines))`** — both stages composed in one expression.
- Only error-free items count toward revenue: 1150 + 2760 = **3910**.

---

## Expected Output

```
ORD-1 net=1000 total=1150
skip: parse "ORD-2,abc": strconv.Atoi: parsing "abc": invalid syntax
ORD-3 net=2400 total=2760
skip: parse "broken-line": missing comma
revenue: 3910
```

## Key Takeaways

1. **A channel is a pipeline** — each stage is one function in, one function out.
2. **`<-chan` return** — sources can only read/write, so composition stays easy.
3. **`defer close` ends the chain** — the consumer's `range` terminates.
4. **Error as data** — an `err` field in the struct; a bad item doesn't stop any stage.
5. **`continue`** — skip the bad record, keep processing the rest.