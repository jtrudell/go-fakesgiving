package controller

import (
	"github.com/jtrudell/go-webapp/viewmodel"
	"html/template"
	"log"
	"net/http"
)

type signup struct {
	signupTemplate *template.Template
}

func (s signup) registerRoutes() {
	http.HandleFunc("/signup", s.handleSignup)
}

func (s signup) handleSignup(w http.ResponseWriter, r *http.Request) {
	viewData := viewmodel.NewSignup()

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println("Something went wrong:", err)
		}
		name := r.PostFormValue("name")
		food := r.PostFormValue("food")
		viewData.Name, viewData.Food = name, food
	}
	s.signupTemplate.Execute(w, viewData)
}
