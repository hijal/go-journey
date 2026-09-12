# service-config-constructor

Go-তে **constructor validation + embedded struct (promoted fields) + sentinel error (`errors.Is`/`%w`)** শেখার ছোট example — `NewServerConfig` দিয়ে config validate করে তৈরি।

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
	"errors"
	"fmt"
)
```

- `errors` — `errors.New`, `errors.Is`।
- `fmt` — `Errorf`, `Printf`, `Println`।

### Lines 8–11

```go
type Timeout struct {
	Read  int
	Write int
}
```

`Timeout` — read/write timeout (সেকেন্ড)।

### Lines 13–17

```go
type ServerConfig struct {
	Timeout
	Host string
	Port int
}
```

`ServerConfig`:

- `Timeout` — **anonymous (embedded) struct field** — field-name ছাড়া embed।
- `Host`, `Port` — পরের fields।

**Promoted fields:** embedded-এর field-গুলো (`Read`, `Write`) parent-এ **promote** হয় — `cfg.Read`, `cfg.Write` দিয়ে সরাসরি access, মাঝে `cfg.Timeout.Read` লিখতে হয় না।

### Line 19

```go
var ErrInvalidConfig = errors.New("invalid configs")
```

**Sentinel error** — validation failure-র reference।

### Lines 21–38

```go
func NewServerConfig(host string, port, readTo, writeTo int) (*ServerConfig, error) {
	if host == "" {
		return nil, fmt.Errorf("%w: host is required", ErrInvalidConfig)
	}

	if port < 1 || port > 65535 {
		return nil, fmt.Errorf("%w: port %d out of range", ErrInvalidConfig, port)
	}

	return &ServerConfig{
		Timeout: Timeout{
			Read:  readTo,
			Write: writeTo,
		},
		Host: host,
		Port: port,
	}, nil
}
```

`NewServerConfig` — validated constructor:

- **`*ServerConfig` + `error` multiple return** — success-এ config, fail-এ error।
- **Guard 1:** `host == ""` → `%w` দিয়ে sentinel wrap + message ("host is required")।
- **Guard 2:** `port < 1 || port > 65535` → port range check fail।
- **Success:** composite literal-এ embedded `Timeout{...}` + `Host`, `Port` set → `nil` error।

**`%w` pattern:** original `ErrInvalidConfig`-টা chain-এ থাকে — caller `errors.Is(err, ErrInvalidConfig)` দিয়ে verify করতে পারে।

### Lines 40–46

```go
func main() {
	cfg, err := NewServerConfig("", 8080, 5, 10)
	if err != nil {
		fmt.Println("rejected:", err)
		fmt.Println("is config error:", errors.Is(err, ErrInvalidConfig))
		return
	}
```

- `NewServerConfig("", ...)` — `host == ""` → error।
- `rejected: invalid configs: host is required` — wrapped message।
- `is config error: true` — `errors.Is` sentinel match-টাই clarify ✓।
- **`return`** — এখানেই শেষ, তাই নিচের positive-path চলে না।

### Lines 48–55

```go
	cfg, err = NewServerConfig("localhost", 8080, 5, 10)
	if err != nil {
		fmt.Println("unexpected:", err)
		return
	}

	fmt.Printf("listening on %s:%d (read %ds / write %ds)\n",
		cfg.Host, cfg.Port, cfg.Read, cfg.Write)
```

- এই branch-টা এখানে **reach হয় না** (উপরে `return`) — কিন্তু "localhost" দিলে চালত।
- `cfg.Read`, `cfg.Write` — **promoted fields** (embedded-এর সরাসরি)।

---

## Expected Output

```
rejected: invalid configs: host is required
is config error: true
```

> `main`-এ প্রথম কল fail → `return`, তাই বৈধ config-এর `listening on ...` output-টা এই রানে দেখায় না।

## মূল শিক্ষা / Key Takeaways

1. **Validated constructor** — config-কে func-এর বাইরে validate, invalid object কখনো build হয় না।
2. **Embedded struct** — anonymous embed + promoted fields (`cfg.Read`).
3. **Sentinel error + `%w`** — wrap + `errors.Is` check।
4. **Multiple return** — `(*ServerConfig, error)`।
5. **Range guard** — `port < 1 || port > 65535`।

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
	"errors"
	"fmt"
)
```

- `errors` — for `errors.New`, `errors.Is`.
- `fmt` — for `Errorf`, `Printf`, `Println`.

### Lines 8–11

```go
type Timeout struct {
	Read  int
	Write int
}
```

`Timeout` — read/write timeouts (seconds).

### Lines 13–17

```go
type ServerConfig struct {
	Timeout
	Host string
	Port int
}
```

`ServerConfig`:

- `Timeout` — an **anonymous (embedded) struct field** — embedded without a field name.
- `Host`, `Port` — the other fields.

**Promoted fields:** the embedded struct's fields (`Read`, `Write`) are **promoted** to the parent — accessed directly as `cfg.Read`, `cfg.Write`, without `cfg.Timeout.Read`.

### Line 19

```go
var ErrInvalidConfig = errors.New("invalid configs")
```

**Sentinel error** — the reference for validation failures.

### Lines 21–38

```go
func NewServerConfig(host string, port, readTo, writeTo int) (*ServerConfig, error) {
	if host == "" {
		return nil, fmt.Errorf("%w: host is required", ErrInvalidConfig)
	}

	if port < 1 || port > 65535 {
		return nil, fmt.Errorf("%w: port %d out of range", ErrInvalidConfig, port)
	}

	return &ServerConfig{
		Timeout: Timeout{
			Read:  readTo,
			Write: writeTo,
		},
		Host: host,
		Port: port,
	}, nil
}
```

`NewServerConfig` — a validated constructor:

- **`*ServerConfig` + `error` multiple return** — the config on success, an error on failure.
- **Guard 1:** `host == ""` → wraps the sentinel with `%w` + a message ("host is required").
- **Guard 2:** `port < 1 || port > 65535` → port range check fails.
- **Success:** a composite literal setting the embedded `Timeout{...}`, `Host`, `Port` → `nil` error.

**The `%w` pattern:** the original `ErrInvalidConfig` stays in the chain so callers can check `errors.Is(err, ErrInvalidConfig)`.

### Lines 40–46

```go
func main() {
	cfg, err := NewServerConfig("", 8080, 5, 10)
	if err != nil {
		fmt.Println("rejected:", err)
		fmt.Println("is config error:", errors.Is(err, ErrInvalidConfig))
		return
	}
```

- `NewServerConfig("", ...)` — `host == ""` → error.
- `rejected: invalid configs: host is required` — the wrapped message.
- `is config error: true` — `errors.Is` confirms the sentinel match ✓.
- **`return`** — ends the program here, so the positive path below doesn't run.

### Lines 48–55

```go
	cfg, err = NewServerConfig("localhost", 8080, 5, 10)
	if err != nil {
		fmt.Println("unexpected:", err)
		return
	}

	fmt.Printf("listening on %s:%d (read %ds / write %ds)\n",
		cfg.Host, cfg.Port, cfg.Read, cfg.Write)
```

- This branch **isn't reached** here (the `return` above), but it would run with `"localhost"`.
- `cfg.Read`, `cfg.Write` — **promoted fields** (accessed directly from the embedded struct).

---

## Expected Output

```
rejected: invalid configs: host is required
is config error: true
```

> The first call in `main` fails and `return`s, so the valid config's `listening on ...` output doesn't appear in this run.

## Key Takeaways

1. **Validated constructor** — validate before building; invalid objects never exist.
2. **Embedded struct** — anonymous embed + promoted fields (`cfg.Read`).
3. **Sentinel error + `%w`** — wrapping + `errors.Is` checks.
4. **Multiple return** — `(*ServerConfig, error)`.
5. **Range guard** — `port < 1 || port > 65535`.