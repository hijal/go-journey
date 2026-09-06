# validate-config

Go-তে **config validation + named return** শেখার ছোট example -- host/port-এর validity যাচাই করে status message return করা.

**ভাষা নির্বাচন করুন / Choose language:**

[বাংলা](#bangla) * [English](#english)

---

<a name="bangla"></a>

## বাংলা সংস্করণ

### Lines 1-7

```go
package main

import (
    "errors"
    "fmt"
    "strings"
)
```

- `package main` -- executable program.
- `errors`, `fmt`, `strings` -- error, formatting, string utilities.

### Lines 9-20

```go
func validateConfig(host string, port int) (status string, err error) {
    if strings.TrimSpace(host) == "" {
        return "", errors.New("host is empty")
    }

    if port < 1 || port > 65535 {
        return "", fmt.Errorf("port %d out of range", port)
    }

    status = fmt.Sprintf("config OK: %s:%d", host, port)
    return status, nil
}
```

`validateConfig` named return values ব্যবহার করে:

- `(status string, err error)` -- return parameters named upfront.
- `host` empty বা whitespace-only হলে error.
- `port` valid range (1-65535) check করে; `fmt.Errorf` dynamic error message দেয়.
- সব ঠিক থাকলে success message format করে named variable-এ assign.
- explicit `return status, nil` -- named returns-এ naked return না করে values সহ return.

### Lines 22-43

```go
func main() {
    status, err := validateConfig("api.example.com", 8080)
    if err != nil {
        fmt.Println("invalid:", err)
    } else {
        fmt.Println(status)
    }

    status, err = validateConfig("", 8080)
    if err != nil {
        fmt.Println("invalid:", err)
    } else {
        fmt.Println(status)
    }

    status, err = validateConfig("api.example.com", 99999)
    if err != nil {
        fmt.Println("invalid:", err)
    } else {
        fmt.Println(status)
    }
}
```

তিনটা call:

- `("api.example.com", 8080)` -- valid -> "config OK: api.example.com:8080"
- `("", 8080)` -- empty host -> "host is empty"
- `("api.example.com", 99999)` -- port out of range -> "port 99999 out of range"

---

## Expected Output

```
config OK: api.example.com:8080
invalid: host is empty
invalid: port 99999 out of range
```

## মূল শিক্ষা / Key Takeaways

1. **Named return values** -- return parameters named করলে function-এর ভিতরে সরাসরি use করা যায়.
2. **Guard clauses** -- invalid input early return করা clean pattern.
3. **`fmt.Errorf`** -- dynamic error message with formatted values (`%d`).
4. **Range validation** -- port range 1-65535, যা standard TCP/UDP port space.
5. **Multiple validation paths** -- host + port আলাদা আলাদা check, specific error message.

---

---

<a name="english"></a>

## English Version

### Lines 1-7

```go
package main

import (
    "errors"
    "fmt"
    "strings"
)
```

- `package main` -- an executable program.
- `errors`, `fmt`, `strings` -- error, formatting, and string utilities.

### Lines 9-20

```go
func validateConfig(host string, port int) (status string, err error) {
    if strings.TrimSpace(host) == "" {
        return "", errors.New("host is empty")
    }

    if port < 1 || port > 65535 {
        return "", fmt.Errorf("port %d out of range", port)
    }

    status = fmt.Sprintf("config OK: %s:%d", host, port)
    return status, nil
}
```

`validateConfig` uses named return values:

- `(status string, err error)` -- return parameters are named upfront.
- empty or whitespace-only `host` returns an error.
- `port` is checked against the valid range (1-65535); `fmt.Errorf` produces a dynamic error message.
- on success, a status message is formatted into the named `status` variable.
- an explicit `return status, nil` returns named values (rather than a naked return).

### Lines 22-43

```go
func main() {
    status, err := validateConfig("api.example.com", 8080)
    if err != nil {
        fmt.Println("invalid:", err)
    } else {
        fmt.Println(status)
    }

    status, err = validateConfig("", 8080)
    if err != nil {
        fmt.Println("invalid:", err)
    } else {
        fmt.Println(status)
    }

    status, err = validateConfig("api.example.com", 99999)
    if err != nil {
        fmt.Println("invalid:", err)
    } else {
        fmt.Println(status)
    }
}
```

Three calls:

- `("api.example.com", 8080)` -- valid -> "config OK: api.example.com:8080"
- `("", 8080)` -- empty host -> "host is empty"
- `("api.example.com", 99999)` -- port out of range -> "port 99999 out of range"

---

## Expected Output

```
config OK: api.example.com:8080
invalid: host is empty
invalid: port 99999 out of range
```

## Key Takeaways

1. **Named return values** -- naming return parameters lets you use them directly inside the function.
2. **Guard clauses** -- early-return on invalid input keeps code clean.
3. **`fmt.Errorf`** -- dynamic error messages with formatted values (`%d`).
4. **Range validation** -- port range 1-65535 corresponds to the standard TCP/UDP port space.
5. **Multiple validation paths** -- host and port are checked separately with specific error messages.
