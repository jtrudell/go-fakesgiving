package controller

import (
	"html/template"
	"net/http"
)

var (
	indexController  index
	signupController signup
)

func Init(templates map[string]*template.Template) {
	indexController.indexTemplate = templates["index.gohtml"]
	indexController.registerRoutes()
	signupController.signupTemplate = templates["signup.gohtml"]
	signupController.registerRoutes()
	http.Handle("/css/", http.FileServer(http.Dir("public")))
}
