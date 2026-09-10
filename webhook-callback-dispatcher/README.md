# webhook-callback-dispatcher

Go-তে **event bus / callback dispatcher** শেখার ছোট example — `Subscribe`/`Publish` pattern, `sync.RWMutex` দিয়ে concurrent-safe listener map, একটা event-এ একাধিক callback।

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
	"sync"
)
```

- `fmt` — `Printf`।
- `sync` — `sync.RWMutex`।

### Lines 8–11

```go
type Event struct {
	Name    string
	Payload map[string]any
}
```

`Event` — event-এর গঠন:

- `Name` — event-এর নাম (যেমন `"user_register"`)।
- `Payload map[string]any` — key-value payload (email, username...)।

### Lines 13–16

```go
type EventBus struct {
	mu        sync.RWMutex
	listeners map[string][]func(Event)
}
```

`EventBus` — event dispatch center:

- `mu sync.RWMutex` — map-টা একাধিক goroutine থেকে access করলে protect করে।
- `listeners` — event name → **callback slice**। একটা event-এ একাধিক listener থাকতে পারে।

### Lines 18–22

```go
func NewEventBus() *EventBus {
	return &EventBus{
		listeners: make(map[string][]func(Event)),
	}
}
```

`NewEventBus` — constructor: `listeners` map initialize করে (nil map-এ append panic করত)।

### Lines 24–28

```go
func (eb *EventBus) Subscribe(eventName string, callback func(Event)) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.listeners[eventName] = append(eb.listeners[eventName], callback)
}
```

`Subscribe` — একটা event-এ callback-টা register করে:

- `eb.mu.Lock()` — map-টা mutate করছি, তাই **write lock**।
- `defer eb.mu.Unlock()` — function শেষে unlock নিশ্চিত।
- `append` — callback-টা event-এর listener list-এ।

### Lines 30–38

```go
func (eb *EventBus) Publish(event Event) {
	eb.mu.RLock()
	defer eb.mu.RUnlock()
	if callbacks, ok := eb.listeners[event.Name]; ok {
		for _, cb := range callbacks {
			cb(event)
		}
	}
}
```

`Publish` — event-টা সব matching listeners-এ পাঠায়:

- `eb.mu.RLock()` — শুধু পড়ছি, তাই **read lock** (একাধিক reader একসাথে allowed, writer block)।
- `map lookup`: `callbacks, ok := eb.listeners[event.Name]` — event-এর নাম-এ matching listeners।
- `ok` true হলে: সব callback-কে `cb(event)` দিয়ে run — এটাই **dispatch**।

### Lines 40–49

```go
func main() {
	bus := NewEventBus()

	bus.Subscribe("user_register", func(e Event) {
		fmt.Printf("[Webhook 1] sending welcome email to %v\n", e.Payload["email"])
	})

	bus.Subscribe("user_register", func(e Event) {
		fmt.Printf("[Webhook 2] Notifying slack channel about user %v\n", e.Payload["username"])
	})
```

- `NewEventBus()` — bus তৈরি।
- ২টা callback একই event `"user_register"`-এ:
  - Webhook 1 → welcome email
  - Webhook 2 → slack notification
- একই event-এ একাধিক listener = **fan-out**।

### Lines 51–58

```go
	bus.Publish(Event{
		Name: "user_register",
		Payload: map[string]any{
			"username": "johndoe",
			"email":    "john@example.com",
		},
	})
}
```

`Publish` — event-টা পাঠানো হয়। Name-টা Subscribe-র **exact match** করতে হবে, নাহলে কোনো listener-ই triggered হবে না।

---

## Expected Output

```
[Webhook 1] sending welcome email to john@example.com
[Webhook 2] Notifying slack channel about user johndoe
```

## মূল শিক্ষা / Key Takeaways

1. **Subscribe/Publish pattern** — decoupled event dispatch।
2. **Multi-listener fan-out** — `map[string][]func(Event)`, এক event-এ একাধিক callback।
3. **`sync.RWMutex`** — write-তে `Lock`, read-তে `RLock` — concurrent-safe।
4. **Exact name match** — publish-র event name subscribe-র সাথে হুবহু মিলতে হবে।
5. **Constructor initialization** — `NewEventBus` দিয়ে map zero-value nil trap এড়ানো।

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
	"sync"
)
```

- `fmt` — for `Printf`.
- `sync` — for `sync.RWMutex`.

### Lines 8–11

```go
type Event struct {
	Name    string
	Payload map[string]any
}
```

`Event` — the structure of an event:

- `Name` — the event's name (e.g. `"user_register"`).
- `Payload map[string]any` — a key-value payload (email, username, ...).

### Lines 13–16

```go
type EventBus struct {
	mu        sync.RWMutex
	listeners map[string][]func(Event)
}
```

`EventBus` — an event dispatch center:

- `mu sync.RWMutex` — protects the map when accessed from multiple goroutines.
- `listeners` — event name → **slice of callbacks**. One event can have many listeners.

### Lines 18–22

```go
func NewEventBus() *EventBus {
	return &EventBus{
		listeners: make(map[string][]func(Event)),
	}
}
```

`NewEventBus` — a constructor: initializes the `listeners` map (appending to a nil map would panic).

### Lines 24–28

```go
func (eb *EventBus) Subscribe(eventName string, callback func(Event)) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.listeners[eventName] = append(eb.listeners[eventName], callback)
}
```

`Subscribe` — registers a callback for an event:

- `eb.mu.Lock()` — we're mutating the map, so a **write lock**.
- `defer eb.mu.Unlock()` — guarantees unlocking when the function ends.
- `append` — adds the callback to the event's listener list.

### Lines 30–38

```go
func (eb *EventBus) Publish(event Event) {
	eb.mu.RLock()
	defer eb.mu.RUnlock()
	if callbacks, ok := eb.listeners[event.Name]; ok {
		for _, cb := range callbacks {
			cb(event)
		}
	}
}
```

`Publish` — sends the event to all matching listeners:

- `eb.mu.RLock()` — we're only reading, so a **read lock** (multiple readers allowed concurrently, blocks writers).
- Map lookup: `callbacks, ok := eb.listeners[event.Name]` — matching listeners for the event name.
- If `ok`: run every callback with `cb(event)` — this is the **dispatch**.

### Lines 40–49

```go
func main() {
	bus := NewEventBus()

	bus.Subscribe("user_register", func(e Event) {
		fmt.Printf("[Webhook 1] sending welcome email to %v\n", e.Payload["email"])
	})

	bus.Subscribe("user_register", func(e Event) {
		fmt.Printf("[Webhook 2] Notifying slack channel about user %v\n", e.Payload["username"])
	})
```

- `NewEventBus()` — creates the bus.
- Two callbacks for the same event `"user_register"`:
  - Webhook 1 → welcome email
  - Webhook 2 → slack notification
- Multiple listeners for one event = **fan-out**.

### Lines 51–58

```go
	bus.Publish(Event{
		Name: "user_register",
		Payload: map[string]any{
			"username": "johndoe",
			"email":    "john@example.com",
		},
	})
}
```

`Publish` — sends the event. The name must **exactly match** the subscribed ones, or no listener will trigger.

---

## Expected Output

```
[Webhook 1] sending welcome email to john@example.com
[Webhook 2] Notifying slack channel about user johndoe
```

## Key Takeaways

1. **Subscribe/Publish pattern** — decoupled event dispatch.
2. **Multi-listener fan-out** — `map[string][]func(Event)`, many callbacks per event.
3. **`sync.RWMutex`** — `Lock` on writes, `RLock` on reads — concurrent-safe.
4. **Exact name match** — the published event name must exactly equal the subscribed one.
5. **Constructor initialization** — `NewEventBus` avoids the nil-map zero-value trap.