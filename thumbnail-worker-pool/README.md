# thumbnail-worker-pool

Go-তে **worker-pool pattern** — 4 worker, job channel + result channel; 12 upload থেকে thumbnail আর ~150ms-এ।

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

### Lines 3–9

```go
import (
	"cmp"
	"fmt"
	"slices"
	"sync"
	"time"
)
```

`slices` (sort) + `cmp` (compare) + `sync`।

### Lines 11–20

```go
type job struct {
	id   int
	file string
}

type result struct {
	id       int
	workerID int
	thumb    string
}
```

নামধারী task struct — job-এ ইনপুট, result-এ আউটপুট + কোন worker-টা।

### Lines 22–27

```go
func worker(id int, jobs <-chan job, results chan<- result) {
	for j := range jobs {
		time.Sleep(50 * time.Millisecond)
		results <- result{id: j.id, workerID: id, thumb: j.file + ".thumb.jpg"}
	}
}
```

**Worker loop** — `jobs` channel close-হওয়া অবধি নেয়, কাজ শেষে `results`-এ পাঠায়। Channel-টা closed হলে `range` নিজেই শেষ।

### Lines 29–53

```go
const workers = 4

jobs := make(chan job)
results := make(chan result, 12)

for w := range workers {
	wg.Go(func() { worker(w+1, jobs, results) })
}

go func() {
	defer close(jobs)
	for i := range 12 {
		jobs <- job{id: i + 1, file: fmt.Sprintf("upload-%02d", i+1)}
	}
}()

go func() {
	wg.Wait()
	close(results)
}()
```

- 4 worker goroutine।
- **Producer goroutine** — 12 job পাঠিয়ে `close(jobs)` (no more jobs signal)।
- **Closer goroutine** — সব worker শেষে `close(results)` (no more results signal)।

### Lines 55–68

```go
var done []result
for r := range results {
	done = append(done, r)
}

slices.SortFunc(done, func(a, b result) int {
	return cmp.Compare(a.id, b.id)
})

for _, r := range done {
	fmt.Printf("job %2d -> %s (worker %d)\n", r.id, r.thumb, r.workerID)
}
```

Result-order nondeterministic — **id-ক্রমানুসারে sort** করে print। Timing: 12 job / 4 worker x 50ms = **~150ms**।

---

## Expected Output (worker IDs vary)

```
job  1 -> upload-01.thumb.jpg (worker 3)
job  2 -> upload-02.thumb.jpg (worker 1)
...
job 12 -> upload-12.thumb.jpg (worker 2)
processed 12 uploads with 4 workers in 150ms
```

## মূল শিক্ষা / Key Takeaways

1. **Worker-pool** — reusable goroutines, job-এ continuous stream।
2. **`close(jobs)`** — producer-done signal; worker loop gracefully শেষ।
3. **`close(results)`** — main-র receiver range শেষ করার cue।
4. **Buffered results** — work-ভারে producer/fan-in unbind।
5. **`cmp.Compare` sort** — id-স্থির output।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Lines 3–9

```go
import (
	"cmp"
	"fmt"
	"slices"
	"sync"
	"time"
)
```

`slices` (sort) + `cmp` (compare) + `sync`.

### Lines 11–20

```go
type job struct {
	id   int
	file string
}

type result struct {
	id       int
	workerID int
	thumb    string
}
```

Structured data — a `job` as input, a `result` with the output + which worker.

### Lines 22–27

```go
func worker(id int, jobs <-chan job, results chan<- result) {
	for j := range jobs {
		time.Sleep(50 * time.Millisecond)
		results <- result{id: j.id, workerID: id, thumb: j.file + ".thumb.jpg"}
	}
}
```

**The worker loop** — pulls jobs until the channel closes, then pushes a result. `range` ends itself once closed.

### Lines 29–53

```go
const workers = 4

jobs := make(chan job)
results := make(chan result, 12)

for w := range workers {
	wg.Go(func() { worker(w+1, jobs, results) })
}

go func() {
	defer close(jobs)
	for i := range 12 {
		jobs <- job{id: i + 1, file: fmt.Sprintf("upload-%02d", i+1)}
	}
}()

go func() {
	wg.Wait()
	close(results)
}()
```

- 4 worker goroutines.
- **Producer goroutine** — pushes 12 jobs then `close(jobs)` (no more jobs signal).
- **Closer goroutine** — `close(results)` once all workers are done (no more results signal).

### Lines 55–68

```go
var done []result
for r := range results {
	done = append(done, r)
}

slices.SortFunc(done, func(a, b result) int {
	return cmp.Compare(a.id, b.id)
})

for _, r := range done {
	fmt.Printf("job %2d -> %s (worker %d)\n", r.id, r.thumb, r.workerID)
}
```

Result arrival is nondeterministic — **sort by id** before printing. Timing: 12 jobs / 4 workers x 50ms = **~150ms**.

---

## Expected Output (worker IDs vary)

```
job  1 -> upload-01.thumb.jpg (worker 3)
job  2 -> upload-02.thumb.jpg (worker 1)
...
job 12 -> upload-12.thumb.jpg (worker 2)
processed 12 uploads with 4 workers in 150ms
```

## Key Takeaways

1. **Worker-pool** — reusable goroutines over a continuous job stream.
2. **`close(jobs)`** — the producer-done signal; workers end their loop gracefully.
3. **`close(results)`** — the cue for the main receiver to finish ranging.
4. **Buffered results** — unbind producer/fan-in from bursty work.
5. **`cmp.Compare` sort** — deterministic id-ordered output.