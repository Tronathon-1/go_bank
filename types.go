package main

import "math/rand/v2"

type Account struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Number    int64  `json:"number"`
	Balance   int64  `json:"balance"`
}

func NewAccount(firstname, lastname string) *Account {
	return &Account{
		ID:        rand.IntN(10000),
		FirstName: firstname,
		LastName:  lastname,
		Number:    rand.Int64N(100000000),
	}
}
