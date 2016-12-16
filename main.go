package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var templates *template.Template
var homeTemplate *template.Template

func main() {
	PopulateTemplate()
	http.Handle("/static/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/login/", LoginFunc)
	http.HandleFunc("/", HomeFunc)
	fmt.Println("Server running on 8081")
	http.ListenAndServe("0.0.0.0:8081", nil)
}

// PopulateTemplate reads the ./templates folder and parses all the html files inside it
// and it stores it in the templates variable which will be lookedup by other variables.
func PopulateTemplate() {
	templates, err := template.ParseGlob("./templates/*.html")

	if err != nil {
		fmt.Println("main.go: PopulateTemplate: ", err)
		os.Exit(1)
	}

	homeTemplate = templates.Lookup("tasks.html")
}

// HomeFunc handles the / URL and asks the name of the user in German.
func HomeFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		type Task struct {
			Title   string
			Content string
			Created string
		}
		task1 := Task{Title: "Title of First Task", Content: "Content of first task", Created: "16 Dec 2016"}

		homeTemplate.Execute(w, task1)
	}
}

// LoginFunc handles the /login URL and shows the profile page of the logged in user on a GET request
// handles the Login process on the POST request.
func LoginFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You are on the profile page!!")
}
