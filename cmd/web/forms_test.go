package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func Test_application_form(t *testing.T) {
	tests := []struct {
		name   string
		errors errors
		field  string
		want   string
	}{
		{
			name: "Field has one error",
			errors: errors{
				"email": {"Email is required"},
			},
			field: "email",
			want:  "Email is required",
		},
		{
			name: "Field has multiple errors",
			errors: errors{
				"email": {"Email is required", "Email must be valid"},
			},
			field: "email",
			want:  "Email is required",
		},
		{
			name: "Field has no errors",
			errors: errors{
				"email": {},
			},
			field: "email",
			want:  "",
		},
		{
			name: "Field does not exist",
			errors: errors{
				"email": {"Email is required"},
			},
			field: "username",
			want:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.errors.Get(tt.field); got != tt.want {
				t.Errorf("errors.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestForm_Has(t *testing.T) {
	form := NewForm(nil)
	has := form.Has("whatever")

	if has {
		t.Error("form shows has field whe it should not")
	}
	postedData := url.Values{}
	postedData.Add("a", "a")
	form = NewForm(postedData)
	has = form.Has("a")
	if !has {
		t.Error("shows form does not have a field when it should")
	}
}
func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := NewForm(r.PostForm)

	form.Required("a", "b", "c")

	if form.Valid() {
		t.Error("form shows valid when required fields are missing")
	}
	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")
	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData

	form = NewForm(r.PostForm)

	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows post does not have required fields, when it does")
	}
}
func TestForm_Check(t *testing.T) {
	form := NewForm(nil)
	form.Check(false, "possword", "password is required")
	if form.Valid() {
		t.Error("valid() returns false, and it should be true when calling check")
	}
}
func TestForm_ErrorGet(t *testing.T) {
	form := NewForm(nil)
	form.Check(false, "password", "password is required")

	s := form.Errors.Get("password")

	if len(s) == 0 {
		t.Error("should have an error returned from get, but do not")
	}
	s = form.Errors.Get("whatever")

	if len(s) != 0 {
		t.Error("should not get error but it does")
	}
}
