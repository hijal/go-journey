# cicd-job-grouping

Go-তে **anonymous struct slice + `map[string][]string` grouping** শেখার ছোট example — CI/CD job-গুলো status-ভিত্তিক গ্রুপ করা (deterministic sorted report)।

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

### Lines 3–7

```go
import (
	"fmt"
	"maps"
	"slices"
)
```

- `fmt` — `Printf`, `Println`।
- `maps` — `Keys` (Go 1.21+)।
- `slices` — `Sorted` (Go 1.21+)।

### Lines 9–20

```go
func main() {
	jobs := []struct {
		Name   string
		Status string
	}{
		{"build-api", "passed"},
		{"lint", "failed"},
		{"unit-tests", "passed"},
		{"integration-tests", "failed"},
		{"deploy-staging", "pending"},
		{"build-web", "passed"},
	}
```

**Anonymous struct slice** — `struct{Name, Status string}` field inline (টাইপ-নাম declare না-করেই)। ৬টা job: passed×3, failed×2, pending×1।

### Lines 22–26

```go
	byStatus := make(map[string][]string)

	for _, j := range jobs {
		byStatus[j.Status] = append(byStatus[j.Status], j.Name)
	}
```

**Grouping idiom:**

- `map[string][]string` — status → job-name list।
- `append(byStatus[j.Status], j.Name)` — zero-value-idempotent: missing status-এ প্রথম `append` খালি slice-এ (missing-key result nil-এ append ঠিকই কাজ করে)।

### Lines 28–34

```go
	for _, status := range slices.Sorted(maps.Keys(byStatus)) {
		fmt.Printf("%s (%d):\n", status, len(byStatus[status]))

		for _, name := range byStatus[status] {
			fmt.Println("  -", name)
		}
	}
```

**Deterministic grouped report:**

- `slices.Sorted(maps.Keys(...))` — status আলফা-সর্ট (failed, passed, pending — random map-order নয়)।
- `len(...)` per-group + inner-loop group members।

---

## Expected Output

```
failed (2):
  - lint
  - integration-tests
passed (3):
  - build-api
  - unit-tests
  - build-web
pending (1):
  - deploy-staging
```

## মূল শিক্ষা / Key Takeaways

1. **Anonymous struct slice** — `[]struct{...}` inline।
2. **Grouping map** — `map[string][]string` append-idiom।
3. **Zero-value append** — missing-key-এ append ঠিকঠাক।
4. **Sorted keys** — `slices.Sorted(maps.Keys(...))` deterministic।
5. **Nested report loop** — group → members।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Lines 3–7

```go
import (
	"fmt"
	"maps"
	"slices"
)
```

- `fmt` — for `Printf`, `Println`.
- `maps` — for `Keys` (Go 1.21+).
- `slices` — for `Sorted` (Go 1.21+).

### Lines 9–20

```go
func main() {
	jobs := []struct {
		Name   string
		Status string
	}{
		{"build-api", "passed"},
		{"lint", "failed"},
		{"unit-tests", "passed"},
		{"integration-tests", "failed"},
		{"deploy-staging", "pending"},
		{"build-web", "passed"},
	}
```

An **anonymous struct slice** — `struct{Name, Status string}` fields inline (no named type). 6 jobs: passed ×3, failed ×2, pending ×1.

### Lines 22–26

```go
	byStatus := make(map[string][]string)

	for _, j := range jobs {
		byStatus[j.Status] = append(byStatus[j.Status], j.Name)
	}
```

**The grouping idiom:**

- `map[string][]string` — status → job-name list.
- `append(byStatus[j.Status], j.Name)` — zero-value-tolerant: a first append into a missing key (nil slice) works fine.

### Lines 28–34

```go
	for _, status := range slices.Sorted(maps.Keys(byStatus)) {
		fmt.Printf("%s (%d):\n", status, len(byStatus[status]))

		for _, name := range byStatus[status] {
			fmt.Println("  -", name)
		}
	}
```

**Deterministic grouped report:**

- `slices.Sorted(maps.Keys(...))` — statuses alpha-sorted (failed, passed, pending — not the random map order).
- `len(...)` per group + an inner loop for the members.

---

## Expected Output

```
failed (2):
  - lint
  - integration-tests
passed (3):
  - build-api
  - unit-tests
  - build-web
pending (1):
  - deploy-staging
```

## Key Takeaways

1. **Anonymous struct slice** — `[]struct{...}` inline.
2. **Grouping map** — the `map[string][]string` append idiom.
3. **Zero-value append** — append into a missing key works.
4. **Sorted keys** — `slices.Sorted(maps.Keys(...))` for determinism.
5. **Nested report loop** — group → members.