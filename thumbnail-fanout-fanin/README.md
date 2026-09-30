# thumbnail-fanout-fanin

Go-তে **fan-out / fan-in worker pool** — jobs/results দুটো channel, `WaitGroup`-এর পর `close(results)`, `SortFunc`-এ deterministic output।

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

### Lines 3–10

```go
import (
	"cmp"
	"errors"
	"fmt"
	"slices"
	"sync"
	"time"
)
```

`cmp` (Compare), `slices` (SortFunc), `sync` (WaitGroup), `time` (Sleep/Since)।

### Line 12

```go
var errCorrupt = errors.New("corrupt image")
```

**Sentinel error** — `errors.Is`/`%w` দিয়ে শনাক্তযোগ্য।

### Lines 14–23

```go
type thumbJob struct {
	id   int
	file string
}

type thumbResult struct {
	id   int
	path string
	err  error
}
```

**দুইটি struct** — ইনপুট (job) ও আউটপুট (result) আলাদা, আর আউটপুটে `id` বহন করা যাতে completion-এর ক্রম অপ্রাসঙ্গিক হলেও সঠিক জোড়া থাকে।

### Lines 25–32

```go
func makeThumbnail(j thumbJob) (string, error) {
	time.Sleep(50 * time.Millisecond)
	if j.file == "" {
		return "", fmt.Errorf("thumbnail for product %d: %w", j.id, errCorrupt)
	}

	return "thumbs/" + j.file, nil
}
```

- **`time.Sleep(50ms)`** — I/O-এর নকল mock (বাস্তবে disk/network latency)।
- খালি file → `%w`-wrapped sentinel error।

### Lines 34–41

```go
func worker(jobs <-chan thumbJob, results chan<- thumbResult, wg *sync.WaitGroup) {
	defer wg.Done()

	for j := range jobs {
		path, err := makeThumbnail(j)
		results <- thumbResult{id: j.id, path: path, err: err}
	}
}
```

- **`defer wg.Done()`** — panic-সহ যেকোনো ক্ষেত্রে counter কমবে।
- **`for j := range jobs`** — channel বন্ধ হলে worker নিজে থেকেই শেষ।
- **Directional channels** — শুধু jobs থেকে পড়ে, results-এ লেখে।
- **Error-ও পাঠায়** — job বাদ দেয় না, ফলে main প্রতিটি ফলাফল পায়।

### Lines 43–68

```go
func main() {
	files := []string{"shoe.jpg", "bag.jpg", "", "watch.jpg", "belt.jpg", "cap.jpg"}
	const numWorkers = 3

	jobs := make(chan thumbJob)
	results := make(chan thumbResult)

	var wg sync.WaitGroup

	for range numWorkers {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	go func() {
		defer close(jobs)

		for i, f := range files {
			jobs <- thumbJob{id: i + 1, file: f}
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()
```

- **Producer goroutine** — ৬টা job পাঠিয়ে `close(jobs)` (সব worker শেষ হবে)।
- **Closer goroutine** — `wg.Wait()`-এর পর `close(results)`, তাই main-এর `range results` সঠিকভাবে শেষ হয়।
- **Fan-out** — ৬ job → ৩ worker, তাই সমান্তরাল।

### Lines 70–94

```go
	start := time.Now()

	var collected []thumbResult

	for r := range results {
		collected = append(collected, r)
	}

	elapsed := time.Since(start)

	slices.SortFunc(collected, func(a, b thumbResult) int {
		return cmp.Compare(a.id, b.id)
	})

	var failed int

	for _, r := range collected {
		if r.err != nil {
			failed++
			fmt.Println("product", r.id, "error:", r.err)
			continue
		}
		fmt.Println("product", r.id, "->", r.path)
	}
	fmt.Printf("done: %d ok, %d failed, faster than sequential: %v\n", len(collected)-failed, failed, elapsed < 250*time.Millisecond)
}
```

- **`SortFunc`** — সম্পন্নের ক্রম অনির্ধারিত, তাই `id` অনুযায়ী সাজালে output-ই deterministic।
- **`failed++` + `continue`** — error আলাদা গুনে হিসাব রাখা।
- **`elapsed < 250ms`** — sequential হলে ৬×50 = 300ms হতো, তাই ৩ worker-এ ~100ms → `true`।

---

## Expected Output

```
product 1 -> thumbs/shoe.jpg
product 2 -> thumbs/bag.jpg
product 3 error: thumbnail for product 3: corrupt image
product 4 -> thumbs/watch.jpg
product 5 -> thumbs/belt.jpg
product 6 -> thumbs/cap.jpg
done: 5 ok, 1 failed, faster than sequential: true
```

## মূল শিক্ষা / Key Takeaways

1. **Fan-out** — এক producer, অনেক worker।
2. **Fan-in** — অনেক worker, এক consumer (`range results`)।
3. **`close` করার দায়িত্ব** — jobs producer বন্ধ করে, results closer goroutine (worker শেষ হলে)।
4. **`SortFunc`** — nondeterministic completion-কে deterministic presentation-এ।
5. **Concurrency gain** — ~100ms বনাম 300ms sequential।
6. **`wg.Add/Done` ছাড়া ভুল হয়** — counter ছাড়া `range results` কখনো শেষ হতো না।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Lines 3–10

```go
import (
	"cmp"
	"errors"
	"fmt"
	"slices"
	"sync"
	"time"
)
```

`cmp` (Compare), `slices` (SortFunc), `sync` (WaitGroup), `time` (Sleep/Since).

### Line 12

```go
var errCorrupt = errors.New("corrupt image")
```

A **sentinel error** — recognisable through `errors.Is`/`%w`.

### Lines 14–23

```go
type thumbJob struct {
	id   int
	file string
}

type thumbResult struct {
	id   int
	path string
	err  error
}
```

**Two structs** — the input (job) and the output (result) are separate, and the result carries `id` so that even though completion order is arbitrary, every result stays paired with its job.

### Lines 25–32

```go
func makeThumbnail(j thumbJob) (string, error) {
	time.Sleep(50 * time.Millisecond)
	if j.file == "" {
		return "", fmt.Errorf("thumbnail for product %d: %w", j.id, errCorrupt)
	}

	return "thumbs/" + j.file, nil
}
```

- **`time.Sleep(50ms)`** — a mock for I/O (in real life disk/network latency).
- An empty file yields the `%w`-wrapped sentinel error.

### Lines 34–41

```go
func worker(jobs <-chan thumbJob, results chan<- thumbResult, wg *sync.WaitGroup) {
	defer wg.Done()

	for j := range jobs {
		path, err := makeThumbnail(j)
		results <- thumbResult{id: j.id, path: path, err: err}
	}
}
```

- **`defer wg.Done()`** — the counter drops even on a panic.
- **`for j := range jobs`** — the worker exits by itself once the channel closes.
- **Directional channels** — it only reads jobs and writes results.
- **Errors are sent too** — no job is dropped, so main sees every outcome.

### Lines 43–68

```go
func main() {
	files := []string{"shoe.jpg", "bag.jpg", "", "watch.jpg", "belt.jpg", "cap.jpg"}
	const numWorkers = 3

	jobs := make(chan thumbJob)
	results := make(chan thumbResult)

	var wg sync.WaitGroup

	for range numWorkers {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	go func() {
		defer close(jobs)

		for i, f := range files {
			jobs <- thumbJob{id: i + 1, file: f}
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()
```

- **Producer goroutine** — pushes 6 jobs then `close(jobs)`, which ends every worker.
- **Closer goroutine** — `close(results)` after `wg.Wait()`, so main's `range results` terminates correctly.
- **Fan-out** — 6 jobs across 3 workers, run in parallel.

### Lines 70–94

```go
	start := time.Now()

	var collected []thumbResult

	for r := range results {
		collected = append(collected, r)
	}

	elapsed := time.Since(start)

	slices.SortFunc(collected, func(a, b thumbResult) int {
		return cmp.Compare(a.id, b.id)
	})

	var failed int

	for _, r := range collected {
		if r.err != nil {
			failed++
			fmt.Println("product", r.id, "error:", r.err)
			continue
		}
		fmt.Println("product", r.id, "->", r.path)
	}
	fmt.Printf("done: %d ok, %d failed, faster than sequential: %v\n", len(collected)-failed, failed, elapsed < 250*time.Millisecond)
}
```

- **`SortFunc`** — completion order is nondeterministic, so sorting by `id` makes the printed output deterministic.
- **`failed++` + `continue`** — counting errors separately from successes.
- **`elapsed < 250ms`** — sequential would be 6×50 = 300ms, so 3 workers take ~100ms → `true`.

---

## Expected Output

```
product 1 -> thumbs/shoe.jpg
product 2 -> thumbs/bag.jpg
product 3 error: thumbnail for product 3: corrupt image
product 4 -> thumbs/watch.jpg
product 5 -> thumbs/belt.jpg
product 6 -> thumbs/cap.jpg
done: 5 ok, 1 failed, faster than sequential: true
```

## Key Takeaways

1. **Fan-out** — one producer, many workers.
2. **Fan-in** — many workers, one consumer (`range results`).
3. **Who closes what** — jobs by the producer, results by the closer goroutine once workers finish.
4. **`SortFunc`** — turns nondeterministic completion into deterministic presentation.
5. **The concurrency gain** — ~100ms versus 300ms sequential.
6. **`wg.Add`/`Done` is mandatory** — without the counter, `range results` would never end.