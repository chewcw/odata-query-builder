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
