package main

import "math/rand"

type Account struct {
	ID        int
	FirstName string
	LastName  string
	Number    int64
	balance   int64 // FLOat ?
}

func NewAccount(firstName string, lastName string) *Account {
	return &Account{
		ID:        rand.Intn(10000),
		FirstName: firstName,
		LastName:  lastName,
		Number:    rand.Int63(),
		// balance goes to 0 normally with go
	}
}
