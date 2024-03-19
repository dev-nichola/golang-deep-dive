package main

import "fmt"

type AccountRepository struct {
	AccountNumber string
}

func (a *AccountRepository) setAccountNumber(accNumber string) {
	a.AccountNumber = accNumber
}

func (a AccountRepository) getAccountNumber() string {
	return a.AccountNumber
}

func main() {
	accountRepo := &AccountRepository{}
	accountRepo.setAccountNumber("08213983947329")

	fmt.Println(accountRepo.getAccountNumber())
}
