# priority-job-queue

Go-তে **`container/heap` + max-priority queue + tie-break + `heap.Fix` escalation** শেখার ছোট example — job queue with priority scheduling।

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
	"container/heap"
	"fmt"
	"sync"
	"time"
)
```

- `container/heap` — generic heap machinery।
- `sync` — `Mutex`।
- `time` — `Time`, `Since`।

### Lines 10–15

```go
type Job struct {
	ID         string
	Priority   int
	EnqueuedAt time.Time
	index      int
}
```

**Job** — priority + enqueue-time (+ **`index`** — heap-position tracking, internal)।

### Lines 17–50

```go
type JobHeap []*Job

func (h JobHeap) Len() int { return len(h) }

func (h JobHeap) Less(i, j int) bool {
	if h[i].Priority != h[j].Priority {
		return h[i].Priority > h[j].Priority
	}
	return h[i].EnqueuedAt.Before(h[j].EnqueuedAt)
}

func (h JobHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *JobHeap) Push(x any) {
	job := x.(*Job)
	job.index = len(*h)
	*h = append(*h, job)
}

func (h *JobHeap) Pop() any {
	old := *h
	n := len(old)
	job := old[n-1]
	old[n-1] = nil
	job.index = -1
	*h = old[:n-1]
	return job
}
```

**`heap.Interface`** (Len/Less/Swap/Push/Pop):

- `Less` — **max-heap**: priority বড় প্রথম; tie-তে enqueue-time আগে।
- `Swap` — `index` sync।
- `Push`/`Pop` — `index` bookkeeping (Pop-এ `-1` = dequeued marker)।

### Lines 52–71

```go
type JobQueue struct {
	mu sync.Mutex
	h  JobHeap
}

func (q *JobQueue) Enqueue(j *Job) {
	q.mu.Lock()
	defer q.mu.Unlock()
	heap.Push(&q.h, j)
}

func (q *JobQueue) Dequeue() (*Job, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.h.Len() == 0 {
		return nil, false
	}
	return heap.Pop(&q.h).(*Job), true
}
```

**Concurrent queue** — mutex + `heap.Push`/`heap.Pop`।

### Lines 73–83

```go
func (q *JobQueue) Escalate(j *Job, newPriority int) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if j.index < 0 {
		return
	}

	j.Priority = newPriority
	heap.Fix(&q.h, j.index)
}
```

**Priority escalation** — `j.index` দিয়ে `heap.Fix` (dequeued job-এ no-op)।

### Lines 85–104

```go
	base := time.Date(2026, 9, 16, 9, 0, 0, 0, time.UTC)
	q := &JobQueue{}

	report := &Job{ID: "generate-monthly-report", Priority: 1, EnqueuedAt: base}
	q.Enqueue(report)
	q.Enqueue(&Job{ID: "send-password-reset", Priority: 10, EnqueuedAt: base.Add(time.Second)})
	q.Enqueue(&Job{ID: "resize-avatar", Priority: 5, EnqueuedAt: base.Add(2 * time.Second)})
	q.Enqueue(&Job{ID: "send-receipt", Priority: 10, EnqueuedAt: base.Add(3 * time.Second)})

	q.Escalate(report, 7)

	for {
		job, ok := q.Dequeue()
		if !ok {
			break
		}
		fmt.Printf("running %-24s priority=%d\n", job.ID, job.Priority)
	}
```

4টা job enqueue → `report`-কে 1→7 escalate → max-priority order-এ drain।

---

## Expected Output

```
running send-password-reset      priority=10
running send-receipt             priority=10
running generate-monthly-report  priority=7
running resize-avatar            priority=5
```

## মূল শিক্ষা / Key Takeaways

1. **`container/heap`** — স্ট্যান্ডার্ড heap machinery (Len/Less/Swap/Push/Pop)।
2. **Max-heap + tie-break** — priority desc, time asc।
3. **`heap.Fix`** — changed-priority reposition (index-tracked)।
4. **Internal `index`** — Pop-এ `-1`, dequeued marker।
5. **Mutex queue** — concurrent Enqueue/Dequeue।

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
	"container/heap"
	"fmt"
	"sync"
	"time"
)
```

- `container/heap` — the generic heap machinery.
- `sync` — `Mutex`.
- `time` — `Time`.

### Lines 10–15

```go
type Job struct {
	ID         string
	Priority   int
	EnqueuedAt time.Time
	index      int
}
```

**Job** — priority + enqueue time (+ **`index`** — heap-position tracking, internal).

### Lines 17–50

```go
type JobHeap []*Job

func (h JobHeap) Len() int { return len(h) }

func (h JobHeap) Less(i, j int) bool {
	if h[i].Priority != h[j].Priority {
		return h[i].Priority > h[j].Priority
	}
	return h[i].EnqueuedAt.Before(h[j].EnqueuedAt)
}

func (h JobHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *JobHeap) Push(x any) {
	job := x.(*Job)
	job.index = len(*h)
	*h = append(*h, job)
}

func (h *JobHeap) Pop() any {
	old := *h
	n := len(old)
	job := old[n-1]
	old[n-1] = nil
	job.index = -1
	*h = old[:n-1]
	return job
}
```

**`heap.Interface`** (Len/Less/Swap/Push/Pop):

- `Less` — **max-heap**: bigger priority first; ties go to the earlier enqueue time.
- `Swap` — syncs `index`.
- `Push`/`Pop` — `index` bookkeeping (`-1` on Pop = the dequeued marker).

### Lines 52–71

```go
type JobQueue struct {
	mu sync.Mutex
	h  JobHeap
}

func (q *JobQueue) Enqueue(j *Job) {
	q.mu.Lock()
	defer q.mu.Unlock()
	heap.Push(&q.h, j)
}

func (q *JobQueue) Dequeue() (*Job, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.h.Len() == 0 {
		return nil, false
	}
	return heap.Pop(&q.h).(*Job), true
}
```

**A concurrent queue** — a mutex + `heap.Push`/`heap.Pop`.

### Lines 73–83

```go
func (q *JobQueue) Escalate(j *Job, newPriority int) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if j.index < 0 {
		return
	}

	j.Priority = newPriority
	heap.Fix(&q.h, j.index)
}
```

**Priority escalation** — `heap.Fix` at `j.index` (a no-op for already-dequeued jobs).

### Lines 85–104

```go
	base := time.Date(2026, 9, 16, 9, 0, 0, 0, time.UTC)
	q := &JobQueue{}

	report := &Job{ID: "generate-monthly-report", Priority: 1, EnqueuedAt: base}
	q.Enqueue(report)
	q.Enqueue(&Job{ID: "send-password-reset", Priority: 10, EnqueuedAt: base.Add(time.Second)})
	q.Enqueue(&Job{ID: "resize-avatar", Priority: 5, EnqueuedAt: base.Add(2 * time.Second)})
	q.Enqueue(&Job{ID: "send-receipt", Priority: 10, EnqueuedAt: base.Add(3 * time.Second)})

	q.Escalate(report, 7)

	for {
		job, ok := q.Dequeue()
		if !ok {
			break
		}
		fmt.Printf("running %-24s priority=%d\n", job.ID, job.Priority)
	}
```

4 jobs enqueued → `report` escalated 1→7 → drained in max-priority order.

---

## Expected Output

```
running send-password-reset      priority=10
running send-receipt             priority=10
running generate-monthly-report  priority=7
running resize-avatar            priority=5
```

## Key Takeaways

1. **`container/heap`** — the standard heap machinery (Len/Less/Swap/Push/Pop).
2. **Max-heap + tie-break** — priority desc, time asc.
3. **`heap.Fix`** — repositioning after a priority change (index-tracked).
4. **Internal `index`** — `-1` on Pop = the dequeued marker.
5. **A mutex queue** — concurrent Enqueue/Dequeue.