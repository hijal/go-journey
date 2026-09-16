# analytics-interface-embedding

Go-তে **interface embedding (Composite interface) + compile-time satisfaction (`var _`) + concrete method access** শেখার ছোট example — metrics reporter।

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

`fmt` — `Println`।

### Lines 5–16

```go
type Counter interface {
	Incr(name string)
}

type Gauge interface {
	Set(name string, value float64)
}

type Reporter interface {
	Counter
	Gauge
}
```

**Interface embedding** — `Reporter` = `Counter` + `Gauge` (দুটো method-set একসাথে)। Composite interface।

### Lines 18–34

```go
type memReporter struct {
	counters map[string]int
	gauges   map[string]float64
}

func newMemReporter() *memReporter { ... }

func (m *memReporter) Incr(name string)     { m.counters[name]++ }
func (m *memReporter) Set(name string, v float64) { m.gauges[name] = v }
func (m *memReporter) Snapshot() (map[string]int, map[string]float64) {
	return m.counters, m.gauges
}
```

`*memReporter` — `Incr` + `Set` (interface-এর method) + **additional** `Snapshot` (concrete-মাত্র)।

### Line 36

```go
func recordCheckoutAttempt(c Counter) { c.Incr("checkout_attempts") }
```

Counter interface param — শুধু `Incr` দরকার।

### Lines 38–47

```go
	rep := newMemReporter()

	recordCheckoutAttempt(rep)
	recordCheckoutAttempt(rep)
	recordCheckoutAttempt(rep)

	var g Gauge = rep
	g.Set("active_carts", 12.5)

	var _ Reporter = rep
```

- `recordCheckoutAttempt(rep)` — `*memReporter` → `Counter` (structural), ৩বার incr।
- `var g Gauge = rep` — interface-এর scope-এ শুধু Set।
- `var _ Reporter = rep` — blank-identifier compile-time check (Reporter satisfies)।

### Lines 50–53

```go
	counters, gauges := rep.Snapshot()

	fmt.Println("counters:", counters)
	fmt.Println("gauges:   ", gauges)
```

**Concrete `rep`** — `Snapshot()` interface-এ নাই, তাই concrete type থেকেই call (interface-এ নেই-জিনিস interface-value-তে call করা যায় না → compile error)।

---

## Expected Output

```
counters: map[checkout_attempts:3]
gauges:   map[active_carts:12.5]
```

## মূল শিক্ষা / Key Takeaways

1. **Interface embedding** — composite method-set।
2. **`var _ T = v`** — compile-time satisfaction check।
3. **Interface vs concrete** — interface-এ নেই-জিনিস concrete-এ-ই call।
4. **Structural typing** — no `implements` keyword।
5. **Map closure + methods** — `Incr`/`Set` state mutate।

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

`fmt` — for `Println`.

### Lines 5–16

```go
type Counter interface {
	Incr(name string)
}

type Gauge interface {
	Set(name string, value float64)
}

type Reporter interface {
	Counter
	Gauge
}
```

**Interface embedding** — `Reporter` = `Counter` + `Gauge` (both method sets in one). This is a composite interface.

### Lines 18–34

```go
type memReporter struct {
	counters map[string]int
	gauges   map[string]float64
}

func newMemReporter() *memReporter { ... }

func (m *memReporter) Incr(name string)     { m.counters[name]++ }
func (m *memReporter) Set(name string, v float64) { m.gauges[name] = v }
func (m *memReporter) Snapshot() (map[string]int, map[string]float64) {
	return m.counters, m.gauges
}
```

`*memReporter` — `Incr` + `Set` (the interface methods) plus an **extra** `Snapshot` (concrete-only).

### Line 36

```go
func recordCheckoutAttempt(c Counter) { c.Incr("checkout_attempts") }
```

A Counter-interface param — only `Incr` is needed.

### Lines 38–47

```go
	rep := newMemReporter()

	recordCheckoutAttempt(rep)
	recordCheckoutAttempt(rep)
	recordCheckoutAttempt(rep)

	var g Gauge = rep
	g.Set("active_carts", 12.5)

	var _ Reporter = rep
```

- `recordCheckoutAttempt(rep)` — `*memReporter` → `Counter` (structurally), incremented 3 times.
- `var g Gauge = rep` — within the interface scope, only `Set` is visible.
- `var _ Reporter = rep` — a blank-identifier compile-time check (Reporter is satisfied).

### Lines 50–53

```go
	counters, gauges := rep.Snapshot()

	fmt.Println("counters:", counters)
	fmt.Println("gauges:   ", gauges)
```

**The concrete `rep`** — `Snapshot()` isn't on the interface, so it's called on the concrete type (methods absent from an interface can't be called through the interface — a compile error).

---

## Expected Output

```
counters: map[checkout_attempts:3]
gauges:   map[active_carts:12.5]
```

## Key Takeaways

1. **Interface embedding** — a composite method set.
2. **`var _ T = v`** — compile-time satisfaction check.
3. **Interface vs concrete** — call concrete-only methods on the concrete value.
4. **Structural typing** — no `implements` keyword.
5. **Map closures + methods** — `Incr`/`Set` mutate state.