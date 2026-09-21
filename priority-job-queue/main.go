package main

import (
	"container/heap"
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID         string
	Priority   int
	EnqueuedAt time.Time
	index      int
}

type JobHeap []*Job

func (h JobHeap) Len() int {
	return len(h)
}

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

func (q *JobQueue) Escalate(j *Job, newPriority int) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if j.index < 0 {
		return
	}

	j.Priority = newPriority
	heap.Fix(&q.h, j.index)
}

func main() {
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
}
