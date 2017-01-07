package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/thewhitetulip/Tasks-new/db"
	"github.com/thewhitetulip/Tasks-new/sessions"
	"github.com/thewhitetulip/Tasks-new/types"
)

var templates *template.Template
var homeTemplate *template.Template
var loginTemplate *template.Template

var tasks types.Tasks

// PopulateTemplate reads the ./templates folder and parses all the html files inside it
// and it stores it in the templates variable which will be lookedup by other variables.
func PopulateTemplate() {
	templates, err := template.ParseGlob("./templates/*.html")

	if err != nil {
		fmt.Println("main.go: PopulateTemplate: ", err)
		os.Exit(1)
	}

	homeTemplate = templates.Lookup("tasks.html")
	loginTemplate = templates.Lookup("login.html")
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
		username := sessions.GetCurrentUserName(r)
		context, _ := db.GetPendingTasks(username)
		context.CSRFToken = "abcd"
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "csrftoken", Value: "abcd", Expires: expiration}
		http.SetCookie(w, &cookie)

		homeTemplate.Execute(w, context)
	}
}

// CompleteTaskFunc handles the /complete/<id> URL and marks the task with ID passed as <id> to status complete
func CompleteTaskFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Redirect(w, r, "/", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(r.URL.Path[len("/complete/"):])
	log.Println(id)
	if err != nil {
		log.Println(err)
	} else {
		username := sessions.GetCurrentUserName(r)
		err = db.CompleteTask(username, id)
		if err != nil {
			message := "Complete task failed"
			log.Println(message)
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}

}
