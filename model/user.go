package model

import (
	"log"
	"time"
)

type User struct {
	ID        int
	Name      string
	Food      string
	CreatedAt *time.Time
}

func NewUser(name, food string) User {
	return User{
		Name: name,
		Food: food,
	}
}

func AllUsers() []User {
	users := []User{}
	err := db.Ping()
	if err != nil {
		log.Fatalln("Database not available %v", err)
	}

	rows, err := db.Query(`SELECT * FROM users;`)
	if err != nil {
		log.Fatalln("Could not get Users %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var food string
		var createdAt *time.Time
		if err := rows.Scan(&id, &name, &food, &createdAt); err != nil {
			log.Fatal(err)
		}
		users = append(users, User{id, name, food, createdAt})
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return users
}

func (u User) Save() error {
	err := db.Ping()
	if err != nil {
		log.Fatalln("Database not available %v", err)
	}

	_, err = db.Exec(`
		INSERT INTO users (name, food)
		VALUES ($1, $2);`, u.Name, u.Food)

	if err != nil {
		log.Printf("Could not save user with name %v: %v", u.Name, err)
	}
	return err
}
