package account

import (
	"atm/internal/bank/domain/enum"
	"errors"
)

type Account struct {
	accountNumber string
	accountType   enum.AccountType
	balance       float64
}

func NewAccount(number string, accType enum.AccountType) Account {
	return Account{
		accountNumber: number,
		accountType:   accType,
		balance:       0.0,
	}
}

func (this *Account) GetAccountNumber() string {
	return this.accountNumber
}

func (this *Account) GetType() enum.AccountType {
	return this.accountType
}

func (this *Account) GetBalance() float64 {
	return this.balance
}

func (this *Account) Debit(amount float64) error {
	if amount <= 0.0 {
		return errors.New("negative or zero amount")
	}

	if amount > this.balance {
		return errors.New("cannot debit more than account balance")
	}

	this.balance -= amount
	return nil
}

func (this *Account) Credit(amount float64) error {
	if amount <= 0.0 {
		return errors.New("negative or zero amount")
	}

	this.balance += amount
	return nil
}
