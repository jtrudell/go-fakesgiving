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
	indexController.indexTemplate = templates["index.html"]
	indexController.registerRoutes()
	signupController.signupTemplate = templates["signup.html"]
	signupController.registerRoutes()
	http.HandleFunc("/favicon.ico", handleFavicon)
	http.Handle("/css/", http.FileServer(http.Dir("public")))
}

func handleFavicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/public/favicon.ico")
}
