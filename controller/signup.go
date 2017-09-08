package controller

import (
	"github.com/jtrudell/go-fakesgiving/model"
	"github.com/jtrudell/go-fakesgiving/viewmodel"
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
		viewData.Food = food

		user := model.NewUser(name)
		viewData.User = user
		err = user.Save()
		if err != nil {
			log.Println("Something went wrong:", err)
		}
	}
	s.signupTemplate.Execute(w, viewData)
}
