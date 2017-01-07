package main

import (
	"fmt"
	"net/http"

	"github.com/thewhitetulip/Tasks-new/views"
)

func main() {
	views.PopulateTemplate()

	http.Handle("/static/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/completed/", views.RequiresLogin(views.ShowCompletedTasksFunc))
	http.HandleFunc("/complete/", views.CompleteTaskFunc)
	http.HandleFunc("/add/", views.AddTaskFunc)
	http.HandleFunc("/login/", views.LoginFunc)
	http.HandleFunc("/logout/", views.LogoutFunc)
	http.HandleFunc("/", views.RequiresLogin(views.HomeFunc))
	fmt.Println("Server running on 8081")
	err := http.ListenAndServe("0.0.0.0:8081", nil)
	if err != nil {
		fmt.Println("main.go: ", err)
	}
}
