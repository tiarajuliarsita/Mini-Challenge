package main

import (
	"errors"
	"html/template"
	"net/http"
)

type Person struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Address  string `json:"address"`
	Age      int    `json:"age"`
}

var persons = []Person{
	{Username: "tiara", Email: "tiara@gmail.com", Address: "palu", Age: 15},
	{Username: "yaya", Email: "yaya@gmail.com", Address: "soul", Age: 15},
	{Username: "ya", Email: "ya@gmail.com", Address: "palu", Age: 15},
}

func main() {
	http.HandleFunc("/login", Login)
	http.ListenAndServe(":8080", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.ServeFile(w, r, "home.html")
		return
	}

	for _, person := range persons {
		if person.Username == username {
			tpl, err := template.ParseFiles("index.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			tpl.Execute(w, person)
			return
		}
	}

	// http.NotFound(w, r)
	err := errors.New("User Not Found")
	http.Error(w, err.Error(), http.StatusNotFound)
}
