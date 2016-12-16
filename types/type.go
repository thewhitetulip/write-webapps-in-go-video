package types

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

type Comments []Comment
type Tasks []Task
