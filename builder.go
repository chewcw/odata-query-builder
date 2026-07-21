package odata

import "strings"

type QueryBuilder struct {
	selects []string
	filters []filterCond
}

func New() *QueryBuilder {
	return &QueryBuilder{}
}

func (qb *QueryBuilder) Select(fields ...string) *QueryBuilder {
	qb.selects = append(qb.selects, fields...)
	return qb
}

func (qb *QueryBuilder) And(fields ...string) *QueryBuilder {
	return qb.Select(fields...)
}

func (qb *QueryBuilder) Filter(field string) *FilterBuilder {
	return &FilterBuilder{
		qb:    qb,
		field: field,
		logic: "and",
	}
}

func (qb *QueryBuilder) OrFilter(field string) *FilterBuilder {
	return &FilterBuilder{
		qb:    qb,
		field: field,
		logic: "or",
	}
}

func (qb *QueryBuilder) Build() string {
	var parts []string
	if len(qb.selects) > 0 {
		parts = append(parts, "$select="+strings.Join(qb.selects, ","))
	}
	if len(qb.filters) > 0 {
		var exprs []string
		for i, c := range qb.filters {
			expr := formatFilterCond(c)
			if i > 0 {
				expr = c.logic + " " + expr
			}
			exprs = append(exprs, expr)
		}
		parts = append(parts, "$filter="+strings.Join(exprs, " "))
	}
	return strings.Join(parts, "&")
}
