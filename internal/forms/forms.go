package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

type Form struct {
	url.Values
	Errors errors
}

//New initializes a form structure
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

//Required validated whether the required form fields are empty or not
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "Field cannot be blank")
		}
	}
}

//Has checks if given field is in form and if it is empty or not
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		return false
	}
	return true

}

//Valid returns true if no errors and false otherwise
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

//MinLength checks if field is minimum characters long
func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	x := f.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be atleast %d characters long", length))
		return false
	}
	return true
}

//Email Validator
func (f *Form) EmailCheck(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid Email")
	}
}
