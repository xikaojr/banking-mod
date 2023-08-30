package models

type BankAccount struct {
	AccountNumber int
	AccountType   string
	Balance       float64
	Bonus         float64
	Person        Person
}

type Person struct {
	Name string
}
