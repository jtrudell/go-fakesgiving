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
)

func main() {
	templates := populateTemplates()
	db := connectToDatabase()
	defer db.Close()
	controller.Init(templates)
	http.ListenAndServe(":8080", nil)
}

// populateTemplates returns a map of file name keys and template values
// it maps over the templates in the content folder, copying the _layout template for each of them
func populateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "templates"
	layout := template.Must(template.ParseFiles(basePath + "/_layout.gohtml"))

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

func connectToDatabase() *sql.DB {
	db, err := sql.Open("postgres", "user=jtrudell dbname=fakesgiving sslmode=disable")
	if err != nil {
		log.Fatalln("Unable to connect to database:", err)
	}
	model.SetDatabase(db)
	return db
}
