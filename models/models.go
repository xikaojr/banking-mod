package models

import "github.com/google/uuid"

type BankAccount struct {
	AccountNumber int
	AccountType   *AccountType
	Balance       float64
	Bonus         float64
	Person        Person
}

type AccountType struct {
	Id   uuid.UUID
	Name string
}

type Person struct {
	Name string
}
