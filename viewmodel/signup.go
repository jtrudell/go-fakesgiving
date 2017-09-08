package viewmodel

import "github.com/jtrudell/go-fakesgiving/model"

type Signup struct {
	Title   string
	Heading string
	Food    string
	User    model.User
}

func NewSignup() Signup {

	return Signup{
		Title:   NewBase().Title,
		Heading: "Sign up for some stuff",
	}
}
