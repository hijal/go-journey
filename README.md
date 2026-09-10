# go-journey

My daily Go practice. Each folder is a small, self-contained example from my learning.

## Run an example

```bash
go run ./payment-tracker
```

## Examples

| Folder | Concept | Docs |
| --- | --- | --- |
| `api-backend` | Struct + constructor method | [README](./api-backend/README.md) |
| `api-rate-limiter` | Closure factory + shared counter rate limit | [README](./api-rate-limiter/README.md) |
| `anonymous-function` | Anonymous function + IIFE pattern | [README](./anonymous-function/README.md) |
| `apply-discount` | Pass-by-value: function param is a copy | [README](./apply-discount/README.md) |
| `attendance` | `min`/`max` builtins + overtime split | [README](./attendance/README.md) |
| `background-jobs` | Loop, error handling, map lookup | [README](./background-jobs/README.md) |
| `bangla-text-processing` | `len` bytes vs `utf8.RuneCountInString` + `for range` byte index | [README](./bangla-text-processing/README.md) |
| `basic-types` | Basic types + rune as numeric code point | [README](./basic-types/README.md) |
| `bill-split` | Float constants + explicit cast + `%.2f` | [README](./bill-split/README.md) |
| `byte-size-constants` | iota + bit shift for KB/MB/GB constants | [README](./byte-size-constants/README.md) |
| `cart-discount` | Tagless `switch` tier discount + `Printf` width/precision | [README](./cart-discount/README.md) |
| `cart-methods-receivers` | Value vs pointer receivers in shopping cart | [README](./cart-methods-receivers/README.md) |
| `cart-stock` | Case-sensitive variables (`quantity` vs `Quantity`) | [README](./cart-stock/README.md) |
| `checkout-cart` | Float arithmetic + type conversion | [README](./checkout-cart/README.md) |
| `closure-fn` | Closure: function capturing outer variable | [README](./closure-fn/README.md) |
| `concurrent-webhook` | Goroutine + channel + WaitGroup | [README](./concurrent-webhook/README.md) |
| `config-loader` | Builtin `len` shadowing | [README](./config-loader/README.md) |
| `config-loading-iife` | IIFE + env vars with defaults | [README](./config-loading-iife/README.md) |
| `config-parsing-multi-return` | JSON parse + `(*Config, error)` + `runtime.Caller` | [README](./config-parsing-multi-return/README.md) |
| `custom-logger` | Variadic `...any` + `fmt.Sprint` custom logger | [README](./custom-logger/README.md) |
| `customer-tier-discount` | Function-as-value conditional assignment | [README](./customer-tier-discount/README.md) |
| `data-validation-pipeline` | Function type + variadic validator pipeline | [README](./data-validation-pipeline/README.md) |
| `devops-config-loader` | Package-level vars, const, funcs | [README](./devops-config-loader/README.md) |
| `deploy-tool` | `defer` for audit/timing log | [README](./deploy-tool/README.md) |
| `deterministic-map-report` | Map + `slices.Sort` keys for deterministic report | [README](./deterministic-map-report/README.md) |
| `devops-tooling` | Multiple assignment + `strconv.Atoi` | [README](./devops-tooling/README.md) |
| `double-entry-ledger` | Methods with pointer receiver, insufficient balance error | [README](./double-entry-ledger/README.md) |
| `dynamic-path-url-joiner` | Variadic + `strings.Trim`/`Join` for path join | [README](./dynamic-path-url-joiner/README.md) |
| `devide-zero` | Error handling + division by zero guard | [README](./devide-zero/README.md) |
| `emi-loan` | Numeric underscore + aligned `%.2f` table | [README](./emi-loan/README.md) |
| `env-variable-parsing` | `strconv` parsing with defaults | [README](./env-variable-parsing/README.md) |
| `find-max-num` | Variadic `...int` + linear max scan | [README](./find-max-num/README.md) |
| `fingerprint-config` | SHA-256 checksum + string/byte immutability | [README](./fingerprint-config/README.md) |
| `fn-value` | Function type + function as value | [README](./fn-value/README.md) |
| `fixed-tax-discount` | Constants + `fmt.Printf` formatting | [README](./fixed-tax-discount/README.md) |
| `free-shipping` | Boolean logic (`||` `&&` `!`) | [README](./free-shipping/README.md) |
| `functional-optional-pattern` | Functional options pattern + closures | [README](./functional-optional-pattern/README.md) |
| `group-variable` | Grouped `const`/`iota` + grouped `var` | [README](./group-variable/README.md) |
| `helpDesk` | Function pipeline of string transformers | [README](./helpDesk/README.md) |
| `higher-order-retry-helper` | Higher-order retry + `errors.Is` + `%w` | [README](./higher-order-retry-helper/README.md) |
| `http-handler-closure` | Closure over a map in HTTP handlers | [README](./http-handler-closure/README.md) |
| `http-status-code` | Unexported constants + `switch` for status mapping | [README](./http-status-code/README.md) |
| `iife-payment` | IIFE validation + message build | [README](./iife-payment/README.md) |
| `invoice-batch-scanner` | `continue` / `break` in loop | [README](./invoice-batch-scanner/README.md) |
| `invoice-json` | JSON marshal + unexported field | [README](./invoice-json/README.md) |
| `inventory-stock` | Unsigned underflow + safe compare | [README](./inventory-stock/README.md) |
| `insurance-eligibility-rule` | Nested struct + boolean eligibility | [README](./insurance-eligibility-rule/README.md) |
| `interget-overflow` | Integer overflow (`int32` vs `int64`) | [README](./interget-overflow/README.md) |
| `ip-subnet` | Bitwise ops + IP subnet membership | [README](./ip-subnet/README.md) |
| `ledger-entry` | Custom type (`Money`) | [README](./ledger-entry/README.md) |
| `leaderboard-tie-break` | `cmp.Compare` + nested switch tie-break | [README](./leaderboard-tie-break/README.md) |
| `merge-multiple-slices` | Variadic slices + `append(...)` merge | [README](./merge-multiple-slices/README.md) |
| `money-conversion-testing` | Table-driven unit test + `math.Round` | [README](./money-conversion-testing/README.md) |
| `monitoring-agent` | Float precision + epsilon compare | [README](./monitoring-agent/README.md) |
| `monitoring-system-temp` | Builtin `max` + `switch` with initializer | [README](./monitoring-system-temp/README.md) |
| `multi-error-aggregation` | Custom error type + nil-filter combine | [README](./multi-error-aggregation/README.md) |
| `multi-gateway-payment-proccess` | Interface + polymorphism | [README](./multi-gateway-payment-proccess/README.md) |
| `network-retry-mechanism` | Linear backoff + jitter retry with `%w` | [README](./network-retry-mechanism/README.md) |
| `new-account` | Zero values (`string`/`int64`/`bool`) | [README](./new-account/README.md) |
| `order-status-enum` | `iota` enum + `String()` method | [README](./order-status-enum/README.md) |
| `package-local-variable` | Package-level vs local variable scope | [README](./package-local-variable/README.md) |
| `parallel-health-check` | Goroutines + `sync.WaitGroup` parallel health checks | [README](./parallel-health-check/README.md) |
| `paybill` | Bill payment with custom error + guard clauses | [README](./paybill/README.md) |
| `paginated-fetching` | `for` + `break` paginated fetch + slice spread `append(... )` | [README](./paginated-fetching/README.md) |
| `pagination` | Ceiling division + offset pagination | [README](./pagination/README.md) |
| `payment-amount` | Custom type (`PaymentAmount`) + paisa-based money | [README](./payment-amount/README.md) |
| `payment-fee-calculator` | Constant + function | [README](./payment-fee-calculator/README.md) |
| `payment-retry` | Sentinel error + `errors.Is` + `%w` + exponential backoff | [README](./payment-retry/README.md) |
| `payment-risk-flagging` | `switch` with `fallthrough` | [README](./payment-risk-flagging/README.md) |
| `payment-status-poll` | `select` + channel timeout | [README](./payment-status-poll/README.md) |
| `payment-tracker` | Variable naming / underscores | [README](./payment-tracker/README.md) |
| `plan-based-rate-limit` | Maps + `slices`/`maps` stdlib (Go 1.21+) | [README](./plan-based-rate-limit/README.md) |
| `predicate-generics-filter` | Generics `Filter[T]` / `Map[T,U]` + predicates | [README](./predicate-generics-filter/README.md) |
| `product-sort-comparator` | `slices.SortFunc` + `cmp.Compare` multi-key sort | [README](./product-sort-comparator/README.md) |
| `rate-limit-config` | Exported vs unexported constants + unit suffix naming | [README](./rate-limit-config/README.md) |
| `named-naked-return` | Named return values + naked return | [README](./named-naked-return/README.md) |
| `request-validation` | Input validation with guard clauses + `strings` | [README](./request-validation/README.md) |
| `retail-loyalty-program` | `switch` + `fallthrough` tier benefits | [README](./retail-loyalty-program/README.md) |
| `retail-pos` | Closure with state: independent per-register totals | [README](./retail-pos/README.md) |
| `role-permission-authorization` | Variadic roles + membership check | [README](./role-permission-authorization/README.md) |
| `rune-byte` | Byte vs rune, UTF-8 decoding | [README](./rune-byte/README.md) |
| `retry-backoff-timing` | Linear backoff + `time.Duration` constants | [README](./retry-backoff-timing/README.md) |
| `retry-policy` | Retry loop + backoff with `time.Sleep` | [README](./retry-policy/README.md) |
| `sent-notification` | Variadic + slice spread for alert channels | [README](./sent-notification/README.md) |
| `settlement-reconciliation` | `goto` retry + labeled `break` | [README](./settlement-reconciliation/README.md) |
| `shared-wallet-closure` | Multiple closures sharing one captured balance | [README](./shared-wallet-closure/README.md) |
| `shipping-cost-tier` | `if/else if` tier-based shipping cost | [README](./shipping-cost-tier/README.md) |
| `signup-validator` | Function type + variadic validators | [README](./signup-validator/README.md) |
| `sql-clause-builder` | Variadic `...any` + SQL `IN` placeholders | [README](./sql-clause-builder/README.md) |
| `structured-logging-variadic` | Variadic `...any` key-value structured logger | [README](./structured-logging-variadic/README.md) |
| `switch-group` | Multi-value `case` log routing | [README](./switch-group/README.md) |
| `test-log-router` | Log destination routing + table test | [README](./test-log-router/README.md) |
| `testing` | `go/token.IsIdentifier` unit test | [README](./testing/README.md) |
| `token-validation` | Sentinel error + `errors.Is` + `%w` wrap | [README](./token-validation/README.md) |
| `transaction-fraud-screening` | `for range` + `continue` fraud flags | [README](./transaction-fraud-screening/README.md) |
| `transaction-rollback` | `defer` for rollback pattern | [README](./transaction-rollback/README.md) |
| `transaction-rollback-defer` | `defer`-based rollback on failed transfer | [README](./transaction-rollback-defer/README.md) |
| `unix-permission-bits` | Bit flags (rwx permission bitmask) | [README](./unix-permission-bits/README.md) |
| `validate-config` | Config validation with named return values | [README](./validate-config/README.md) |
| `variadic-function` | Variadic function (`...int`) | [README](./variadic-function/README.md) |
| `withdrawal-atm` | `if/else` balance validation | [README](./withdrawal-atm/README.md) |

## Notes

- Go version: 1.26
