# failed-job-csv-report

Go-তে **`strings.Fields` + `strings.Cut` key=value parse + filtered CSV building** শেখার ছোট example — log-to-CSV pipeline।

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
	"strings"
)
```

- `fmt` — `Println`, `Print`।
- `strings` — `Cut`, `Fields`, `Split`, `Contains`, `Builder`, `TrimSpace`।

### Lines 8–15

```go
func extractField(line, key string) (string, bool) {
	for _, field := range strings.Fields(line) {
		if k, v, ok := strings.Cut(field, "="); ok && k == key {
			return v, true
		}
	}
	return "", false
}
```

**Generic key=value extractor:**

- `strings.Fields(line)` — whitespace-tokenize লাইন।
- `strings.Cut(field, "=")` — `key=value` থেকে key ও value আলাদা।
- key match হলে value return (true)। না-হলে false।

*("duration=640" → Cut → k="duration", v="640"।)*

### Lines 17–21

```go
	jobLog := `2026-01-05T09:00:01 job=invoice-gen status=done duration=812
2026-01-05T09:00:04 job=report-mail status=failed duration=640
2026-01-05T09:00:09 job=invoice-gen status=failed duration=905
2026-01-05T09:00:12 job=db-cleanup status=done duration=430`
```

৪ লাইন job log — ২টা done, ২টা failed।

### Lines 23–26

```go
	var report strings.Builder
	report.WriteString("job,duration_ms\n")
	for _, line := range strings.Split(jobLog, "\n") {
		if !strings.Contains(line, "status=failed") {
			continue
		}
```

- `strings.Builder` — efficient CSV concat।
- Header লিখে, প্রতিটি line-এ `Contains("status=failed")` filter — done গুলো skip।

### Lines 30–36

```go
		job, jobOK := extractField(line, "job")
		duration, durationOK := extractField(line, "duration")

		if !jobOK || !durationOK {
			continue
		}
		report.WriteString(job)
		report.WriteString(",")
		report.WriteString(duration)
		report.WriteString("\n")
```

- extractField দিয়ে `job` ও `duration` বের করে CSV row তৈরি।
- field missing → skip (safety)।

### Lines 38–40

```go
	fmt.Println("failed jobs (CSV):")
	fmt.Print(strings.TrimSpace(report.String()))
}
```

`TrimSpace` — trailing newline মুছে output।

---

## Expected Output

```
failed jobs (CSV):
job,duration_ms
report-mail,640
invoice-gen,905
```

## মূল শিক্ষা / Key Takeaways

1. **`strings.Cut`** — key=value field parse।
2. **`strings.Contains`** — filter condition।
3. **`strings.Fields`** — line tokenize।
4. **`strings.Builder`** — efficient output concat।
5. **Skip (continue) logic** — done-গুলো + missing-field-গুলো skip।

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
	"strings"
)
```

- `fmt` — for `Println`, `Print`.
- `strings` — for `Cut`, `Fields`, `Split`, `Contains`, `Builder`, `TrimSpace`.

### Lines 8–15

```go
func extractField(line, key string) (string, bool) {
	for _, field := range strings.Fields(line) {
		if k, v, ok := strings.Cut(field, "="); ok && k == key {
			return v, true
		}
	}
	return "", false
}
```

**A generic key=value extractor:**

- `strings.Fields(line)` — tokenizes the line by whitespace.
- `strings.Cut(field, "=")` — splits `key=value` into k and v.
- On a key match the value is returned (true). Otherwise false.

*("duration=640" → Cut → k="duration", v="640".)*

### Lines 17–21

```go
	jobLog := `2026-01-05T09:00:01 job=invoice-gen status=done duration=812
2026-01-05T09:00:04 job=report-mail status=failed duration=640
2026-01-05T09:00:09 job=invoice-gen status=failed duration=905
2026-01-05T09:00:12 job=db-cleanup status=done duration=430`
```

4 job-log lines — 2 done, 2 failed.

### Lines 23–26

```go
	var report strings.Builder
	report.WriteString("job,duration_ms\n")
	for _, line := range strings.Split(jobLog, "\n") {
		if !strings.Contains(line, "status=failed") {
			continue
		}
```

- `strings.Builder` — efficient CSV concatenation.
- Writes the header, then `Contains("status=failed")` filters each line — done lines are skipped.

### Lines 30–36

```go
		job, jobOK := extractField(line, "job")
		duration, durationOK := extractField(line, "duration")

		if !jobOK || !durationOK {
			continue
		}
		report.WriteString(job)
		report.WriteString(",")
		report.WriteString(duration)
		report.WriteString("\n")
```

- extractField pulls out `job` and `duration` to build a CSV row.
- Missing field → skip (safety).

### Lines 38–40

```go
	fmt.Println("failed jobs (CSV):")
	fmt.Print(strings.TrimSpace(report.String()))
}
```

`TrimSpace` strips the trailing newline for clean output.

---

## Expected Output

```
failed jobs (CSV):
job,duration_ms
report-mail,640
invoice-gen,905
```

## Key Takeaways

1. **`strings.Cut`** — key=value field parsing.
2. **`strings.Contains`** — a filter condition.
3. **`strings.Fields`** — line tokenization.
4. **`strings.Builder`** — efficient output concatenation.
5. **Skip (continue) logic** — skips done lines + missing fields.