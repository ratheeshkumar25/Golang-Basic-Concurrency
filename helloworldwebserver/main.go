package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	fmt.Fprint(w, "Hello World - Welcome to world of go progrmaming , go is introduced by google\n, currently one of popular langauage by unique feature of concurrency\n, apart from go is first choice of cloud native application langauage")
}

func main() {
	fmt.Println("Server connecting to localhost :8080/hello")
	http.HandleFunc("/hello", HelloServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
