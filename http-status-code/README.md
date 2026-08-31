# http-status-code

Go-তে **exported vs unexported constants**, **`switch`** statement আর HTTP status code-কে readable string-এ convert করা শেখার ছোট example।

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

### Lines 5–11

```go
const (
	statusOk            = 200
	statusBadRequest    = 400
	statusUnauthorized  = 401
	statusNotFound      = 404
	statusInternalError = 500
)
```

HTTP status code-গুলো **unexported constants**:

- **নামে lowercase দিয়ে শুরু** (`statusOk`) → package-এর **বাইরে** access করা যায় না।
- মানগুলো পরিচিত: 200 OK, 400 Bad Request, 401 Unauthorized, 404 Not Found, 500 Internal Server Error।

### Lines 13–19

```go
const (
	msgOk            = "OK"
	msgBadRequest    = "Bad Request"
	msgUnauthorized  = "Unauthorized"
	msgNotFound      = "Not Found"
	msgInternalError = "Internal Server Error"
)
```

সংশ্লিষ্ট **message strings** — আবার unexported constants। প্রতিটি status code-এর নাম।

### Lines 21–36

```go
func statusText(code int) string {
	switch code {
	case statusOk:
		return msgOk
	case statusBadRequest:
		return msgBadRequest
	case statusUnauthorized:
		return msgUnauthorized
	case statusNotFound:
		return msgNotFound
	case statusInternalError:
		return msgInternalError
	default:
		return "Unknown Status"
	}
}
```

`statusText` — status code int-কে message string-এ convert করে:

- `switch code` — code-এর উপর switch। প্রতিটা `case`-এ constant-টা ধরে (known values)।
- `return msgOk` — corresponding message string return।
- `default` — অজানা code-এর জন্য `"Unknown Status"`।

**দেখলে** যে: `switch`-এর `case`-তে exported constant ব্যবহার করা হয়েছে, তবে সব unexported (`statusOk` not `StatusOk`)। কারণ সব একই package-এ।

### Lines 38–45

```go
func main() {
	codes := []int{
		statusOk, statusUnauthorized, statusNotFound, 503,
	}

	for _, code := range codes {
		fmt.Printf("%d -> %s\n", code, statusText(code))
	}
}
```

- `codes` — চারটা int: তিনটা known, একটা `503` (switch-তে নেই, `default`-এ যাবে)।
- loop-তে `statusText(code)` call করে প্রতিটা code-এর message print।

### Expected Output

```
200 -> OK
401 -> Unauthorized
404 -> Not Found
503 -> Unknown Status
```

## মূল শিক্ষা / Key Takeaways

1. **Unexported constant** — lowercase দিয়ে শুরু; package-এর ভেতরেই সীমাবদ্ধ।
2. **Switch statement** — value-এর উপর direct dispatch; `default` শেষে যায়।
3. **Readable code maps** — magic numbers (200, 404) বদলে named constants (`statusNotFound`)।
4. **Parallel const blocks** — code/message pair-কে আলাদা block-এ রাখা — map-like pattern।
5. **Unmatched case** — `503` switch-তে match করে না, `default` fallback।

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

### Lines 5–11

```go
const (
	statusOk            = 200
	statusBadRequest    = 400
	statusUnauthorized  = 401
	statusNotFound      = 404
	statusInternalError = 500
)
```

HTTP status codes as **unexported constants**:

- **Names start with lowercase** (`statusOk`) → inaccessible outside the package.
- The familiar values: 200 OK, 400 Bad Request, 401 Unauthorized, 404 Not Found, 500 Internal Server Error.

### Lines 13–19

```go
const (
	msgOk            = "OK"
	msgBadRequest    = "Bad Request"
	msgUnauthorized  = "Unauthorized"
	msgNotFound      = "Not Found"
	msgInternalError = "Internal Server Error"
)
```

The corresponding **message strings**, also as unexported constants — the human-readable name for each status code.

### Lines 21–36

```go
func statusText(code int) string {
	switch code {
	case statusOk:
		return msgOk
	case statusBadRequest:
		return msgBadRequest
	case statusUnauthorized:
		return msgUnauthorized
	case statusNotFound:
		return msgNotFound
	case statusInternalError:
		return msgInternalError
	default:
		return "Unknown Status"
	}
}
```

`statusText` — converts an int status code to a message string:

- `switch code` — switches on the code. Each `case` tests against a constant (known values).
- `return msgOk` — returns the corresponding message string.
- `default` — `"Unknown Status"` for an unrecognized code.

Note: the `case` labels use unexported constants (not `StatusOk`), because they're all in the same package.

### Lines 38–45

```go
func main() {
	codes := []int{
		statusOk, statusUnauthorized, statusNotFound, 503,
	}

	for _, code := range codes {
		fmt.Printf("%d -> %s\n", code, statusText(code))
	}
}
```

- `codes` — four ints: three known, one `503` (not in the switch, hits `default`).
- The loop prints `statusText(code)` for each.

---

## Expected Output

```
200 -> OK
401 -> Unauthorized
404 -> Not Found
503 -> Unknown Status
```

## Key Takeaways

1. **Unexported constants** — start with lowercase; confined to the package.
2. **Switch statement** — direct dispatch on a value; `default` is the fallback.
3. **Readable code maps** — named constants replace magic numbers (200, 404).
4. **Parallel const blocks** — code/message pairs in separate blocks as a map-like pattern.
5. **Unmatched case** — `503` isn't in any `case`, so `default` handles it.