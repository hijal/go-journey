# maps-basics

Go-তে **map basics** — `make`, literal, key-value struct-মান।

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

`fmt` (print)।

### Lines 5–10

```go
type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex
var mp map[string]Vertex
```

`map[string]Vertex` — string key → struct value।

### Lines 12–26

```go
func main() {
	m = make(map[string]Vertex)

	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	mp = map[string]Vertex{
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}

	fmt.Println(mp)
}
```

- **`make` mapping** — key-এ struct assign।
- **Map literal** — `key: value` pairs।
- Access → value, পুরো map print → `map[...]`।

---

## Expected Output

```
{40.68433 -74.39967}
map[Google:{37.42202 -122.08408}]
```

## মূল শিক্ষা / Key Takeaways

1. **`make(map[...]...)`** — usable map।
2. **Map literal** — inline data।
3. **Struct value** — যেকোনো type value হতে পারে।
4. **`m[key]` access** — default format।

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

`fmt` (print).

### Lines 5–10

```go
type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex
var mp map[string]Vertex
```

`map[string]Vertex` — string key → struct value.

### Lines 12–26

```go
func main() {
	m = make(map[string]Vertex)

	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	mp = map[string]Vertex{
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}

	fmt.Println(mp)
}
```

- **`make` map** — assign a struct to a key.
- **Map literal** — `key: value` pairs.
- Read a key, and print the whole map as `map[...]`.

---

## Expected Output

```
{40.68433 -74.39967}
map[Google:{37.42202 -122.08408}]
```

## Key Takeaways

1. **`make(map[...]...)`** — a usable map.
2. **Map literal** — inline data.
3. **Struct value** — any type can be a value.
4. **`m[key]` access** — default map formatting.