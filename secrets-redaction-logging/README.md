# secrets-redaction-logging

Go-তে **secrets redaction (`fmt.Stringer` + `GoStringer` + `slog.LogValuer`)** শেখার ছোট example — password কখনো logger/formatter-এ plain আসে না।

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
	"log/slog"
	"os"
)
```

- `log/slog` — structured logging (JSON handler)।
- `os` — `Stdout`।
- `fmt` — `Sprintf`, `Printf`।

### Lines 9–27

```go
type Secret string

const redacted = "[REDACTED]"

func (Secret) String() string {
	return redacted
}

func (Secret) GoString() string {
	return `Secret("` + redacted + `")`
}

func (Secret) LogValue() slog.Value {
	return slog.StringValue(redacted)
}

func (s Secret) Reveal() string {
	return string(s)
}
```

**দুই-স্তরের type** (`Secret`) — ৩টা redaction hook:

- `String()` — `fmt` `%v`/`%+v` → `[REDACTED]`।
- `GoString()` — `%#v` → `Secret("[REDACTED]")`।
- `LogValue()` — slog dump → `[REDACTED]`।
- `Reveal()` — **একমাত্র deliberate path** (ইচ্ছাকৃত use-এ)।

### Lines 29–47

```go
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password Secret
}

func (c DatabaseConfig) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("host", c.Host),
		slog.Int("port", c.Port),
		slog.String("user", c.User),
		slog.Any("password", c.Password),
	)
}

func (c DatabaseConfig) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/app", c.User, c.Password.Reveal(), c.Host, c.Port)
}
```

- **`LogValue`** — struct whole-ভাবে slog-এ group; `slog.Any("password")` → inner `Secret.LogValue()` → redacted।
- **`DSN()`** — legitimately `Reveal()` ব্যবহার (real password, কিন্তু এটা deliberate)।

### Lines 49–54

```go
	cfg := DatabaseConfig{Host: "db.internal", Port: 5432, User: "app", Password: "abc123"}

	fmt.Printf("%v\n", cfg)
	fmt.Printf("%+v\n", cfg)
	fmt.Printf("%#v\n", cfg)
```

৩টা fmt verb — প্রতিটা-তে `Password` **redacted**।

### Lines 56–65

```go
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return a
		},
	}))

	logger.Info("connecting to database", "db", cfg)
```

- JSON handler + `ReplaceAttr` — time-key remove (clean output)।
- `logger.Info(..., "db", cfg)` — struct-এর `LogValue` → group JSON, password `[REDACTED]`।

### Lines 67

```go
	fmt.Println("dsn length:", len(cfg.DSN()))
```

`DSN()` — আসল credential-সহ ব্যবহার, কিন্তু output-এ শুধু **length**।

---

## Expected Output

```
{db.internal 5432 app [REDACTED]}
{Host:db.internal Port:5432 User:app Password:[REDACTED]}
main.DatabaseConfig{Host:"db.internal", Port:5432, User:"app", Password:Secret("[REDACTED]")}
{"level":"INFO","msg":"connecting to database","db":{"host":"db.internal","port":5432,"user":"app","password":"[REDACTED]"}}
dsn length: 42
```

## মূল শিক্ষা / Key Takeaways

1. **`String()`** — `%v`/`%+v` redact।
2. **`GoString()`** — `%#v` redact।
3. **`LogValue()`** — slog dump redact (struct-level `GroupValue` সহ)।
4. **`Reveal()`** — deliberate escape-hatch, blind-use নয়।
5. **Default behavior** — secret type-এ method দিলেই সব formatter তারকা খেলবে।

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
	"log/slog"
	"os"
)
```

- `log/slog` — structured logging (JSON handler).
- `os` — for `Stdout`.
- `fmt` — `Sprintf`, `Printf`.

### Lines 9–27

```go
type Secret string

const redacted = "[REDACTED]"

func (Secret) String() string {
	return redacted
}

func (Secret) GoString() string {
	return `Secret("` + redacted + `")`
}

func (Secret) LogValue() slog.Value {
	return slog.StringValue(redacted)
}

func (s Secret) Reveal() string {
	return string(s)
}
```

**A two-tier type** (`Secret`) — 3 redaction hooks:

- `String()` — `fmt` `%v`/`%+v` → `[REDACTED]`.
- `GoString()` — `%#v` → `Secret("[REDACTED]")`.
- `LogValue()` — slog dump → `[REDACTED]`.
- `Reveal()` — the **only deliberate path** (intentional use).

### Lines 29–47

```go
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password Secret
}

func (c DatabaseConfig) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("host", c.Host),
		slog.Int("port", c.Port),
		slog.String("user", c.User),
		slog.Any("password", c.Password),
	)
}

func (c DatabaseConfig) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/app", c.User, c.Password.Reveal(), c.Host, c.Port)
}
```

- **`LogValue`** — the whole struct logs as a group; `slog.Any("password")` → the inner `Secret.LogValue()` → redacted.
- **`DSN()`** — legitimately uses `Reveal()` (the real password, but this is deliberate).

### Lines 49–54

```go
	cfg := DatabaseConfig{Host: "db.internal", Port: 5432, User: "app", Password: "abc123"}

	fmt.Printf("%v\n", cfg)
	fmt.Printf("%+v\n", cfg)
	fmt.Printf("%#v\n", cfg)
```

All 3 fmt verbs — `Password` is **redacted** in every one.

### Lines 56–65

```go
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return a
		},
	}))

	logger.Info("connecting to database", "db", cfg)
```

- A JSON handler + `ReplaceAttr` removing the time key (clean output).
- `logger.Info(..., "db", cfg)` — the struct's `LogValue` → a JSON group, password `[REDACTED]`.

### Line 67

```go
	fmt.Println("dsn length:", len(cfg.DSN()))
```

`DSN()` — uses the real credential, but only the **length** is printed.

---

## Expected Output

```
{db.internal 5432 app [REDACTED]}
{Host:db.internal Port:5432 User:app Password:[REDACTED]}
main.DatabaseConfig{Host:"db.internal", Port:5432, User:"app", Password:Secret("[REDACTED]")}
{"level":"INFO","msg":"connecting to database","db":{"host":"db.internal","port":5432,"user":"app","password":"[REDACTED]"}}
dsn length: 42
```

## Key Takeaways

1. **`String()`** — redacts `%v`/`%+v`.
2. **`GoString()`** — redacts `%#v`.
3. **`LogValue()`** — redacts slog dumps (incl. struct-level `GroupValue`).
4. **`Reveal()`** — a deliberate escape hatch, not blind use.
5. **Default behavior** — add the methods to the secret type and every formatter plays along.