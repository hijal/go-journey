# role-permission-authorization

Go-তে **variadic `...string`** আর **early return** দিয়ে role-based authorization check শেখার ছোট example।

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

Console-এ output print করার জন্য `fmt` package import করা হয়।

### Lines 5–12

```go
func HasRole(userRole string, allowedRoles ...string) bool {
	for _, role := range allowedRoles {
		if userRole == role {
			return true
		}
	}
	return false
}
```

`HasRole` — user-এর role allowed roles-এর list-এ আছে কি না:

- `userRole string` — user-এর role।
- `allowedRoles ...string` — **variadic**: allowed roles-গুলো (একটা `[]string`)।
- Loop-এ প্রতিটা `role`:
  - `if userRole == role` — match হলে **`return true` immediately** (early return — বাকি layer-এ যেতে হবে না)।
- Loop শেষ — কোনো match নেই → **`return false`**।

**Why variadic:** caller-কে allowed roles একটা slice/array-তে গুছিয়ে দিতে হয় না; সহজভাবে comma-দিয়ে pass করা যায়: `HasRole("editor", "admin", "moderator")`। আর `eval`-এর মতো symetric role set — permission checks-এ common pattern।

### Line 14

```go
func main() {
```

Program-এর entry point।

### Line 15

```go
fmt.Println(HasRole("editor", "admin", "moderator", "editor"))
```

`HasRole("editor", "admin", "moderator", "editor")` — allowed: admin, moderator, editor। `userRole` = editor, যা allowed-এ আছে। Loop-এর `role`-টি match → `true`। Output: `true`।

### Line 16

```go
fmt.Println(HasRole("guest", "admin", "moderator"))
```

`HasRole("guest", "admin", "moderator")` — allowed: admin, moderator। `guest` allowed roles-এ নেই → `false`। Output: `false`।

### Line 17

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
true
false
```

## মূল শিক্ষা / Key Takeaways

1. **Variadic `...string`** — allowed roles-কে comma-সহ সহজে pass করা।
2. **Early return** — match-এ সাথে-সাথে `true` — বাকি loop-এ যেতে হয় না।
3. **Membership test** — slice-এ element-টা আছে কি না ("list contains" pattern)।
4. **`==` string comparison** — exact role match।

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

Imports the `fmt` package for console output.

### Lines 5–12

```go
func HasRole(userRole string, allowedRoles ...string) bool {
	for _, role := range allowedRoles {
		if userRole == role {
			return true
		}
	}
	return false
}
```

`HasRole` — whether the user's role is among the allowed roles:

- `userRole string` — the user's role.
- `allowedRoles ...string` — **variadic**: the allowed roles (a `[]string`).
- The loop over each `role`:
  - `if userRole == role` — on a match, **`return true` immediately** (early return — no need to check the rest).
- Loop ends — no match → **`return false`**.

**Why variadic:** the caller doesn't need to wrap the allowed roles in a slice/array; they can be passed comma-separated: `HasRole("editor", "admin", "moderator")`. It's a common pattern for permission checks — a clean role-set membership test.

### Line 14

```go
func main() {
```

Program entry point.

### Line 15

```go
fmt.Println(HasRole("editor", "admin", "moderator", "editor"))
```

`HasRole("editor", "admin", "moderator", "editor")` — allowed: admin, moderator, editor. `userRole` = editor, which is in the list. A `role` matches → `true`. Output: `true`.

### Line 16

```go
fmt.Println(HasRole("guest", "admin", "moderator"))
```

`HasRole("guest", "admin", "moderator")` — allowed: admin, moderator. `guest` isn't among them → `false`. Output: `false`.

### Line 17

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
true
false
```

## Key Takeaways

1. **Variadic `...string`** — pass allowed roles simply, comma-separated.
2. **Early return** — return `true` immediately on a match — no need to finish the loop.
3. **Membership test** — is an element in a slice ("list contains" pattern).
4. **`==` string comparison** — exact role matching.