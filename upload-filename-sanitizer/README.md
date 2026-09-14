# upload-filename-sanitizer

Go-তে **filename sanitization (`ToLower` + `ReplaceAll`) + `HasSuffix` allowlist + `LastIndex` split** শেখার ছোট example — upload validator।

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

- `fmt` — `Println`।
- `strings` — `ReplaceAll`, `ToLower`, `HasSuffix`, `LastIndex`, `Join`।

### Line 8

```go
var allowedExtensions = []string{".tar.gz", ".zip", ".log"}
```

Allowlist — `.tar.gz` compound-extension সহ (সবশেষে যাচাই করা হয়)।

### Lines 10–16

```go
func splitExt(name string) (string, string) {
	i := strings.LastIndex(name, ".")
	if i < 0 {
		return name, ""
	}
	return name[:i], name[i:]
}
```

**Last-dot split:**

- `strings.LastIndex(name, ".")` — শেষ `.`-এর position।
- dot না-থাকলে → (`name`, "")।
- `name[:i]`, `name[i:]` — base ও ext।

*(`LastIndex` — শেষ dot, তাই `"backup_jan.tar.gz"` → base `"backup_jan.tar"`, ext `".gz"`।)*

### Lines 18–19

```go
	uploaded := []string{"Backup JAN.tar.gz", "db Dump.ZIP", "notes.txt"}
```

৩টা upload — mixed-case + space-সহ।

### Lines 21–23

```go
	for _, f := range uploaded {
		clean := strings.ToLower(strings.ReplaceAll(f, " ", "_"))
```

**Sanitize:** space → `_`, সব lowercase। `"db Dump.ZIP"` → `db_dump.zip`।

### Lines 24–31

```go
		allowed := false

		for _, ext := range allowedExtensions {
			if strings.HasSuffix(clean, ext) {
				allowed = true
				break
			}
		}
```

`HasSuffix` দিয়ে allowlist-এর যেকোনো extension-এ শেষ হলে `allowed = true` (early break)।

### Lines 33–40

```go
		base, ext := splitExt(clean)

		if allowed {
			fmt.Println(clean, "-> ACCEPTED  (base:", base, "| ext:", ext+")")
		} else {
			fmt.Println(clean, "-> REJECTED (allowed:", strings.Join(allowedExtensions, ", ")+")")
		}
```

- accepted → base+ext ফরম্যাটে।
- rejected → allowlist-টা দেখায় (`strings.Join`), parens-alignment এর জন্য ACCEPTED-এ 2-space / REJECTED-এ 1-space।

---

## Expected Output

```
backup_jan.tar.gz -> ACCEPTED  (base: backup_jan.tar | ext: .gz)
db_dump.zip -> ACCEPTED  (base: db_dump | ext: .zip)
notes.txt -> REJECTED (allowed: .tar.gz, .zip, .log)
```

## মূল শিক্ষা / Key Takeaways

1. **Sanitize-first** — `ToLower` + `ReplaceAll(" ", "_")`।
2. **`strings.HasSuffix`** — allowlist whitelist-check।
3. **`strings.LastIndex`** — শেষ-dot split (`splitExt`)।
4. **Early `break`** — প্রথম ম্যাচে allowed।
5. **`strings.Join`** — extension list print।

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

- `fmt` — for `Println`.
- `strings` — for `ReplaceAll`, `ToLower`, `HasSuffix`, `LastIndex`, `Join`.

### Line 8

```go
var allowedExtensions = []string{".tar.gz", ".zip", ".log"}
```

An allowlist — including the compound `.tar.gz` (checked last).

### Lines 10–16

```go
func splitExt(name string) (string, string) {
	i := strings.LastIndex(name, ".")
	if i < 0 {
		return name, ""
	}
	return name[:i], name[i:]
}
```

**Last-dot split:**

- `strings.LastIndex(name, ".")` — the position of the last `.`.
- No dot → return (`name`, "").
- `name[:i]`, `name[i:]` — base and ext.

*(`LastIndex` uses the final dot, so `"backup_jan.tar.gz"` → base `"backup_jan.tar"`, ext `".gz"`.)*

### Lines 18–19

```go
	uploaded := []string{"Backup JAN.tar.gz", "db Dump.ZIP", "notes.txt"}
```

3 uploads — mixed case + spaces.

### Lines 21–23

```go
	for _, f := range uploaded {
		clean := strings.ToLower(strings.ReplaceAll(f, " ", "_"))
```

**Sanitize:** spaces → `_`, all lowercase. `"db Dump.ZIP"` → `db_dump.zip`.

### Lines 24–31

```go
		allowed := false

		for _, ext := range allowedExtensions {
			if strings.HasSuffix(clean, ext) {
				allowed = true
				break
			}
		}
```

`HasSuffix` — if the name ends with any allowlisted extension, set `allowed = true` (early break).

### Lines 33–40

```go
		base, ext := splitExt(clean)

		if allowed {
			fmt.Println(clean, "-> ACCEPTED  (base:", base, "| ext:", ext+")")
		} else {
			fmt.Println(clean, "-> REJECTED (allowed:", strings.Join(allowedExtensions, ", ")+")")
		}
```

- accepted → printed in base+ext form.
- rejected → shows the allowlist (`strings.Join`); parens alignment comes from 2 spaces (ACCEPTED) vs 1 space (REJECTED).

---

## Expected Output

```
backup_jan.tar.gz -> ACCEPTED  (base: backup_jan.tar | ext: .gz)
db_dump.zip -> ACCEPTED  (base: db_dump | ext: .zip)
notes.txt -> REJECTED (allowed: .tar.gz, .zip, .log)
```

## Key Takeaways

1. **Sanitize-first** — `ToLower` + `ReplaceAll(" ", "_")`.
2. **`strings.HasSuffix`** — an allowlist whitelist-check.
3. **`strings.LastIndex`** — last-dot split (`splitExt`).
4. **Early `break`** — allowed fires on the first match.
5. **`strings.Join`** — printing the extension list.