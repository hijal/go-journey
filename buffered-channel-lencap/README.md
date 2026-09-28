# buffered-channel-lencap

Go-তে **buffered channel** — `make(chan T, n)`-এ `len`/`cap`, buffered send-এ block হয় না।

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

### Lines 5–16

```go
func main() {
	jobs := make(chan string, 3)

	jobs <- "resize-image"
	jobs <- "send-invoice"

	fmt.Println("len:", len(jobs), "cap:", cap(jobs))

	fmt.Println(<-jobs)
	fmt.Println(<-jobs)

	fmt.Println("len:", len(jobs))
}
```

- **`make(chan string, 3)`** — buffer 3-এর channel, মাত্র ২টা value রাখা।

  - `len` = এখন buffer-এ কয়টা (2)
  - `cap` = buffer-এর সর্বোচ্চ ধারণ (3)

- দুটো `jobs <- ...` send — buffer খালি থাকায় **সঙ্গে সঙ্গে** হয়, block হয় না।
- দুটো `<-jobs` receive → value ছাড়ে।
- শেষে `len: 0` — buffer খালি।

---

## Expected Output

```
len: 2 cap: 3
resize-image
send-invoice
len: 0
```

## মূল শিক্ষা / Key Takeaways

1. **`make(chan T, n)`** — buffer-সহ channel।
2. **`len` vs `cap`** — বর্তমান element বনাম সর্বোচ্চ ক্ষমতা।
3. **Buffered send** — receiver ছাড়াই এগোয় (buffer খালি থাকলে)।
4. **Unbuffered vs buffered** — scheduling-এর পার্থক্য।

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

### Lines 5–16

```go
func main() {
	jobs := make(chan string, 3)

	jobs <- "resize-image"
	jobs <- "send-invoice"

	fmt.Println("len:", len(jobs), "cap:", cap(jobs))

	fmt.Println(<-jobs)
	fmt.Println(<-jobs)

	fmt.Println("len:", len(jobs))
}
```

- **`make(chan string, 3)`** — a channel with a buffer of 3, holding only 2 values.

  - `len` = how many are in the buffer right now (2)
  - `cap` = the buffer's maximum capacity (3)

- The two `jobs <- ...` sends complete **immediately** because the buffer has room; no blocking.
- Two `<-jobs` receives pull the values out.
- Finally `len: 0` — the buffer is empty.

---

## Expected Output

```
len: 2 cap: 3
resize-image
send-invoice
len: 0
```

## Key Takeaways

1. **`make(chan T, n)`** — a buffered channel.
2. **`len` vs `cap`** — current elements vs maximum capacity.
3. **Buffered send** — doesn't wait for a receiver while the buffer has room.
4. **Unbuffered vs buffered** — the scheduling difference.