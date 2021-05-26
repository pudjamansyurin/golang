package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls Hello with a name, checking for a valid return value
func TestHelloName(t *testing.T) {
	name := "Pudja"
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello(%q) = %q, %v, want match for %#q, nil`, name, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	name := ""
	msg, err := Hello(name)
	if msg != "" || err == nil {
		t.Fatalf(`Hello(%q) = %q, %v, want %q, error`, name, msg, err, name)
	}
}
