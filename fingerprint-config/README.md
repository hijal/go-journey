# fingerprint-config

Go-তে **`crypto/sha256`** দিয়ে config-এর **fingerprint (checksum)** বানানো এবং **`string` vs `[]byte`**-এর immutability বোঝার ছোট example।

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
	"crypto/sha256"
	"fmt"
)
```

দুটো package import করা হয়:

- `crypto/sha256` — SHA-256 hash function; config-এর uniqueness check করতে।
- `fmt` — console-এ output print করার জন্য।

### Line 8

```go
func main() {
```

Program-এর entry point।

### Lines 9–13

```go
config := `
server:
  port: 8080
  timeout: 30s
`
```

`config` — একটা **raw string literal** (backtick `` ` `` দিয়ে)। প্রথম character নতুন line (`\n`) — তাই `config[:5]`-এর আগে blank line আসে। এটা অ্যাপের config file/content-এর stand-in।

### Line 14

```go
rawBytes := []byte(config)
```

`config` (string) থেকে **conversion**: `[]byte(config)` — string-এর byte-গুলো কপি করে একটা byte slice বানায়। 

> **গুরুত্বপূর্ণ:** conversion-এ **copy** হয়। `rawBytes` mutate করলে মূল `config` string-টা অপরিবর্তিত থাকে — string আর byte slice আলাদা জিনিস।

### Line 16

```go
var checksum [32]byte = sha256.Sum256(rawBytes)
```

- `sha256.Sum256(rawBytes)` — byte slice-এর উপর SHA-256 hash compute করে; result হলো **fixed-size `[32]byte` array** (256 bit = 32 byte)।
- `var checksum [32]byte = ...` — সেই array-টা `checksum`-এ রাখে।
- এটাই config-এর **fingerprint** — মূল config-এ যেকোনো এক bit-ও বদলালে এই checksum পুরো বদলে যায়। তাই দূরের কোনো system-কে বলে দেওয়া যায় "এই config-টা দেখো identical কিনা"।

### Line 17

```go
fmt.Printf("Config fingerprint: %x\n", checksum)
```

`%x` verb দিয়ে checksum-এর byte-গুলো hex-এ print করে: `Config fingerprint: 1a1a...ba9`। (32 byte → 64 hex digit।)

### Line 19

```go
rawBytes[1] = 'X'
```

`rawBytes`-এর **2nd byte**-টা (`X`) বদলে দেয়। মজার কথা: এটা legal — `rawBytes` mutable (modifiable) slice। কিন্তু মূল `config` string-এর উপর এর কোনো effect নেই (copy হওয়ায়)।

### Lines 20–21

```go
fmt.Println("Mutated bytes did not change the original string")
fmt.Println("  original string still starts with:", config[:5])
```

- প্রথম print: "byte mutate করলেও মূল string বদলায়নি" — immutability-র proof।
- দ্বিতীয় print: `config[:5]` — config-এর প্রথম ৫ character। মনে রেখো config raw string-টা **newline দিয়ে শুরু**, তাই output-এ আগে একটা blank line দেখাবে, তারপর `serv` (config-এর `server:`-এর শুরু)।

### Lines 23–24

```go
backToString := string(rawBytes)
fmt.Println("Bytes converted back to string, first line differs:", backToString[:5])
```

`string(rawBytes)` — mutate-করা byte slice-কে আবার **string-এ convert** করে। এবার result-টা আলাদা: `X` বসেছে 2nd character-এ, তাই `backToString[:5]`-এ দেখায় `Xerv`। এটা প্রমাণ করে `rawBytes` mutate করা embed-এ মূল string-টা বদলে না, কিন্তু slice-টা নিজেই আগে যা ছিল সেটা নয়।

### Line 25

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
Config fingerprint: 1a1a6c537e55486e56b04e6d45cc25f9bedd00addb704f6f0da2d0203f490ba9
Mutated bytes did not change the original string
  original string still starts with: 
serv
Bytes converted back to string, first line differs: 
Xerv
```

## মূল শিক্ষা / Key Takeaways

1. **`sha256.Sum256`** — fixed-size `[32]byte` checksum; content integrity verify করতে।
2. **`%x` formatting** — byte array-কে hex-এ print করতে।
3. **`string` → `[]byte` copy** — conversion-এ data কপি হয়; তাই mutate করলে আসল string বদলায় না।
4. **`[]byte` mutable, `string` immutable** — slice edit করা যায়, string edit করা যায় না।
5. **Fingerprint use-case** — config-এর পরিবর্তন detect করা, version comparing করা।

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
	"crypto/sha256"
	"fmt"
)
```

Two packages are imported:

- `crypto/sha256` — the SHA-256 hash function; for checking config uniqueness.
- `fmt` — for printing output to the console.

### Line 8

```go
func main() {
```

Program entry point.

### Lines 9–13

```go
config := `
server:
  port: 8080
  timeout: 30s
`
```

`config` — a **raw string literal** (delimited by backticks). Its first character is a newline (`\n`), so `config[:5]` has a leading blank line. It stands in for an app's config file/content.

### Line 14

```go
rawBytes := []byte(config)
```

Converts `config` (a string) to a **byte slice**: `[]byte(config)` copies the string's bytes into a slice.

> **Important:** the conversion **copies** the bytes. Mutating `rawBytes` leaves the original `config` string untouched — a string and a byte slice are different things.

### Line 16

```go
var checksum [32]byte = sha256.Sum256(rawBytes)
```

- `sha256.Sum256(rawBytes)` — computes the SHA-256 hash of the byte slice; the result is a **fixed-size `[32]byte` array** (256 bits = 32 bytes).
- `var checksum [32]byte = ...` — stores that array in `checksum`.
- This is the config's **fingerprint** — flip a single bit of the config and the checksum changes completely. So you can tell a remote system "check if this config matches".

### Line 17

```go
fmt.Printf("Config fingerprint: %x\n", checksum)
```

The `%x` verb prints the checksum's bytes in hex: `Config fingerprint: 1a1a...ba9`. (32 bytes → 64 hex digits.)

### Line 19

```go
rawBytes[1] = 'X'
```

Replaces the slice's **2nd byte** with `X`. The interesting part: this is legal — `rawBytes` is a mutable slice. But it has no effect on the original `config` string (because of the copy).

### Lines 20–21

```go
fmt.Println("Mutated bytes did not change the original string")
fmt.Println("  original string still starts with:", config[:5])
```

- First print: "mutating the bytes didn't change the original string" — proof of immutability.
- Second print: `config[:5]` — the first 5 characters of config. Remember the raw string starts with a newline, so the output shows a blank line first, then `serv` (the start of config's `server:`).

### Lines 23–24

```go
backToString := string(rawBytes)
fmt.Println("Bytes converted back to string, first line differs:", backToString[:5])
```

`string(rawBytes)` — converts the mutated byte slice **back to a string**. This time the result differs: an `X` sits in the 2nd character, so `backToString[:5]` shows `Xerv`. This proves mutating `rawBytes` doesn't change the original string, but the slice itself is no longer what it started as.

### Line 25

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
Config fingerprint: 1a1a6c537e55486e56b04e6d45cc25f9bedd00addb704f6f0da2d0203f490ba9
Mutated bytes did not change the original string
  original string still starts with: 
serv
Bytes converted back to string, first line differs: 
Xerv
```

## Key Takeaways

1. **`sha256.Sum256`** — fixed-size `[32]byte` checksum for verifying content integrity.
2. **`%x` formatting** — prints a byte array in hex.
3. **`string` → `[]byte` copies** — the conversion copies data, so mutating doesn't change the original string.
4. **`[]byte` mutable vs `string` immutable** — slices can be edited, strings can't.
5. **Fingerprint use-case** — detect config changes, compare versions.