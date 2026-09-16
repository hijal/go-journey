# event-queue-type-switch

Go-তে **`[]any` heterogeneous queue + type switch (`switch ev.(type)`) + `%T` + default dead-letter** শেখার ছোট example — event dispatcher।

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

### Line 3

```go
import "fmt"
```

`fmt` — `Printf`।

### Lines 5–18

```go
type OrderPlaced struct {
	OrderID string
	Amount  int64
}

type PaymentFailed struct {
	OrderID string
	Reason  string
}

type OrderShipped struct {
	OrderID string
	Courier string
}
```

৩টা ধরণের event — একে-অপরের থেকে আলাদা struct।

### Lines 20–31

```go
func handle(ev any) {
	switch e := ev.(type) {
	case OrderPlaced:
		fmt.Printf("reserve stock for %s (৳%.2f)\n", e.OrderID, float64(e.Amount)/100)
	case PaymentFailed:
		fmt.Printf("alert ops: %s failed — %s\n", e.OrderID, e.Reason)
	case OrderShipped:
		fmt.Printf("email tracking link for %s via %s\n", e.OrderID, e.Courier)
	default:
		fmt.Printf("unknown event %T — route to dead-letter queue\n", ev)
	}
}
```

**Type switch** — `switch e := ev.(type)`:

- প্রতিটি `case`-এ `e` **typed value** (`e.OrderID`, `e.Reason` — field access-এর প্রয়োজনীয় narrow)।
- `default` — অজানা type → **dead-letter** routing; `%T` dynamic type print (`string`)।
- `ev any` — যেকোনো type নেয়।

### Lines 33–39

```go
	queue := []any{
		OrderPlaced{OrderID: "ORD-7", Amount: 42000},
		PaymentFailed{OrderID: "ORD-8", Reason: "card declined"},
		OrderShipped{OrderID: "ORD-7", Courier: "Pathao"},
		"bad-message",
	}
```

**Heterogeneous queue** — `[]any`-তে নানা struct + plain string।

### Lines 41–44

```go
	for _, ev := range queue {
		handle(ev)
	}
```

প্রতিটি event dispatch — type অনুযায়ী branch।

---

## Expected Output

```
reserve stock for ORD-7 (৳420.00)
alert ops: ORD-8 failed — card declined
email tracking link for ORD-7 via Pathao
unknown event string — route to dead-letter queue
```

## মূল শিক্ষা / Key Takeaways

1. **`[]any`** — mixed-type queue/stream।
2. **Type switch** — `switch ev.(type)` + typed `e`।
3. **`%T`** — dynamic type print।
4. **Default dead-letter** — unknown event degrade gracefully।
5. **Cents formatting** — `int64/100` + `%.2f`।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Line 3

```go
import "fmt"
```

`fmt` — for `Printf`.

### Lines 5–18

```go
type OrderPlaced struct {
	OrderID string
	Amount  int64
}

type PaymentFailed struct {
	OrderID string
	Reason  string
}

type OrderShipped struct {
	OrderID string
	Courier string
}
```

3 kinds of events — distinct structs.

### Lines 20–31

```go
func handle(ev any) {
	switch e := ev.(type) {
	case OrderPlaced:
		fmt.Printf("reserve stock for %s (৳%.2f)\n", e.OrderID, float64(e.Amount)/100)
	case PaymentFailed:
		fmt.Printf("alert ops: %s failed — %s\n", e.OrderID, e.Reason)
	case OrderShipped:
		fmt.Printf("email tracking link for %s via %s\n", e.OrderID, e.Courier)
	default:
		fmt.Printf("unknown event %T — route to dead-letter queue\n", ev)
	}
}
```

**The type switch** — `switch e := ev.(type)`:

- In each `case`, `e` is the **typed value** (`e.OrderID`, `e.Reason` — the narrow field access you need).
- `default` — unknown types route to **dead-letter**; `%T` prints the dynamic type (`string`).
- `ev any` — accepts any type.

### Lines 33–39

```go
	queue := []any{
		OrderPlaced{OrderID: "ORD-7", Amount: 42000},
		PaymentFailed{OrderID: "ORD-8", Reason: "card declined"},
		OrderShipped{OrderID: "ORD-7", Courier: "Pathao"},
		"bad-message",
	}
```

**A heterogeneous queue** — `[]any` mixing structs and a plain string.

### Lines 41–44

```go
	for _, ev := range queue {
		handle(ev)
	}
```

Each event dispatched — branching by type.

---

## Expected Output

```
reserve stock for ORD-7 (৳420.00)
alert ops: ORD-8 failed — card declined
email tracking link for ORD-7 via Pathao
unknown event string — route to dead-letter queue
```

## Key Takeaways

1. **`[]any`** — a mixed-type queue/stream.
2. **Type switch** — `switch ev.(type)` + the typed `e`.
3. **`%T`** — dynamic type printing.
4. **Default dead-letter** — graceful degradation for unknown events.
5. **Cents formatting** — `int64/100` + `%.2f`.