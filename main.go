package main

import "fmt"

func main() {
	server := NewAPIServer(":3000")
	println("Server running on %s", server.listenAddr)
	server.Run()

	//fmt.Println("Go REST API v1.0")
	//
	//mux := http.NewServeMux()
	//
	//// OLD
	//mux.HandleFunc("GET /comment", func(w http.ResponseWriter, r *http.Request) {
	//
	//	fmt.Fprint(w, "return all comment")
	//})
	//
	//mux.HandleFunc("GET /comment{id}", func(w http.ResponseWriter, r *http.Request) {
	//	id := r.PathValue("id")
	//	fmt.Fprint(w, "return a single comment with id: %s", id)
	//})
	//
	//mux.HandleFunc("POST /comment", func(w http.ResponseWriter, r *http.Request) {
	//
	//	fmt.Fprint(w, "Post a comment")
	//})
	//
	//if err := http.ListenAndServe("localhost:8080", mux); err != nil {
	//	fmt.Println(err.Error())
	//}

	fmt.Println("YO")
}
