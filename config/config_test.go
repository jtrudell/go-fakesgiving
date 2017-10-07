package config

import "testing"

func TestSetEnv(t *testing.T) {
	envLines := []string{"PORT=XXXX", "DBNAME=name", "DBUSER=user"}
	actual := setEnv(envLines)
	expected := Env{"XXXX", "name", "user"}
	if actual != expected {
		t.Errorf("env %v did not equal expected env %v", actual, expected)
	}
}
