# deadlock-goroutine-leak

Go-তে **deadlock by design** — unbuffered channel-এ receiver ছাড়াই send → runtime fatal error।

> **Note** — এই example ইচ্ছাকৃতভাবে crash করে। / This example **crashes by design.**

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

### Line 3

```go
import "fmt"
```

`fmt` (print)।

### Lines 5–10

```go
func main() {
	orders := make(chan string)

	orders <- "ORD-1"

	fmt.Println(<-orders)
}
```

- **Unbuffered channel** — `orders <- "ORD-1"` এখানেই আটকে যায়, কারণ কোনো receiver নেই।
- `fmt.Println(<-orders)` — এই line-ই আসল সমস্যার জায়গা, কিন্তু কোড-order-এ এটা আসে **পরে**, আগে send-ই থেমে যায়।
- Runtime দেখে সব goroutine আটকে → **`fatal error: all goroutines are asleep - deadlock!`** এবং exit code 1।

### Deadlock-এর নিয়ম

Unbuffered channel-এ send-এর জন্য **receiver** দরকার, receive-এর জন্য **sender**। দুই পাশেই যদি কেউ না থাকে, কোনো goroutine আর এগোতে পারে না।

---

## Expected Output (fatal error, exit 1)

```
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
	.../main.go:8 +0x36
```

## মূল শিক্ষা / Key Takeaways

1. **Unbuffered send blocks** — receiver ছাড়া অপেক্ষা।
2. **Runtime deadlock detector** — সব goroutine asleep হলে fatal।
3. **`+0x36`** — stack-এ instruction offset।
4. **Fix** — buffered channel, `select`+`default`, বা goroutine-তে receive।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Line 3

```go
import "fmt"
```

`fmt` (print).

### Lines 5–10

```go
func main() {
	orders := make(chan string)

	orders <- "ORD-1"

	fmt.Println(<-orders)
}
```

- **Unbuffered channel** — `orders <- "ORD-1"` blocks right here, because there is no receiver.
- `fmt.Println(<-orders)` is where the real fix belongs, but in code order it comes **after** the blocking send.
- The runtime sees every goroutine asleep → **`fatal error: all goroutines are asleep - deadlock!`** with exit code 1.

### The deadlock rule

An unbuffered send needs a **receiver**; a receive needs a **sender**. If neither side has a partner, no goroutine can proceed.

---

## Expected Output (fatal error, exit 1)

```
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
	.../main.go:8 +0x36
```

## Key Takeaways

1. **An unbuffered send blocks** — waiting for a receiver.
2. **Runtime deadlock detection** — fatal once every goroutine sleeps.
3. **`+0x36`** — an instruction offset in the stack trace.
4. **The fixes** — a buffered channel, `select`+`default`, or receiving in a goroutine.