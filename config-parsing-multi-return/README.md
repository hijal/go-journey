# config-parsing-multi-return

Go-তে **JSON file parse** + **multiple return values** (`(*Config, error)`) শেখার ছোট example — `os.ReadFile` + `json.Unmarshal` + error wrapping, আর `runtime.Caller` দিয়ে config file-এর path খুঁজে।

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

### Lines 3–8

```go
import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)
```

- `encoding/json` — `json.Unmarshal`।
- `fmt` — `Errorf`, `Printf`।
- `os` — `os.ReadFile`।
- `path/filepath` — path join/dir।
- `runtime` — `runtime.Caller`।

### Lines 10–13

```go
type Config struct {
	Port   int    `json:"port"`
	Env    string `json:"env"`
	DbHost string `json:"db_host"`
}
```

`Config` struct — JSON keys-এর সাথে struct tags:

- `json:"port"` — JSON-এ `"port"` field-টা `Port`-এ ম্যাপ হবে।
- `json:"env"`, `json:"db_host"` — একইভাবে।

### Lines 15–20

```go
func configPath() string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return "config.json"
	}
	return filepath.Join(filepath.Dir(file), "config.json")
}
```

`configPath` — config file-এর **absolute path** source file-এর পাশে থেকে বের করা:

- `runtime.Caller(0)` — চালানো **current source file-এর path** return করে (0 = এই function-টাই)।
- `filepath.Dir(file)` — source-এর directory।
- `filepath.Join(dir, "config.json")` — হলে সেই directory-র ভেতরে `config.json`-এর পথ।

**কেন:** `go run ./folder` repo root থেকে চালালে `os.ReadFile("config.json")` (relative) root-এ খুঁজত — fail। `runtime.Caller`-এ path-টা কোথা থেকে run করলেও কাজ করে (cwd-নির্বিশেষে)।

### Lines 22–36

```go
func ParseConfig(filePath string) (*Config, error) {
	data, err := os.ReadFile(filePath)

	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	var cfg Config

	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return &cfg, nil
}
```

`ParseConfig` — একটা file path নেয়, result হিসেবে **দুটো value**:

- `*Config` — parsed config (success এ)।
- `error` — ব্যর্থ হলে (else `nil`)।
- `os.ReadFile(filePath)` — file-এর সম্পূর্ণ content byte-এ read।
  - fail-এ: `nil, fmt.Errorf("failed to read file %s: %w", ...)` — **`%w`** দিয়ে original error wrap-করা। Multiple return-এ **দুটো value** — `nil` pointer + error।
- `json.Unmarshal(data, &cfg)` — JSON bytes-কে `Config` struct-এ parse (pointer দিয়ে populate)।
  - fail-এ: `nil, wrapped error` ("failed to parse JSON")।
- success: `&cfg, nil` — pointer + no error।

**Pattern (Go idiom):** fail-এ `(nil, err)` return, অন্যথায় `(value, nil)`। Caller-কে আগে `err` check করতে হয়।

### Lines 38–46

```go
func main() {
	cfg, err := ParseConfig(configPath())

	if err != nil {
		fmt.Printf("fatal error loading config: %v\n", err)
		return
	}

	fmt.Printf("loaded config for env: %v\n", cfg)
}
```

- `ParseConfig(configPath())` — অটো-পথে config parse।
- `if err != nil` — ব্যর্থ হলে `fatal error loading config: ...` print + `return` (fail-fast)।
- success: `loaded config for env: &{8080 development postgres://localhost:5432/app}`। `%v` struct-টা field-order-এ print করে।

---

## Expected Output

```
loaded config for env: &{8080 development postgres://localhost:5432/app}
```

`config.json`-এর content:

```json
{
  "port": 8080,
  "env": "development",
  "db_host": "postgres://localhost:5432/app"
}
```

## মূল শিক্ষা / Key Takeaways

1. **Multiple return** — `(*Config, error)` — Go-র standard error-handling signature।
2. **`%w` wrapping** — `fmt.Errorf` দিয়ে original error-টা চেইন-এ রাখা।
3. **Struct tags** — `json:"port"` দিয়ে field-ম্যাপিং।
4. **`json.Unmarshal`** — bytes → struct (pointer দিয়ে)।
5. **`runtime.Caller(0)`** — cwd-নির্বিশেষে source file-এর পাশে config খোঁজা।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Lines 3–8

```go
import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)
```

- `encoding/json` — for `json.Unmarshal`.
- `fmt` — for `Errorf`, `Printf`.
- `os` — for `os.ReadFile`.
- `path/filepath` — for path join/dir.
- `runtime` — for `runtime.Caller`.

### Lines 10–13

```go
type Config struct {
	Port   int    `json:"port"`
	Env    string `json:"env"`
	DbHost string `json:"db_host"`
}
```

`Config` struct — struct tags mapping JSON keys:

- `json:"port"` — the JSON field `"port"` maps to `Port`.
- `json:"env"`, `json:"db_host"` — similarly.

### Lines 15–20

```go
func configPath() string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return "config.json"
	}
	return filepath.Join(filepath.Dir(file), "config.json")
}
```

`configPath` — derives an **absolute path** to the config file next to the source file:

- `runtime.Caller(0)` — returns the **path of the currently running source file** (0 = this function).
- `filepath.Dir(file)` — the source's directory.
- `filepath.Join(dir, "config.json")` — builds the path to `config.json` in that directory.

**Why:** running `go run ./folder` from the repo root, `os.ReadFile("config.json")` (relative) would look in the root and fail. `runtime.Caller` makes the path work regardless of where you run it (cwd-independent).

### Lines 22–36

```go
func ParseConfig(filePath string) (*Config, error) {
	data, err := os.ReadFile(filePath)

	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	var cfg Config

	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return &cfg, nil
}
```

`ParseConfig` — takes a file path, returns **two values**:

- `*Config` — the parsed config (on success).
- `error` — the failure (or `nil`).
- `os.ReadFile(filePath)` — reads the file's whole content into bytes.
  - On failure: `nil, fmt.Errorf("failed to read file %s: %w", ...)` — wraps the original error with **`%w`**. Two values on the return: a `nil` pointer + error.
- `json.Unmarshal(data, &cfg)` — parses the JSON bytes into the `Config` struct (populating via pointer).
  - On failure: `nil, wrapped error` ("failed to parse JSON").
- On success: `&cfg, nil` — a pointer + no error.

**Pattern (Go idiom):** on failure return `(nil, err)`, otherwise `(value, nil)`. The caller must check `err` first.

### Lines 38–46

```go
func main() {
	cfg, err := ParseConfig(configPath())

	if err != nil {
		fmt.Printf("fatal error loading config: %v\n", err)
		return
	}

	fmt.Printf("loaded config for env: %v\n", cfg)
}
```

- `ParseConfig(configPath())` — parses the config using the auto-derived path.
- `if err != nil` — on failure, prints `fatal error loading config: ...` and `return`s (fail-fast).
- On success: `loaded config for env: &{8080 development postgres://localhost:5432/app}`. `%v` prints the struct in field order.

---

## Expected Output

```
loaded config for env: &{8080 development postgres://localhost:5432/app}
```

The contents of `config.json`:

```json
{
  "port": 8080,
  "env": "development",
  "db_host": "postgres://localhost:5432/app"
}
```

## Key Takeaways

1. **Multiple return** — `(*Config, error)` — the standard Go error-handling signature.
2. **`%w` wrapping** — keeping the original error in the chain with `fmt.Errorf`.
3. **Struct tags** — field mapping via `json:"port"`.
4. **`json.Unmarshal`** — bytes → struct (via pointer).
5. **`runtime.Caller(0)`** — finds the config next to the source, independent of the cwd.