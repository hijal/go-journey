# unix-permission-bits

Go-তে **bitwise operators** (`|`, `&`, `|=`, `&^=`) আর **bit flags** দিয়ে Unix permission (rwx bitmask) শেখার ছোট example।

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
const (
	read    = 1 << 2
	write   = 1 << 1
	execute = 1 << 0
)
```

তিনটা **bit flag**, প্রতিটা আলাদা একটা bit-এ:

- `read` = `1 << 2` = `100` binary = **4**
- `write` = `1 << 1` = `010` binary = **2**
- `execute` = `1 << 0` = `001` binary = **1**

Unix permission-এর মতো (rwx → 4+2+1)। প্রতিটা flag একটা নির্দিষ্ট bit-position দখল করে।

### Line 11

```go
func main() {
```

Program-এর entry point।

### Line 12

```go
perm := read | write
```

**Bitwise OR (`|`)** — flag combine করে:

- `perm` = `read | write` = 4 | 2 = 6 = binary `110`

মানে: r+w, কিন্তু x নয়। একটা **bitmask** — ৩টা bit-এ কোন কোন permission on।

### Lines 13–14

```go
fmt.Println("permission:", perm)
fmt.Println("can read:", perm&read != 0)
```

- `permission: 6`
- `perm & read` = `110 & 100` = `100` = 4, যা `!= 0` → **true** — read permission on।

**Bitwise AND (`&`) + `!= 0`** — flag test-এর standard pattern: mask-এর সাথে AND করলে ওই bit on থাকলে nonzero হয়।

### Line 15

```go
fmt.Println("can execute:", perm&execute != 0)
```

`perm & execute` = `110 & 001` = `000` = 0 → `0 != 0` = **false** — execute এখনো off। Output: `can execute: false`।

### Lines 17–18

```go
perm |= execute
fmt.Println("after chmod +x:", perm, "can execute:", perm&execute != 0)
```

**`|=`** (OR-assign) — flag-টা **set** (add) করে:

- `perm |= execute` → `perm = 6 | 1 = 7` = binary `111`
- ওপরের print: `after chmod +x: 7 can execute: true`

এটা Unix-এর `chmod +x`-এর মতো — execute bit যোগ করা হয়।

### Lines 20–21

```go
perm &^= write
fmt.Println("after chmod -w:", perm, "can write:", perm&write != 0)
```

**`&^=`** (AND-NOT) — Go-র বিশেষ অপারেটর। `a &^ b` = `a & ^b` — `b`-তে যেসব bit on, `a`-র সেই bits **clear** (remove) করে:

- `perm &^= write` → `perm = 7 &^ 2` = `111`-হতে `101`-এ, write bit remove → `5`
- Output check: `perm & write` = `101 & 010` = `000` = 0 → `can write: false`। Unix-এর `chmod -w`-এর মতো।

### Line 22

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
permission: 6
can read: true
can execute: false
after chmod +x: 7 can execute: true
after chmod -w: 5 can write: false
```

## মূল শিক্ষা / Key Takeaways

1. **Bit flags** — `1 << n` দিয়ে আলাদা bit-এ permission অবস্থান।
2. **`|` combine / `|=` set** — flag যোগ করা (chmod +x)।
3. **`&` test** — `perm & flag != 0` — flag on কি না।
4. **`&^` clear** — `&^=` দিয়ে bit remove (chmod -w)।
5. **Bitmask** — একটা integer-এ অনেক boolean state pack করা।

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
const (
	read    = 1 << 2
	write   = 1 << 1
	execute = 1 << 0
)
```

Three **bit flags**, each occupying its own bit:

- `read` = `1 << 2` = `100` binary = **4**
- `write` = `1 << 1` = `010` binary = **2**
- `execute` = `1 << 0` = `001` binary = **1**

Like Unix permissions (rwx → 4+2+1). Each flag owns a distinct bit position.

### Line 11

```go
func main() {
```

Program entry point.

### Line 12

```go
perm := read | write
```

**Bitwise OR (`|`)** — combines flags:

- `perm` = `read | write` = 4 | 2 = 6 = binary `110`

Meaning: r+w, but no x. A **bitmask** — which of 3 bits are on.

### Lines 13–14

```go
fmt.Println("permission:", perm)
fmt.Println("can read:", perm&read != 0)
```

- `permission: 6`
- `perm & read` = `110 & 100` = `100` = 4, which is `!= 0` → **true** — read is set.

**Bitwise AND (`&`) + `!= 0`** — the standard flag test pattern: ANDing with the mask gives nonzero if that bit is set.

### Line 15

```go
fmt.Println("can execute:", perm&execute != 0)
```

`perm & execute` = `110 & 001` = `000` = 0 → `0 != 0` = **false** — execute not set yet. Output: `can execute: false`.

### Lines 17–18

```go
perm |= execute
fmt.Println("after chmod +x:", perm, "can execute:", perm&execute != 0)
```

**`|=`** (OR-assign) — **sets** (adds) the flag:

- `perm |= execute` → `perm = 6 | 1 = 7` = binary `111`
- prints: `after chmod +x: 7 can execute: true`

Like Unix `chmod +x` — adds the execute bit.

### Lines 20–21

```go
perm &^= write
fmt.Println("after chmod -w:", perm, "can write:", perm&write != 0)
```

**`&^=`** (AND-NOT) — a Go-specific operator. `a &^ b` = `a & ^b` — clears (removes) the bits of `b` from `a`:

- `perm &^= write` → `perm = 7 &^ 2` = from `111` to `101`, removing the write bit → `5`
- Output check: `perm & write` = `101 & 010` = `000` = 0 → `can write: false`. Like Unix `chmod -w`.

### Line 22

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
permission: 6
can read: true
can execute: false
after chmod +x: 7 can execute: true
after chmod -w: 5 can write: false
```

## Key Takeaways

1. **Bit flags** — `1 << n` positions each permission on a distinct bit.
2. **`|` combine / `|=` set** — adding a flag (chmod +x).
3. **`&` test** — `perm & flag != 0` — whether the flag is set.
4. **`&^` clear** — `&^=` removes bits (chmod -w).
5. **Bitmask** — pack many boolean states into one integer.