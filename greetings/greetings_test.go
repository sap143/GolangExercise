package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "salman"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("salman")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("salman")= %q,%v want match for %#q,nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q,%v want "",error`, msg, err)
	}
}
