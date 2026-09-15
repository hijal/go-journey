# buffer-pooling-sync

Go-তে **`sync.Pool` + `bytes.Buffer` reuse + `defer Put` + type assertion** শেখার ছোট example — zero-alloc JSON builder।

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
	"bytes"
	"fmt"
	"sync"
)
```

- `bytes` — `Buffer`।
- `fmt` — `Println`, `Sprintf`।
- `sync` — `Pool`।

### Lines 9–13

```go
var bufferpool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}
```

Package-level pool — `New` func নতুন `*bytes.Buffer` তৈরি করে যখন খালি।

### Lines 15–24

```go
func renderResponse(message string) string {
	buf := bufferpool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufferpool.Put(buf)

	buf.WriteString(`{"message":"`)
	buf.WriteString(message)
	buf.WriteString(`"}`)
	return buf.String()
}
```

**Pool reuse flow:**

1. `bufferpool.Get()` — pool-থেকে নাও (interface{} return)।
2. `.(*bytes.Buffer)` — **type assertion**: interface{} → `*bytes.Buffer`।
3. `buf.Reset()` — আগের ব্যবহারের ডেটা মুছে দাও (reuse)।
4. `defer bufferpool.Put(buf)` — function end-এ ফেরত দাও।
5. `WriteString` chains — JSON বানাও।
6. `buf.String()` — চূড়ান্ত string।

*(`Reset()` + `defer Put` — প্রতিটি call-এর জন্য clean slate + automatic return।)*

### Lines 26–30

```go
func main() {
	for i := 0; i < 3; i++ {
		response := renderResponse(fmt.Sprintf("request #%d handled", i))
		fmt.Println(response)
	}
}
```

৩বার render — একই buffer pool reuse হবে (zero-alloc)।

---

## Expected Output

```
{"message":"request #0 handled"}
{"message":"request #1 handled"}
{"message":"request #2 handled"}
```

## মূল শিক্ষা / Key Takeaways

1. **`sync.Pool`** — reuse buffer (GC pressure কমায়)।
2. **Type assertion** — `interface{}` → `*bytes.Buffer`।
3. **`Reset()` + `defer Put()`** — clean + automatic return।
4. **Backtick strings** — no escaping; careful with quotes inside JSON।
5. **`-race` safe** — pool-এর internal lock concurrency handle করে।

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
	"bytes"
	"fmt"
	"sync"
)
```

- `bytes` — for `Buffer`.
- `fmt` — for `Println`, `Sprintf`.
- `sync` — for `Pool`.

### Lines 9–13

```go
var bufferpool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}
```

A package-level pool — the `New` function creates a fresh `*bytes.Buffer` when the pool is empty.

### Lines 15–24

```go
func renderResponse(message string) string {
	buf := bufferpool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufferpool.Put(buf)

	buf.WriteString(`{"message":"`)
	buf.WriteString(message)
	buf.WriteString(`"}`)
	return buf.String()
}
```

**The pool reuse flow:**

1. `bufferpool.Get()` — grab a buffer from the pool (returns `interface{}`).
2. `.(*bytes.Buffer)` — **type assertion**: interface{} → `*bytes.Buffer`.
3. `buf.Reset()` — wipe the previous data (reusable).
4. `defer bufferpool.Put(buf)` — return it at function end.
5. `WriteString` chains — build the JSON.
6. `buf.String()` — the final string.

*(`Reset()` + `defer Put` — clean slate each call + automatic return.)*

### Lines 26–30

```go
func main() {
	for i := 0; i < 3; i++ {
		response := renderResponse(fmt.Sprintf("request #%d handled", i))
		fmt.Println(response)
	}
}
```

3 renders — the same buffer is reused (zero-alloc).

---

## Expected Output

```
{"message":"request #0 handled"}
{"message":"request #1 handled"}
{"message":"request #2 handled"}
```

## Key Takeaways

1. **`sync.Pool`** — buffer reuse (less GC pressure).
2. **Type assertion** — `interface{}` → `*bytes.Buffer`.
3. **`Reset()` + `defer Put()`** — clean slate + automatic return.
4. **Backtick strings** — no escaping; watch for quotes inside JSON.
5. **`-race` safe** — pool's internal lock handles concurrency.