# config-merge-copy

Go-তে **`slices.Clone` copy semantics** শেখার ছোট example — base config-এর ক্ষতি এড়িয়ে replica-set merge করা।

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
- `slices` — `Clone`, `Equal` (Go 1.21+)।

### Lines 8–14

```go
func buildReplicaSet(base []string, replicas int) []string {
	merged := slices.Clone(base)
	for i := range replicas {
		merged = append(merged, fmt.Sprintf("pod-%02d", i))
	}
	return merged
}
```

**Replica builder:**

- `merged := slices.Clone(base)` — base-এর **নিজস্ব copy**। `clone` শুধু slice-header নয়, পুরো backing array-ও copy করে — অর্থাৎ `merged`-এ যোগ/পরিবর্তন **কখনো base-কে বদলাবে না**।
- `for i := range replicas` — `0..replicas-1`।
- `fmt.Sprintf("pod-%02d", i)` — **zero-padded**: `pod-00`, `pod-01`… (২ ডিজিট)।
- `append` + return — merged slice।

> **কেন Clone দরকার?** `append(defaultConfig, ...)` capacity-থাকলে base-এর পেছনে লিখতো এবং সেটাকে mutate করত। Clone সেই ভয় দূর করে — `defaultConfig` untouched থাকবে।

### Lines 16–17

```go
func main() {
	defaultConfig := []string{"--port=8080", "--log-level=info"}
```

Base config — 2 flag। এর integrative-তা test হবে।

### Lines 19–20

```go
	prod := buildReplicaSet(defaultConfig, 3)
	prod[0] = "--port=9090"
```

- `prod` — base copy + ৩টা pod flag।
- `prod[0] = "--port=9090"` — production-এ পোর্ট বদল — Clone-এর সুবাদে base untouched।

### Lines 22–24

```go
	fmt.Println("default config:", defaultConfig)
	fmt.Println("production:", prod)
	fmt.Println("equal?", slices.Equal(defaultConfig, []string{"--port=8080", "--log-level=info"}))
```

- `defaultConfig` — এখনো `--port=8080` (mutate হয়নি)।
- `prod` — `--port=9090` + pod flags।
- `slices.Equal` — content-wise compare → `true`।

---

## Expected Output

```
default config: [--port=8080 --log-level=info]
production: [--port=9090 --log-level=info pod-00 pod-01 pod-02]
equal? true
```

## মূল শিক্ষা / Key Takeaways

1. **`slices.Clone`** — backing array সহ full copy — base-সুরক্ষা।
2. **Shared-backing hazard** — `append`-এ direct base pass avoided।
3. **Zero-pad** — `%02d` format।
4. **Range-over-int** — `for i := range replicas` (Go 1.22+)।
5. **`slices.Equal`** — content-wise slice compare।

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
- `slices` — for `Clone`, `Equal` (Go 1.21+).

### Lines 8–14

```go
func buildReplicaSet(base []string, replicas int) []string {
	merged := slices.Clone(base)
	for i := range replicas {
		merged = append(merged, fmt.Sprintf("pod-%02d", i))
	}
	return merged
}
```

**Replica builder:**

- `merged := slices.Clone(base)` — a **private copy** of base. `Clone` copies the whole backing array, not just the slice header — so appends/mutations to `merged` will **never change base**.
- `for i := range replicas` — `0..replicas-1`.
- `fmt.Sprintf("pod-%02d", i)` — **zero-padded**: `pod-00`, `pod-01`… (2 digits).
- `append` + return — the merged slice.

> **Why Clone?** With spare capacity, `append(defaultConfig, ...)` would write right after base and mutate it. Clone removes that fear — `defaultConfig` stays untouched.

### Lines 16–17

```go
func main() {
	defaultConfig := []string{"--port=8080", "--log-level=info"}
```

Base config — 2 flags. Its integrity is what gets tested.

### Lines 19–20

```go
	prod := buildReplicaSet(defaultConfig, 3)
	prod[0] = "--port=9090"
```

- `prod` — a base copy + 3 pod flags.
- `prod[0] = "--port=9090"` — changing the port in production; thanks to Clone, base stays untouched.

### Lines 22–24

```go
	fmt.Println("default config:", defaultConfig)
	fmt.Println("production:", prod)
	fmt.Println("equal?", slices.Equal(defaultConfig, []string{"--port=8080", "--log-level=info"}))
```

- `defaultConfig` — still `--port=8080` (not mutated).
- `prod` — `--port=9090` + pod flags.
- `slices.Equal` — content-wise compare → `true`.

---

## Expected Output

```
default config: [--port=8080 --log-level=info]
production: [--port=9090 --log-level=info pod-00 pod-01 pod-02]
equal? true
```

## Key Takeaways

1. **`slices.Clone`** — full copy including the backing array — base protection.
2. **Shared-backing hazard** — avoided by not passing base straight into `append`.
3. **Zero-pad** — the `%02d` format.
4. **Range-over-int** — `for i := range replicas` (Go 1.22+).
5. **`slices.Equal`** — content-wise slice comparison.