# welcome-email-queue

Go-তে **producer/consumer + completion signal** — buffered queue, `close(emails)` শেষ, `done` channel-এ সম্পন্ন-সংকেত।

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

### Lines 5–15

```go
func main() {
	emails := make(chan string, 10)
	done := make(chan struct{})

	go func() {
		defer close(done)
		for add := range emails {
			fmt.Println("sending welcome email to", add)
		}
		fmt.Println("queue drained, worker exiting")
	}()
```

- **`make(chan string, 10)`** — ১০-slot queue, তাই producer (main) receive ছাড়াই সব address পাঠাতে পারে।
- **`done := make(chan struct{})`** — zero-size signal channel, শুধু "শেষ হয়েছে" জানাতে।
- **`defer close(done)`** — worker যেকোনোভাবেই শেষ হোক, `done` বন্ধ হবে।
- **`for add := range emails`** — channel বন্ধ হওয়া পর্যন্ত value পড়ে, তারপর লুপ শেষ।
- লুপের বাইরে `queue drained...` — মানে সব email পাঠানো হয়েছে।

### Lines 17–25

```go
	signups := []string{"rina@shop.test", "tanvir@shop.test", "nadia@shop.test"}

	for _, add := range signups {
		emails <- add
	}
	close(emails)

	<-done
	fmt.Println("enqueued", len(signups), "emails; shutting down")
}
```

- **FIFO** — address-গুলো queue-তে যেই ক্রমে ঢোকে, সেই ক্রমেই আসে।
- **`close(emails)`** — "আর কোনো email আসবে না" — এটাই worker-কে লুপ থেকে বের করে।
- **`<-done`** — worker-পুরো শেষ না হওয়া পর্যন্ত main-ও থামে না (নাহলে program শেষ হয়ে যেত)।

---

## Expected Output

```
sending welcome email to rina@shop.test
sending welcome email to tanvir@shop.test
sending welcome email to nadia@shop.test
queue drained, worker exiting
enqueued 3 emails; shutting down
```

## মূল শিক্ষা / Key Takeaways

1. **Buffered queue** — producer block হয় না (10 > 3)।
2. **`close` = শেষ সংকেত** — consumer-এর `range` এখানেই শেষ।
3. **`chan struct{}` + `close(done)`** — completion signal।
4. **দুই ধরনের close** — ডেটা শেষ (`emails`) বনাম কাজ শেষ (`done`)।

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

### Lines 5–15

```go
func main() {
	emails := make(chan string, 10)
	done := make(chan struct{})

	go func() {
		defer close(done)
		for add := range emails {
			fmt.Println("sending welcome email to", add)
		}
		fmt.Println("queue drained, worker exiting")
	}()
```

- **`make(chan string, 10)`** — a 10-slot queue, so the producer (main) can push every address without waiting for a receive.
- **`done := make(chan struct{})`** — a zero-size signal channel used only to say "finished".
- **`defer close(done)`** — however the worker ends, `done` gets closed.
- **`for add := range emails`** — reads until the channel closes, then the loop ends.
- The `queue drained...` line sits after the loop — meaning every email was sent.

### Lines 17–25

```go
	signups := []string{"rina@shop.test", "tanvir@shop.test", "nadia@shop.test"}

	for _, add := range signups {
		emails <- add
	}
	close(emails)

	<-done
	fmt.Println("enqueued", len(signups), "emails; shutting down")
}
```

- **FIFO** — addresses come out in the order they went in.
- **`close(emails)`** — "no more emails are coming", and that's what frees the worker from its loop.
- **`<-done`** — main won't finish before the worker has (otherwise the program would exit early).

---

## Expected Output

```
sending welcome email to rina@shop.test
sending welcome email to tanvir@shop.test
sending welcome email to nadia@shop.test
queue drained, worker exiting
enqueued 3 emails; shutting down
```

## Key Takeaways

1. **A buffered queue** — the producer never blocks (10 > 3).
2. **`close` as an end-of-data signal** — the consumer's `range` stops there.
3. **`chan struct{}` + `close(done)`** — a completion signal.
4. **Two different closes** — end of data (`emails`) vs end of work (`done`).