# go-journey

My daily Go practice. Each folder is a small, self-contained example from my learning.

## Run an example

```bash
go run ./payment-tracker
```

## Examples

| Folder | Concept | Docs |
| --- | --- | --- |
| `account-number-masking` | `strings.Repeat` + tail-slice mask + edge guard | [README](./account-number-masking/README.md) |
| `api-backend` | Struct + constructor method | [README](./api-backend/README.md) |
| `api-pagination-slicing` | Offset formula + bounds-guard + clamp subslice | [README](./api-pagination-slicing/README.md) |
| `api-rate-limiter` | Closure factory + shared counter rate limit | [README](./api-rate-limiter/README.md) |
| `analytics-interface-embedding` | Interface embedding + `var _` compile-time check | [README](./analytics-interface-embedding/README.md) |
| `anonymous-function` | Anonymous function + IIFE pattern | [README](./anonymous-function/README.md) |
| `apply-discount` | Pass-by-value: function param is a copy | [README](./apply-discount/README.md) |
| `atomic-counter` | `atomic.Int64` lock-free safe counter | [README](./atomic-counter/README.md) |
| `attendance` | `min`/`max` builtins + overtime split | [README](./attendance/README.md) |
| `background-jobs` | Loop, error handling, map lookup | [README](./background-jobs/README.md) |
| `batch-payment-grouping` | Threshold split + pre-allocated grouping | [README](./batch-payment-grouping/README.md) |
| `bank-account-basics` | Struct zero value + `%+v` + field mutation | [README](./bank-account-basics/README.md) |
| `bank-balance-update` | Pointer-receiver method + validation + balance mutation | [README](./bank-balance-update/README.md) |
| `bangla-text-processing` | `len` bytes vs `utf8.RuneCountInString` + `for range` byte index | [README](./bangla-text-processing/README.md) |
| `basic-array` | Fixed-size `[4]int` + 4 range-loop variations | [README](./basic-array/README.md) |
| `basic-types` | Basic types + rune as numeric code point | [README](./basic-types/README.md) |
| `bill-split` | Float constants + explicit cast + `%.2f` | [README](./bill-split/README.md) |
| `buffer-pooling-sync` | `sync.Pool` reuse + `bytes.Buffer` + type assertion | [README](./buffer-pooling-sync/README.md) |
| `byte-size-constants` | iota + bit shift for KB/MB/GB constants | [README](./byte-size-constants/README.md) |
| `card-authorization-error` | Custom error type + `errors.As` typed classification | [README](./card-authorization-error/README.md) |
| `cart-discount` | Tagless `switch` tier discount + `Printf` width/precision | [README](./cart-discount/README.md) |
| `cart-methods-receivers` | Value vs pointer receivers in shopping cart | [README](./cart-methods-receivers/README.md) |
| `cart-stock` | Case-sensitive variables (`quantity` vs `Quantity`) | [README](./cart-stock/README.md) |
| `checkout-cart` | Float arithmetic + type conversion | [README](./checkout-cart/README.md) |
| `cicd-job-grouping` | Anonymous struct + status grouping map | [README](./cicd-job-grouping/README.md) |
| `cli-flag-parsing` | Standard `flag` package + Bool/Int/String + Parse | [README](./cli-flag-parsing/README.md) |
| `closure-fn` | Closure: function capturing outer variable | [README](./closure-fn/README.md) |
| `compare-versions` | Dotted-numeric compare + missing part = 0 + `max` + `Atoi` | [README](./compare-versions/README.md) |
| `concurrent-webhook` | Goroutine + channel + WaitGroup | [README](./concurrent-webhook/README.md) |
| `config-loader` | Builtin `len` shadowing | [README](./config-loader/README.md) |
| `config-loading-iife` | IIFE + env vars with defaults | [README](./config-loading-iife/README.md) |
| `config-merge-copy` | `slices.Clone` copy + zero-pad + `slices.Equal` | [README](./config-merge-copy/README.md) |
| `config-parsing-multi-return` | JSON parse + `(*Config, error)` + `runtime.Caller` | [README](./config-parsing-multi-return/README.md) |
| `currency-conversion-map` | Global rate map + comma-ok error + spread side-effect | [README](./currency-conversion-map/README.md) |
| `custom-io-writer` | `io.Writer` interface + custom `Write([]byte)` method | [README](./custom-io-writer/README.md) |
| `custom-logger` | Variadic `...any` + `fmt.Sprint` custom logger | [README](./custom-logger/README.md) |
| `customer-profile-nested` | Nested structs + slice field + dot-chain | [README](./customer-profile-nested/README.md) |
| `customer-tier-discount` | Function-as-value conditional assignment | [README](./customer-tier-discount/README.md) |
| `data-race-detector` | Deliberate `DATA RACE` demo (`-race` catches it) | [README](./data-race-detector/README.md) |
| `data-race-detector-solution` | `sync.Mutex`-guarded fix (always 1000) | [README](./data-race-detector-solution/README.md) |
| `data-validation-pipeline` | Function type + variadic validator pipeline | [README](./data-validation-pipeline/README.md) |
| `devops-config-loader` | Package-level vars, const, funcs | [README](./devops-config-loader/README.md) |
| `deploy-set-tooling` | Map-as-set deploy filter + `maps.Keys`/`slices.Sorted` | [README](./deploy-set-tooling/README.md) |
| `deploy-tool` | `defer` for audit/timing log | [README](./deploy-tool/README.md) |
| `deterministic-map-report` | Map + `slices.Sort` keys for deterministic report | [README](./deterministic-map-report/README.md) |
| `devops-tooling` | Multiple assignment + `strconv.Atoi` | [README](./devops-tooling/README.md) |
| `double-entry-ledger` | Methods with pointer receiver, insufficient balance error | [README](./double-entry-ledger/README.md) |
| `duplicate-order-detector` | Map seen-set + `slices.Contains` dedupe | [README](./duplicate-order-detector/README.md) |
| `dynamic-path-url-joiner` | Variadic + `strings.Trim`/`Join` for path join | [README](./dynamic-path-url-joiner/README.md) |
| `devide-zero` | Error handling + division by zero guard | [README](./devide-zero/README.md) |
| `emi-loan` | Numeric underscore + aligned `%.2f` table | [README](./emi-loan/README.md) |
| `env-variable-parsing` | `strconv` parsing with defaults | [README](./env-variable-parsing/README.md) |
| `error-http-status-mapping` | `errors.Is`/`As` + custom error → HTTP status | [README](./error-http-status-mapping/README.md) |
| `event-queue-type-switch` | `[]any` queue + type switch + `%T` dead-letter | [README](./event-queue-type-switch/README.md) |
| `failed-job-csv-report` | `Cut` key=value + `Contains` filter + `Builder` CSV | [README](./failed-job-csv-report/README.md) |
| `feature-flags-map` | Nil map read + comma-ok + `delete` | [README](./feature-flags-map/README.md) |
| `find-max-num` | Variadic `...int` + linear max scan | [README](./find-max-num/README.md) |
| `fingerprint-config` | SHA-256 checksum + string/byte immutability | [README](./fingerprint-config/README.md) |
| `fn-value` | Function type + function as value | [README](./fn-value/README.md) |
| `fixed-tax-discount` | Constants + `fmt.Printf` formatting | [README](./fixed-tax-discount/README.md) |
| `fluent-http-builder` | Fluent builder pattern + chain setters | [README](./fluent-http-builder/README.md) |
| `free-shipping` | Boolean logic (`||` `&&` `!`) | [README](./free-shipping/README.md) |
| `functional-optional-pattern` | Functional options pattern + closures | [README](./functional-optional-pattern/README.md) |
| `goroutine-ordering` | Non-deterministic goroutine order + main-exit kill | [README](./goroutine-ordering/README.md) |
| `grid-load-monitoring` | Parallel arrays + one-pass peak/low scan | [README](./grid-load-monitoring/README.md) |
| `group-variable` | Grouped `const`/`iota` + grouped `var` | [README](./group-variable/README.md) |
| `helpDesk` | Function pipeline of string transformers | [README](./helpDesk/README.md) |
| `higher-order-retry-helper` | Higher-order retry + `errors.Is` + `%w` | [README](./higher-order-retry-helper/README.md) |
| `http-client-functional-options` | Functional options (`Option func(*T)`) | [README](./http-client-functional-options/README.md) |
| `http-handler-closure` | Closure over a map in HTTP handlers | [README](./http-handler-closure/README.md) |
| `http-handler-interface` | `http.Handler` interface + ServeMux + slog | [README](./http-handler-interface/README.md) |
| `http-middleware-chain` | `Middleware` chain + wrap + statusRecorder | [README](./http-middleware-chain/README.md) |
| `http-status-code` | Unexported constants + `switch` for status mapping | [README](./http-status-code/README.md) |
| `iife-payment` | IIFE validation + message build | [README](./iife-payment/README.md) |
| `invoice-batch-scanner` | `continue` / `break` in loop | [README](./invoice-batch-scanner/README.md) |
| `invoice-line-totals` | Accumulated subtotal + `slices.Max` + float avg | [README](./invoice-line-totals/README.md) |
| `invoice-json` | JSON marshal + unexported field | [README](./invoice-json/README.md) |
| `inventory-cache-pointers` | `map[string]*T` cache + pointer dereference mutate | [README](./inventory-cache-pointers/README.md) |
| `inventory-stock` | Unsigned underflow + safe compare | [README](./inventory-stock/README.md) |
| `insurance-eligibility-rule` | Nested struct + boolean eligibility | [README](./insurance-eligibility-rule/README.md) |
| `interget-overflow` | Integer overflow (`int32` vs `int64`) | [README](./interget-overflow/README.md) |
| `ip-subnet` | Bitwise ops + IP subnet membership | [README](./ip-subnet/README.md) |
| `job-queue-drain` | `slices.DeleteFunc` in-place drain + retry list | [README](./job-queue-drain/README.md) |
| `job-worker-interface` | Interface consumer + sentinel error + `errors.Is` + `%w` | [README](./job-worker-interface/README.md) |
| `json-patch-optional` | Optional `*T` pointer fields + nil-guard partial update | [README](./json-patch-optional/README.md) |
| `key-value-config-parser` | `strings.Cut` + Trim/HasPrefix mini-INI parser | [README](./key-value-config-parser/README.md) |
| `ledger-entry` | Custom type (`Money`) | [README](./ledger-entry/README.md) |
| `leaderboard-tie-break` | `cmp.Compare` + nested switch tie-break | [README](./leaderboard-tie-break/README.md) |
| `linked-list` | Singly-linked list + pointer walk + `Stringer` | [README](./linked-list/README.md) |
| `log-analysis-slicing` | Slicing + shared backing array + `IndexFunc` | [README](./log-analysis-slicing/README.md) |
| `matrix-slice-report` | `[][]int` matrix + row/column aggregation | [README](./matrix-slice-report/README.md) |
| `merge-multiple-slices` | Variadic slices + `append(...)` merge | [README](./merge-multiple-slices/README.md) |
| `money-conversion-testing` | Table-driven unit test + `math.Round` | [README](./money-conversion-testing/README.md) |
| `monitoring-agent` | Float precision + epsilon compare | [README](./monitoring-agent/README.md) |
| `monitoring-system-temp` | Builtin `max` + `switch` with initializer | [README](./monitoring-system-temp/README.md) |
| `moving-average-monitor` | Builtin `max` clamp + sliding-window moving avg | [README](./moving-average-monitor/README.md) |
| `multi-error-aggregation` | Custom error type + nil-filter combine | [README](./multi-error-aggregation/README.md) |
| `multi-gateway-payment-proccess` | Interface + polymorphism | [README](./multi-gateway-payment-proccess/README.md) |
| `multi-payment-checkout` | Interface polymorphism + cents-to-currency + `[]PaymentMethod` | [README](./multi-payment-checkout/README.md) |
| `network-retry-mechanism` | Linear backoff + jitter retry with `%w` | [README](./network-retry-mechanism/README.md) |
| `new-account` | Zero values (`string`/`int64`/`bool`) | [README](./new-account/README.md) |
| `order-service-di` | Constructor DI + interfaces + `%w` wrapping | [README](./order-service-di/README.md) |
| `order-status-enum` | `iota` enum + `String()` method | [README](./order-status-enum/README.md) |
| `otp-verification` | Array `==` element-wise compare | [README](./otp-verification/README.md) |
| `package-local-variable` | Package-level vs local variable scope | [README](./package-local-variable/README.md) |
| `parallel-health-check` | Goroutines + `sync.WaitGroup` parallel health checks | [README](./parallel-health-check/README.md) |
| `paybill` | Bill payment with custom error + guard clauses | [README](./paybill/README.md) |
| `paginated-fetching` | `for` + `break` paginated fetch + slice spread `append(... )` | [README](./paginated-fetching/README.md) |
| `pagination` | Ceiling division + offset pagination | [README](./pagination/README.md) |
| `payment-amount` | Custom type (`PaymentAmount`) + paisa-based money | [README](./payment-amount/README.md) |
| `payment-fee-calculator` | Constant + function | [README](./payment-fee-calculator/README.md) |
| `payment-gateway-abstraction` | Interface abstraction + sentinel error + `%w` | [README](./payment-gateway-abstraction/README.md) |
| `payment-gateway-strategy` | Function-type strategy (Stripe/PayPal swap) | [README](./payment-gateway-strategy/README.md) |
| `payment-retry` | Sentinel error + `errors.Is` + `%w` + exponential backoff | [README](./payment-retry/README.md) |
| `payment-risk-flagging` | `switch` with `fallthrough` | [README](./payment-risk-flagging/README.md) |
| `payment-status-poll` | `select` + channel timeout | [README](./payment-status-poll/README.md) |
| `payment-tracker` | Variable naming / underscores | [README](./payment-tracker/README.md) |
| `permission-set-ops` | Map-as-set membership + clone/union + intersection | [README](./permission-set-ops/README.md) |
| `plan-based-rate-limit` | Maps + `slices`/`maps` stdlib (Go 1.21+) | [README](./plan-based-rate-limit/README.md) |
| `predicate-generics-filter` | Generics `Filter[T]` / `Map[T,U]` + predicates | [README](./predicate-generics-filter/README.md) |
| `priority-job-queue` | `container/heap` max-priority + tie-break + `heap.Fix` | [README](./priority-job-queue/README.md) |
| `product-sort-comparator` | `slices.SortFunc` + `cmp.Compare` multi-key sort | [README](./product-sort-comparator/README.md) |
| `product-value-pointer` | Value vs pointer struct mutation + copy | [README](./product-value-pointer/README.md) |
| `rate-limit-config` | Exported vs unexported constants + unit suffix naming | [README](./rate-limit-config/README.md) |
| `named-naked-return` | Named return values + naked return | [README](./named-naked-return/README.md) |
| `recent-activity-feed` | Builtin `min` tail-trim + `Clone`+`Reverse` | [README](./recent-activity-feed/README.md) |
| `request-validation` | Input validation with guard clauses + `strings` | [README](./request-validation/README.md) |
| `retail-loyalty-program` | `switch` + `fallthrough` tier benefits | [README](./retail-loyalty-program/README.md) |
| `retail-pos` | Closure with state: independent per-register totals | [README](./retail-pos/README.md) |
| `role-permission-authorization` | Variadic roles + membership check | [README](./role-permission-authorization/README.md) |
| `rune-byte` | Byte vs rune, UTF-8 decoding | [README](./rune-byte/README.md) |
| `retry-backoff-timing` | Linear backoff + `time.Duration` constants | [README](./retry-backoff-timing/README.md) |
| `retry-policy` | Retry loop + backoff with `time.Sleep` | [README](./retry-policy/README.md) |
| `sales-report-sort` | `slices.SortFunc` multi-key + `MaxFunc` + anonymous struct | [README](./sales-report-sort/README.md) |
| `search-query-matcher` | `Fields` normalize + `Contains` AND-match + case-insensitive | [README](./search-query-matcher/README.md) |
| `secrets-redaction-logging` | `Stringer`+`GoStringer`+`LogValuer` secret redaction | [README](./secrets-redaction-logging/README.md) |
| `select-channels` | `select` + `time.After` timeout + async fetch | [README](./select-channels/README.md) |
| `sent-notification` | Variadic + slice spread for alert channels | [README](./sent-notification/README.md) |
| `server-health-slice` | `append` + pre-allocation + dynamic growth | [README](./server-health-slice/README.md) |
| `service-config-constructor` | Embedded struct + validated constructor + sentinel | [README](./service-config-constructor/README.md) |
| `settlement-reconciliation` | `goto` retry + labeled `break` | [README](./settlement-reconciliation/README.md) |
| `shared-wallet-closure` | Multiple closures sharing one captured balance | [README](./shared-wallet-closure/README.md) |
| `shipping-cost-tier` | `if/else if` tier-based shipping cost | [README](./shipping-cost-tier/README.md) |
| `shopping-cart-slice` | Slice `len` + index access + range sum | [README](./shopping-cart-slice/README.md) |
| `signup-validator` | Function type + variadic validators | [README](./signup-validator/README.md) |
| `sku-price-lookup` | Map + comma-ok lookup present/absent path | [README](./sku-price-lookup/README.md) |
| `sku-validation` | Staged guard clauses + `Split` + `Atoi` + `Errorf` | [README](./sku-validation/README.md) |
| `sliding-window-limiter` | Ring-buffer sliding window + threshold limit | [README](./sliding-window-limiter/README.md) |
| `sms-template-renderer` | `strings.NewReplacer` multi-placeholder template render | [README](./sms-template-renderer/README.md) |
| `sql-clause-builder` | Variadic `...any` + SQL `IN` placeholders | [README](./sql-clause-builder/README.md) |
| `statement-sort-interface` | `sort.Interface` vs `slices.SortFunc` + `cmp.Compare` | [README](./statement-sort-interface/README.md) |
| `status-code-analytics` | Map count + `Sorted(Keys())` + max-scan | [README](./status-code-analytics/README.md) |
| `streaming-upload-checksum` | `TeeReader`+`LimitReader` streaming upload + sha256 | [README](./streaming-upload-checksum/README.md) |
| `structured-logging-variadic` | Variadic `...any` key-value structured logger | [README](./structured-logging-variadic/README.md) |
| `swap-function` | Pointer swap + parallel assignment + slice-element address | [README](./swap-function/README.md) |
| `switch-group` | Multi-value `case` log routing | [README](./switch-group/README.md) |
| `test-log-router` | Log destination routing + table test | [README](./test-log-router/README.md) |
| `testing` | `go/token.IsIdentifier` unit test | [README](./testing/README.md) |
| `thread-safe-counter` | `sync.Mutex` + `defer` unlock + `WaitGroup` | [README](./thread-safe-counter/README.md) |
| `ticket-category-count` | Map frequency-count zero-value idiom | [README](./ticket-category-count/README.md) |
| `token-validation` | Sentinel error + `errors.Is` + `%w` wrap | [README](./token-validation/README.md) |
| `topk-latency-dashboard` | Sort + tail-slice top-K + p95 approximation | [README](./topk-latency-dashboard/README.md) |
| `transaction-fraud-screening` | `for range` + `continue` fraud flags | [README](./transaction-fraud-screening/README.md) |
| `transaction-rollback` | `defer` for rollback pattern | [README](./transaction-rollback/README.md) |
| `transaction-rollback-defer` | `defer`-based rollback on failed transfer | [README](./transaction-rollback-defer/README.md) |
| `transaction-stringer` | `String()` method + `fmt.Stringer` + `%-7s` align | [README](./transaction-stringer/README.md) |
| `unbuffered-channel` | Unbuffered channel sync + `close` + `range` | [README](./unbuffered-channel/README.md) |
| `undo-stack-editor` | Snapshot-based undo stack + `slices.Clone` | [README](./undo-stack-editor/README.md) |
| `unix-permission-bits` | Bit flags (rwx permission bitmask) | [README](./unix-permission-bits/README.md) |
| `upload-filename-sanitizer` | Sanitize + `HasSuffix` allowlist + `LastIndex` split | [README](./upload-filename-sanitizer/README.md) |
| `validate-config` | Config validation with named return values | [README](./validate-config/README.md) |
| `variadic-function` | Variadic function (`...int`) | [README](./variadic-function/README.md) |
| `waitgroup-basics` | `sync.WaitGroup` classic Add/Done/Wait + modern `wg.Go` | [README](./waitgroup-basics/README.md) |
| `warehouse-inventory-receiver` | Sentinel error + `%w` wrap + pointer receiver | [README](./warehouse-inventory-receiver/README.md) |
| `webhook-callback-dispatcher` | EventBus Subscribe/Publish + `RWMutex` fan-out | [README](./webhook-callback-dispatcher/README.md) |
| `wrap-text-exercise` | Greedy word-wrap + `Fields` + `Builder` + width tracking | [README](./wrap-text-exercise/README.md) |
| `withdrawal-atm` | `if/else` balance validation | [README](./withdrawal-atm/README.md) |

## Notes

- Go version: 1.26
