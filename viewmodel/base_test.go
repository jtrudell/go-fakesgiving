package viewmodel

import (
	"strconv"
	"testing"
	"time"
)

func TestNewBase(t *testing.T) {
	actual := NewBase()
	year := time.Now().Year()
	imageurl := "/images/thanksgiving_giphy.gif"

	expected := Base{"Fakesgiving " + strconv.Itoa(year), imageurl}
	if actual != expected {
		t.Errorf("got %v, expected %v", actual, expected)
	}
}
