# odata-query-builder

Zero-dependency Go library for building OData v4 query parameters with a fluent builder API.

## Installation

```bash
go get github.com/chewcw/odata-query-builder
```

## Quick Start

```go
package main

import (
	"fmt"
	"github.com/chewcw/odata-query-builder"
)

func main() {
	query := odata.New().
		Select("name", "email").
		Filter("status").Eq("active").
		Top(10).
		Build()

	fmt.Println(query)
	// Output: $select=name,email&$filter=status eq 'active'&$top=10
}
```

## Features

- **Zero dependencies** — stdlib only
- **Fluent API** — chainable builder methods
- **Full OData v4 support** — `$select`, `$filter`, `$expand`, `$orderby`, `$top`, `$skip`, `$count`, `$search`
- **Nested `$expand`** — with sub-query callbacks
- **Type-safe values** — strings auto-quoted/escaped, numbers/booleans unquoted
- **Nil-safe** — nil pointers for `$top`, `$skip`, `$count` are omitted (no `$top=0` noise)

## API Reference

### QueryBuilder

```go
func New() *QueryBuilder
```

Creates a new query builder.

#### Selection

```go
func (qb *QueryBuilder) Select(fields ...string) *QueryBuilder
func (qb *QueryBuilder) And(fields ...string) *QueryBuilder
```

Add fields to `$select`. `And` is an alias for `Select`.

```go
odata.New().Select("name", "email").Build()
// $select=name,email
```

#### Filtering

```go
func (qb *QueryBuilder) Filter(field string) *FilterBuilder
func (qb *QueryBuilder) OrFilter(field string) *FilterBuilder
```

Start a filter condition. Subsequent calls to `Filter` are implicitly ANDed; use `OrFilter` for OR logic.

```go
odata.New().
    Filter("status").Eq("active").
    Filter("age").Gt(18).
    Build()
// $filter=status eq 'active' and age gt 18

odata.New().
    Filter("status").Eq("active").
    OrFilter("role").Eq("admin").
    Build()
// $filter=status eq 'active' or role eq 'admin'
```

#### Filter Operators (on `FilterBuilder`)

| Method | OData Operator | Example |
|--------|----------------|---------|
| `Eq(value)` | `eq` | `Filter("status").Eq("active")` |
| `Ne(value)` | `ne` | `Filter("status").Ne("deleted")` |
| `Gt(value)` | `gt` | `Filter("price").Gt(100)` |
| `Ge(value)` | `ge` | `Filter("price").Ge(0)` |
| `Lt(value)` | `lt` | `Filter("price").Lt(50.5)` |
| `Le(value)` | `le` | `Filter("price").Le(99.99)` |
| `Contains(s)` | `contains` | `Filter("name").Contains("john")` |
| `StartsWith(s)` | `startswith` | `Filter("name").StartsWith("A")` |
| `EndsWith(s)` | `endswith` | `Filter("email").EndsWith(".com")` |
| `In(values...)` | `in` | `Filter("id").In(1, 2, 3)` |
| `Has(value)` | `has` | `Filter("style").Has("Sales.Color'Yellow'")` |
| `Not()` | `not` | `Filter("status").Not().Eq("deleted")` |

```go
// String values are auto-quoted and escaped (O'Brien -> 'O''Brien')
Filter("name").Eq("O'Brien")  // name eq 'O''Brien'

// Numbers and booleans are unquoted
Filter("age").Eq(42)          // age eq 42
Filter("enabled").Eq(true)    // enabled eq true
```

#### Expand

```go
func (qb *QueryBuilder) Expand(prop string, nested func(*QueryBuilder)) *QueryBuilder
```

Add `$expand` with optional nested query builder.

```go
// Simple expand
odata.New().Expand("Orders", nil).Build()
// $expand=Orders

// Nested expand with sub-query
odata.New().Expand("Orders", func(sq *odata.QueryBuilder) {
    sq.Select("id", "price").Top(5)
}).Build()
// $expand=Orders($select=id,price;$top=5)

// Multiple expands
odata.New().
    Expand("Orders", nil).
    Expand("Profile", func(sq *odata.QueryBuilder) {
        sq.Select("name", "avatar")
    }).Build()
// $expand=Orders,Profile($select=name,avatar)
```

#### Ordering

```go
func (qb *QueryBuilder) OrderBy(fields ...string) *QueryBuilder
func (qb *QueryBuilder) OrderByDesc(fields ...string) *QueryBuilder
```

```go
odata.New().OrderBy("name", "age").Build()
// $orderby=name,age

odata.New().OrderByDesc("created").Build()
// $orderby=created desc

odata.New().OrderBy("name").OrderByDesc("age").Build()
// $orderby=name,age desc
```

#### Pagination

```go
func (qb *QueryBuilder) Top(n int) *QueryBuilder
func (qb *QueryBuilder) Skip(n int) *QueryBuilder
```

```go
odata.New().Top(10).Skip(20).Build()
// $top=10&$skip=20
```

Nil pointers are omitted — `Top(0)` produces `$top=0`, but `Top(-1)` is not validated (serializes as-is).

#### Count & Search

```go
func (qb *QueryBuilder) Count() *QueryBuilder
func (qb *QueryBuilder) Search(term string) *QueryBuilder
```

```go
odata.New().Count().Build()
// $count=true

odata.New().Search("john").Build()
// $search=john
```

#### Build

```go
func (qb *QueryBuilder) Build() string
```

Serializes all clauses into a URL query fragment (no leading `?`, no URL encoding — caller handles that).

```go
query := odata.New().Select("name").Filter("status").Eq("active").Build()
// $select=name&$filter=status eq 'active'

// Use with net/url
u := url.URL{Path: "/api/users", RawQuery: query}
// /api/users?$select=name&$filter=status%20eq%20%27active%27
```

## Full Example

```go
query := odata.New().
    Select("name", "email", "age").
    Filter("status").Eq("active").
    Filter("age").Gt(18).
    OrFilter("role").Eq("admin").
    OrderBy("name").
    OrderByDesc("age").
    Top(25).
    Skip(50).
    Count().
    Search("john").
    Build()

// $select=name,email,age&$filter=status eq 'active' and age gt 18 or role eq 'admin'&$search=john&$orderby=name,age desc&$top=25&$skip=50&$count=true
```

## Testing

```bash
go test ./... -v -cover
```

## License

MIT