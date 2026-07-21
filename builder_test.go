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
