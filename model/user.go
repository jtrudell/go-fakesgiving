package model

import (
	"log"
	"time"
)

type User struct {
	ID        int
	Name      string
	CreatedAt *time.Time
}

func NewUser(name string) User {
	return User{
		Name: name,
	}
}

func AllUsers() []User {
	users := []User{}
	rows, err := db.Query(`SELECT * FROM users;`)
	if err != nil {
		log.Fatalln("Could not get Users %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var id int
		var createdAt *time.Time
		if err := rows.Scan(&id, &name, &createdAt); err != nil {
			log.Fatal(err)
		}
		users = append(users, User{id, name, createdAt})
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return users
}

func (u User) Save() error {
	_, err := db.Exec(`
	INSERT INTO users (name)
	VALUES ($1);`, u.Name)

	if err != nil {
		log.Printf("Could not save user with name %v: %v", u.Name, err)
	}
	AllUsers()
	return err
}
