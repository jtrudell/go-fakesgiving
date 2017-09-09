package viewmodel

import "github.com/jtrudell/go-fakesgiving/model"

type Signup struct {
	Title   string
	Heading string
	User    model.User
	Users   []model.User
}

func NewSignup() Signup {

	return Signup{
		Title:   NewBase().Title,
		Heading: "Sign up for some stuff",
	}
}
