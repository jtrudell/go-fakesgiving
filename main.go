package main

import (
	"database/sql"
	"github.com/jtrudell/go-fakesgiving/controller"
	"github.com/jtrudell/go-fakesgiving/model"
	_ "github.com/lib/pq"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// type Config struct {
// 	Port   string
// 	DBName string
// 	DBUser string
// }

const (
	envFile = ".env"
)

func main() {
	setEnvironment()
	templates := populateTemplates()
	db := connectToDatabase()
	defer db.Close()
	controller.Init(templates)
	port := os.Getenv("PORT")
	http.ListenAndServe(port, nil)
}

// populateTemplates returns a map of file name keys and template values
// it maps over the templates in the content folder, copying the _layout template for each of them
func populateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "templates"
	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))

	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template/content directory: " + err.Error())
	}

	filesInfo, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory:" + err.Error())
	}

	for _, fi := range filesInfo {
		file, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to open template " + fi.Name())
		}

		content, err := ioutil.ReadAll(file)
		if err != nil {
			panic("Failed to read content from file" + fi.Name())
		}
		file.Close()

		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of " + fi.Name() + " as template")
		}

		result[fi.Name()] = tmpl
	}

	return result
}

// TODO: move to its own package; use config struct
func setEnvironment() {
	file, err := os.Open(envFile)
	if err != nil {
		panic("Failed to open " + envFile)
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic("Failed to read content from " + envFile)
	}
	envVars := strings.Split(string(content), "\n")
	// config := Config{}
	for _, envVar := range envVars {
		x := strings.Split(envVar, "=")
		if len(x) > 1 {
			err := os.Setenv(x[0], x[1])
			if err != nil {
				panic("Could not set environment variable" + x[0])
			}
		}
	}
	file.Close()
}

// connectToDatabase opens a postgres database named fakesgiving
// and creates the users table if it doesn't exist. If then sets
// the fakesgiving database as the db for the model to use
func connectToDatabase() *sql.DB {
	db := openDatabase()
	createUsersTable(db)
	model.SetDatabase(db)
	return db
}

func openDatabase() *sql.DB {
	user := os.Getenv("USER")
	dbname := os.Getenv("DBNAME")
	db, err := sql.Open("postgres", "user="+user+" dbname="+dbname+" sslmode=disable")
	if err != nil {
		log.Fatalln("Unable to connect to database:", err)
	}
	return db
}

func createUsersTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS USERS(
  ID SERIAL UNIQUE PRIMARY KEY NOT NULL,
  NAME TEXT NOT NULL,
  FOOD TEXT NOT NULL,
  CREATED_AT TIMESTAMP DEFAULT NOW() NOT NULL);`)

	if err != nil {
		log.Fatalln("Could not create Users table: %v", err)
	}
}
