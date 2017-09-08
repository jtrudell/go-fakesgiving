package viewmodel

type Signup struct {
	Title   string
	Heading string
	Name    string
	Food    string
}

func NewSignup() Signup {

	return Signup{
		Title:   NewBase().Title,
		Heading: "Sign up for some stuff",
	}
}
