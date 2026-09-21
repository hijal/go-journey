# order-service-di

Go-তে **Dependency Injection (constructor-injected interface dependencies) + `%w` error wrapping + duplicate detection** শেখার ছোট example — order service (in-memory repo + notifier)।

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
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
)
```

- `context` — `Background` + ctx propagation।
- `errors` — `New`।
- `sync` — `Mutex`।
- `fmt` — `Errorf`, `Sprintf`, `Printf`।
- `log` — `Fatal`, `Printf`।

### Lines 11–17

```go
type Order struct {
	ID         string
	CustomerID string
	Email      string
	TotalCents int64
	Status     string
}
```

Domain model — order সারি।

### Lines 19–25

```go
type OrderRepository interface {
	Save(ctx context.Context, o Order) error
}

type Notifier interface {
	Notify(ctx context.Context, to, message string) error
}
```

**DI interfaces** — `OrderService` concrete-নির্দিষ্ট নয়, interface dependency চায়।

### Lines 27–34

```go
type OrderService struct {
	repo     OrderRepository
	notifier Notifier
}

func NewOrderService(repo OrderRepository, notifier Notifier) *OrderService {
	return &OrderService{repo: repo, notifier: notifier}
}
```

**Constructor injection** — dependency constructor-constraint মাধ্যমে ঢোকা (মানক DI pattern)।

### Line 36

```go
var ErrEmptyOrder = errors.New("order total must be positive")
```

Sentinel error।

### Lines 38–53

```go
func (s *OrderService) PlaceOrder(ctx context.Context, o Order) error {
	if o.TotalCents <= 0 {
		return ErrEmptyOrder
	}

	o.Status = "placed"
	if err := s.repo.Save(ctx, o); err != nil {
		return fmt.Errorf("save order %s: %w", o.ID, err)
	}
	msg := fmt.Sprintf("Your order %s has been placed.", o.ID)

	if err := s.notifier.Notify(ctx, o.Email, msg); err != nil {
		log.Printf("notify %s: %v", o.Email, err)
	}
	return nil
}
```

- Validate → sentinel error।
- `Status` set → `repo.Save`; fail-এ `%w` wrap।
- Notify fail **non-fatal** — log-only (order save হয়েছে তাই)।

### Lines 55–73

```go
type MemoryOrderRepo struct {
	mu     sync.Mutex
	orders map[string]Order
}

func (r *MemoryOrderRepo) Save(_ context.Context, o Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.orders[o.ID]; exists {
		return fmt.Errorf("order %s already exist", o.ID)
	}
	r.orders[o.ID] = o
	return nil
}
```

**Concurrent-safe in-memory repo** — mutex + duplicate check (২য় বারে error)।

### Lines 75–80

```go
type ConsoleNotifier struct{}

func (ConsoleNotifier) Notify(_ context.Context, to, message string) error {
	fmt.Printf("Email to=%s: %s\n", to, message)
	return nil
}
```

Stub notifier — console-এ "email" print।

### Lines 82–94

```go
	svc := NewOrderService(NewMemoryOrderRepo(), ConsoleNotifier{})
	ctx := context.Background()

	order := Order{ID: "ord-1001", CustomerID: "c-7", Email: "example@example.com", TotalCents: 1466}

	if err := svc.PlaceOrder(ctx, order); err != nil {
		log.Fatal(err)
	}

	if err := svc.PlaceOrder(ctx, order); err != nil {
		fmt.Println("error:", err)
	}
```

- প্রথম call: saves + notifies।
- দ্বিতীয় call: duplicate → `error: save order ord-1001: order ord-1001 already exist`।

---

## Expected Output

```
Email to=example@example.com: Your order ord-1001 has been placed.
error: save order ord-1001: order ord-1001 already exist
```

## মূল শিক্ষা / Key Takeaways

1. **Constructor DI** — dependencies interface-রূপে inject।
2. **Interface-based testing** — repo/notifier সহজ replace।
3. **`%w` wrap** — error context + cause chain।
4. **Non-fatal deps** — notify fail-এ log-only।
5. **`sync.Mutex` + map** — concurrent-safe store।

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
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
)
```

- `context` — `Background` + ctx propagation.
- `errors` — `New`.
- `sync` — `Mutex`.
- `fmt` — `Errorf`, `Sprintf`, `Printf`.
- `log` — `Fatal`, `Printf`.

### Lines 11–17

```go
type Order struct {
	ID         string
	CustomerID string
	Email      string
	TotalCents int64
	Status     string
}
```

The domain model — the order fields.

### Lines 19–25

```go
type OrderRepository interface {
	Save(ctx context.Context, o Order) error
}

type Notifier interface {
	Notify(ctx context.Context, to, message string) error
}
```

**DI interfaces** — `OrderService` depends on interfaces, not concrete types.

### Lines 27–34

```go
type OrderService struct {
	repo     OrderRepository
	notifier Notifier
}

func NewOrderService(repo OrderRepository, notifier Notifier) *OrderService {
	return &OrderService{repo: repo, notifier: notifier}
}
```

**Constructor injection** — dependencies are injected through the constructor (the canonical DI pattern).

### Line 36

```go
var ErrEmptyOrder = errors.New("order total must be positive")
```

A sentinel error.

### Lines 38–53

```go
func (s *OrderService) PlaceOrder(ctx context.Context, o Order) error {
	if o.TotalCents <= 0 {
		return ErrEmptyOrder
	}

	o.Status = "placed"
	if err := s.repo.Save(ctx, o); err != nil {
		return fmt.Errorf("save order %s: %w", o.ID, err)
	}
	msg := fmt.Sprintf("Your order %s has been placed.", o.ID)

	if err := s.notifier.Notify(ctx, o.Email, msg); err != nil {
		log.Printf("notify %s: %v", o.Email, err)
	}
	return nil
}
```

- Validate → sentinel error.
- `Status` is set → `repo.Save`; on failure it's wrapped with `%w`.
- Notify failures are **non-fatal** — log-only (the order was already saved).

### Lines 55–73

```go
type MemoryOrderRepo struct {
	mu     sync.Mutex
	orders map[string]Order
}

func (r *MemoryOrderRepo) Save(_ context.Context, o Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.orders[o.ID]; exists {
		return fmt.Errorf("order %s already exist", o.ID)
	}
	r.orders[o.ID] = o
	return nil
}
```

**A concurrent-safe in-memory repo** — a mutex + duplicate check (errors on the 2nd call).

### Lines 75–80

```go
type ConsoleNotifier struct{}

func (ConsoleNotifier) Notify(_ context.Context, to, message string) error {
	fmt.Printf("Email to=%s: %s\n", to, message)
	return nil
}
```

A stub notifier — prints the "email" to the console.

### Lines 82–94

```go
	svc := NewOrderService(NewMemoryOrderRepo(), ConsoleNotifier{})
	ctx := context.Background()

	order := Order{ID: "ord-1001", CustomerID: "c-7", Email: "example@example.com", TotalCents: 1466}

	if err := svc.PlaceOrder(ctx, order); err != nil {
		log.Fatal(err)
	}

	if err := svc.PlaceOrder(ctx, order); err != nil {
		fmt.Println("error:", err)
	}
```

- First call: saves + notifies.
- Second call: duplicate → `error: save order ord-1001: order ord-1001 already exist`.

---

## Expected Output

```
Email to=example@example.com: Your order ord-1001 has been placed.
error: save order ord-1001: order ord-1001 already exist
```

## Key Takeaways

1. **Constructor DI** — dependencies injected as interfaces.
2. **Interface-based testing** — repo/notifier swap easily.
3. **`%w` wrapping** — error context + an intact cause chain.
4. **Non-fatal deps** — notify failures are log-only.
5. **`sync.Mutex` + map** — a concurrent-safe store.