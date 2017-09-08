package model

import (
	"log"
	"time"
)

type User struct {
	Name      string
	ID        int
	CreatedAt *time.Time
}

func NewUser(name string) User {
	return User{
		Name: name,
	}
}

func (u User) Save() error {
	_, err := db.Exec(`
	INSERT INTO users (name)
	VALUES ($1);`, u.Name)

	if err != nil {
		log.Printf("Could not save user with name %v: %v", u.Name, err)
	}
	return err
}
