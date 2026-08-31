# rate-limit-config

Go-তে **exported vs unexported constants**, `fmt.Printf` আর descriptive naming দিয়ে config constants শেখার ছোট example।

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

### Lines 5–10

```go
const (
	MaxRequestsPerMin = 120
	IdleTimeoutSec    = 30
	ReadTimeoutSec    = 5
	defaultZone       = "us-east-1"
)
```

**Exported আর unexported constant-এর পার্থক্য:**

- `MaxRequestsPerMin`, `IdleTimeoutSec`, `ReadTimeoutSec` — **Exported** (Uppercase দিয়ে শুরু): যেকোনো অন্য package থেকে `ratelimit.MaxRequestsPerMin`-এর মতো access করা যায়।
- `defaultZone` — **Unexported** (lowercase): শুধু এই package-এর ভেতরে।

**Type:** কোনো explicit type দেওয়া হয়নি। Go যাচাই করে:

- `MaxRequestsPerMin = 120` → untyped integer (default `int`)
- `IdleTimeoutSec = 30` → untyped integer
- `ReadTimeoutSec = 5` → untyped integer
- `defaultZone = "us-east-1"` → untyped string

**Unit in the name:** `IdleTimeoutSec`, `ReadTimeoutSec` — "Sec" নামেই unit hint দেয় (seconds)। `MaxRequestsPerMin` — "PerMin" বোঝায়। Naming convention-এ unit suffix রাখলে constant-টা নিজেই documentation।

### Line 12

```go
func main() {
```

Program-এর entry point।

### Lines 13–15

```go
fmt.Printf("Rate limit %d req/min\n", MaxRequestsPerMin)
fmt.Printf("Idle timeout %ds | read timeout: %ds\n", IdleTimeoutSec, ReadTimeoutSec)
fmt.Printf("Default deployment zone: %s\n", defaultZone)
```

Output:

- `Rate limit 120 req/min`
- `Idle timeout 30s | read timeout: 5s`
- `Default deployment zone: us-east-1`

`%d` int-এর জন্য, `%s` string-এর জন্য।

### Line 16

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
Rate limit 120 req/min
Idle timeout 30s | read timeout: 5s
Default deployment zone: us-east-1
```

## মূল শিক্ষা / Key Takeaways

1. **Exported constant (uppercase)** — `MaxRequestsPerMin` বাইরের package থেকে access-যোগ্য।
2. **Unexported constant (lowercase)** — `defaultZone` শুধু ভেতরের জন্য।
3. **Unit suffix naming** — `IdleTimeoutSec` নামেই seconds hint; code documentation-এর কাজ করে।
4. **`fmt.Printf` formatting** — `%d` (int), `%s` (string)।
5. **Untyped constants** — explicit type না দিলে Go value দেখে type infer করে।

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

### Lines 5–10

```go
const (
	MaxRequestsPerMin = 120
	IdleTimeoutSec    = 30
	ReadTimeoutSec    = 5
	defaultZone       = "us-east-1"
)
```

**Exported vs unexported constants:**

- `MaxRequestsPerMin`, `IdleTimeoutSec`, `ReadTimeoutSec` — **exported** (start uppercase): accessible from other packages as `ratelimit.MaxRequestsPerMin`.
- `defaultZone` — **unexported** (lowercase): only visible inside this package.

**Type:** no explicit type is given. Go infers from the value:

- `MaxRequestsPerMin = 120` → untyped integer (default `int`)
- `IdleTimeoutSec = 30` → untyped integer
- `ReadTimeoutSec = 5` → untyped integer
- `defaultZone = "us-east-1"` → untyped string

**Unit in the name:** `IdleTimeoutSec`, `ReadTimeoutSec` — the "Sec" suffix is a unit hint (seconds). `MaxRequestsPerMin` — "PerMin" signals the unit. Putting unit suffixes in constant names makes the constant self-documenting.

### Line 12

```go
func main() {
```

Program entry point.

### Lines 13–15

```go
fmt.Printf("Rate limit %d req/min\n", MaxRequestsPerMin)
fmt.Printf("Idle timeout %ds | read timeout: %ds\n", IdleTimeoutSec, ReadTimeoutSec)
fmt.Printf("Default deployment zone: %s\n", defaultZone)
```

Output:

- `Rate limit 120 req/min`
- `Idle timeout 30s | read timeout: 5s`
- `Default deployment zone: us-east-1`

`%d` for ints, `%s` for strings.

### Line 16

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
Rate limit 120 req/min
Idle timeout 30s | read timeout: 5s
Default deployment zone: us-east-1
```

## Key Takeaways

1. **Exported constant (uppercase)** — `MaxRequestsPerMin` is accessible from outside packages.
2. **Unexported constant (lowercase)** — `defaultZone` is internal only.
3. **Unit suffix naming** — `IdleTimeoutSec` carries a seconds hint; acts as documentation.
4. **`fmt.Printf` formatting** — `%d` for int, `%s` for string.
5. **Untyped constants** — without an explicit type, Go infers from the value.