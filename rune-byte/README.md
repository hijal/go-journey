# rune-byte

Go-তে **`byte` vs `rune`** বোঝার ছোট example — UTF-8-এ koto byte, koto character, আর `range` দিয়ে rune-by-rune decode। Bengali, accented নাম, plain ASCII — নাম validation-এ character count।

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
	"unicode/utf8"
)
```

- `fmt` — output print।
- `unicode/utf8` — `utf8.RuneCountInString`: string-এ কতগুলো **rune** (characters) আছে, অথবা বাইট-গোনা।

### Lines 8–19

```go
func validateUsername(name string) error {
	const maxChars = 15
	byteLen := len(name)
	runeLen := utf8.RuneCountInString(name)
	fmt.Printf("%q -> bytes=%d, runes=%d\n", name, byteLen, runeLen)

	if runeLen > maxChars {
		return fmt.Errorf("username %q has %d characters, max is %d", name, runeLen, maxChars)
	}

	return nil
}
```

`validateUsername` — একটা username-কে check করে এবং `error` return করে:

- `const maxChars = 15` — অনুমোদিত সর্বোচ্চ **character** সংখ্যা।
- `byteLen := len(name)` — `len()` string-এর **byte**-সংখ্যা দেয় (অক্ষর নয়)।
- `runeLen := utf8.RuneCountInString(name)` — প্রকৃত **rune** (Unicode character) সংখ্যা।
- `%q` দিয়ে নাম, তারপর bytes/runes print।
- `runeLen > maxChars` হলে error (characters-র উপর ভিত্তি করে validate — byte-তে হলে `বাংলা`-র মতো multibyte নাম ভুলভাবে reject হতো)।

### Lines 21–23

```go
func main() {
	names := []string{
		"GoLang",
		"বাংলা",
		"José_García",
	}
```

`main`-এ ৩টা নাম test:

- `"GoLang"` — plain ASCII (5 ASCII char, প্রতি-1 byte)।
- `"বাংলা"` — Bengali; প্রতি অক্ষর UTF-8-এ 3 byte।
- `"José_García"` — `é` আর `í` accented; প্রতিটা 2 byte।

### Lines 24–34

```go
for _, name := range names {
	if err := validateUsername(name); err != nil {
		fmt.Println(" ->", err)
	} else {
		fmt.Println(" -> accepted")
	}
}
```

প্রতিটা নাম `validateUsername`-এ পাঠায়, error থাকলে print, নাহলে "accepted":

- `"GoLang"` → bytes=6, runes=6 → accepted
- `"বাংলা"` → bytes=15, runes=5 → accepted (৫টা character, 15 byte)
- `"José_García"` → bytes=13, runes=11 → accepted

লক্ষ্য করো: `বাংলা`-র byte-count (15) ঠিক max-এর সমান, কিন্তু rune-count (5) অনেক কম — তাই rune-based validate-এ pass করে। byte-based হলে borderline reject হতো।

### Line 36

```go
	fmt.Println("\nDecoding GoLang rune by rune:")
```

এখন UTF-8 decode দেখানো হয়। `\n` — খালি লাইন।

### Lines 37–39

```go
for i, r := range "GoLang" {
	fmt.Printf("  byte offset %2d -> rune %q (code point U+%04X)\n", i, r, r)
}
```

`range`-এ string-এর উপর গেলে প্রতিটা **rune**-তে iteration হয়:

- `i` — বাইট-offset (rune শুরু হওয়া জায়গা)।
- `r` — rune (Unicode code-point)।
- `%q` — প্রিন্টযোগ্য rune-টা, `U+%04X` — hex code-point।

ASCII-তে প্রতি rune 1 byte, তাই offset 0,1,2,3,4,5 — প্রতিটা এক একটা letter।

### Lines 41–44

```go
fmt.Println("\nDecoding বাংলা rune by rune:")
for i, r := range "বাংলা" {
	fmt.Printf("  byte offset %2d -> rune %q (code point U+%04X)\n", i, r, r)
}
```

`বাংলা`-তে প্রতি Bengali অক্ষর ৩ byte। তাই offset-গুলো: 0, 3, 6, 9, 12 — প্রতিটা rune-র **শুরু** জায়গা। দেখা যায়:

- `ব` U+09AC, `া` U+09BE, `ং` U+0982, `ল` U+09B2, `া` U+09BE

(5টা rune, 15 byte — কেন rune-count দরকারি সেটার live demonstration।)

### Lines 46–49

```go
fmt.Println("\nDecoding José_García rune by rune:")
for i, r := range "José_García" {
	fmt.Printf("  byte offset %2d -> rune %q (code point U+%04X)\n", i, r, r)
}
```

`José_García`-তে `é` (U+00E9) আর `í` (U+00ED) multibyte — `é`-এর পরে offset 5-এ jump। 13 byte, 11 rune।

### Line 50

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
"GoLang" -> bytes=6, runes=6
 -> accepted
"বাংলা" -> bytes=15, runes=5
 -> accepted
"José_García" -> bytes=13, runes=11
 -> accepted

Decoding GoLang rune by rune:
  byte offset  0 -> rune 'G' (code point U+0047)
  byte offset  1 -> rune 'o' (code point U+006F)
  byte offset  2 -> rune 'L' (code point U+004C)
  byte offset  3 -> rune 'a' (code point U+0061)
  byte offset  4 -> rune 'n' (code point U+006E)
  byte offset  5 -> rune 'g' (code point U+0067)

Decoding বাংলা rune by rune:
  byte offset  0 -> rune 'ব' (code point U+09AC)
  byte offset  3 -> rune 'া' (code point U+09BE)
  byte offset  6 -> rune 'ং' (code point U+0982)
  byte offset  9 -> rune 'ল' (code point U+09B2)
  byte offset 12 -> rune 'া' (code point U+09BE)

Decoding José_García rune by rune:
  byte offset  0 -> rune 'J' (code point U+004A)
  byte offset  1 -> rune 'o' (code point U+006F)
  byte offset  2 -> rune 's' (code point U+0073)
  byte offset  3 -> rune 'é' (code point U+00E9)
  byte offset  5 -> rune '_' (code point U+005F)
  byte offset  6 -> rune 'G' (code point U+0047)
  byte offset  7 -> rune 'a' (code point U+0061)
  byte offset  8 -> rune 'r' (code point U+0072)
  byte offset  9 -> rune 'c' (code point U+0063)
  byte offset 10 -> rune 'í' (code point U+00ED)
  byte offset 12 -> rune 'a' (code point U+0061)
```

## মূল শিক্ষা / Key Takeaways

1. **`byte` vs `rune`** — `byte` (1 byte, ASCII), `rune` (Unicode code point, 1–4 byte UTF-8)।
2. **`len(string)` = bytes** — character-সংখ্যা নয়; multibyte text-এ ভুল।
3. **`utf8.RuneCountInString`** — সঠিক character count।
4. **`range` over string** — rune-by-rune iterate করে বাইট-offset সহ।
5. **Validation by runes** — username/limit-এ characters-র উপর ইউজ-ভ্যালিডেট, না-হলে বহু-বাইট নাম ভুল reject হয়।

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
	"unicode/utf8"
)
```

- `fmt` — for output.
- `unicode/utf8` — for `utf8.RuneCountInString`: how many **runes** (characters) vs bytes are in a string.

### Lines 8–19

```go
func validateUsername(name string) error {
	const maxChars = 15
	byteLen := len(name)
	runeLen := utf8.RuneCountInString(name)
	fmt.Printf("%q -> bytes=%d, runes=%d\n", name, byteLen, runeLen)

	if runeLen > maxChars {
		return fmt.Errorf("username %q has %d characters, max is %d", name, runeLen, maxChars)
	}

	return nil
}
```

`validateUsername` — checks a username and returns an `error`:

- `const maxChars = 15` — max allowed **characters**.
- `byteLen := len(name)` — `len()` returns the **byte** count (not characters).
- `runeLen := utf8.RuneCountInString(name)` — the actual **rune** (Unicode character) count.
- Prints the name with `%q`, then bytes/runes.
- Returns an error if `runeLen > maxChars` (validates by characters — counting bytes would wrongly reject multibyte names like `বাংলা`).

### Lines 21–23

```go
func main() {
	names := []string{
		"GoLang",
		"বাংলা",
		"José_García",
	}
```

Three names tested in `main`:

- `"GoLang"` — plain ASCII (all 1 byte per char).
- `"বাংলা"` — Bengali; each letter is 3 bytes in UTF-8.
- `"José_García"` — accented `é` and `í`; each is 2 bytes.

### Lines 24–34

```go
for _, name := range names {
	if err := validateUsername(name); err != nil {
		fmt.Println(" ->", err)
	} else {
		fmt.Println(" -> accepted")
	}
}
```

Each name goes through `validateUsername`; on error print it, else "accepted":

- `"GoLang"` → bytes=6, runes=6 → accepted
- `"বাংলা"` → bytes=15, runes=5 → accepted (5 characters, 15 bytes)
- `"José_García"` → bytes=13, runes=11 → accepted

Note that `বাংলা`'s byte count (15) equals max exactly while its rune count (5) is far below — so rune-based validation passes. Byte-based validation would be a borderline reject.

### Line 36

```go
	fmt.Println("\nDecoding GoLang rune by rune:")
```

Now the UTF-8 decoding demo begins. `\n` — a blank line.

### Lines 37–39

```go
for i, r := range "GoLang" {
	fmt.Printf("  byte offset %2d -> rune %q (code point U+%04X)\n", i, r, r)
}
```

Ranging over a string iterates **runes**:

- `i` — the byte offset (where the rune starts).
- `r` — the rune (Unicode code point).
- `%q` — printable rune; `U+%04X` — hex code point.

In ASCII every rune is 1 byte, so offsets go 0,1,2,3,4,5 — one per letter.

### Lines 41–44

```go
fmt.Println("\nDecoding বাংলা rune by rune:")
for i, r := range "বাংলা" {
	fmt.Printf("  byte offset %2d -> rune %q (code point U+%04X)\n", i, r, r)
}
```

Each Bengali letter is 3 bytes, so offsets skip: 0, 3, 6, 9, 12 — the **start** of each rune. You can see:

- `ব` U+09AC, `া` U+09BE, `ং` U+0982, `ল` U+09B2, `া` U+09BE

(5 runes, 15 bytes — a live demo of why rune counts matter.)

### Lines 46–49

```go
fmt.Println("\nDecoding José_García rune by rune:")
for i, r := range "José_García" {
	fmt.Printf("  byte offset %2d -> rune %q (code point U+%04X)\n", i, r, r)
}
```

In `José_García`, `é` (U+00E9) and `í` (U+00ED) are multibyte — after `é` the offset jumps to 5. 13 bytes, 11 runes.

### Line 50

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
"GoLang" -> bytes=6, runes=6
 -> accepted
"বাংলা" -> bytes=15, runes=5
 -> accepted
"José_García" -> bytes=13, runes=11
 -> accepted

Decoding GoLang rune by rune:
  byte offset  0 -> rune 'G' (code point U+0047)
  byte offset  1 -> rune 'o' (code point U+006F)
  byte offset  2 -> rune 'L' (code point U+004C)
  byte offset  3 -> rune 'a' (code point U+0061)
  byte offset  4 -> rune 'n' (code point U+006E)
  byte offset  5 -> rune 'g' (code point U+0067)

Decoding বাংলা rune by rune:
  byte offset  0 -> rune 'ব' (code point U+09AC)
  byte offset  3 -> rune 'া' (code point U+09BE)
  byte offset  6 -> rune 'ং' (code point U+0982)
  byte offset  9 -> rune 'ল' (code point U+09B2)
  byte offset 12 -> rune 'া' (code point U+09BE)

Decoding José_García rune by rune:
  byte offset  0 -> rune 'J' (code point U+004A)
  byte offset  1 -> rune 'o' (code point U+006F)
  byte offset  2 -> rune 's' (code point U+0073)
  byte offset  3 -> rune 'é' (code point U+00E9)
  byte offset  5 -> rune '_' (code point U+005F)
  byte offset  6 -> rune 'G' (code point U+0047)
  byte offset  7 -> rune 'a' (code point U+0061)
  byte offset  8 -> rune 'r' (code point U+0072)
  byte offset  9 -> rune 'c' (code point U+0063)
  byte offset 10 -> rune 'í' (code point U+00ED)
  byte offset 12 -> rune 'a' (code point U+0061)
```

## Key Takeaways

1. **`byte` vs `rune`** — `byte` (1 byte, ASCII), `rune` (Unicode code point, 1–4 bytes UTF-8).
2. **`len(string)` = bytes** — not character count; wrong for multibyte text.
3. **`utf8.RuneCountInString`** — the correct character count.
4. **`range` over string** — iterates rune by rune with the byte offset.
5. **Validate by runes** — count characters for username/limits, or multibyte names get wrongly rejected.