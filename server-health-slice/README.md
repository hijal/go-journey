# server-health-slice

Go-তে **slice `append` + dynamic growth + pre-allocation** শেখার ছোট example — DevOps monitoring: server-এ ping পাঠিয়ে fail-গুলো আলাদা slice-এ জমানো।

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

### Lines 3–6

```go
import (
	"fmt"
	"math/rand"
)
```

- `fmt` — `Println`।
- `math/rand` — healthy/fail simulate-র জন্য random।

### Line 9

```go
	servers := []string{"api-1", "api-2", "db-1", "cache-1", "worker-1"}
```

`servers` — 5টা server-এর নাম (slice literal)। এগুলোই check হবে।

### Line 11

```go
	failed := make([]string, 0, len(servers))
```

**Pre-allocation** — `make([]string, 0, 5)`:

- `len` = **0** — এখনো কোনো element নেই।
- `cap` = **5** — পেছনের capacity-টা আগে থেকেই 5।
- সুবিধা: সাইজ জানা (5 server) → capacity reserve → পরের `append`-গুলো **re-allocation ছাড়া** কাজ করে (costly copy এড়ায়)। এরই একটা best practice যখন maximum সাইজ prediction করা যায়।

### Lines 13–18

```go
	for _, server := range servers {
		healthy := rand.Intn(100) < 80
		if !healthy {
			failed = append(failed, server)
		}
	}
```

- `rand.Intn(100) < 80` — **80% chance** healthy।
- `if !healthy` — fail হলে `append(failed, server)` — slice-এ dynamic জমানো। Capacity শেষ হলে `append` নিজেই bigger array allocate করে (automatic growth) — এখানে 5-capacity থাকায় বাড়তেই হয় না।

### Lines 19–21

```go
	fmt.Println("healthy servers:", len(servers)-len(failed))
	fmt.Println("total servers:", len(servers))
	fmt.Println("failed servers:", failed)
```

- `len(servers)-len(failed)` — healthy = total − failed (fail-slice থেকে derive)।
- `len(servers)` — total।
- `failed` slice print — শুধু দরকারি (fail) সার্ভারগুলো।

---

## Expected Output

(সম্ভাব্য — random-এর জন্য ভিন্ন হতে পারে)

```
healthy servers: 4
total servers: 5
failed servers: [api-2]
```

0-5টা fail হতে পারে — সব healthy-এ run-এ `failed: []` দেখাবে।

## মূল শিক্ষা / Key Takeaways

1. **`make([]string, 0, cap)`** — pre-allocation: zero-length, reserved capacity।
2. **`append`** — dynamic slice build।
3. **Automatic growth** — capacity শেষে append নিজেই reslice।
4. **Derived count** — healthy = total − failed-এর diff।
5. **Random simulation** — `rand.Intn` দিয়ে healthy flag।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Lines 3–6

```go
import (
	"fmt"
	"math/rand"
)
```

- `fmt` — for `Println`.
- `math/rand` — randomness to simulate healthy/fail.

### Line 9

```go
	servers := []string{"api-1", "api-2", "db-1", "cache-1", "worker-1"}
```

`servers` — the names of 5 servers (a slice literal). These are what gets checked.

### Line 11

```go
	failed := make([]string, 0, len(servers))
```

**Pre-allocation** — `make([]string, 0, 5)`:

- `len` = **0** — no elements yet.
- `cap` = **5** — the backing capacity is already 5.
- Benefit: when the size is known (5 servers), reserving capacity means the later `append`s work **without re-allocations** (avoiding costly copies). A best practice when the maximum size can be predicted.

### Lines 13–18

```go
	for _, server := range servers {
		healthy := rand.Intn(100) < 80
		if !healthy {
			failed = append(failed, server)
		}
	}
```

- `rand.Intn(100) < 80` — an **80% chance** to be healthy.
- `if !healthy` — on failure, `append(failed, server)` — building the slice dynamically. When the capacity runs out, `append` allocates a bigger array itself (automatic growth) — here the pre-allocated 5-capacity is enough.

### Lines 19–21

```go
	fmt.Println("healthy servers:", len(servers)-len(failed))
	fmt.Println("total servers:", len(servers))
	fmt.Println("failed servers:", failed)
```

- `len(servers)-len(failed)` — healthy = total − failed (derived from the fail slice).
- `len(servers)` — the total.
- The `failed` slice prints only the servers we care about.

---

## Expected Output

(possible — varies with randomness)

```
healthy servers: 4
total servers: 5
failed servers: [api-2]
```

0–5 failures are possible — an all-healthy run shows `failed: []`.

## Key Takeaways

1. **`make([]string, 0, cap)`** — pre-allocation: zero length, reserved capacity.
2. **`append`** — dynamic slice building.
3. **Automatic growth** — append reslices when capacity runs out.
4. **Derived count** — healthy = the total − failed difference.
5. **Random simulation** — a healthy flag via `rand.Intn`.