package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, World!</h1> <hr> <h3>Version 1.0</h3>")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server starting on port 8080")
	http.ListenAndServe(":8080", nil)
}
