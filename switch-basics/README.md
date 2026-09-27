# switch-basics

Go-তে **expression `switch`** — `runtime.GOOS`-এর ভিত্তিতে OS-নাম, case-গুলো top-down।

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
	"runtime"
)
```

`runtime` (GOOS)।

### Lines 8–18

```go
func main() {
	fmt.Print("Go runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("MacOS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)
	}
}
```

**Switch with init** — `os := runtime.GOOS` switch-এ scoped; value-`os`-এর সঙ্গে case-গুলো। Go-তে **implicit break** — মিলে-গেলেই থেমে যায় (fallthrough নেই), `default` আর সব মেললে।

---

## Expected Output (Linux এ)

```
Go runs on Linux
```

## মূল শিক্ষা / Key Takeaways

1. **`switch` with init** — scoped short variable।
2. **Top-down matching** — প্রথম matching case-ই চলে।
3. **No explicit break** — Go নিজেই break করে।
4. **`default`** — কোনো মেল না-পেলে fallback।

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
	"runtime"
)
```

`runtime` (GOOS).

### Lines 8–18

```go
func main() {
	fmt.Print("Go runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("MacOS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)
	}
}
```

**A switch with an init** — `os := runtime.GOOS` is scoped to the switch; the cases compare against `os`. In Go there's an **implicit break** — the first matching case wins (no fallthrough), and `default` catches the rest.

---

## Expected Output (on Linux)

```
Go runs on Linux
```

## Key Takeaways

1. **`switch` with init** — a scoped short variable.
2. **Top-down matching** — only the first match runs.
3. **No explicit break** — Go breaks on its own.
4. **`default`** — the fallback when nothing matches.