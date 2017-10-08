package viewmodel

import (
	"strconv"
	"testing"
	"time"
)

func TestNewSignup(t *testing.T) {
	actual := NewSignup()
	year := time.Now().Year()
	title := "Fakesgiving " + strconv.Itoa(year)
	expected := Signup{Title: title, Heading: "What are you bringing?"}
	if actual.Title != expected.Title {
		t.Errorf("actual title %v did not equal expected title %v", actual.Title, expected.Title)
	}
	if actual.Heading != expected.Heading {
		t.Errorf("actual heading %v did not equal expected heading %v", actual.Heading, expected.Heading)
	}
}
