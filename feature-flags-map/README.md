# feature-flags-map

Go-তে **nil map fade + comma-ok semantics + `delete`** শেখার ছোট example — feature-flag registry।

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

### Lines 5–6

```go
func main() {
	var flags map[string]bool
```

**`var` — nil map** (আরো `make` হয়নি)। Read সেফ, write করলে panic — তাই কথা হলো।

### Lines 8–12

```go
	if enabled, ok := flags["new-checkout"]; ok {
		fmt.Println("new-checkout configured:", enabled)
	} else {
		fmt.Println("new-checkout: no flag configured yet")
	}
```

**Nil-map read:**

- Nil map-এ read → zero values: `enabled=false`, `ok=false`।
- দায়িত্বে **comma-ok**: `ok=false` → flag-টা configure-ই হয়নি।

### Lines 14–18

```go
	flags = make(map[string]bool)

	flags["new-checkout"] = true
	flags["dark-mode"] = true
	flags["legacy-api"] = false
```

- `make` — এখন write-able।
- ৩টা flag (একটা explicit `false`)।

### Lines 20–22

```go
	fmt.Println("dark-mode enabled?  ", flags["dark-mode"])
	fmt.Println("legacy-api enabled? ", flags["legacy-api"])
	fmt.Println("beta-search enabled?", flags["beta-search"])
```

**জিরো-value read:**

- dark-mode → true, legacy-api → false (explicit)।
- `beta-search` — configure-ই নেই, তবু `false` print — **absent-flag == false pitfall**। প্লেইন read configure-false আর absent আলাদা করে না।

### Lines 24–26

```go
	if v, ok := flags["legacy-api"]; ok {
		fmt.Println("legacy-api explicitly configured as", v)
	}
```

**Comma-ok verify** — `legacy-api` সত্যিই configure করা (`ok=true`, value `false`) — absent-টা নয়।

### Lines 28–29

```go
	delete(flags, "legacy-api")
	fmt.Println("Flags remaining:", len(flags))
```

- `delete(map, key)` — entry সরায়।
- `len(flags)` → 2।

---

## Expected Output

```
new-checkout: no flag configured yet
dark-mode enabled?   true
legacy-api enabled?  false
beta-search enabled? false
legacy-api explicitly configured as false
Flags remaining: 2
```

## মূল শিক্ষা / Key Takeaways

1. **Nil map** — `var m map[...]` read-only, actual write panic।
2. **Plain read pitfall** — absent-flag এর `false`।
3. **Comma-ok** — configure-false vs absent আলাদা।
4. **`delete`** — entry removal + `len`।
5. **Initialize `make`** — write করার আগে।

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

### Lines 5–6

```go
func main() {
	var flags map[string]bool
```

**`var` — a nil map** (no `make` yet). Reads are safe, writes would panic — that's the point.

### Lines 8–12

```go
	if enabled, ok := flags["new-checkout"]; ok {
		fmt.Println("new-checkout configured:", enabled)
	} else {
		fmt.Println("new-checkout: no flag configured yet")
	}
```

**Nil-map read:**

- A read on a nil map → zero values: `enabled=false`, `ok=false`.
- The **comma-ok** carries the truth: `ok=false` → the flag was never configured.

### Lines 14–18

```go
	flags = make(map[string]bool)

	flags["new-checkout"] = true
	flags["dark-mode"] = true
	flags["legacy-api"] = false
```

- `make` — now writable.
- 3 flags (one explicit `false`).

### Lines 20–22

```go
	fmt.Println("dark-mode enabled?  ", flags["dark-mode"])
	fmt.Println("legacy-api enabled? ", flags["legacy-api"])
	fmt.Println("beta-search enabled?", flags["beta-search"])
```

**Zero-value reads:**

- dark-mode → true, legacy-api → false (explicit).
- `beta-search` — never configured, yet prints `false` — **the absent-flag == false pitfall**. A plain read can't tell configured-false from absent.

### Lines 24–26

```go
	if v, ok := flags["legacy-api"]; ok {
		fmt.Println("legacy-api explicitly configured as", v)
	}
```

**Comma-ok verify** — `legacy-api` really is configured (`ok=true`, value `false`) — unlike the absent one.

### Lines 28–29

```go
	delete(flags, "legacy-api")
	fmt.Println("Flags remaining:", len(flags))
```

- `delete(map, key)` — removes the entry.
- `len(flags)` → 2.

---

## Expected Output

```
new-checkout: no flag configured yet
dark-mode enabled?   true
legacy-api enabled?  false
beta-search enabled? false
legacy-api explicitly configured as false
Flags remaining: 2
```

## Key Takeaways

1. **Nil map** — `var m map[...]` is read-only; an actual write panics.
2. **Plain read pitfall** — an absent flag's `false`.
3. **Comma-ok** — separates configured-false from absent.
4. **`delete`** — entry removal + `len`.
5. **Initialize with `make`** — before writing.