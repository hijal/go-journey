# config-loading-iife

Go-তে **IIFE (Immediately Invoked Function Expression)** দিয়ে package-level config initialize শেখার ছোট example — `var cfg = func() Config {...}()` — env vars read + defaults।

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
	"os"
	"strconv"
)
```

- `fmt` — print।
- `os` — `os.Getenv`।
- `strconv` — `Atoi`।

### Lines 9–14

```go
type Config struct {
	port  int
	env   string
	dbURL string
	debug bool
}
```

`Config` struct — ৪টা field: port, env, dbURL, debug।

### Lines 16–40

```go
var cfg = func() Config {
	port, err := strconv.Atoi(os.Getenv("APP_PORT"))

	if err != nil {
		port = 8080
	}

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	dbURL := os.Getenv("APP_DB_URL")

	if dbURL == "" {
		dbURL = "postgres://localhost:5432/app"
	}

	return Config{
		port:  port,
		env:   env,
		dbURL: dbURL,
		debug: env == "development",
	}
}()
```

**IIFE pattern:** `var cfg = func() Config {...}()` — একটা **anonymous function declare করে সঙ্গে-সঙ্গে `()` দিয়ে invoke** করা হয়েছে। Result একটা `Config`-এ assign হয় `cfg`-তে।

**Flow:**
- `strconv.Atoi(os.Getenv("APP_PORT"))` — `APP_PORT` env read; invalid/shunno হলে `err` — then `port = 8080` (default)।
- `env` — `APP_ENV`, খালি হলে `"development"`।
- `dbURL` — `APP_DB_URL`, খালি হলে local postgres-এর default।
- `Config{...}` — সব value-সহ struct literal return।
- `debug: env == "development"` — env-টা development হলে true।

**কেন IIFE:** package-level `var`-এ লজিক সরাসরি initialize করা যায় না (multiple statements); IIFE-তে মাল্টি-statement logic চলে তারপর একটা value return — এক লাইনে। Setup একবারই হয় (package initialize-টাইম), তারপর `cfg` immutable।

**এই environment-এ default values:** APP_PORT/APP_ENV/APP_DB_URL set না — সব default:
- port 8080, env "development", dbURL postgres://localhost:5432/app, debug true

### Lines 42–45

```go
func main() {
	fmt.Printf("starting on port %d (%s)\n", cfg.port, cfg.env)
	fmt.Println("db:", cfg.dbURL)
	fmt.Println("debug mode:", cfg.debug)
}
```

Output:
- `starting on port 8080 (development)`
- `db: postgres://localhost:5432/app`
- `debug mode: true`

### Line 46

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
starting on port 8080 (development)
db: postgres://localhost:5432/app
debug mode: true
```

## মূল শিক্ষা / Key Takeaways

1. **IIFE** — `func() T {...}()` — declare + immediately invoke।
2. **Config at package scope** — `var cfg = ...` — setup একবার, generic-access।
3. **Env vars + defaults** — `os.Getenv("X")` + fallback (`if err/empty → default`)।
4. **Derived fields** — `debug` env-value-থেকে computed।
5. **Single initialization** — package load হওয়ার সময়সীমার ভেতর।

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
	"os"
	"strconv"
)
```

- `fmt` — for printing.
- `os` — for `os.Getenv`.
- `strconv` — for `Atoi`.

### Lines 9–14

```go
type Config struct {
	port  int
	env   string
	dbURL string
	debug bool
}
```

The `Config` struct — 4 fields: port, env, dbURL, debug.

### Lines 16–40

```go
var cfg = func() Config {
	port, err := strconv.Atoi(os.Getenv("APP_PORT"))

	if err != nil {
		port = 8080
	}

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	dbURL := os.Getenv("APP_DB_URL")

	if dbURL == "" {
		dbURL = "postgres://localhost:5432/app"
	}

	return Config{
		port:  port,
		env:   env,
		dbURL: dbURL,
		debug: env == "development",
	}
}()
```

**IIFE pattern:** `var cfg = func() Config {...}()` — an **anonymous function declared and immediately invoked with `()`**. The result is assigned to `cfg` as a `Config`.

**Flow:**
- `strconv.Atoi(os.Getenv("APP_PORT"))` — reads `APP_PORT`; if invalid/empty, `err` is set — then `port = 8080` (default).
- `env` — reads `APP_ENV`, default `"development"` if empty.
- `dbURL` — reads `APP_DB_URL`, defaulting to the local postgres URL.
- Returns a `Config{...}` struct literal with all values.
- `debug: env == "development"` — true when env is development.

**Why IIFE:** a package-level `var` can't hold multi-statement logic directly; an IIFE runs the multi-statement setup then returns a single value — in one expression. The setup runs once (at package init), then `cfg` is effectively read-only.

**In this environment:** APP_PORT/APP_ENV/APP_DB_URL aren't set — all defaults apply:
- port 8080, env "development", dbURL postgres://localhost:5432/app, debug true

### Lines 42–45

```go
func main() {
	fmt.Printf("starting on port %d (%s)\n", cfg.port, cfg.env)
	fmt.Println("db:", cfg.dbURL)
	fmt.Println("debug mode:", cfg.debug)
}
```

Output:
- `starting on port 8080 (development)`
- `db: postgres://localhost:5432/app`
- `debug mode: true`

### Line 46

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
starting on port 8080 (development)
db: postgres://localhost:5432/app
debug mode: true
```

## Key Takeaways

1. **IIFE** — `func() T {...}()` — declare and immediately invoke.
2. **Config at package scope** — `var cfg = ...` — setup once, globally accessible.
3. **Env vars + defaults** — `os.Getenv("X")` with a fallback (`if err/empty → default`).
4. **Derived fields** — `debug` computed from the env value.
5. **Single initialization** — runs at package load time.