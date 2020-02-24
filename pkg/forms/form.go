package forms

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

var UrlRX = regexp.MustCompile(`^(https?)://[\\.a-zA-Z0-9-]+(:[0-9]+)*$`)

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

func (f *Form) MatchesPattern(field string, pattern *regexp.Regexp) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if !pattern.MatchString(value) {
		f.Errors.Add(field, "This field is invalid")
	}
}


func (f *Form) ValidUrl(field string) {
	value := f.Get(field)
	_, err := url.Parse(value)
	if err != nil {
		f.Errors.Add(field, fmt.Sprintf("This field is not a valid URL %s", err))
	}
}



func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
