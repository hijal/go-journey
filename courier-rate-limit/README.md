# courier-rate-limit

Go-তে **semaphore দিয়ে rate limit + atomic peak tracking** — `chan struct{}` (max 2), `atomic.Int32`, CAS loop, index-aligned write।

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
	"sync/atomic"
	"time"
)
```

`sync` (WaitGroup), `sync/atomic` (Int32), `time` (Sleep)।

### Line 10

```go
const maxConcurrent = 2
```

কতগুলো courier API একসঙ্গে (simultaneously) বলা যাবে।

### Lines 12–15

```go
func quoteShipping(parcelKg int) int {
	time.Sleep(300 * time.Millisecond)
	return 60 + parcelKg*20
}
```

- **`time.Sleep(300ms)`** — rate-limited ওয়েবসার্ভারের নকল latency (এই কারণেই concurrency সহায়ক)।
- দাম = ৬০ + ওজন×২০ টাকা।

### Lines 17–24

```go
func main() {
	parcels := []int{1, 3, 2, 5, 1, 4}
	quotes := make([]int, len(parcels))

	sem := make(chan struct{}, maxConcurrent)

	var inFlight, peak atomic.Int32
	var wg sync.WaitGroup
```

- **`quotes := make([]int, len(parcels))`** — ফলাফল আগে থেকে বানানো, index অনুযায়ী লেখা হয় (order-independent)।
- **`sem := make(chan struct{}, 2)`** — buffer 2-এর semaphore channel।
- **`inFlight`/`peak`** — atomic counter, কোনো mutex ছাড়া।

### Lines 26–46

```go
	for i, kg := range parcels {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem <- struct{}{}
			defer func() {
				<-sem
			}()

			n := inFlight.Add(1)

			for {
				p := peak.Load()
				if n <= p || peak.CompareAndSwap(p, n) {
					break
				}
			}
			quotes[i] = quoteShipping(kg)
			inFlight.Add(-1)
		}()
	}
```

1. **`sem <- struct{}{}`** — খালি slot নিলে ঢোকে, না পেলে এখানে অপেক্ষা (এটাই rate limit)।
2. **`defer func(){ <-sem }()`** — শেষে slot ফেরত দেয়; `defer wg.Done()`-এর **পরে** ঘটে (LIFO), তাই work শেষ হওয়ার পরেই slot মুক্ত হয়।
3. **`n := inFlight.Add(1)`** — এই মুহূর্তে active call সংখ্যা।
4. **CAS loop** — `peak` শুধু বাড়ানো হয়: `n <= peak` হলে থেমে যায়, নয়তো `CompareAndSwap` দিয়ে replace করে আবার চেষ্টা করে। (কেউ `peak` কমালে data race হতো।)
5. **`quotes[i] = ...`** — প্রতিটি goroutine আলাদা index-এ লেখে, তাই **কোনো race নেই** (mutex লাগে না)।
6. **`inFlight.Add(-1)`** — শেষ হলে কমাও।

### Lines 48–54

```go
	wg.Wait()

	for i, q := range quotes {
		fmt.Printf("parcel %d (%dkg): %d BDT\n", i+1, parcels[i], q)
	}

	fmt.Println("peak concurrent API calls:", peak.Load())
}
```

সব শেষ হয়ে ফল ছাপা হয় → order সবসময় একই। `peak` = **2** (সীমা ছাড়ায় ২-এর বেশি যায় না)।

---

## Expected Output

```
parcel 1 (1kg): 80 BDT
parcel 2 (3kg): 120 BDT
parcel 3 (2kg): 100 BDT
parcel 4 (5kg): 160 BDT
parcel 5 (1kg): 80 BDT
parcel 6 (4kg): 140 BDT
peak concurrent API calls: 2
```

## পর্যবেক্ষণ

`maxConcurrent=2` হওয়ায় ৬টা কাজ ৩ batch-এ শেষ হয় → মোট **~0.9s** (সীমা ছাড়া ৬×300ms = 1.8s হতো)।

## মূল শিক্ষা / Key Takeaways

1. **`chan struct{}` semaphore** — concurrency cap সবচেয়ে সহজ উপায়।
2. **`atomic.Int32` + CAS loop** — lock-free max ট্র্যাকিং।
3. **আলাদা index-এ লেখা race-free** — result slice-এর `quotes[i]`।
4. **`defer` LIFO** — resource release-এর সঠিক ক্রম।
5. **`maxConcurrent` কমালে** runtime বাড়ে (trade-off)।

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
	"sync/atomic"
	"time"
)
```

`sync` (WaitGroup), `sync/atomic` (Int32), `time` (Sleep).

### Line 10

```go
const maxConcurrent = 2
```

How many courier API calls may be in flight at the same time.

### Lines 12–15

```go
func quoteShipping(parcelKg int) int {
	time.Sleep(300 * time.Millisecond)
	return 60 + parcelKg*20
}
```

- **`time.Sleep(300ms)`** — mocked latency of the rate-limited web server (exactly what makes concurrency worthwhile).
- The price is 60 + weight×20 taka.

### Lines 17–24

```go
func main() {
	parcels := []int{1, 3, 2, 5, 1, 4}
	quotes := make([]int, len(parcels))

	sem := make(chan struct{}, maxConcurrent)

	var inFlight, peak atomic.Int32
	var wg sync.WaitGroup
```

- **`quotes := make([]int, len(parcels))`** — the results are allocated up front and written by index (order-independent).
- **`sem := make(chan struct{}, 2)`** — a semaphore channel with a buffer of 2.
- **`inFlight`/`peak`** — atomic counters, no mutex involved.

### Lines 26–46

```go
	for i, kg := range parcels {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem <- struct{}{}
			defer func() {
				<-sem
			}()

			n := inFlight.Add(1)

			for {
				p := peak.Load()
				if n <= p || peak.CompareAndSwap(p, n) {
					break
				}
			}
			quotes[i] = quoteShipping(kg)
			inFlight.Add(-1)
		}()
	}
```

1. **`sem <- struct{}{}`** — enters on a free slot, otherwise waits here (that *is* the rate limit).
2. **`defer func(){ <-sem }()`** — returns the slot at the end; it runs **after** `defer wg.Done()` (LIFO), so the slot frees only once the work is done.
3. **`n := inFlight.Add(1)`** — the number of active calls at this moment.
4. **The CAS loop** — `peak` only ever grows: break when `n <= peak`, otherwise replace it with `CompareAndSwap` and retry. (A plain read-then-write would be a data race.)
5. **`quotes[i] = ...`** — each goroutine writes a distinct index, so there is **no race** and no mutex is needed.
6. **`inFlight.Add(-1)`** — decrements when finished.

### Lines 48–54

```go
	wg.Wait()

	for i, q := range quotes {
		fmt.Printf("parcel %d (%dkg): %d BDT\n", i+1, parcels[i], q)
	}

	fmt.Println("peak concurrent API calls:", peak.Load())
}
```

Results print only after everything finishes, so the order is always identical. `peak` = **2** (never above the limit).

---

## Expected Output

```
parcel 1 (1kg): 80 BDT
parcel 2 (3kg): 120 BDT
parcel 3 (2kg): 100 BDT
parcel 4 (5kg): 160 BDT
parcel 5 (1kg): 80 BDT
parcel 6 (4kg): 140 BDT
peak concurrent API calls: 2
```

## Observation

With `maxConcurrent=2`, the 6 jobs finish in 3 batches → about **0.9s** total (without the limit it would be 6×300ms = 1.8s).

## Key Takeaways

1. **A `chan struct{}` semaphore** — the simplest concurrency cap.
2. **`atomic.Int32` + a CAS loop** — lock-free max tracking.
3. **Writing distinct indices is race-free** — `quotes[i]` needs no lock.
4. **`defer` LIFO order** — the correct order for resource release.
5. **Lowering `maxConcurrent` costs time** — the trade-off.