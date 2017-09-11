package model

import "testing"

func TestNewUser(t *testing.T) {
	actual := NewUser("Bob", "Turkey")
	expected := User{Name: "Bob", Food: "Turkey"}
	if actual != expected {
		t.Errorf("user %v did not equal expected user %v", actual, expected)
	}
}
