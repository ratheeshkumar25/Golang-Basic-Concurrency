package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users []User

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	json.NewDecoder(r.Body).Decode(&newUser)
	users = append(users, newUser)
	w.WriteHeader(201)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
	w.WriteHeader(200)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id ", 400)
		return
	}
	var updateUser User

	json.NewDecoder(r.Body).Decode(&updateUser)

	for i, user := range users {
		if user.Id == id {
			users[i].Name = updateUser.Name
			users[i].Email = updateUser.Email
			w.WriteHeader(200)
			return
		}
	}
	http.Error(w, "User not found", 404)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id", 404)
		return
	}

	for i, user := range users {
		if user.Id == id {
			users = append(users[:i], users[i+1:]...)
			w.WriteHeader(204)
			return
		}
	}
	http.Error(w, "User not found", 404)

}

func main() {
	http.HandleFunc("/user", createUser)
	http.HandleFunc("/user/list", getUser)
	http.HandleFunc("/edit", updateUser)
	http.HandleFunc("/delete", deleteUser)

	err := http.ListenAndServe(":8080", nil)
	fmt.Println("Server running", err)
}
