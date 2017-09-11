package controller

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootGetsCorrectTemplate(t *testing.T) {
	h := new(index)
	expected := "index template"
	h.indexTemplate, _ = template.New("").Parse(expected)

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	h.handleIndex(w, r)

	actual, _ := ioutil.ReadAll(w.Result().Body)
	if string(actual) != expected {
		t.Errorf("Got %v, was expecting %v", actual, expected)
	}
}
