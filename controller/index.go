package controller

import (
	"html/template"
	"net/http"

	"github.com/jtrudell/go-fakesgiving/viewmodel"
)

type index struct {
	indexTemplate *template.Template
}

func (i index) registerRoutes() {
	http.HandleFunc("/", i.handleIndex)
}

func (i index) handleIndex(w http.ResponseWriter, r *http.Request) {
	viewData := viewmodel.NewBase()
	i.indexTemplate.Execute(w, viewData)
}
