# streaming-upload-checksum

Go-তে **streaming I/O pipeline (`io.TeeReader` + `io.LimitReader` + `sha256`) + counting writer + upload size limit** শেখার ছোট example — file upload verification।

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

### Lines 3–12

```go
import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
)
```

- `bytes` — `Buffer`, `NewReader`।
- `crypto/sha256` — hasher।
- `encoding/hex` — hash→string।
- `io` — `LimitReader`, `TeeReader`, `Copy`, `MultiWriter`।
- `log` — `Fatal`।

### Lines 14–23

```go
type CountingWriter struct {
	W io.Writer
	N int64
}

func (c *CountingWriter) Write(p []byte) (int, error) {
	n, err := c.W.Write(p)
	c.N += int64(n)
	return n, err
}
```

**CountingWriter** — `io.Writer` implement (যেকোনো writer-wrap যেকোনো place-এ insert)। **Accumulate** (`+=`) — একাধিক chunk-এ total byte count।

### Line 25

```go
const maxUploadBytes = 1 << 20 // 1MB
```

**1MB limit** (`1 << 20`)।

### Line 27

```go
var ErrTooLarge = errors.New("upload exceeds size limit")
```

Sentinel error।

### Lines 29–45

```go
func SaveUpload(dst io.Writer, src io.Reader) (checksum string, size int64, err error) {
	hasher := sha256.New()
	limited := io.LimitReader(src, maxUploadBytes+1)

	tee := io.TeeReader(limited, hasher)

	counter := &CountingWriter{W: dst}
	if _, err := io.Copy(counter, tee); err != nil {
		return "", 0, fmt.Errorf("copy upload : %w", err)
	}

	if counter.N > maxUploadBytes {
		return "", 0, ErrTooLarge
	}

	return hex.EncodeToString(hasher.Sum(nil)), counter.N, nil
}
```

**Streaming pipeline** — একই read-এ সব:

1. `LimitReader(src, max+1)` — সর্বোচ্চ 1MB+1 byte।
2. `TeeReader(limited, hasher)` — store-এ write করার সাথে সাথে hash।
3. `io.Copy(counter, tee)` — counter-এর মাধ্যমে copy।
4. `counter.N > max` → `ErrTooLarge`।
5. `hex(hasher.Sum(nil))` — sha256 checksum।

কোথাও পুরো file memory-তে রাখা হয়নি — **O(1) memory streaming**।

### Lines 47–63

```go
	body := strings.NewReader("invoice-2026-09.pdf contents...")
	var storage bytes.Buffer

	sum, n, err := SaveUpload(&storage, body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("stored %d bytes, sha256=%s\n", n, sum)
	fmt.Printf("storage now holds: %q\n", storage.String())

	huge := bytes.NewReader(make([]byte, maxUploadBytes+10))

	_, _, err = SaveUpload(io.MultiWriter(&bytes.Buffer{}), huge)
	fmt.Println("large upload:", err)
```

- Acceptable upload → checksum + size + stored content।
- Oversized upload (`1MB+10`) → `upload exceeds size limit`।

---

## Expected Output

```
stored 31 bytes, sha256=e3885136e772f8a9b768105343acc81f23690c931ab02442262ca7c569a9eaeb
storage now holds: "invoice-2026-09.pdf contents..."
large upload: upload exceeds size limit
```

## মূল শিক্ষা / Key Takeaways

1. **`io.TeeReader`** — এক stream-এ দুই consumer (store + hash)।
2. **`io.LimitReader`** — সীমা-বাঁধা read (`max+1` trick)।
3. **Writer wrap (`+=`!)** — chunk-counting-এ accumulate দরকার।
4. **Streaming** — O(1) memory, bulk লোড হয় না।
5. **Sentinel error + `%w`** — `ErrTooLarge` + `errors.Is`।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Lines 3–12

```go
import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
)
```

- `bytes` — `Buffer`, `NewReader`.
- `crypto/sha256` — the hasher.
- `encoding/hex` — hash→string.
- `io` — `LimitReader`, `TeeReader`, `Copy`, `MultiWriter`.
- `log` — `Fatal`.

### Lines 14–23

```go
type CountingWriter struct {
	W io.Writer
	N int64
}

func (c *CountingWriter) Write(p []byte) (int, error) {
	n, err := c.W.Write(p)
	c.N += int64(n)
	return n, err
}
```

`CountingWriter` — implements `io.Writer` (wraps any writer, insertable anywhere). It **accumulates** (`+=`) — the total across multiple chunks.

### Line 25

```go
const maxUploadBytes = 1 << 20 // 1MB
```

The **1MB limit** (`1 << 20`).

### Line 27

```go
var ErrTooLarge = errors.New("upload exceeds size limit")
```

A sentinel error.

### Lines 29–45

```go
func SaveUpload(dst io.Writer, src io.Reader) (checksum string, size int64, err error) {
	hasher := sha256.New()
	limited := io.LimitReader(src, maxUploadBytes+1)

	tee := io.TeeReader(limited, hasher)

	counter := &CountingWriter{W: dst}
	if _, err := io.Copy(counter, tee); err != nil {
		return "", 0, fmt.Errorf("copy upload : %w", err)
	}

	if counter.N > maxUploadBytes {
		return "", 0, ErrTooLarge
	}

	return hex.EncodeToString(hasher.Sum(nil)), counter.N, nil
}
```

**A streaming pipeline** — everything in one pass:

1. `LimitReader(src, max+1)` — at most 1MB+1 bytes.
2. `TeeReader(limited, hasher)` — hashes while writing to the store.
3. `io.Copy(counter, tee)` — copies through the counter.
4. `counter.N > max` → `ErrTooLarge`.
5. `hex(hasher.Sum(nil))` — the sha256 checksum.

The whole file is never buffered in memory — **O(1)-memory streaming**.

### Lines 47–63

```go
	body := strings.NewReader("invoice-2026-09.pdf contents...")
	var storage bytes.Buffer

	sum, n, err := SaveUpload(&storage, body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("stored %d bytes, sha256=%s\n", n, sum)
	fmt.Printf("storage now holds: %q\n", storage.String())

	huge := bytes.NewReader(make([]byte, maxUploadBytes+10))

	_, _, err = SaveUpload(io.MultiWriter(&bytes.Buffer{}), huge)
	fmt.Println("large upload:", err)
```

- An acceptable upload → checksum + size + stored content.
- An oversized upload (`1MB+10`) → `upload exceeds size limit`.

---

## Expected Output

```
stored 31 bytes, sha256=e3885136e772f8a9b768105343acc81f23690c931ab02442262ca7c569a9eaeb
storage now holds: "invoice-2026-09.pdf contents..."
large upload: upload exceeds size limit
```

## Key Takeaways

1. **`io.TeeReader`** — two consumers from one stream (store + hash).
2. **`io.LimitReader`** — a bounded read (the `max+1` trick).
3. **Writer wrapping (`+=`!)** — chunk counting must accumulate.
4. **Streaming** — O(1) memory, no bulk loading.
5. **Sentinel error + `%w`** — `ErrTooLarge` + `errors.Is`.