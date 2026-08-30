# env-variable-parsing

Go-তে **environment-ভেরিয়েবল parsing** শেখার ছোট example — string value-কে `strconv` দিয়ে সঠিক type-এ convert করা এবং missing/invalid value-র জন্য **default** দেওয়া।

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
	"strconv"
)
```

দুটো package import করা হয়:

- `fmt` — console-এ output print করার জন্য।
- `strconv` — string-কে অন্য type-এ convert করার জন্য (`strconv.ParseBool`, `strconv.Atoi`)।

### Line 8

```go
func main() {
```

Program-এর entry point।

### Lines 9–13

```go
rawEnv := map[string]string{
	"APP_DEBUG":           "true",
	"MAX_WORKERS":         "12",
	"REQUEST_TIMEOUT_SEC": "",
}
```

`rawEnv` — একটা map (`map[string]string`) যা environment variable-গুলো `string` হিসেবে ধারণ করে (real app-এ `os.Getenv`-এর মতো source থেকে আসে)। লক্ষ্য করো: value-গুলো সবই string — `"true"`, `"12"`, আর `""` (empty)। Empty হওয়াটা conscious: `REQUEST_TIMEOUT_SEC` set করা হয়নি বোঝায়।

### Line 15

```go
var debugMode bool
```

`var debugMode bool` — `debugMode` **zero value** (`false`) দিয়ে declare হয়। Default সেটিং: parse fail বা missing হলে `false`-ই থাকবে।

### Lines 17–24

```go
if v, ok := rawEnv["APP_DEBUG"]; ok {
	parsed, err := strconv.ParseBool(v)
	if err != nil {
		fmt.Println("invalid APP_DEBUG, keeping default false")
	} else {
		debugMode = parsed
	}
}
```

`APP_DEBUG` parse করা হয়:

- `if v, ok := rawEnv["APP_DEBUG"]; ok` — map lookup-এর compound syntax: `v` = value (`"true"`), `ok` = key আছে কিনা।
- `strconv.ParseBool(v)` — string-কে bool-এ convert করে (`"true"` → `true`)।
- `if err != nil` — invalid format হলে (যেমন `"yes"` বা `"1"`) error আসে; তখন message print করে default `false` ধরে রাখা হয়।
- Parse সফল হলে `debugMode = parsed` — মানে `true`।

### Lines 26–29

```go
maxWorkers, err := strconv.Atoi(rawEnv["MAX_WORKERS"])
if err != nil {
	maxWorkers = 4
}
```

- `strconv.Atoi` — ASCII string-কে `int`-এ convert করে (`"12"` → `12`)।
- `rawEnv["MAX_WORKERS"]` — key missing বা value bad হলে `err` non-nil হয়; তখন `maxWorkers = 4` (default) বসে। এখানে `"12"` valid, তাই `maxWorkers` = `12`।

### Lines 31–36

```go
var timeoutSec int
if v, ok := rawEnv["REQUEST_TIMEOUT_SEC"]; ok && v != "" {
	timeoutSec, _ = strconv.Atoi(v)
} else {
	timeoutSec = 30
}
```

`timeoutSec`-এর জন্য **default-সহ safe parse**:

- `var timeoutSec int` — zero value (`0`) দিয়ে শুরু।
- `ok && v != ""` — key আছে **এবং** শূন্য না — এই দুটো condition-ই মিললে-শুধু `Atoi` করা হয়। 
- `timeoutSec, _ = strconv.Atoi(v)` — underscore `_` দিয়ে error discard করা হয় (এখানে non-empty guaranteed, তাই safe)।
- অন্যথায় `timeoutSec = 30` — default।

`REQUEST_TIMEOUT_SEC`-এর value `""` হওয়ায় `else` branch চলে → `timeoutSec` = `30`।

### Lines 38–40

```go
fmt.Println("debugMode  :", debugMode)
fmt.Println("maxWorkers :", maxWorkers)
fmt.Println("timeoutSec :", timeoutSec)
```

তিনটা final setting print করে:

- `debugMode  : true`
- `maxWorkers : 12`
- `timeoutSec : 30`

### Line 41

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
debugMode  : true
maxWorkers : 12
timeoutSec : 30
```

## মূল শিক্ষা / Key Takeaways

1. **`strconv.ParseBool` / `Atoi`** — string-কে bool/int-এ convert করার standard-library function।
2. **Map lookup compound syntax** — `if v, ok := m[k]; ok` — key presence-check সহ value পড়া।
3. **Defaults** — invalid/missing value-র জন্য sensible default রাখা (production-grade behavior)।
4. **Empty string check** — `v != ""` দিয়ে unset-কে আলাদা করা।
5. **Zero value as default** — `var debugMode bool` আগে থেকেই নিরাপদ default ধরে রাখে।

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
	"strconv"
)
```

Two packages are imported:

- `fmt` — for printing output to the console.
- `strconv` — for converting strings to other types (`strconv.ParseBool`, `strconv.Atoi`).

### Line 8

```go
func main() {
```

Program entry point.

### Lines 9–13

```go
rawEnv := map[string]string{
	"APP_DEBUG":           "true",
	"MAX_WORKERS":         "12",
	"REQUEST_TIMEOUT_SEC": "",
}
```

`rawEnv` — a map (`map[string]string`) holding environment variables as strings (in a real app these come from something like `os.Getenv`). Note the values are all strings: `"true"`, `"12"`, and `""` (empty). The empty is deliberate — it means `REQUEST_TIMEOUT_SEC` was not set.

### Line 15

```go
var debugMode bool
```

`var debugMode bool` — `debugMode` is declared with its **zero value** (`false`). Default: if parsing fails or is missing, it stays `false`.

### Lines 17–24

```go
if v, ok := rawEnv["APP_DEBUG"]; ok {
	parsed, err := strconv.ParseBool(v)
	if err != nil {
		fmt.Println("invalid APP_DEBUG, keeping default false")
	} else {
		debugMode = parsed
	}
}
```

Parsing `APP_DEBUG`:

- `if v, ok := rawEnv["APP_DEBUG"]; ok` — the compound map-lookup syntax: `v` = the value (`"true"`), `ok` = whether the key exists.
- `strconv.ParseBool(v)` — converts the string to a bool (`"true"` → `true`).
- `if err != nil` — on an invalid format (e.g. `"yes"` or `"1"`), print a message and keep the default `false`.
- On success, `debugMode = parsed` — i.e. `true`.

### Lines 26–29

```go
maxWorkers, err := strconv.Atoi(rawEnv["MAX_WORKERS"])
if err != nil {
	maxWorkers = 4
}
```

- `strconv.Atoi` — converts an ASCII string to an `int` (`"12"` → `12`).
- If the key is missing or the value is bad, `err` is non-nil; then `maxWorkers = 4` (default) applies. Here `"12"` is valid, so `maxWorkers` = `12`.

### Lines 31–36

```go
var timeoutSec int
if v, ok := rawEnv["REQUEST_TIMEOUT_SEC"]; ok && v != "" {
	timeoutSec, _ = strconv.Atoi(v)
} else {
	timeoutSec = 30
}
```

Safe parsing with a default for `timeoutSec`:

- `var timeoutSec int` — starts at its zero value (`0`).
- `ok && v != ""` — only convert with `Atoi` if the key exists **and** isn't empty.
- `timeoutSec, _ = strconv.Atoi(v)` — the underscore `_` discards the error (non-empty is guaranteed, so it's safe).
- Otherwise `timeoutSec = 30` — the default.

Because `REQUEST_TIMEOUT_SEC` is `""`, the `else` branch runs → `timeoutSec` = `30`.

### Lines 38–40

```go
fmt.Println("debugMode  :", debugMode)
fmt.Println("maxWorkers :", maxWorkers)
fmt.Println("timeoutSec :", timeoutSec)
```

Prints the three final settings:

- `debugMode  : true`
- `maxWorkers : 12`
- `timeoutSec : 30`

### Line 41

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
debugMode  : true
maxWorkers : 12
timeoutSec : 30
```

## Key Takeaways

1. **`strconv.ParseBool` / `Atoi`** — standard-library helpers to convert strings to bool/int.
2. **Compound map lookup** — `if v, ok := m[k]; ok` reads a value with a presence check.
3. **Defaults** — use sensible defaults for invalid/missing values (production behavior).
4. **Empty-string check** — `v != ""` distinguishes "unset" from other values.
5. **Zero value as default** — `var debugMode bool` already holds the safe default.