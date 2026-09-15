# cli-flag-parsing

Go-তে **standard `flag` package + `flag.Bool/Int/String` + `flag.Parse`** শেখার ছোট example — CLI flag parsing।

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
	"flag"
	"fmt"
)
```

- `flag` — standard flag parser।
- `fmt` — `Println`, `Printf`।

### Lines 8–12

```go
func main() {
	verbose := flag.Bool("verbose", false, "enable verbose logging")
	workers := flag.Int("workers", 4, "number of worker goroutines")
	output := flag.String("output", "stdout", "output destination")
	flag.Parse()
```

**Flag registration** — `flag.Bool/Int/String(নাম, default, help)`:

- প্রতিটা **pointer** return করে (`*bool`, `*int`, `*string`)।
- default: `false`, `4`, `"stdout"`।
- `flag.Parse()` — `os.Args` থেকে flag parse করে pointer-গুলো update করে।

*(`-h`/`--help` automatically usage-তালিকা দেখায়।)*

### Lines 14–16

```go
	if *verbose {
		fmt.Println("Verbose mode enabled")
	}
```

`*verbose` — **pointer dereference**।

### Line 18

```go
	fmt.Printf("starting %d workers, writing to %s\n", *workers, *output)
```

`*workers`, `*output` — parse-করা মান print।

---

## Expected Output

**Default (কোনো flag ছাড়া):**

```
starting 4 workers, writing to stdout
```

**`-verbose -workers=8 -output=file.log`-সহ:**

```
Verbose mode enabled
starting 8 workers, writing to file.log
```

## মূল শিক্ষা / Key Takeaways

1. **`flag.Bool/Int/String`** — name + default + help।
2. **Pointer-return** — value dereference করতে হয়।
3. **`flag.Parse()`** — os.Args process।
4. **Built-in `-h` help** — zero-config।

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
	"flag"
	"fmt"
)
```

- `flag` — the standard flag parser.
- `fmt` — for `Println`, `Printf`.

### Lines 8–12

```go
func main() {
	verbose := flag.Bool("verbose", false, "enable verbose logging")
	workers := flag.Int("workers", 4, "number of worker goroutines")
	output := flag.String("output", "stdout", "output destination")
	flag.Parse()
```

**Flag registration** — `flag.Bool/Int/String(name, default, help)`:

- Each returns a **pointer** (`*bool`, `*int`, `*string`).
- Defaults: `false`, `4`, `"stdout"`.
- `flag.Parse()` — parses `os.Args` and fills the pointers.

*(`-h`/`--help` shows the usage list automatically.)*

### Lines 14–16

```go
	if *verbose {
		fmt.Println("Verbose mode enabled")
	}
```

`*verbose` — **pointer dereference**.

### Line 18

```go
	fmt.Printf("starting %d workers, writing to %s\n", *workers, *output)
```

`*workers`, `*output` — printing the parsed values.

---

## Expected Output

**With no flags (defaults):**

```
starting 4 workers, writing to stdout
```

**With `-verbose -workers=8 -output=file.log`:**

```
Verbose mode enabled
starting 8 workers, writing to file.log
```

## Key Takeaways

1. **`flag.Bool/Int/String`** — name + default + help.
2. **Pointer return** — you must dereference the values.
3. **`flag.Parse()`** — processes `os.Args`.
4. **Built-in `-h` help** — zero configuration.