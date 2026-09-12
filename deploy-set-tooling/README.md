# deploy-set-tooling

Go-তে **map-as-set (`map[string]struct{}`) + pointer-slice targets + lazy map init + `slices`/`maps` stdlib helpers** শেখার ছোট example — deploy plan (allowed envs filter)।

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

### Lines 3–7

```go
import (
	"fmt"
	"maps"
	"slices"
)
```

- `fmt` — `Println`, `Printf`।
- `maps` — `maps.Keys` (Go 1.21+)।
- `slices` — `Sort`, `Sorted` (Go 1.21+)।

### Lines 9–13

```go
type Server struct {
	Host string
	Env  string
	Up   bool
}
```

`Server` struct — host, env, up-flag — deploy-eligible server।

### Lines 15–18

```go
type DeployPlan struct {
	allowedEnvs map[string]struct{}
	targets     []*Server
}
```

`DeployPlan`:

- `allowedEnvs` — **map-as-set**: `map[string]struct{}` — value-টা zero-byte struct। Presence-ই ব্যাপার — এটা Go-র "set" idiom।
- `targets []*Server` — pointer-to-server slice।

### Lines 20–28

```go
func (d *DeployPlan) Allow(envs ...string) {
	if d.allowedEnvs == nil {
		d.allowedEnvs = make(map[string]struct{}, len(envs))
	}

	for _, e := range envs {
		d.allowedEnvs[e] = struct{}{}
	}
}
```

`Allow` — allowed environments register করে:

- **Lazy init:** `if d.allowedEnvs == nil` — map আগে না-থাকলে `make(...)`। nil map-এ write panic করত, তাই এই guard-টা দরকার।
- Variadic `envs ...string` — এক/অনেক env pass।
- `d.allowedEnvs[e] = struct{}{}` — key-টা set-এ add (value placeholder)।

### Lines 30–32

```go
func (d *DeployPlan) Add(s *Server) {
	d.targets = append(d.targets, s)
}
```

`Add` — target server register (`*Server`)। Pointer-এর মানে — বেশি memory copy নয়।

### Lines 34–44

```go
func (d *DeployPlan) Run() []string {
	eligible := []string{}

	for _, s := range d.targets {
		_, envOk := d.allowedEnvs[s.Env]
		if envOk && s.Up {
			eligible = append(eligible, s.Host)
		}
	}
	return eligible
}
```

`Run` — eligible targets বের করে:

- `_, envOk := d.allowedEnvs[s.Env]` — map lookup: key-টা allowed-এ আছে কিনা।
- `if envOk && s.Up` — env allowed **এবং** server up → eligible। Whitelist + health-check filter।

### Lines 46–49

```go
func main() {
	plan := new(DeployPlan)
	plan.Allow("prod", "staging")
```

- `new(DeployPlan)` — zero-value pointer (map এখনো nil — lazy init পরে)।
- Allow: `prod`, `staging`।

### Lines 50–59

```go
	servers := []*Server{
		{Host: "web-01", Env: "prod", Up: true},
		{Host: "web-02", Env: "prod", Up: false}, // down: skipped
		{Host: "dev-01", Env: "dev", Up: true},   // env not allowed: skipped
		{Host: "stg-01", Env: "staging", Up: true},
	}

	for _, s := range servers {
		plan.Add(s)
	}
```

৪টা server: web-01 ✓, web-02 (down ✗), dev-01 (dev allowed-na ✗), stg-01 ✓।

### Lines 61–63

```go
	picked := plan.Run()
	slices.Sort(picked)
	fmt.Println("deploy targets:", picked)
```

- `Run()` → `["web-01", "stg-01"]` → `slices.Sort` asc → `[stg-01 web-01]`।

### Lines 65–78

```go
	done := make(map[string]struct{})

	for _, h := range picked {
		done[h] = struct{}{}
	}

	host := "stg-01"

	if _, ok := done[host]; ok {
		fmt.Printf("%s already deployed, skipping\n", host)
	}

	fmt.Println("all done:", slices.Sorted(maps.Keys(done)))
```

**Deployed-set ম্যানেজমেন্ট:**

- `done` — map-as-set, picked hosts।
- `if _, ok := done[host]; ok` — membership test → `stg-01 already deployed, skipping`।
- `maps.Keys(done)` — keys-এর iterator → `slices.Sorted` → sorted slice: `[stg-01 web-01]`। `maps.Keys` + `slices.Sorted` যুগল ordering দেয়।

---

## Expected Output

```
deploy targets: [stg-01 web-01]
stg-01 already deployed, skipping
all done: [stg-01 web-01]
```

## মূল শিক্ষা / Key Takeaways

1. **Map-as-set** — `map[string]struct{}` — zero-byte value-ভিত্তিক membership।
2. **Lazy map init** — nil-check দিয়ে `make`, nil-map write panic এড়ানো।
3. **Pointer-slice targets** — `[]*Server`।
4. **Whitelist + health filter** — `envOk && s.Up`।
5. **`maps.Keys` + `slices.Sorted`** — sorted key extraction (Go 1.21+)।

---

---

<a name="english"></a>

## 🇬🇧 English Version

### Line 1

```go
package main
```

Declares an executable program (`main` package), runnable via `go run`.

### Lines 3–7

```go
import (
	"fmt"
	"maps"
	"slices"
)
```

- `fmt` — for `Println`, `Printf`.
- `maps` — for `maps.Keys` (Go 1.21+).
- `slices` — for `Sort`, `Sorted` (Go 1.21+).

### Lines 9–13

```go
type Server struct {
	Host string
	Env  string
	Up   bool
}
```

`Server` struct — host, env, up-flag — a deploy-eligible server.

### Lines 15–18

```go
type DeployPlan struct {
	allowedEnvs map[string]struct{}
	targets     []*Server
}
```

`DeployPlan`:

- `allowedEnvs` — **map-as-set**: `map[string]struct{}` — a zero-byte struct value. Presence matters — Go's "set" idiom.
- `targets []*Server` — a slice of server pointers.

### Lines 20–28

```go
func (d *DeployPlan) Allow(envs ...string) {
	if d.allowedEnvs == nil {
		d.allowedEnvs = make(map[string]struct{}, len(envs))
	}

	for _, e := range envs {
		d.allowedEnvs[e] = struct{}{}
	}
}
```

`Allow` — registers allowed environments:

- **Lazy init:** `if d.allowedEnvs == nil` — `make(...)` when the map doesn't exist yet. Writing to a nil map would panic, so this guard is needed.
- Variadic `envs ...string` — pass one or many envs.
- `d.allowedEnvs[e] = struct{}{}` — adds the key to the set (the value is a placeholder).

### Lines 30–32

```go
func (d *DeployPlan) Add(s *Server) {
	d.targets = append(d.targets, s)
}
```

`Add` — registers a target server (`*Server`). Pointers avoid copying large structs.

### Lines 34–44

```go
func (d *DeployPlan) Run() []string {
	eligible := []string{}

	for _, s := range d.targets {
		_, envOk := d.allowedEnvs[s.Env]
		if envOk && s.Up {
			eligible = append(eligible, s.Host)
		}
	}
	return eligible
}
```

`Run` — collects eligible targets:

- `_, envOk := d.allowedEnvs[s.Env]` — a map lookup: is the key in the allowed set.
- `if envOk && s.Up` — env allowed **and** server up → eligible. A whitelist + health filter.

### Lines 46–49

```go
func main() {
	plan := new(DeployPlan)
	plan.Allow("prod", "staging")
```

- `new(DeployPlan)` — a zero-value pointer (the map is still nil — lazy init later).
- Allow: `prod`, `staging`.

### Lines 50–59

```go
	servers := []*Server{
		{Host: "web-01", Env: "prod", Up: true},
		{Host: "web-02", Env: "prod", Up: false}, // down: skipped
		{Host: "dev-01", Env: "dev", Up: true},   // env not allowed: skipped
		{Host: "stg-01", Env: "staging", Up: true},
	}

	for _, s := range servers {
		plan.Add(s)
	}
```

Four servers: web-01 ✓, web-02 (down ✗), dev-01 (dev not allowed ✗), stg-01 ✓.

### Lines 61–63

```go
	picked := plan.Run()
	slices.Sort(picked)
	fmt.Println("deploy targets:", picked)
```

- `Run()` → `["web-01", "stg-01"]` → `slices.Sort` asc → `[stg-01 web-01]`.

### Lines 65–78

```go
	done := make(map[string]struct{})

	for _, h := range picked {
		done[h] = struct{}{}
	}

	host := "stg-01"

	if _, ok := done[host]; ok {
		fmt.Printf("%s already deployed, skipping\n", host)
	}

	fmt.Println("all done:", slices.Sorted(maps.Keys(done)))
```

**Deployed-set management:**

- `done` — a map-as-set holding the picked hosts.
- `if _, ok := done[host]; ok` — a membership test → `stg-01 already deployed, skipping`.
- `maps.Keys(done)` — an iterator over the keys → `slices.Sorted` → a sorted slice: `[stg-01 web-01]`. The `maps.Keys` + `slices.Sorted` pair gives ordering.

---

## Expected Output

```
deploy targets: [stg-01 web-01]
stg-01 already deployed, skipping
all done: [stg-01 web-01]
```

## Key Takeaways

1. **Map-as-set** — `map[string]struct{}` — zero-byte-value membership.
2. **Lazy map init** — `make` with a nil-check, avoiding the nil-map write panic.
3. **Pointer-slice targets** — `[]*Server`.
4. **Whitelist + health filter** — `envOk && s.Up`.
5. **`maps.Keys` + `slices.Sorted`** — sorted key extraction (Go 1.21+).