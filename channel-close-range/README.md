# channel-close-range

Go-তে **`close` + `range` drain** — channel বন্ধ করলে `range` শেষ, comma-ok false।

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

### Lines 5–18

```go
func main() {
	events := make(chan string, 2)

	events <- "user.signup"
	events <- "user.login"

	close(events)

	for e := range events {
		fmt.Println("event:", e)
	}

	v, ok := <-events
	fmt.Printf("v=%q ok=%v\n", v, ok)
}
```

- **`close(events)`** — আর কোনো value পাঠানো হবে না; buffer-এর ২টা value এখনও পড়া যায়।
- **`for e := range events`** — বন্ধ হওয়া পর্যন্ত বাকি value পড়ে, তারপর লুপ শেষ।
- **`v, ok := <-events`** — বন্ধ channel থেকে receive → `ok=false`, value zero (`""`)।

---

## Expected Output

```
event: user.signup
event: user.login
v="" ok=false
```

## মূল শিক্ষা / Key Takeaways

1. **`close(ch)`** — আর send হবে না, buffered value পড়া যায়।
2. **`range ch`** — close হলে automatic শেষ।
3. **Comma-ok** — `ok=false` মানে channel বন্ধ।
4. **Send on closed = panic** — তাই শুধু sender `close` করে।

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

### Lines 5–18

```go
func main() {
	events := make(chan string, 2)

	events <- "user.signup"
	events <- "user.login"

	close(events)

	for e := range events {
		fmt.Println("event:", e)
	}

	v, ok := <-events
	fmt.Printf("v=%q ok=%v\n", v, ok)
}
```

- **`close(events)`** — no more values will be sent, but the 2 buffered ones are still readable.
- **`for e := range events`** — drains the remaining values, then the loop ends.
- **`v, ok := <-events`** — receiving from a closed channel gives `ok=false` and the zero value (`""`).

---

## Expected Output

```
event: user.signup
event: user.login
v="" ok=false
```

## Key Takeaways

1. **`close(ch)`** — no further sends, buffered values still readable.
2. **`range ch`** — ends automatically once closed.
3. **Comma-ok** — `ok=false` means the channel is closed.
4. **Send on a closed channel panics** — which is why only the sender closes.