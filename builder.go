package odata

import "strings"

type QueryBuilder struct {
	selects []string
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

func (qb *QueryBuilder) Build() string {
	var parts []string
	if len(qb.selects) > 0 {
		parts = append(parts, "$select="+strings.Join(qb.selects, ","))
	}
	return strings.Join(parts, "&")
}
