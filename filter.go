package odata

import (
	"fmt"
	"strings"
)

type filterCond struct {
	field    string
	operator string
	value    string // pre-formatted (quoted if string)
	logic    string // "and" or "or"
	negated  bool
	funcOp   bool // true for contains/startswith/endswith: serialize as op(field, value)
}

type FilterBuilder struct {
	qb      *QueryBuilder
	field   string
	logic   string
	negated bool
}

func (fb *FilterBuilder) addCond(operator string, value any) *QueryBuilder {
	formatted := formatValue(value)
	if formatted == "" {
		// No op, early return
		return fb.qb
	}
	fb.qb.filters = append(fb.qb.filters, filterCond{
		field:    fb.field,
		operator: operator,
		value:    formatted,
		logic:    fb.logic,
		negated:  fb.negated,
	})
	return fb.qb
}

func (fb *FilterBuilder) Not() *FilterBuilder {
	fb.negated = !fb.negated
	return fb
}

func (fb *FilterBuilder) Eq(value any) *QueryBuilder { return fb.addCond("eq", value) }
func (fb *FilterBuilder) Ne(value any) *QueryBuilder { return fb.addCond("ne", value) }
func (fb *FilterBuilder) Gt(value any) *QueryBuilder { return fb.addCond("gt", value) }
func (fb *FilterBuilder) Ge(value any) *QueryBuilder { return fb.addCond("ge", value) }
func (fb *FilterBuilder) Lt(value any) *QueryBuilder { return fb.addCond("lt", value) }
func (fb *FilterBuilder) Le(value any) *QueryBuilder { return fb.addCond("le", value) }

func formatValue(value any) string {
	switch v := value.(type) {
	case nil:
		return ""
	case string:
		escaped := strings.ReplaceAll(v, "'", "''")
		return fmt.Sprintf("'%s'", escaped)
	case bool:
		return fmt.Sprintf("%t", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

func formatFilterCond(c filterCond) string {
	var expr string
	switch {
	case c.funcOp:
		expr = fmt.Sprintf("%s(%s, %s)", c.operator, c.field, c.value)
	case c.operator == "in":
		expr = fmt.Sprintf("%s in %s", c.field, c.value)
	case c.operator == "has":
		expr = fmt.Sprintf("%s has %s", c.field, c.value)
	default:
		expr = fmt.Sprintf("%s %s %s", c.field, c.operator, c.value)
	}
	if c.negated {
		expr = fmt.Sprintf("not (%s)", expr)
	}
	return expr
}

func (fb *FilterBuilder) Contains(value string) *QueryBuilder {
	fb.qb.filters = append(fb.qb.filters, filterCond{
		field:    fb.field,
		operator: "contains",
		value:    formatValue(value),
		funcOp:   true,
		logic:    fb.logic,
		negated:  fb.negated,
	})
	return fb.qb
}

func (fb *FilterBuilder) StartsWith(value string) *QueryBuilder {
	fb.qb.filters = append(fb.qb.filters, filterCond{
		field:    fb.field,
		operator: "startswith",
		value:    formatValue(value),
		funcOp:   true,
		logic:    fb.logic,
		negated:  fb.negated,
	})
	return fb.qb
}

func (fb *FilterBuilder) EndsWith(value string) *QueryBuilder {
	fb.qb.filters = append(fb.qb.filters, filterCond{
		field:    fb.field,
		operator: "endswith",
		value:    formatValue(value),
		funcOp:   true,
		logic:    fb.logic,
		negated:  fb.negated,
	})
	return fb.qb
}

func (fb *FilterBuilder) In(values ...any) *QueryBuilder {
	var formatted []string
	for _, v := range values {
		f := formatValue(v)
		if f != "" {
			formatted = append(formatted, f)
		}
	}
	if len(formatted) == 0 {
		return fb.qb
	}
	fb.qb.filters = append(fb.qb.filters, filterCond{
		field:    fb.field,
		operator: "in",
		value:    "(" + strings.Join(formatted, ", ") + ")",
		logic:    fb.logic,
		negated:  fb.negated,
	})
	return fb.qb
}

func (fb *FilterBuilder) Has(value string) *QueryBuilder {
	fb.qb.filters = append(fb.qb.filters, filterCond{
		field:    fb.field,
		operator: "has",
		value:    value, // raw, no quoting
		logic:    fb.logic,
		negated:  fb.negated,
	})
	return fb.qb
}
