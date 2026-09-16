# job-worker-interface

Go-তে **interface-driven design + sentinel error (`ErrNotFound`) + `errors.Is` + `%w` wrap** শেখার ছোট example — failed-order notifier।

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
	"errors"
	"fmt"
)
```

- `errors` — `New`, `Is`।
- `fmt` — `Println`, `Errorf`।

### Lines 8–13

```go
type OrderStore interface {
	Status(id string) (string, error)
	SaveStatus(id, status string) error
}

var ErrNotFound = errors.New("order not found")
```

- **`OrderStore` interface** — consumer শুধু logic-এ মনোযোগ দেয়।
- **Sentinel error** `ErrNotFound` — specific cause classifier।

### Lines 15–30

```go
type fakeStore struct {
	db map[string]string
}

func (f *fakeStore) Status(id string) (string, error) {
	status, ok := f.db[id]
	if !ok {
		return "", fmt.Errorf("lookup %s: %w", id, ErrNotFound)
	}
	return status, nil
}

func (f *fakeStore) SaveStatus(id, status string) error {
	f.db[id] = status
	return nil
}
```

**Implementation** — `fakeStore` satisfies `OrderStore` (কোনো explicit implements declaration লাগে না; structural)।

- `Status`: comma-ok; miss → `%w` দিয়ে `ErrNotFound` wrap (context সহ but cause-preserved)।
- `SaveStatus`: map update।

### Lines 32–55

```go
func notifyFailedOrders(store OrderStore, ids []string) {
	for _, id := range ids {
		status, err := store.Status(id)

		if err != nil {
			if errors.Is(err, ErrNotFound) {
				fmt.Println("skip", id, "(not found)")
				continue
			}
			fmt.Println("skip", id, "error:", err)
			continue
		}

		if status == "FAILED" {
			if err := store.SaveStatus(id, "NOTIFIED"); err != nil {
				fmt.Println("save failed for", id, ":", err)
				continue
			}
			fmt.Println("notified customer for", id)
		} else {
			fmt.Println("no action for", id, "status:", status)
		}
	}
}
```

**Consumer logic:**

- `OrderStore` param — যেকোনো implementation plug-in।
- `errors.Is(err, ErrNotFound)` — wrapped-ই-হোক cause সনাক্ত => "not found" skip।
- FAILED হলে `NOTIFIED`-তে update + notify; নাহলে no-action।

### Lines 57–64

```go
	store := &fakeStore{db: map[string]string{
		"ORD-1": "FAILED",
		"ORD-2": "SETTLED",
	}}

	notifyFailedOrders(store, []string{"ORD-1", "ORD-2", "ORD-9"})
```

৩টা id: FAILED, SETTLED, not-found।

---

## Expected Output

```
notified customer for ORD-1
no action for ORD-2 status: SETTLED
skip ORD-9 (not found)
```

## মূল শিক্ষা / Key Takeaways

1. **Interface consumer** — `OrderStore` param, যেকোনো implementation।
2. **Sentinel error** + **`%w` wrap** — cause-preserving context।
3. **`errors.Is`** — wrapped error-এও classify।
4. **Structural satisfaction** — No explicit `implements` keyword।
5. **State transition** — FAILED → NOTIFIED।

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
	"errors"
	"fmt"
)
```

- `errors` — for `New`, `Is`.
- `fmt` — for `Println`, `Errorf`.

### Lines 8–13

```go
type OrderStore interface {
	Status(id string) (string, error)
	SaveStatus(id, status string) error
}

var ErrNotFound = errors.New("order not found")
```

- The **`OrderStore` interface** — the consumer focuses on logic alone.
- The **sentinel error** `ErrNotFound` — the specific cause classifier.

### Lines 15–30

```go
type fakeStore struct {
	db map[string]string
}

func (f *fakeStore) Status(id string) (string, error) {
	status, ok := f.db[id]
	if !ok {
		return "", fmt.Errorf("lookup %s: %w", id, ErrNotFound)
	}
	return status, nil
}

func (f *fakeStore) SaveStatus(id, status string) error {
	f.db[id] = status
	return nil
}
```

**The implementation** — `fakeStore` satisfies `OrderStore` (no explicit implements declaration; it's structural).

- `Status`: a comma-ok lookup; a miss wraps `ErrNotFound` with `%w` (context added, cause preserved).
- `SaveStatus`: map update.

### Lines 32–55

```go
func notifyFailedOrders(store OrderStore, ids []string) {
	for _, id := range ids {
		status, err := store.Status(id)

		if err != nil {
			if errors.Is(err, ErrNotFound) {
				fmt.Println("skip", id, "(not found)")
				continue
			}
			fmt.Println("skip", id, "error:", err)
			continue
		}

		if status == "FAILED" {
			if err := store.SaveStatus(id, "NOTIFIED"); err != nil {
				fmt.Println("save failed for", id, ":", err)
				continue
			}
			fmt.Println("notified customer for", id)
		} else {
			fmt.Println("no action for", id, "status:", status)
		}
	}
}
```

**The consumer logic:**

- The `OrderStore` param — any implementation plugs in.
- `errors.Is(err, ErrNotFound)` — detects the cause even when wrapped → "not found" skip.
- FAILED gets updated to `NOTIFIED` + notified; otherwise no action.

### Lines 57–64

```go
	store := &fakeStore{db: map[string]string{
		"ORD-1": "FAILED",
		"ORD-2": "SETTLED",
	}}

	notifyFailedOrders(store, []string{"ORD-1", "ORD-2", "ORD-9"})
```

3 ids: FAILED, SETTLED, not-found.

---

## Expected Output

```
notified customer for ORD-1
no action for ORD-2 status: SETTLED
skip ORD-9 (not found)
```

## Key Takeaways

1. **Interface consumer** — an `OrderStore` param, any implementation works.
2. **Sentinel error** + **`%w` wrap** — cause-preserving context.
3. **`errors.Is`** — classifies even wrapped errors.
4. **Structural satisfaction** — no `implements` keyword.
5. **State transition** — FAILED → NOTIFIED.