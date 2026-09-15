# json-patch-optional

Go-তে **optional pointer-field JSON patch (`*T` + nil-check) + partial update** শেখার ছোট example — PATCH-style update।

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
	"encoding/json"
	"fmt"
)
```

- `json` — `Unmarshal`।
- `fmt` — `Println`, `Printf`।

### Lines 8–12

```go
type User struct {
	Name  string
	Email string
	Age   int
}
```

Exist-করা user model।

### Lines 14–18

```go
type UserUpdateRequest struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
	Age   *int    `json:"age,omitempty"`
}
```

**Optional pointer fields** — প্রতিটা `*T`:

- JSON-তে key **absent** → pointer `nil` (field update হয় না)।
- key **present** → pointer allocated (update হবে)।

*(`omitempty` marshal-time; unmarshal-এ key-এর presence-ই pointer nil/present ঠিক করে — এটাই partial-update-এর ট্রিক।)*

### Lines 20–32

```go
func applyUpdate(existing *User, req UserUpdateRequest) {
	if req.Name != nil {
		existing.Name = *req.Name
	}

	if req.Email != nil {
		existing.Email = *req.Email
	}

	if req.Age != nil {
		existing.Age = *req.Age
	}
}
```

**Nil-guard apply** — প্রতিটি field: `nil` নয়-হলেই dereference (`*ptr`) করে overwrite; নাহলে skip।

### Lines 34–43

```go
	user := &User{Name: "bob", Email: "bob@example.com", Age: 30}

	payload := `{"age": 31}`

	var req UserUpdateRequest
	if err := json.Unmarshal([]byte(payload), &req); err != nil {
		fmt.Println("unmarshal error:", err)
		return
	}
```

- `payload`-এ **শুধু `age`** — Name/Email key নাই → তাদের pointer `nil`।
- Unmarshal error-handle।

### Lines 45–47

```go
	applyUpdate(user, req)
	fmt.Printf("%+v\n", user)
}
```

`%+v` — field-নামসহ struct print; আশা: শুধু Age 31, বাকি অপরিবর্তিত।

---

## Expected Output

```
&{Name:bob Email:bob@example.com Age:31}
```

## মূল শিক্ষা / Key Takeaways

1. **Pointer-field optionality** — absent key → `nil` pointer।
2. **Nil-guard dereference** — `*ptr` শুধু present-এ।
3. **Partial update** — শুধু পাঠানো fields change।
4. **`omitempty`** — marshal-time companion (nil → omitted)।
5. **Unmarshal error-handling** — invalid JSON guard।

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
	"encoding/json"
	"fmt"
)
```

- `json` — for `Unmarshal`.
- `fmt` — for `Println`, `Printf`.

### Lines 8–12

```go
type User struct {
	Name  string
	Email string
	Age   int
}
```

The existing user model.

### Lines 14–18

```go
type UserUpdateRequest struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
	Age   *int    `json:"age,omitempty"`
}
```

**Optional pointer fields** — each `*T`:

- The JSON key **absent** → pointer is `nil` (no update).
- The key **present** → pointer allocated (update happens).

*(`omitempty` matters at marshal time; on unmarshal the key's presence alone decides nil vs set — that is the partial-update trick.)*

### Lines 20–32

```go
func applyUpdate(existing *User, req UserUpdateRequest) {
	if req.Name != nil {
		existing.Name = *req.Name
	}

	if req.Email != nil {
		existing.Email = *req.Email
	}

	if req.Age != nil {
		existing.Age = *req.Age
	}
}
```

**Nil-guarded apply** — each field: overwrite only if `!= nil` (dereference with `*ptr`); otherwise skip.

### Lines 34–43

```go
	user := &User{Name: "bob", Email: "bob@example.com", Age: 30}

	payload := `{"age": 31}`

	var req UserUpdateRequest
	if err := json.Unmarshal([]byte(payload), &req); err != nil {
		fmt.Println("unmarshal error:", err)
		return
	}
```

- `payload` has **only `age`** — Name/Email keys absent → their pointers stay `nil`.
- Unmarshal errors are handled.

### Lines 45–47

```go
	applyUpdate(user, req)
	fmt.Printf("%+v\n", user)
}
```

`%+v` — struct print with field names; expect Age 31, everything else unchanged.

---

## Expected Output

```
&{Name:bob Email:bob@example.com Age:31}
```

## Key Takeaways

1. **Pointer-field optionality** — an absent key yields a `nil` pointer.
2. **Nil-guarded dereference** — `*ptr` only when present.
3. **Partial update** — only the sent fields change.
4. **`omitempty`** — the marshal-time companion (nil → omitted).
5. **Unmarshal error-handling** — guards invalid JSON.