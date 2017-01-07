package db

import (
	"database/sql"
	"github.com/thewhitetulip/Tasks-new/types"
	"html/template"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	md "github.com/shurcooL/github_flavored_markdown"
)

var db *sql.DB
var taskStatus map[string]int
var err error
var status string

func init() {
	db, err = sql.Open("sqlite3", "./tasks.db")
	taskStatus = map[string]int{"COMPLETE": 1, "PENDING": 2, "DELETED": 3}
	if err != nil {
		log.Fatal(err)
	}
}

func Close() {
	db.Close()
}

// GetPendingTasks returns the pending tasks of the username passed as the argument
func GetPendingTasks(username string) (types.Context, error) {
	log.Println("getting tasks for ", status)
	var tasks []types.Task
	var categories []types.Category
	var task types.Task

	var TaskCreated time.Time
	var context types.Context
	var rows *sql.Rows

	basicSQL := "select t.id, title, content, created_date, priority, case when c.name is null then 'NA' else c.name end from task t, status s, user u left outer join  category c on c.id=t.cat_id where u.username=? and s.id=t.task_status_id and u.id=t.user_id and s.status='PENDING' and t.hide!=1 order by t.created_date asc"

	rows, err := db.Query(basicSQL, "suraj")
	if err != nil {
		log.Println("something went wrong in fetching tasks")
		return types.Context{}, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&task.ID, &task.Title, &task.Content, &TaskCreated, &task.Priority, &task.Category)
		task.HTMLContent = template.HTML(md.Markdown([]byte(task.Content)))
		TaskCreated = TaskCreated.Local()
		task.Created = TaskCreated.Format("Jan 2 2006")

		tasks = append(tasks, task)
	}

	categories = GetCategories(username)

	context = types.Context{Tasks: tasks, CSRFToken: "supersecret", Categories: categories}
	return context, nil
}

//GetCategories will return the list of categories to be
//rendered in the template
func GetCategories(username string) []types.Category {
	stmt := "select name from category c, user u where u.id = c.user_id and u.username=?"
	rows, _ := db.Query(stmt, username)
	var categories []types.Category
	var category types.Category

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&category.Name)
		if err != nil {
			log.Println(err)
		}
		categories = append(categories, category)
	}
	return categories
}
