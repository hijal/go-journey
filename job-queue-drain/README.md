# job-queue-drain

Go-তে **`slices.DeleteFunc` in-place drain + positional struct** শেখার ছোট example — failed job-গুলো queue থেকে বাদ দিয়ে retry-list আলাদা।

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
	"slices"
)
```

- `fmt` — `Printf`, `Println`।
- `slices` — `DeleteFunc` (Go 1.21+)।

### Lines 8–12

```go
type Job struct {
	ID      string
	Payload string
	Failed  bool
}
```

`Job` struct — ID, Payload, Failed flag।

### Lines 14–21

```go
func main() {
	queue := []Job{
		{"J1", "send-email:1123", false},
		{"J2", "resize-image:8891", true},
		{"J3", "charge-card:5521", false},
		{"J4", "generate-pdf:3320", true},
		{"J5", "sync-inventory:7788", false},
	}
```

5টা job — **positional struct literal** (field-name ছাড়া)। J2 আর J4-র `Failed: true`।

### Lines 23–25

```go
	queue = slices.DeleteFunc(queue, func(j Job) bool {
		return j.Failed
	})
```

**In-place delete:**

- `slices.DeleteFunc(queue, predicate)` — predicate `true` দিলে element সরায়; বাকিগুলো **একই backing array-তে** compact হয় (অনেকটা drain)।
- Return-টা resliced-modified slice — আবার `queue`-তে assign করা হলো (লেন-বদল হওয়ায় প্রয়োজন)।
- Element-order maintain হয় — J1, J3, J5।

### Lines 27–30

```go
	fmt.Println("remaining in queue:")
	for _, j := range queue {
		fmt.Printf("  %s -> %s\n", j.ID, j.Payload)
	}
```

Remaining queue-টা print — `%s -> %s` format।

### Lines 32–33

```go
	failedIDs := []string{"J2", "J4"}
	fmt.Println("to retry later:", failedIDs)
```

- Failed IDs — আলাদা slice-এ (retry-পরবর্তী)।
- **Drain pattern:** queue থেকে বাদ, তবে retry-র জন্য রেকর্ড রাখা।

---

## Expected Output

```
remaining in queue:
  J1 -> send-email:1123
  J3 -> charge-card:5521
  J5 -> sync-inventory:7788
to retry later: [J2 J4]
```

## মূল শিক্ষা / Key Takeaways

1. **`slices.DeleteFunc`** — predicate-based in-place removal (Go 1.21+)।
2. **Re-assign return** — DeleteFunc-এর result আবার ব্যবহার কর (len বদলায়)।
3. **Positional struct literal** — `{"J1", "send-email:1123", false}`।
4. **Drain pattern** — বাদ দেওয়া + retry-list আলাদা রাখা।
5. **Order preserved** — delete-পরেও original sequence।

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
	"slices"
)
```

- `fmt` — for `Printf`, `Println`.
- `slices` — for `DeleteFunc` (Go 1.21+).

### Lines 8–12

```go
type Job struct {
	ID      string
	Payload string
	Failed  bool
}
```

`Job` struct — ID, Payload, and a Failed flag.

### Lines 14–21

```go
func main() {
	queue := []Job{
		{"J1", "send-email:1123", false},
		{"J2", "resize-image:8891", true},
		{"J3", "charge-card:5521", false},
		{"J4", "generate-pdf:3320", true},
		{"J5", "sync-inventory:7788", false},
	}
```

5 jobs — **positional struct literals** (no field names). J2 and J4 have `Failed: true`.

### Lines 23–25

```go
	queue = slices.DeleteFunc(queue, func(j Job) bool {
		return j.Failed
	})
```

**In-place delete:**

- `slices.DeleteFunc(queue, predicate)` — removes elements for which the predicate returns `true`; the survivors are compacted into the **same backing array** (like a drain).
- The return is the resliced modified slice — reassigned back to `queue` (needed because the length changed).
- Element order stays stable — J1, J3, J5.

### Lines 27–30

```go
	fmt.Println("remaining in queue:")
	for _, j := range queue {
		fmt.Printf("  %s -> %s\n", j.ID, j.Payload)
	}
```

Prints the remaining queue — `%s -> %s` format.

### Lines 32–33

```go
	failedIDs := []string{"J2", "J4"}
	fmt.Println("to retry later:", failedIDs)
```

- Failed IDs — kept in a separate slice (for later retry).
- **Drain pattern:** out of the queue, but recorded for retry.

---

## Expected Output

```
remaining in queue:
  J1 -> send-email:1123
  J3 -> charge-card:5521
  J5 -> sync-inventory:7788
to retry later: [J2 J4]
```

## Key Takeaways

1. **`slices.DeleteFunc`** — predicate-based in-place removal (Go 1.21+).
2. **Re-assign the return** — use DeleteFunc's result again (the length changes).
3. **Positional struct literal** — `{"J1", "send-email:1123", false}`.
4. **Drain pattern** — dropped from the queue, but kept in a retry list.
5. **Order preserved** — the original sequence survives the delete.