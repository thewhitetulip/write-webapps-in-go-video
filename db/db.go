package db

import (
	"github.com/thewhitetulip/Tasks-new/types"
	"html/template"

	md "github.com/shurcooL/github_flavored_markdown"
)

// GetPendingTasks returns the pending tasks of the username passed as the argument
func GetPendingTasks(username string) types.Context {
	var categories []types.Category
	category := types.Category{Name: "Random", TaskCount: "1"}
	categories = append(categories, category)
	category = types.Category{Name: "Very Random", TaskCount: "1"}
	categories = append(categories, category)

	var tasks types.Tasks
	task := types.Task{Title: "Task Title", Content: "Task Content\n_italics_", Category: "Random", Priority: "3", Created: "5 Jan 2017"}

	tasks = append(tasks, task)
	tasks = append(tasks, task)
	tasks = append(tasks, task)

	for i, _ := range tasks {
		tasks[i].HTMLContent = template.HTML(string(md.Markdown([]byte(tasks[i].Content))))
	}

	context := types.Context{Tasks: tasks, CSRFToken: "supersecret", Categories: categories}
	return context
}
