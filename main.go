package main

import "fmt"

func main() {
	server := NewAPIServer(":3000")
	println("Server running on %s", server.listenAddr)
	server.Run()

	fmt.Println("YO")
}
