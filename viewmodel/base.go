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
		Title: fmt.Sprintf("Fakesgiving %v", time.Now().Year()),
	}
}
