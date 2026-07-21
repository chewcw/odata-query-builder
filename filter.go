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
}

type FilterBuilder struct {
	qb      *QueryBuilder
	field   string
	logic   string
	negated bool
}

func (fb *FilterBuilder) addCond(operator string, value any) *QueryBuilder {
	fb.qb.filters = append(fb.qb.filters, filterCond{
		field:    fb.field,
		operator: operator,
		value:    formatValue(value),
		logic:    fb.logic,
		negated:  fb.negated,
	})
	return fb.qb
}

func (fb *FilterBuilder) Not() *FilterBuilder {
	fb.negated = !fb.negated
	return fb
}

func (fb *FilterBuilder) Eq(value any) *QueryBuilder  { return fb.addCond("eq", value) }
func (fb *FilterBuilder) Ne(value any) *QueryBuilder  { return fb.addCond("ne", value) }
func (fb *FilterBuilder) Gt(value any) *QueryBuilder  { return fb.addCond("gt", value) }
func (fb *FilterBuilder) Ge(value any) *QueryBuilder  { return fb.addCond("ge", value) }
func (fb *FilterBuilder) Lt(value any) *QueryBuilder  { return fb.addCond("lt", value) }
func (fb *FilterBuilder) Le(value any) *QueryBuilder  { return fb.addCond("le", value) }

func formatValue(value any) string {
	switch v := value.(type) {
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
	expr := fmt.Sprintf("%s %s %s", c.field, c.operator, c.value)
	if c.negated {
		expr = fmt.Sprintf("not (%s)", expr)
	}
	return expr
}
