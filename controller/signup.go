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
	viewData.Users = model.AllUsers()

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println("Something went wrong:", err)
		}
		name := r.PostFormValue("name")
		food := r.PostFormValue("food")

		user := model.NewUser(name, food)
		err = user.Save()
		if err != nil {
			log.Println("Something went wrong:", err)
		}
		viewData.User = user
	}
	s.signupTemplate.Execute(w, viewData)
}
