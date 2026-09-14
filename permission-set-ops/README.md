# permission-set-ops

Go-তে **map-as-set ops (`map[string]struct{}`)** শেখার ছোট example — membership, `maps.Clone`/`maps.Copy` union, intersection, sorted keys।

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

- `fmt` — `Println`।
- `maps` — `Clone`, `Copy` (Go 1.21+)।
- `slices` — `Sorted` + `maps.Keys` pair (Go 1.21+)।

### Lines 9–12

```go
func has(perms map[string]struct{}, p string) bool {
	_, ok := perms[p]
	return ok
}
```

`has` — **membership helper**: value discard, comma-ok।

### Lines 14–20

```go
func main() {
	adminPerms := map[string]struct{}{
		"read": {}, "write": {}, "delete": {}, "invite": {},
	}
	editorPerms := map[string]struct{}{
		"read": {},
		"write": {},
	}
```

দুটা **set** (`struct{}` = zero-byte value): admin ৪টা, editor ২টা permission।

### Lines 22–23

```go
	fmt.Println("admin can delete?", has(adminPerms, "delete"))
	fmt.Println("editor can delete?", has(editorPerms, "delete"))
```

Membership check — true vs false।

### Lines 25–28

```go
	audit := maps.Clone(adminPerms)

	maps.Copy(audit, map[string]struct{}{"billing": {}})
	fmt.Println("audit can billing ", has(audit, "billing"))
```

- `maps.Clone(adminPerms)` — নিজস্ব copy (set snapshot)।
- `maps.Copy(audit, {billing})` — **union-add**: audit + billing → `has` true।

### Lines 30–36

```go
	baseline := map[string]struct{}{}

	for p := range editorPerms {
		if _, ok := adminPerms[p]; ok {
			baseline[p] = struct{}{}
		}
	}
```

**Intersection** — editorP-এর প্রতিটা permission admin-এ থাকলে baseline-এ — `{read, write}`।

### Lines 38–40

```go
	for _, p := range slices.Sorted(maps.Keys(baseline)) {
		fmt.Println("shared permission:", p)
	}
```

`slices.Sorted(maps.Keys(...))` — deterministic sorted iteration।

---

## Expected Output

```
admin can delete? true
editor can delete? false
audit can billing  true
shared permission: read
shared permission: write
```

## মূল শিক্ষা / Key Takeaways

1. **Map-as-set** — `map[string]struct{}` membership।
2. **`maps.Clone`/`maps.Copy`** — set snapshot + union।
3. **Manual intersection** — nested membership-check।
4. **`Sorted(Keys())`** — deterministic output।
5. **Comma-ok helper** — `has()`।

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

- `fmt` — for `Println`.
- `maps` — for `Clone`, `Copy` (Go 1.21+).
- `slices` — the `Sorted` + `maps.Keys` pair (Go 1.21+).

### Lines 9–12

```go
func has(perms map[string]struct{}, p string) bool {
	_, ok := perms[p]
	return ok
}
```

`has` — a **membership helper**: value discarded, comma-ok.

### Lines 14–20

```go
func main() {
	adminPerms := map[string]struct{}{
		"read": {}, "write": {}, "delete": {}, "invite": {},
	}
	editorPerms := map[string]struct{}{
		"read": {},
		"write": {},
	}
```

Two **sets** (`struct{}` = zero-byte value): admin has 4 perms, editor has 2.

### Lines 22–23

```go
	fmt.Println("admin can delete?", has(adminPerms, "delete"))
	fmt.Println("editor can delete?", has(editorPerms, "delete"))
```

Membership checks — true vs false.

### Lines 25–28

```go
	audit := maps.Clone(adminPerms)

	maps.Copy(audit, map[string]struct{}{"billing": {}})
	fmt.Println("audit can billing ", has(audit, "billing"))
```

- `maps.Clone(adminPerms)` — a private copy (a set snapshot).
- `maps.Copy(audit, {billing})` — **union-add**: audit + billing → `has` is true.

### Lines 30–36

```go
	baseline := map[string]struct{}{}

	for p := range editorPerms {
		if _, ok := adminPerms[p]; ok {
			baseline[p] = struct{}{}
		}
	}
```

**Intersection** — each editor perm present in admin lands in baseline — `{read, write}`.

### Lines 38–40

```go
	for _, p := range slices.Sorted(maps.Keys(baseline)) {
		fmt.Println("shared permission:", p)
	}
```

`slices.Sorted(maps.Keys(...))` — deterministic sorted iteration.

---

## Expected Output

```
admin can delete? true
editor can delete? false
audit can billing  true
shared permission: read
shared permission: write
```

## Key Takeaways

1. **Map-as-set** — `map[string]struct{}` membership.
2. **`maps.Clone`/`maps.Copy`** — set snapshot + union.
3. **Manual intersection** — a nested membership check.
4. **`Sorted(Keys())`** — deterministic output.
5. **Comma-ok helper** — `has()`.