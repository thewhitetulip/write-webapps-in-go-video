package types

type Category struct {
	Name      string
	ID        string
	TaskCount string
}

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
	Priority string
	Category string
	Hidden   string
	Created  string
	Comments []Comment
}

type Context struct {
	Tasks      []Task
	CSRFToken  string
	Categories []Category
}

type Comments []Comment
type Tasks []Task
