package viewmodel

import (
	"fmt"
	"time"
)

type Base struct {
	Title string
}

func NewBase() Base {

	return Base{
		Title: fmt.Sprintf("Thanksgiving %v", time.Now().Year()),
	}
}
