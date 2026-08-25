package transaction

import (
	"atm/internal/bank/domain/enum"
	"time"
)

type Transaction struct {
	id              string
	transactionType enum.TransactionType
	accountNumber   string
	amount          float64
	timestamp       string
}

func NewTransaction(id string, transactionType enum.TransactionType, accountNumber string, amount float64) Transaction {
	return Transaction{
		id:              id,
		transactionType: transactionType,
		accountNumber:   accountNumber,
		amount:          amount,
		timestamp:       time.Now().Format("2006-01-02 15:04:05"),
	}
}

func (this *Transaction) GetID() string {
	return this.id
}

func (this *Transaction) GetType() enum.TransactionType {
	return this.transactionType
}

func (this *Transaction) GetAmount() float64 {
	return this.amount
}

func (this *Transaction) GetAccountNumber() string {
	return this.accountNumber
}
