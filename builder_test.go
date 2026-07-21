package odata

import "testing"

func TestSelectBasic(t *testing.T) {
	got := New().Select("name", "email").Build()
	want := "$select=name,email"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestSelectWithAnd(t *testing.T) {
	got := New().Select("name").And("email").Build()
	want := "$select=name,email"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestEmptyBuild(t *testing.T) {
	got := New().Build()
	if got != "" {
		t.Errorf("expected empty string, got %q", got)
	}
}

func TestFilterEq(t *testing.T) {
	got := New().Filter("status").Eq("active").Build()
	want := "$filter=status eq 'active'"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFilterEqNumber(t *testing.T) {
	got := New().Filter("age").Eq(42).Build()
	want := "$filter=age eq 42"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFilterEqBool(t *testing.T) {
	got := New().Filter("enabled").Eq(true).Build()
	want := "$filter=enabled eq true"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFilterNe(t *testing.T) {
	got := New().Filter("status").Ne("deleted").Build()
	want := "$filter=status ne 'deleted'"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFilterGt(t *testing.T) {
	got := New().Filter("price").Gt(100).Build()
	want := "$filter=price gt 100"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFilterGe(t *testing.T) {
	got := New().Filter("price").Ge(0).Build()
	want := "$filter=price ge 0"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFilterLt(t *testing.T) {
	got := New().Filter("price").Lt(50.5).Build()
	want := "$filter=price lt 50.5"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFilterLe(t *testing.T) {
	got := New().Filter("price").Le(99.99).Build()
	want := "$filter=price le 99.99"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFilterNot(t *testing.T) {
	got := New().Filter("status").Not().Eq("deleted").Build()
	want := "$filter=not (status eq 'deleted')"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFilterMultipleImplicitAnd(t *testing.T) {
	got := New().
		Filter("status").Eq("active").
		Filter("age").Gt(18).
		Build()
	want := "$filter=status eq 'active' and age gt 18"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestSelectAndFilterCombined(t *testing.T) {
	got := New().
		Select("name", "email").
		Filter("status").Eq("active").
		Build()
	want := "$select=name,email&$filter=status eq 'active'"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFilterContains(t *testing.T) {
	got := New().Filter("name").Contains("john").Build()
	want := "$filter=contains(name, 'john')"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFilterStartsWith(t *testing.T) {
	got := New().Filter("name").StartsWith("A").Build()
	want := "$filter=startswith(name, 'A')"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFilterEndsWith(t *testing.T) {
	got := New().Filter("email").EndsWith(".com").Build()
	want := "$filter=endswith(email, '.com')"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFilterIn(t *testing.T) {
	got := New().Filter("status").In("active", "pending").Build()
	want := "$filter=status in ('active', 'pending')"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFilterInNumbers(t *testing.T) {
	got := New().Filter("id").In(1, 2, 3).Build()
	want := "$filter=id in (1, 2, 3)"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFilterHas(t *testing.T) {
	got := New().Filter("style").Has("Sales.Color'Yellow'").Build()
	want := "$filter=style has Sales.Color'Yellow'"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestTop(t *testing.T) {
	got := New().Top(10).Build()
	want := "$top=10"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestSkip(t *testing.T) {
	got := New().Skip(20).Build()
	want := "$skip=20"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestCount(t *testing.T) {
	got := New().Count().Build()
	want := "$count=true"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestSearch(t *testing.T) {
	got := New().Search("john").Build()
	want := "$search=john"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestPagination(t *testing.T) {
	got := New().Top(10).Skip(20).Build()
	want := "$top=10&$skip=20"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFullQuery(t *testing.T) {
	got := New().
		Select("name", "email").
		Filter("status").Eq("active").
		Top(10).
		Build()
	want := "$select=name,email&$filter=status eq 'active'&$top=10"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestOrderBy(t *testing.T) {
	got := New().OrderBy("name").Build()
	want := "$orderby=name"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestOrderByMultiple(t *testing.T) {
	got := New().OrderBy("name", "age").Build()
	want := "$orderby=name,age"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestOrderByDesc(t *testing.T) {
	got := New().OrderByDesc("created").Build()
	want := "$orderby=created desc"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestOrderByMixed(t *testing.T) {
	got := New().OrderBy("name").OrderByDesc("created").Build()
	want := "$orderby=name,created desc"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestExpandSimple(t *testing.T) {
	got := New().Expand("Orders", nil).Build()
	want := "$expand=Orders"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestExpandNested(t *testing.T) {
	got := New().Expand("Orders", func(sq *QueryBuilder) {
		sq.Select("id", "price").Top(5)
	}).Build()
	want := "$expand=Orders($select=id,price;$top=5)"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestExpandMultiple(t *testing.T) {
	got := New().
		Expand("Orders", nil).
		Expand("Profile", func(sq *QueryBuilder) {
			sq.Select("name", "avatar")
		}).
		Build()
	want := "$expand=Orders,Profile($select=name,avatar)"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestExpandDeep(t *testing.T) {
	got := New().Expand("Orders", func(sq *QueryBuilder) {
		sq.Expand("Profile", nil)
	}).Build()
	want := "$expand=Orders($expand=Profile)"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestExpandDeepWithSelect(t *testing.T) {
	got := New().Expand("Orders", func(sq *QueryBuilder) {
		sq.Select("id").Expand("Profile", nil)
	}).Build()
	want := "$expand=Orders($select=id;$expand=Profile)"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestOrFilter(t *testing.T) {
	got := New().
		Filter("status").Eq("active").
		OrFilter("priority").Gt(5).
		Build()
	want := "$filter=status eq 'active' or priority gt 5"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFilterDoubleNot(t *testing.T) {
	got := New().Filter("status").Not().Not().Eq("deleted").Build()
	want := "$filter=status eq 'deleted'"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFilterEscapedQuote(t *testing.T) {
	got := New().Filter("name").Eq("O'Brien").Build()
	want := "$filter=name eq 'O''Brien'"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFilterFloat(t *testing.T) {
	got := New().Filter("score").Eq(3.14).Build()
	want := "$filter=score eq 3.14"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestComplexQuery(t *testing.T) {
	got := New().
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
	want := "$select=name,email,age&$filter=status eq 'active' and age gt 18 or role eq 'admin'&$search=john&$orderby=name,age desc&$top=25&$skip=50&$count=true"
	if got != want {
		t.Errorf("got  %q\nwant %q", got, want)
	}
}

func TestEmptySelect(t *testing.T) {
	got := New().Select().Build()
	if got != "" {
		t.Errorf("expected empty, got %q", got)
	}
}

func TestFilterEmptyValue(t *testing.T) {
	got := New().Filter("name").Eq("").Build()
	want := "$filter=name eq ''"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFilterEmptyField(t *testing.T) {
	got := New().Filter("").Eq("value").Build()
	want := "$filter= eq 'value'"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestNegativeOnePanics(t *testing.T) {
	got := New().Top(-1).Build()
	want := "$top=-1"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
