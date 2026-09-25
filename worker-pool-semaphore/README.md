# worker-pool-semaphore

Go-তে **buffered channel semaphore (`chan struct{}`) দিয়ে concurrency limiting** — 100k image resize job-এ maximum 100 parallel worker।

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

### Lines 3–8

```go
import (
	"fmt"
	"sync"
	"time"
)
```

`sync` (`WaitGroup`) + `time` (simulated resize delay)।

### Lines 9–12

```go
const maxConcurrent = 100

sem := make(chan struct{}, maxConcurrent)
```

**Semaphore** — capacity-`maxConcurrent`-এর buffered channel। `struct{}` (zero-size) token holder হিসেবে।

### Lines 14–24

```go
var wg sync.WaitGroup

for imageID := range 100000 {
	wg.Go(func() {
		sem <- struct{}{}
		defer func() { <-sem }()

		fmt.Println("resizing image", imageID)
		time.Sleep(100 * time.Millisecond)
	})
}
```

**Acquire → work → release:**

- `sem <-` acquire — empty slot-এ token push; সব slot ভর্তি হলে নতুন goroutine block করে (max 100 parallel)।
- `defer <-sem` — কাজ শেষে release; অন্যান্য-এর সুযোগ।
- throughput = 100 job / 100ms — semaphore-এর effect directly observable।

### Lines 26–28

```go
wg.Wait()

fmt.Println("all images resized")
```

সব job শেষে report।

---

## Expected Output (truncated preview)

```
resizing image 0
resizing image 1
...
resizing image 99998
resizing image 99999
all images resized
```

(~100k lines; `maxConcurrent=100` হলে মোট runtime ≈ 100s)

## মূল শিক্ষা / Key Takeaways

1. **`chan struct{}` as semaphore** — buffered capacity = concurrency limit।
2. **`sem <-` acquire / `<-sem` release** — blocking handshake।
3. **No extra lock** — channel-ই synchronization।
4. **`wg.Wait`** — সব worker শেষ হওয়া অবধি block।
5. **Peak concurrency** — কখনো 100-এর বেশি parallel নয়।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Lines 3–8

```go
import (
	"fmt"
	"sync"
	"time"
)
```

`sync` (`WaitGroup`) + `time` (simulated resize delay).

### Lines 9–12

```go
const maxConcurrent = 100

sem := make(chan struct{}, maxConcurrent)
```

**A semaphore** — a buffered channel with capacity `maxConcurrent`. `struct{}` (zero-size) as token holder.

### Lines 14–24

```go
var wg sync.WaitGroup

for imageID := range 100000 {
	wg.Go(func() {
		sem <- struct{}{}
		defer func() { <-sem }()

		fmt.Println("resizing image", imageID)
		time.Sleep(100 * time.Millisecond)
	})
}
```

**Acquire → work → release:**

- `sem <-` acquire — push a token into a free slot; when all slots are full, new goroutines block (max 100 parallel).
- `defer <-sem` — release after the work, giving others a chance.
- throughput = 100 jobs / 100ms — the semaphore's effect is directly observable.

### Lines 26–28

```go
wg.Wait()

fmt.Println("all images resized")
```

Report once every job has finished.

---

## Expected Output (truncated preview)

```
resizing image 0
resizing image 1
...
resizing image 99998
resizing image 99999
all images resized
```

(~100k lines; total runtime ≈ 100s with `maxConcurrent=100`)

## Key Takeaways

1. **`chan struct{}` as semaphore** — buffered capacity = concurrency limit.
2. **`sem <-` acquire / `<-sem` release** — block-backed handshake.
3. **No extra lock** — the channel is the synchronization.
4. **`wg.Wait`** — blocks until every worker finishes.
5. **Peak concurrency** — never more than 100 parallel.