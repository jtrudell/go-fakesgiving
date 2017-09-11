package viewmodel

import (
	"strconv"
	"testing"
	"time"
)

func TestNewBase(t *testing.T) {
	actual := NewBase()
	year := time.Now().Year()
	expected := Base{"Fakesgiving " + strconv.Itoa(year)}
	if actual != expected {
		t.Errorf("got %v, expected %v", actual, expected)
	}
}
