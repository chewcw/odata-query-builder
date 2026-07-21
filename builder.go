package odata

import (
	"fmt"
	"strings"
)

type QueryBuilder struct {
	selects []string
	filters []filterCond
	search  string
	top     *int
	skip    *int
	count   *bool
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

func (qb *QueryBuilder) Top(n int) *QueryBuilder {
	qb.top = &n
	return qb
}

func (qb *QueryBuilder) Skip(n int) *QueryBuilder {
	qb.skip = &n
	return qb
}

func (qb *QueryBuilder) Count() *QueryBuilder {
	t := true
	qb.count = &t
	return qb
}

func (qb *QueryBuilder) Search(term string) *QueryBuilder {
	qb.search = term
	return qb
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
	if qb.search != "" {
		parts = append(parts, "$search="+qb.search)
	}
	if qb.top != nil {
		parts = append(parts, fmt.Sprintf("$top=%d", *qb.top))
	}
	if qb.skip != nil {
		parts = append(parts, fmt.Sprintf("$skip=%d", *qb.skip))
	}
	if qb.count != nil && *qb.count {
		parts = append(parts, "$count=true")
	}
	return strings.Join(parts, "&")
}
