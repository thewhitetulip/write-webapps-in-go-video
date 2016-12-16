package views

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/thewhitetulip/Tasks-new/types"
)

var templates *template.Template
var homeTemplate *template.Template

var comments types.Comments

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

// ShowCompletedTasksFunc populates the /completed URL for the GET request for the currently loggedin user.
func ShowCompletedTasksFunc(w http.ResponseWriter, r *http.Request) {
	var tasks types.Tasks
	comment := types.Comment{ID: "1", Content: "This is a comment!", Author: "@thewhitetulip", Created: "16 Dec 2016"}
	comment1 := types.Comment{ID: "2", Content: "This is a comment!", Author: "@thewhitetulip", Created: "16 Dec 2016"}

	comments = append(comments, comment)
	comments = append(comments, comment1)

	task1 := types.Task{ID: "1", Title: "Title of First Task", Content: "Content of first task", Created: "16 Dec 2016", Comments: comments}

	task2 := types.Task{ID: "2", Title: "Title of second Task", Content: "Content of second task", Created: "16 Dec 2016", Comments: comments}
	task3 := types.Task{ID: "3", Title: "Title of third Task", Content: "Content of third task", Created: "16 Dec 2016", Comments: comments}

	tasks = append(tasks, task1)
	tasks = append(tasks, task2)
	tasks = append(tasks, task3)

	context := types.Context{Tasks: tasks}

	homeTemplate.Execute(w, context)
}

// HomeFunc handles the / URL and asks the name of the user in German.
func HomeFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		var tasks types.Tasks
		comment := types.Comment{ID: "1", Content: "This is a comment!", Author: "@thewhitetulip", Created: "16 Dec 2016"}
		comment1 := types.Comment{ID: "2", Content: "This is a comment!", Author: "@thewhitetulip", Created: "16 Dec 2016"}

		comments = append(comments, comment)
		comments = append(comments, comment1)

		task1 := types.Task{ID: "1", Title: "Title of First Task", Content: "Content of first task", Created: "16 Dec 2016", Comments: comments}

		task2 := types.Task{ID: "2", Title: "Title of second Task", Content: "Content of second task", Created: "16 Dec 2016", Comments: comments}
		task3 := types.Task{ID: "3", Title: "Title of third Task", Content: "Content of third task", Created: "16 Dec 2016", Comments: comments}

		tasks = append(tasks, task1)
		tasks = append(tasks, task2)
		tasks = append(tasks, task3)

		context := types.Context{Tasks: tasks}

		homeTemplate.Execute(w, context)
	}
}

// LoginFunc handles the /login URL and shows the profile page of the logged in user on a GET request
// handles the Login process on the POST request.
func LoginFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You are on the profile page!!")
}
