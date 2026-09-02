# token-validation

Go-তে **custom sentinel error**, **`fmt.Errorf` / `%w` wrapping**, **`errors.Is`** আর **`errors.New`** দিয়ে token validation শেখার ছোট example।

**📖 ভাষা নির্বাচন করুন / Choose language:**

[🇧🇩 বাংলা](#bangla) • [🇬🇧 English](#english)

---

<a name="bangla"></a>

## 🇧🇩 বাংলা সংস্করণ

### Lines 1–7

```go
package main

import (
    "errors"
    "fmt"
    "strings"
)
```

- `package main` — একটা executable program।
- `errors` — error তৈরি (`errors.New`) আর error compare (`errors.Is`) করার জন্য।
- `fmt` — output ও formatted error-এর জন্য।
- `strings` — string-এর সাথে কাজ করা (`HasPrefix`, `TrimPrefix`)।

### Line 9

```go
var errTokenExpired = errors.New("Token expired")
```

একটা **package-level variable** `errTokenExpired`, যা একটা predefined (sentinel) error compare করছে `errors.New("Token expired")`। এটাকে বলে **sentinel error** — একটা known, identifiable error value, যেটা পরে `errors.Is` দিয়ে compare করা হয়।

### Lines 11–22

```go
func validateToken(token string) (string, error) {
    if !strings.HasPrefix(token, "Bearer ") {
        return "", fmt.Errorf("malformed token %q: missing Bearer prefix", token)
    }
    payload := strings.TrimPrefix(token, "Bearer ")

    if payload == "expired.jwt" {
        return "", fmt.Errorf("token check failed: %w", errTokenExpired)
    }

    return payload, nil
}
```

`validateToken` function — একটা token (string) নেয় এবং দুটো value return করে: `payload` (string) আর `error`।

- `if !strings.HasPrefix(token, "Bearer ")` — token-টা `"Bearer "` দিয়ে শুরু না হলে → `fmt.Errorf("malformed token %q: missing Bearer prefix", token)` — একটা formatted error (input-টা `%q`-তে quote সহ দেখায়)। `("", error)` return করা হয়।
- `strings.TrimPrefix(token, "Bearer ")` — আগের `"Bearer "` অংশটা সরিয়ে দেয়, বাকিটা `payload`-এ রাখে।
- `if payload == "expired.jwt"` — payload-টা expired হলে → `fmt.Errorf("token check failed: %w", errTokenExpired)` — এখানে **`%w`** দিয়ে `errTokenExpired`-কে error-এর ভেতরে **wrap** করা হয়। এর মানে: error message-এ টেক্সট ("token check failed:") + মূল sentinel error-টা। Wrap করলে `errors.Is` দিয়ে পুরোনো error-ও match করা যায়।
- নাহলে `return payload, nil` — সফল।

### Lines 24–34

```go
func handleRequest(authHeader string) {
    if userID, err := validateToken(authHeader); err != nil {
        if errors.Is(err, errTokenExpired) {
            fmt.Println("401: please refresh your token")
        } else {
            fmt.Println("401:", err)
        }
    } else {
        fmt.Printf("200: serving data for user %q\n", userID)
    }
}
```

`handleRequest` function — একটা HTTP auth header string নেয়।

- `if userID, err := validateToken(authHeader); err != nil` — **`if` with initializer**: আগে `validateToken` call করে `userID` ও `err` পাওয়া যায়, তারপর `err != nil` (error আছে কিনা) check হয়।
- Error থাকলে:
  - `if errors.Is(err, errTokenExpired)` — **`errors.Is`** দিয়ে wrap-করা error-টাও unpack করে দেখে এটা আসলেই `errTokenExpired` কিনা। যদি হ্যাঁ → `"401: please refresh your token"` (user-কে token refresh করতে বলার user-friendly message)।
  - না হলে → `"401:", err` (generic error message)।
- Error না থাকলে → `fmt.Printf("200: serving data for user %q\n", userID)` — সফল response (userID quote সহ)।

### Line 36

```go
func main() {
```

Program-এর entry point।

### Lines 37–40

```go
handleRequest("Bearer user-42.jwt")
handleRequest("Bearer expired.jwt")
handleRequest("Basic abc")
```

তিনটা case test:

1. `"Bearer user-42.jwt"` — valid → `200: serving data for user "user-42.jwt"`।
2. `"Bearer expired.jwt"` — valid prefix, কিন্তু expired payload → `401: please refresh your token` (কারণ `errors.Is` match করে)।
3. `"Basic abc"` — `"Bearer "` prefix নেই → malformed → `401: malformed token "Basic abc": missing Bearer prefix`।

---

## Expected Output

```
200: serving data for user "user-42.jwt"
401: please refresh your token
401: malformed token "Basic abc": missing Bearer prefix
```

## মূল শিক্ষা / Key Takeaways

1. **Sentinel error** — package-level predefined error value (`errors.New`), তুলনার জন্য।
2. **`%w` wrapping** — `fmt.Errorf("...: %w", err)` দিয়ে পুরোনো error-কে wrap করা।
3. **`errors.Is`** — wrap-করা error chain-এও মূল error-কে match করা।
4. **`strings.HasPrefix` / `TrimPrefix`** — prefix check ও remove।
5. **`if` with initializer** — call + error-check একসাথে।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Lines 1–7

```go
package main

import (
    "errors"
    "fmt"
    "strings"
)
```

- `package main` — an executable program.
- `errors` — for creating errors (`errors.New`) and comparing errors (`errors.Is`).
- `fmt` — for output and formatted errors.
- `strings` — for string utilities (`HasPrefix`, `TrimPrefix`).

### Line 9

```go
var errTokenExpired = errors.New("Token expired")
```

A **package-level variable** `errTokenExpired` holding a predefined error created with `errors.New("Token expired")`. This is a **sentinel error** — a known, identifiable error value that can later be compared with `errors.Is`.

### Lines 11–22

```go
func validateToken(token string) (string, error) {
    if !strings.HasPrefix(token, "Bearer ") {
        return "", fmt.Errorf("malformed token %q: missing Bearer prefix", token)
    }
    payload := strings.TrimPrefix(token, "Bearer ")

    if payload == "expired.jwt" {
        return "", fmt.Errorf("token check failed: %w", errTokenExpired)
    }

    return payload, nil
}
```

`validateToken` function — takes a token (string) and returns two values: `payload` (string) and an `error`.

- `if !strings.HasPrefix(token, "Bearer ")` — if the token doesn't start with `"Bearer "` → `fmt.Errorf("malformed token %q: missing Bearer prefix", token)` — a formatted error (showing the input quoted with `%q`). Returns `("", error)`.
- `strings.TrimPrefix(token, "Bearer ")` — strips the leading `"Bearer "` part and keeps the rest in `payload`.
- `if payload == "expired.jwt"` — if the payload is expired → `fmt.Errorf("token check failed: %w", errTokenExpired)` — here **`%w` wraps** `errTokenExpired` inside the error. This means: the error message text ("token check failed:") plus the original sentinel error. By wrapping it, the original error can still be matched with `errors.Is`.
- Otherwise, `return payload, nil` — success.

### Lines 24–34

```go
func handleRequest(authHeader string) {
    if userID, err := validateToken(authHeader); err != nil {
        if errors.Is(err, errTokenExpired) {
            fmt.Println("401: please refresh your token")
        } else {
            fmt.Println("401:", err)
        }
    } else {
        fmt.Printf("200: serving data for user %q\n", userID)
    }
}
```

`handleRequest` function — takes an HTTP auth header string.

- `if userID, err := validateToken(authHeader); err != nil` — **`if` with initializer**: first calls `validateToken` to get `userID` and `err`, then checks `err != nil` (whether there's an error).
- If there's an error:
  - `if errors.Is(err, errTokenExpired)` — **`errors.Is`** unwraps the wrapped error to check if it's really `errTokenExpired`. If yes → `"401: please refresh your token"` (a user-friendly message asking to refresh the token).
  - Otherwise → `"401:", err` (generic error message).
- If there's no error → `fmt.Printf("200: serving data for user %q\n", userID)` — a success response (userID quoted).

### Line 36

```go
func main() {
```

Program entry point.

### Lines 37–40

```go
handleRequest("Bearer user-42.jwt")
handleRequest("Bearer expired.jwt")
handleRequest("Basic abc")
```

Tests three cases:

1. `"Bearer user-42.jwt"` — valid → `200: serving data for user "user-42.jwt"`.
2. `"Bearer expired.jwt"` — valid prefix, but expired payload → `401: please refresh your token` (because `errors.Is` matches).
3. `"Basic abc"` — no `"Bearer "` prefix → malformed → `401: malformed token "Basic abc": missing Bearer prefix`.

---

## Expected Output

```
200: serving data for user "user-42.jwt"
401: please refresh your token
401: malformed token "Basic abc": missing Bearer prefix
```

## Key Takeaways

1. **Sentinel error** — a package-level predefined error value (`errors.New`) for comparison.
2. **`%w` wrapping** — wrapping an old error with `fmt.Errorf("...: %w", err)`.
3. **`errors.Is`** — matching the original error even through a wrapped chain.
4. **`strings.HasPrefix` / `TrimPrefix`** — prefix check and removal.
5. **`if` with initializer** — calling and error-checking in one step.
