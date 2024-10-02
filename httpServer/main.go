package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func decoderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid json data", http.StatusBadRequest)
		return
	}

	response := fmt.Sprintf("%s is % d year old", user.Name, user.Age)
	w.Header().Set("Content-type", "text/plain")
	fmt.Fprintln(w, response)
}

func encodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid json data", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "Error encoding json", http.StatusInternalServerError)
		return
	}
}

func main() {
	add := ":8080"
	fmt.Print("Starting server at", add)

	http.HandleFunc("/decode", decoderHandler)
	http.HandleFunc("/encode", encodeHandler)
	if err := http.ListenAndServe(add, nil); err != nil {
		log.Fatalf("Could not start server %s", err.Error())
	}
}
