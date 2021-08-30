package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestFormValid(t *testing.T) {
	r := httptest.NewRequest("POST", "/Idontknowwheretogo", nil)
	form := New(r.PostForm)
	isValid := form.Valid()
	if !isValid {
		t.Error("Got invalid when it should have been valid")
	}
}
func TestFormRequired(t *testing.T) {
	r := httptest.NewRequest("POST", "/Idontknowwheretogo", nil)
	form := New(r.PostForm)
	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("Got valid even though form is missign fields")
	}
	postData := url.Values{}
	postData.Add("a", "a")
	postData.Add("b", "b")
	postData.Add("c", "c")
	r, _ = http.NewRequest("POST", "/Idontknowwheretogo", nil)
	r.PostForm = postData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("Shows that doesn't have required fields even when it does")
	}
}
func TestFormHas(t *testing.T) {
	postData := url.Values{}

	form := New(postData)
	res := form.Has("wrongfield")
	if res != false {
		t.Error("Returned true for non-existent field")
	}
	postData = url.Values{}
	postData.Add("a", "a")

	form = New(postData)
	res = form.Has("a")

	if res != true {
		t.Error("Returned false for existent field")
	}
}

func TestFormMinLength(t *testing.T) {
	postData := url.Values{}
	form := New(postData)
	form.MinLength("a", 3)
	if form.Valid() {
		t.Error("Form shows min length for non-existent field")
	}
	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("Should give error but didnt get one")
	}
	postData = url.Values{}
	postData.Add("field1", "small val")

	form = New(postData)
	form.MinLength("field1", 15)
	if form.Valid() {
		t.Error("Shows true for fields smaller than required length")
	}
	postData = url.Values{}
	postData.Add("field2", "bigger val")

	form = New(postData)
	form.MinLength("field2", 5)
	if !form.Valid() {
		t.Error("Shows false for fields which have required length")
	}
	isError = form.Errors.Get("field2")
	if isError != "" {
		t.Error("Should not give error but  get one")
	}

}

func TestFormEmailCheck(t *testing.T) {
	postData := url.Values{}

	form := New(postData)
	form.EmailCheck("phone")
	if form.Valid() {
		t.Error("Form shows valid email for non-existent field")
	}
	postData = url.Values{}
	postData.Add("email", "vedantb@gmail.com")

	form = New(postData)
	form.EmailCheck("email")
	if !form.Valid() {
		t.Error("Shows invalid for valid email")
	}
	postData = url.Values{}
	postData.Add("email", "vedantb@gmail")

	form = New(postData)
	form.EmailCheck("email")
	if form.Valid() {
		t.Error("Shows valid for invalid email")
	}
}
