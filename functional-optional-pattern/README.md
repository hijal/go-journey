# functional-optional-pattern

Go-তে **functional options pattern** শেখার ছোট example — `type Option func(*T)`, variadic options, constructor-এর সাথে।

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

### Lines 5–9

```go
type Server struct {
	Port    int
	Timeout int
	IsHTTPS bool
}
```

`Server` struct — ৩টা field: `Port`, `Timeout`, `IsHTTPS`। struct-টা constructor-এর মাধ্যমে setup হবে।

### Line 11

```go
type Option func(*Server)
```

**`Option` type** — একটা function type: pointer-to-`Server` নেয়, কিছু return করে না। প্রতিটা option-ই একটা function যা `Server`-কে mutate করে।

### Lines 13–29

```go
func WithPort(port int) Option {
	return func(s *Server) {
		s.Port = port
	}
}

func WithTimeout(t int) Option {
	return func(s *Server) {
		s.Timeout = t
	}
}

func WithHTTPS(isHTTPS bool) Option {
	return func(s *Server) {
		s.IsHTTPS = isHTTPS
	}
}
```

তিনটা **functional option** — প্রতিটা "with" function একটা `Option` return করে:

- `WithPort(port int)` — একটা closure return করে যা `s.Port = port` set করে।
- `WithTimeout(t int)` — `s.Timeout = t` set করে।
- `WithHTTPS(isHTTPS bool)` — `s.IsHTTPS = isHTTPS` set করে।

**কী করছে:** প্রতিটা একটা **closure** — `port`/`t`/`isHTTPS`-কে capture করে; পরে option-টা call করলে ওই field mutate হয়। এটাই functional options-এর মূল: build-time-এ parameter capture, ব্যবহার-সময়ে apply।

### Lines 31–40

```go
func NewServer(opts ...Option) *Server {
	s := &Server{
		Port:    8080,
		Timeout: 30,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}
```

`NewServer` — constructor:

- **Default values** (`&Server{Port: 8080, Timeout: 30}`) — caller option না দিলে এগুলো।
- `opts ...Option` — **variadic** options।
- Loop-এ প্রতিটা `opt(s)` — pointer-এ apply, fields override।
- Pointer `*Server` return।

**Fluency:** caller-কে overriding fields-এর জন্য builder-এর মতো flexibility। Go-তে struct-এ optional config-এর standard idioms — positional arguments নয়, নাম-করা options।

### Lines 42–50

```go
func main() {
	srv := NewServer(WithPort(9090), WithHTTPS(true), WithTimeout(50))
	fmt.Printf("%+v\n", srv)

	srv = NewServer()
	fmt.Printf("%+v\n", srv)

	srv = NewServer(WithTimeout(100))
	fmt.Printf("%+v\n", srv)
}
```

তিনটা construction case:

1. সব options apply: `WithPort(9090), WithHTTPS(true), WithTimeout(50)` → `&{Port:9090 Timeout:50 IsHTTPS:true}`
2. No options: `NewServer()` → defaults → `&{Port:8080 Timeout:30 IsHTTPS:false}`
3. Only timeout: `WithTimeout(100)` → `&{Port:8080 Timeout:100 IsHTTPS:false}`

**`%+v`** — struct-কে field-name-সহ print।

### Line 51

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
&{Port:9090 Timeout:50 IsHTTPS:true}
&{Port:8080 Timeout:30 IsHTTPS:false}
&{Port:8080 Timeout:100 IsHTTPS:false}
```

## মূল শিক্ষা / Key Takeaways

1. **Functional options** — `type Option func(*T)` — config-container-এর signature।
2. **Closures** — প্রতিটা option-টা constructor-এর সময় capture-করা value apply করে।
3. **Defaults + overrides** — constructor-এ default, options দিয়ে selectively modify।
4. **`%+v`** — field names-সহ struct print।
5. **Variadic options** — `NewServer(WithX(...), ...)` — flexible, readable API।

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

### Lines 5–9

```go
type Server struct {
	Port    int
	Timeout int
	IsHTTPS bool
}
```

The `Server` struct — 3 fields: `Port`, `Timeout`, `IsHTTPS`. It will be set up via the constructor.

### Line 11

```go
type Option func(*Server)
```

**`Option` type** — a function type: takes a pointer-to-`Server`, returns nothing. Each option is a function that mutates the `Server`.

### Lines 13–29

```go
func WithPort(port int) Option {
	return func(s *Server) {
		s.Port = port
	}
}

func WithTimeout(t int) Option {
	return func(s *Server) {
		s.Timeout = t
	}
}

func WithHTTPS(isHTTPS bool) Option {
	return func(s *Server) {
		s.IsHTTPS = isHTTPS
	}
}
```

Three **functional options** — each "with" function returns an `Option`:

- `WithPort(port int)` — returns a closure that sets `s.Port = port`.
- `WithTimeout(t int)` — sets `s.Timeout = t`.
- `WithHTTPS(isHTTPS bool)` — sets `s.IsHTTPS = isHTTPS`.

**What's happening:** each returns a **closure** capturing `port`/`t`/`isHTTPS`; when the option is later called, that field is mutated. That's the essence of functional options: capture parameters at build time, apply them at use time.

### Lines 31–40

```go
func NewServer(opts ...Option) *Server {
	s := &Server{
		Port:    8080,
		Timeout: 30,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}
```

`NewServer` — the constructor:

- **Default values** (`&Server{Port: 8080, Timeout: 30}`) — used when the caller passes no options.
- `opts ...Option` — **variadic** options.
- The loop applies each `opt(s)` on the pointer, overriding fields.
- Returns `*Server`.

**Fluency:** gives the caller builder-like flexibility to override fields. This is the standard Go idiom for optional config on a struct — named options instead of positional arguments.

### Lines 42–50

```go
func main() {
	srv := NewServer(WithPort(9090), WithHTTPS(true), WithTimeout(50))
	fmt.Printf("%+v\n", srv)

	srv = NewServer()
	fmt.Printf("%+v\n", srv)

	srv = NewServer(WithTimeout(100))
	fmt.Printf("%+v\n", srv)
}
```

Three construction cases:

1. All options applied: `WithPort(9090), WithHTTPS(true), WithTimeout(50)` → `&{Port:9090 Timeout:50 IsHTTPS:true}`
2. No options: `NewServer()` → defaults → `&{Port:8080 Timeout:30 IsHTTPS:false}`
3. Timeout only: `WithTimeout(100)` → `&{Port:8080 Timeout:100 IsHTTPS:false}`

**`%+v`** — prints a struct with field names.

### Line 51

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
&{Port:9090 Timeout:50 IsHTTPS:true}
&{Port:8080 Timeout:30 IsHTTPS:false}
&{Port:8080 Timeout:100 IsHTTPS:false}
```

## Key Takeaways

1. **Functional options** — `type Option func(*T)` — the config-container signature.
2. **Closures** — each option applies a value captured at construction time.
3. **Defaults + overrides** — constructor defaults, options selectively modify.
4. **`%+v`** — prints a struct with field names.
5. **Variadic options** — `NewServer(WithX(...), ...)` — a flexible, readable API.