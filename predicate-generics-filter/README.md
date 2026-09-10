# predicate-generics-filter

Go-তে **Generics** (`Filter[T]`, `Map[T, U]`) + **predicate function** শেখার ছোট example — type-safe reusable filter/map helpers।

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

- `fmt` — `Printf`।
- `strings` — `strings.ToUpper`।

### Lines 8–13

```go
type User struct {
	ID       int
	Name     string
	Role     string
	IsActive bool
}
```

`User` struct — ID, Name, Role, IsActive — filter-এর test case।

### Lines 15–25

```go
func Filter[T any](items []T, predicate func(T) bool) []T {
	var result []T

	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}
```

`Filter` — **generic** helper:

- `[T any]` — type parameter; function-টা যেকোনো type-এ কাজ করে।
- `items []T` + `predicate func(T) bool` — predicate-টা প্রতিটা item-এ true/false return করে।
- Loop: predicate true হলে `result`-এ append।
- Result `[]T` — একই type-এর filtered slice।

**Type-safe:** compile-time-এ নিশ্চিত — non-predicate লজিক reuse, কোনো type-ই বিশেষায়িত না।

### Lines 27–35

```go
func Map[T any, U any](items []T, transform func(T) U) []U {
	result := make([]U, len(items))

	for i, item := range items {
		result[i] = transform(item)
	}

	return result
}
```

`Map` — generic transform:

- **দুটো type parameter** `[T any, U any]` — input type `T`, output type `U` (ভিন্নও হতে পারে)।
- `result := make([]U, len(items))` — preallocate, transform-গুলো index দিয়ে fill করে।
- প্রতিটা item → `transform(item)` → `U`-array-তে।

### Lines 37–43

```go
func main() {
	users := []User{
		{1, "Alice", "admin", true},
		{2, "Bob", "user", false},
		{3, "Charlie", "user", true},
		{4, "Diana", "admin", true},
	}
```

৪ টা user:

- Alice — admin, active
- Bob — user, inactive
- Charlie — user, active
- Diana — admin, active

### Lines 45–49

```go
	isActive := func(u User) bool { return u.IsActive }
	isAdmin := func(u User) bool { return u.Role == "admin" }

	activeAdmins := Filter(Filter(users, isActive), isAdmin)
	fmt.Printf("active admins: %v\n", activeAdmins)
```

- `isActive`, `isAdmin` — দুটো **predicate closure**।
- `Filter(Filter(users, isActive), isAdmin)` — **compose**: আগে active-slice, তারপর admin-filter।
  - Active: Alice, Charlie, Diana → admin: Alice, Diana।
- Output: `[{1 Alice admin true} {4 Diana admin true}]`।

### Lines 51–55

```go
	names := Map(activeAdmins, func(u User) string {
		return strings.ToUpper(u.Name)
	})

	fmt.Printf("active admin names: %v\n", names)
```

- `Map` — `User` → `string` (`U = string`)। Name uppercase: `ALICE`, `DIANA`।

---

## Expected Output

```
active admins: [{1 Alice admin true} {4 Diana admin true}]
active admin names: [ALICE DIANA]
```

## মূল শিক্ষা / Key Takeaways

1. **Generics** — `Filter[T any]`, `Map[T any, U any]` — reusable, type-safe।
2. **Predicate** — `func(T) bool`, যেকোনো condition-এর জন্য।
3. **Function composition** — `Filter(Filter(items, p1), p2)`।
4. **Two type params** — input/output ভিন্ন type (`User` → `string`)।
5. **Preallocated result** — `make([]U, len(items))` + index write।

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

- `fmt` — for `Printf`.
- `strings` — for `strings.ToUpper`.

### Lines 8–13

```go
type User struct {
	ID       int
	Name     string
	Role     string
	IsActive bool
}
```

`User` struct — ID, Name, Role, IsActive — the test subject for filtering.

### Lines 15–25

```go
func Filter[T any](items []T, predicate func(T) bool) []T {
	var result []T

	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}
```

`Filter` — a **generic** helper:

- `[T any]` — a type parameter; the function works on any type.
- `items []T` + `predicate func(T) bool` — the predicate returns true/false for each item.
- Loop: append to `result` when the predicate is true.
- Returns `[]T` — a filtered slice of the same type.

**Type-safe:** guaranteed at compile time — reusable filtering logic without specializing per type.

### Lines 27–35

```go
func Map[T any, U any](items []T, transform func(T) U) []U {
	result := make([]U, len(items))

	for i, item := range items {
		result[i] = transform(item)
	}

	return result
}
```

`Map` — a generic transform:

- **Two type parameters** `[T any, U any]` — input type `T`, output type `U` (they can differ).
- `result := make([]U, len(items))` — preallocated, filled by index.
- Each item → `transform(item)` → stored in the `U`-array.

### Lines 37–43

```go
func main() {
	users := []User{
		{1, "Alice", "admin", true},
		{2, "Bob", "user", false},
		{3, "Charlie", "user", true},
		{4, "Diana", "admin", true},
	}
```

Four users:

- Alice — admin, active
- Bob — user, inactive
- Charlie — user, active
- Diana — admin, active

### Lines 45–49

```go
	isActive := func(u User) bool { return u.IsActive }
	isAdmin := func(u User) bool { return u.Role == "admin" }

	activeAdmins := Filter(Filter(users, isActive), isAdmin)
	fmt.Printf("active admins: %v\n", activeAdmins)
```

- `isActive`, `isAdmin` — two **predicate closures**.
- `Filter(Filter(users, isActive), isAdmin)` — **composition**: first the active slice, then the admin filter.
  - Active: Alice, Charlie, Diana → admin: Alice, Diana.
- Output: `[{1 Alice admin true} {4 Diana admin true}]`.

### Lines 51–55

```go
	names := Map(activeAdmins, func(u User) string {
		return strings.ToUpper(u.Name)
	})

	fmt.Printf("active admin names: %v\n", names)
```

- `Map` — `User` → `string` (`U = string`). Name uppercased: `ALICE`, `DIANA`.

---

## Expected Output

```
active admins: [{1 Alice admin true} {4 Diana admin true}]
active admin names: [ALICE DIANA]
```

## Key Takeaways

1. **Generics** — `Filter[T any]`, `Map[T any, U any]` — reusable and type-safe.
2. **Predicate** — `func(T) bool`, for any condition.
3. **Function composition** — `Filter(Filter(items, p1), p2)`.
4. **Two type params** — differing input/output types (`User` → `string`).
5. **Preallocated result** — `make([]U, len(items))` + index writes.