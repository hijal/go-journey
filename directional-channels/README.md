# directional-channels

Go-তে **directional channel type** — `chan<- int` (send-only), `<-chan int` (receive-only)।

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
func produce(out chan<- int) {
	for i := range 10 {
		out <- i * 100
	}
	close(out)
}
```

**`chan<- int`** — শুধু **send** করা যায়। Parameter-এ direction দিয়ে API সীমাবদ্ধ ও নিজে থেকে `close` করে signal দেয়।

### Lines 12–16

```go
func consume(in <-chan int) {
	for n := range in {
		fmt.Println("number:", n)
	}
}
```

**`<-chan int`** — শুধু **receive**। `consume` `close` করতে পারে না।

### Lines 18–22

```go
func main() {
	ch := make(chan int)

	go produce(ch)
	consume(ch)
}
```

`ch` দুই পাশেই pass হলেও প্রত্যেক function নিজের direction মেনে চলে।

---

## Expected Output

```
number: 0
number: 100
number: 200
...
number: 900
```

## মূল শিক্ষা / Key Takeaways

1. **`chan<- T`** — send-only (producer)।
2. **`<-chan T`** — receive-only (consumer)।
3. **Producer closes** — শেষ হওয়ার signal।
4. **সুরক্ষিত API** — ভুল দিকে ব্যবহার compile error।

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
func produce(out chan<- int) {
	for i := range 10 {
		out <- i * 100
	}
	close(out)
}
```

**`chan<- int`** — **send only**. A direction in the parameter restricts the API, and the producer signals completion by closing.

### Lines 12–16

```go
func consume(in <-chan int) {
	for n := range in {
		fmt.Println("number:", n)
	}
}
```

**`<-chan int`** — **receive only**. `consume` cannot close it.

### Lines 18–22

```go
func main() {
	ch := make(chan int)

	go produce(ch)
	consume(ch)
}
```

`ch` is passed to both sides, yet each function honours its own direction.

---

## Expected Output

```
number: 0
number: 100
number: 200
...
number: 900
```

## Key Takeaways

1. **`chan<- T`** — send-only (producer).
2. **`<-chan T`** — receive-only (consumer).
3. **The producer closes** — that's the completion signal.
4. **A safer API** — using the wrong direction is a compile error.