package repository

import (
	"atm/internal/bank/domain/model/account"
	"atm/internal/bank/domain/model/card"
	"atm/internal/bank/domain/model/transaction"
)

type BankRepository interface {
	CreateAccount(account account.Account)
	FindAccountByNumber(accountNumber string) (account.Account, error)
	CreateCard(card card.Card)
	FindCardByCardNumber(cardNumber string) (card.Card, error)
	FindCardByAccountNumber(accountNumber string) (card.Card, error)
	UpdateAccount(accountNumber string, account account.Account) error
	CreateTransaction(transaction transaction.Transaction)
	FindTransactionByID(trxnID string) (transaction.Transaction, error)
	FindAllTransactions(accountNumber string) []transaction.Transaction
}
