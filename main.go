package main

import (
	"database/sql"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jtrudell/go-fakesgiving/config"
	"github.com/jtrudell/go-fakesgiving/controller"
	"github.com/jtrudell/go-fakesgiving/model"
	_ "github.com/lib/pq"
)

var port string
var dbname string
var dbuser string

func init() {
	env := config.Setup()
	port, dbname, dbuser = env.Port, env.DBName, env.DBUser
}

func main() {
	templates := populateTemplates()
	db := connectToDatabase()
	defer db.Close()
	controller.Init(templates)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Server failed to start:", err)
	}
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
	var db *sql.DB
	var err error

	dburl := os.Getenv("DATABASE_URL")
	if dburl == "" {
		db, err = sql.Open("postgres", "user="+dbuser+" dbname="+dbname+" sslmode=disable")
	} else {
		db, err = sql.Open("postgres", dburl)
	}

	if err != nil {
		log.Fatalln("Unable to connect to database:", err)
	}
	return db
}

func createUsersTable(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatalln("Database not available %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS USERS(
		ID SERIAL UNIQUE PRIMARY KEY NOT NULL,
		NAME TEXT NOT NULL,
		FOOD TEXT NOT NULL,
		CREATED_AT TIMESTAMP DEFAULT NOW() NOT NULL);`)

	if err != nil {
		log.Fatalln("Could not create Users table: %v", err)
	}
}
