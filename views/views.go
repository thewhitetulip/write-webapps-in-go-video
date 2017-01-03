package views

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	md "github.com/shurcooL/github_flavored_markdown"

	"github.com/thewhitetulip/Tasks-new/types"
)

var templates *template.Template
var homeTemplate *template.Template

var comments types.Comments
var category types.Category
var categories []types.Category
var tasks types.Tasks

func init() {
	category = types.Category{Name: "Random", TaskCount: "1"}
	categories = append(categories, category)
	category = types.Category{Name: "Very Random", TaskCount: "1"}
	categories = append(categories, category)

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

// AddTaskFunc adds a new task to the array and populates the /add/ URL. Later it is going to insert data in database.
func AddTaskFunc(w http.ResponseWriter, r *http.Request) {
	var task types.Task

	if r.Method == "POST" {
		r.ParseForm()

		csrf := r.FormValue("CSRFToken")
		if csrf == "supersecret" {
			task.Title = r.FormValue("title")
			task.Content = r.FormValue("content")
			task.Category = r.FormValue("category")
			task.Priority = r.FormValue("priority")
			task.Hidden = r.FormValue("hide")

			if task.Hidden == "" {
				task.Hidden = "off"
			}

			if task.Hidden == "" || task.Hidden != "on" {
				tasks = append(tasks, task)
			}

			task.Created = "17 Dec 2016"

			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			fmt.Println("CSRF Mismatch")
			http.Redirect(w, r, "/", http.StatusBadRequest)
		}
	}
}

// ShowCompletedTasksFunc populates the /completed URL for the GET request for the currently loggedin user.
func ShowCompletedTasksFunc(w http.ResponseWriter, r *http.Request) {
	context := types.Context{Tasks: nil}

	homeTemplate.Execute(w, context)
}

// HomeFunc handles the / URL and asks the name of the user in German.
func HomeFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		for i, _ := range tasks {
			tasks[i].HTMLContent = template.HTML(string(md.Markdown([]byte(tasks[i].Content))))
		}

		context := types.Context{Tasks: tasks, CSRFToken: "supersecret", Categories: categories}

		homeTemplate.Execute(w, context)
	}
}

// LoginFunc handles the /login URL and shows the profile page of the logged in user on a GET request
// handles the Login process on the POST request.
func LoginFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You are on the profile page!!")
}
