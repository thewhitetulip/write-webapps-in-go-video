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
		type Comment struct {
			ID      string
			Content string
			Author  string
			Created string
		}

		type Task struct {
			ID       string
			Title    string
			Content  string
			Created  string
			Comments []Comment
		}

		type Context struct {
			Tasks []Task
		}

		comment := Comment{ID: "1", Content: "This is a comment!", Author: "@thewhitetulip", Created: "16 Dec 2016"}
		comment1 := Comment{ID: "2", Content: "This is a comment!", Author: "@thewhitetulip", Created: "16 Dec 2016"}

		var comments []Comment
		comments = append(comments, comment)
		comments = append(comments, comment1)

		task1 := Task{ID: "1", Title: "Title of First Task", Content: "Content of first task", Created: "16 Dec 2016", Comments: comments}
		task2 := Task{ID: "2", Title: "Title of second Task", Content: "Content of second task", Created: "16 Dec 2016", Comments: comments}
		task3 := Task{ID: "3", Title: "Title of third Task", Content: "Content of third task", Created: "16 Dec 2016", Comments: comments}

		var tasks []Task

		tasks = append(tasks, task1)
		tasks = append(tasks, task2)
		tasks = append(tasks, task3)

		context := Context{Tasks: tasks}

		homeTemplate.Execute(w, context)
	}
}

// LoginFunc handles the /login URL and shows the profile page of the logged in user on a GET request
// handles the Login process on the POST request.
func LoginFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You are on the profile page!!")
}
