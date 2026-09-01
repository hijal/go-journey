# ip-subnet

Go-তে **bitwise operators** (`<<`, `|`, `&`) আর `uint32` দিয়ে IP-কে integer-এ pack করে **subnet membership** check শেখার ছোট example।

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

### Lines 5–7

```go
func ipToUnit32(a, b, c, d byte) uint32 {
	return uint32(a)<<24 | uint32(b)<<16 | uint32(c)<<8 | uint32(d)
}
```

`ipToUnit32` — ৪টা byte (IPv4 octet)-কে একটা `uint32`-এ **pack** করে:

- `a<<24` — প্রথম octet-কে 24 bit বামে shift (`192 << 24`)
- `b<<16` — দ্বিতীয় octet 16 bit
- `c<<8` — তৃতীয় octet 8 bit
- `d` — চতুর্থ octet unchanged
- `|` (OR) — এগুলো combine করে একটা 32-bit integer

IP `192.168.1.55` → `0xc0a80137` (hex). **কেন:** IP-কে numeric-এ ধরা bitwise mask-এর সাথে সহজে compare করা যায়।

### Line 9

```go
func main() {
```

Program-এর entry point।

### Lines 10–12

```go
ip := ipToUnit32(192, 168, 1, 55)
network := ipToUnit32(192, 168, 1, 0)
mask := uint32(0xFFFFFF00)
```

- `ip` = 192.168.1.55-এর packed value
- `network` = 192.168.1.0 (network address)
- `mask` = `/24` subnet mask = 255.255.255.0 (আগের 24 bit = network, শেষ 8 bit = host)। `0xFFFFFF00` hex literal।

### Line 14

```go
inSubnet := ip&mask == network&mask
```

**Bitwise AND (`&`)**:

- `ip & mask` — ip-এর host bits (last 8) zero করা — শুধু network part রাখে। 192.168.1.55 & 255.255.255.0 = 192.168.1.0
- `network & mask` — network-ও mask করা (same = 192.168.1.0)
- দুটো **equal** → `true`। **মানে:** ip-এর network portion-টা network-address-এর সাথে match করছে — তাই ip subnet-এ আছে।

### Line 15

```go
fmt.Printf("ip: %08x, masked: %08x\n", ip, ip&mask)
```

`%08x` — hex format, minimum 8 digit, leading zero-fill: `ip: c0a80137, masked: c0a80100`।

### Line 16

```go
fmt.Println("in subnet 192.168.1.0/24:", inSubnet)
```

Output: `in subnet 192.168.1.0/24: true`। 192.168.1.55 /24 subnet-এ আছে (same first 24 bits)।

### Lines 18–19

```go
otherIP := ipToUnit32(192, 168, 7, 55)
fmt.Println("192.168.7.55 in subnet:", otherIP&mask == network&mask)
```

`otherIP` = 192.168.7.55। `otherIP & 0xFFFFFF00` = 192.168.7.0, যা `network & mask` (= 192.168.1.0)-এর **সাথে match করে না** (এটা 1.x, ওটা 7.x) → `false`। Output: `192.168.7.55 in subnet: false`। তৃতীয় octet-এর পার্থক্যই কারণ।

### Line 20

```go
}
```

Closing brace — `main` function শেষ হয়।

---

## Expected Output

```
ip: c0a80137, masked: c0a80100
in subnet 192.168.1.0/24: true
192.168.7.55 in subnet: false
```

## মূল শিক্ষা / Key Takeaways

1. **Bit shifting `<<`** — octet-গুলোকে correct position-এ shift।
2. **Bitwise OR `|`** — ৪টা octet combine করে একটা `uint32`।
3. **Bitwise AND `&`** — mask প্রয়োগ করে network part-টা আলাদা করা।
4. **Subnet test** — `ip & mask == network & mask` — দুটো IP-র network part compare।
5. **Hex literal** — `0xFFFFFF00` — /24 mask-কে readable প্রকাশ।

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

### Lines 5–7

```go
func ipToUnit32(a, b, c, d byte) uint32 {
	return uint32(a)<<24 | uint32(b)<<16 | uint32(c)<<8 | uint32(d)
}
```

`ipToUnit32` — **packs** 4 bytes (IPv4 octets) into one `uint32`:

- `a<<24` — shift the first octet 24 bits left (`192 << 24`)
- `b<<16` — second octet 16 bits
- `c<<8` — third octet 8 bits
- `d` — fourth octet unchanged
- `|` (OR) — combines them into a single 32-bit integer

IP `192.168.1.55` → `0xc0a80137` (hex). **Why:** holding the IP numerically makes it easy to compare with a bitwise mask.

### Line 9

```go
func main() {
```

Program entry point.

### Lines 10–12

```go
ip := ipToUnit32(192, 168, 1, 55)
network := ipToUnit32(192, 168, 1, 0)
mask := uint32(0xFFFFFF00)
```

- `ip` = the packed value of 192.168.1.55
- `network` = 192.168.1.0 (the network address)
- `mask` = the /24 subnet mask = 255.255.255.0 (first 24 bits are network, last 8 are host). A hex literal `0xFFFFFF00`.

### Line 14

```go
inSubnet := ip&mask == network&mask
```

**Bitwise AND (`&`)**:

- `ip & mask` — zero out the ip's host bits (last 8) — keep only the network part. 192.168.1.55 & 255.255.255.0 = 192.168.1.0
- `network & mask` — mask the network too (same = 192.168.1.0)
- The two are **equal** → `true`. **Meaning:** the ip's network portion matches the network address — so it's in the subnet.

### Line 15

```go
fmt.Printf("ip: %08x, masked: %08x\n", ip, ip&mask)
```

`%08x` — hex format, at least 8 digits, zero-padded: `ip: c0a80137, masked: c0a80100`.

### Line 16

```go
fmt.Println("in subnet 192.168.1.0/24:", inSubnet)
```

Output: `in subnet 192.168.1.0/24: true`. 192.168.1.55 is in the /24 subnet (same first 24 bits).

### Lines 18–19

```go
otherIP := ipToUnit32(192, 168, 7, 55)
fmt.Println("192.168.7.55 in subnet:", otherIP&mask == network&mask)
```

`otherIP` = 192.168.7.55. `otherIP & 0xFFFFFF00` = 192.168.7.0, which **doesn't match** `network & mask` (= 192.168.1.0) (it's 1.x vs 7.x) → `false`. Output: `192.168.7.55 in subnet: false`. The third octet difference is why.

### Line 20

```go
}
```

Closing brace — ends the `main` function.

---

## Expected Output

```
ip: c0a80137, masked: c0a80100
in subnet 192.168.1.0/24: true
192.168.7.55 in subnet: false
```

## Key Takeaways

1. **Bit shifting `<<`** — moves octets into the correct position.
2. **Bitwise OR `|`** — combines 4 octets into one `uint32`.
3. **Bitwise AND `&`** — applies the mask to isolate the network part.
4. **Subnet test** — `ip & mask == network & mask` — compare the network parts of two IPs.
5. **Hex literal** — `0xFFFFFF00` — expresses the /24 mask readably.