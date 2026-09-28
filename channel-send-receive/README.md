# channel-send-receive

Go-তে **unbuffered channel sync** — goroutine-এ send, main-এ receive।

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

### Lines 5–13

```go
func main() {
	receipts := make(chan string)

	go func() {
		receipts <- "RCPT-9001"
	}()

	r := <-receipts
	fmt.Println("got receipt:", r)
}
```

- **`make(chan string)`** — unbuffered channel: send ও receive একসাথে না হলে দুই পাশই অপেক্ষা করে।
- goroutine-টা send করতে ready, তখন main-এর `<-receipts` receive তাকে সঙ্গে-সঙ্গে ছাড়ে।
- এই handshake-ই synchronization — কোনো `WaitGroup` বা `sleep` লাগে না।

---

## Expected Output

```
got receipt: RCPT-9001
```

## মূল শিক্ষা / Key Takeaways

1. **Unbuffered `make(chan T)`** — buffer নেই, শুধু handoff।
2. **Send waits for receive** — দুই পাশ একে অপরের জন্য অপেক্ষা।
3. **Synchronization primitive** — data race ছাড়াই handoff।
4. **Goroutine শেষ হওয়ার আগেই** main-এর value পাওয়া যায়।

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

### Lines 5–13

```go
func main() {
	receipts := make(chan string)

	go func() {
		receipts <- "RCPT-9001"
	}()

	r := <-receipts
	fmt.Println("got receipt:", r)
}
```

- **`make(chan string)`** — an unbuffered channel: both sides wait unless a send and a receive happen together.
- The goroutine becomes ready to send, and main's `<-receipts` receive releases it at the same moment.
- That handshake *is* the synchronization — no `WaitGroup` or `sleep` needed.

---

## Expected Output

```
got receipt: RCPT-9001
```

## Key Takeaways

1. **Unbuffered `make(chan T)`** — no buffer, only a handoff.
2. **A send waits for a receive** — both sides wait on each other.
3. **A synchronization primitive** — hands off data with no race.
4. **The value arrives** before the goroutine can finish.