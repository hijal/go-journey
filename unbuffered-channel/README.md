# unbuffered-channel

Go-তে **unbuffered channel (sync send/receive) + `close` + `range`-over-channel** — producer/consumer pattern।

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

`fmt` — `Println`।

### Line 6

```go
	payments := make(chan int)
```

**`make(chan int)`** — **unbuffered**: capacity 0। একটা `send` block হয় যতক্ষণ না একটা `receive` তার সাথে মোতায়েক; এটা goroutine-গুলোর মধ্যে **synchronization** তৈরি করে।

### Lines 8–13

```go
	go func() {
		for _, amount := range []int{500, 1200, 300} {
			payments <- amount
		}
		close(payments)
	}()
```

**Sender goroutine** — sequential থেকে ৩টা value পাঠায়; শেষে **`close(payments)`** — আর value আসবে না।

### Lines 15–19

```go
	total := 0

	for amount := range payments {
		total += amount
	}
```

**`range` over channel** — close হওয়া পর্যন্ত প্রতি value পড়ে; close হলেই loop শেষ।

### Lines 21–22

```go
	fmt.Println("total received:", total)
```

500 + 1200 + 300 = **2000**।

> Note: এটা **unbuffered** (sync) channel। Buffer চাইলে `make(chan int, n)` — send-ব্লক-এ delay ছাড়া n-value hold করে।

---

## Expected Output

```
total received: 2000
```

## মূল শিক্ষা / Key Takeaways

1. **Unbuffered channel** — `send`↔`receive` synchronous (handshake)।
2. **`close`** — sender-এর signal: আর data নাই।
3. **`range` + channel** — close পর্যন্ত consume।
4. **Sorted send-receive** — producer-এ sequential guaranteed।
5. **No buffer** — block-ভিত্তিক flow control সহজাত।

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

`fmt` — `Println`.

### Line 6

```go
	payments := make(chan int)
```

**`make(chan int)`** — **unbuffered**: capacity 0. A `send` blocks until a `receive` pairs up with it; this creates **synchronization** between the goroutines.

### Lines 8–13

```go
	go func() {
		for _, amount := range []int{500, 1200, 300} {
			payments <- amount
		}
		close(payments)
	}()
```

**The sender goroutine** — sends 3 values sequentially; then **`close(payments)`** — signaling no more values.

### Lines 15–19

```go
	total := 0

	for amount := range payments {
		total += amount
	}
```

**`range` over a channel** — reads every value until it's closed; the loop ends on close.

### Lines 21–22

```go
	fmt.Println("total received:", total)
```

500 + 1200 + 300 = **2000**.

> Note: this is **unbuffered** (synchronous). For a buffer use `make(chan int, n)` — it holds n values without blocking the sender.

---

## Expected Output

```
total received: 2000
```

## Key Takeaways

1. **Unbuffered channel** — `send`↔`receive` are synchronous (a handshake).
2. **`close`** — the sender's signal: no more data.
3. **`range` + channel** — consume until closed.
4. **Ordered send-receive** — sequential producer guaranteed.
5. **No buffer** — inherent block-based flow control.