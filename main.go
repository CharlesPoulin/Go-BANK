package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Go REST API v1.0")

	mux := http.NewServeMux()

	// OLD
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println(err.Error())
	}
}
