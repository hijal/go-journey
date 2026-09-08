# bangla-text-processing

Go-তে **byte vs rune**, **`len` vs `utf8.RuneCountInString`**, আর **`for range`-এ byte position** শেখার ছোট example.

**ভাষা নির্বাচন করুন / Choose language:**

[বাংলা](#bangla) * [English](#english)

---

<a name="bangla"></a>

## বাংলা সংস্করণ

### Lines 1-6

```go
package main

import (
	"fmt"
	"unicode/utf8"
)
```

- `package main` -- একটা executable program.
- `fmt` -- output print করার জন্য.
- `unicode/utf8` -- UTF-8 string-এ কয়টা character (rune) আছে গোনার জন্য (`RuneCountInString`).

### Line 9

```go
sms := "গো শিখি"
```

`:=` দিয়ে `sms` নামে string variable বানানো হয়. Value-টা বাংলা text -- মনে রাখতে হবে Go-র string হলো **byte-এর sequence**, character-এর নয়. বাংলা প্রতিটা যুক্তাক্ষর/কার UTF-8-এ 3 byte নেয়, আর space 1 byte.

### Lines 10-11

```go
byteLen := len(sms)
runes := utf8.RuneCountInString(sms)
```

- `len(sms)` -- string-এ কয়টা **byte** আছে দেয়. এখানে 19.
- `utf8.RuneCountInString(sms)` -- কয়টা **character (rune)** আছে দেয়. এখানে 7 (`গ`, `ো`, ` `, `শ`, `ি`, `খ`, `ি`).

এটাই মূল পার্থক্য: byte count (19) != character count (7).

### Line 13

```go
fmt.Printf("bytes: %d, characters: %d\n", byteLen, runes)
```

`%d` placeholder দিয়ে দুটো সংখ্যা print করে. Output: `bytes: 19, characters: 7`.

### Lines 15-17

```go
for i, r := range sms {
	fmt.Printf("byte position %d: %c\n", i, r)
}
```

String-এর উপর `for range` করলে প্রতিটা iteration-এ দুটো value দেয়:

- `i` -- ওই rune-টা কোন **byte position** থেকে শুরু (0, 3, 6, 7, 10...).
- `r` -- ওই **rune** (character) নিজে, type `rune` (= `int32`).

`%c` দিয়ে rune-টাকে character হিসাবে print করা হয়.

**কেন position লাফায়:** `গ` 3 byte (0-2), তাই পরের rune শুরু 3-এ. Space 1 byte (6), তাই পরেরটা 7 থেকে শুরু. এজন্য index গোনা হয় byte দিয়ে, character দিয়ে নয়.

### Line 18

```go
}
```

Closing brace -- `main` function শেষ হয়.

---

## Expected Output

```
bytes: 19, characters: 7
byte position 0: গ
byte position 3: ো
byte position 6:  
byte position 7: শ
byte position 10: ি
byte position 13: খ
byte position 16: ি
```

## মূল শিক্ষা / Key Takeaways

1. **`len(s)` = byte count** -- character count নয়.
2. **`utf8.RuneCountInString`** -- আসল character (rune) সংখ্যা দেয়.
3. **`for range` on string** -- byte index + rune return করে.
4. **UTF-8 encoding** -- বাংলা অক্ষর 3 byte, space 1 byte; তাই index লাফিয়ে বাড়ে.
5. **`%c`** -- rune-কে character হিসাবে print করে.

---

---

<a name="english"></a>

##  English Version

### Lines 1-6

```go
package main

import (
	"fmt"
	"unicode/utf8"
)
```

- `package main` -- an executable program.
- `fmt` -- for console output.
- `unicode/utf8` -- for counting characters (runes) in a UTF-8 string (`RuneCountInString`).

### Line 9

```go
sms := "গো শিখি"
```

Creates a string variable `sms` with `:=`. The value is Bangla text -- remember a Go string is a **sequence of bytes**, not characters. Each Bangla letter/vowel sign takes 3 bytes in UTF-8, and a space takes 1 byte.

### Lines 10-11

```go
byteLen := len(sms)
runes := utf8.RuneCountInString(sms)
```

- `len(sms)` -- gives the number of **bytes** in the string. Here 19.
- `utf8.RuneCountInString(sms)` -- gives the number of **characters (runes)**. Here 7 (`গ`, `ো`, ` `, `শ`, `ি`, `খ`, `ি`).

That is the key difference: byte count (19) != character count (7).

### Line 13

```go
fmt.Printf("bytes: %d, characters: %d\n", byteLen, runes)
```

Prints both numbers with `%d` placeholders. Output: `bytes: 19, characters: 7`.

### Lines 15-17

```go
for i, r := range sms {
	fmt.Printf("byte position %d: %c\n", i, r)
}
```

A `for range` over a string yields two values per iteration:

- `i` -- the **byte position** where that rune starts (0, 3, 6, 7, 10...).
- `r` -- the **rune** (character) itself, of type `rune` (= `int32`).

`%c` prints the rune as a character.

**Why the position jumps:** `গ` is 3 bytes (0-2), so the next rune starts at 3. A space is 1 byte (6), so the next one starts at 7. The index counts bytes, not characters.

### Line 18

```go
}
```

Closing brace -- ends the `main` function.

---

## Expected Output

```
bytes: 19, characters: 7
byte position 0: গ
byte position 3: ো
byte position 6:  
byte position 7: শ
byte position 10: ি
byte position 13: খ
byte position 16: ি
```

## Key Takeaways

1. **`len(s)` = byte count** -- not the character count.
2. **`utf8.RuneCountInString`** -- gives the real character (rune) count.
3. **`for range` on string** -- returns byte index + rune.
4. **UTF-8 encoding** -- Bangla letters are 3 bytes, space is 1 byte; so the index jumps.
5. **`%c`** -- prints a rune as a character.
