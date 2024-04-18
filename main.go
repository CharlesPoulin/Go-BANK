package main

import (
	"fmt"
	"log"
)

func main() {

	store, err := NewMySQLStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.init(); err != nil {
		log.Fatal(err)
	
	}

	fmt.Printf("%+v/n", store)

	server := NewAPIServer(":3000", store)
	server.Run()

	fmt.Println("it worked")
}
