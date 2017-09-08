package controller

import (
	"github.com/jtrudell/go-fakesgiving/viewmodel"
	"html/template"
	"net/http"
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
