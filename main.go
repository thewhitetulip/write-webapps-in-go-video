package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/profile", ProfileFunc)
	http.HandleFunc("/", HomeFuncß)
	fmt.Println("Server running on 8081")
	http.ListenAndServe("0.0.0.0:8081", nil)
}

// HomeFuncß handles the / URL and asks the name of the user in German.
func HomeFuncß(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hallo Wie heißen Sie?")
}

// ProfileFunc handles the /profile URL and shows the profile page of the logged in user.
func ProfileFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You are on the profile page!!")
}
