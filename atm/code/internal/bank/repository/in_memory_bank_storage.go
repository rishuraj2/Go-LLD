package repository

import (
	"atm/internal/bank/domain/model/account"
	"atm/internal/bank/domain/model/card"
	"atm/internal/bank/domain/model/transaction"
	"errors"
	"sync"
)

type InMemoryBankStorage struct {
	accountData     map[string]account.Account
	cardData        map[string]card.Card
	transactionData map[string]transaction.Transaction
}

var (
	instance *InMemoryBankStorage
	once     sync.Once
)

func NewInMemoryBankStorage() *InMemoryBankStorage {
	once.Do(func() {
		instance = &InMemoryBankStorage{
			accountData:     make(map[string]account.Account),
			cardData:        make(map[string]card.Card),
			transactionData: make(map[string]transaction.Transaction),
		}
	})

	return instance
}

func (this *InMemoryBankStorage) CreateAccount(account account.Account) {
	this.accountData[account.GetAccountNumber()] = account
}

func (this *InMemoryBankStorage) FindAccountByNumber(accountNumber string) (account.Account, error) {
	for accNo, acc := range this.accountData {
		if accNo == accountNumber {
			return acc, nil
		}
	}

	return account.Account{}, errors.New("account not found")
}

func (this *InMemoryBankStorage) CreateCard(card card.Card) {
	this.cardData[card.GetCardNumber()] = card
}

func (this *InMemoryBankStorage) FindCardByCardNumber(cardNumber string) (card.Card, error) {
	for crdNo, crd := range this.cardData {
		if crdNo == cardNumber {
			return crd, nil
		}
	}

	return card.Card{}, errors.New("card not found")
}

func (this *InMemoryBankStorage) FindCardByAccountNumber(accountNumber string) (card.Card, error) {
	for _, crd := range this.cardData {
		if crd.GetAccountNumber() == accountNumber {
			return crd, nil
		}
	}

	return card.Card{}, errors.New("account not found")
}

func (this *InMemoryBankStorage) UpdateAccount(accountNumber string, account account.Account) error {
	for accNo := range this.accountData {
		if accNo == accountNumber {
			this.accountData[accNo] = account
			return nil
		}
	}

	return errors.New("account not found")
}

func (this *InMemoryBankStorage) CreateTransaction(transaction transaction.Transaction) {
	this.transactionData[transaction.GetID()] = transaction
}

func (this *InMemoryBankStorage) FindTransactionByID(trxnID string) (transaction.Transaction, error) {
	for id, trxn := range this.transactionData {
		if id == trxnID {
			return trxn, nil
		}
	}

	return transaction.Transaction{}, errors.New("transaction not found")
}

func (this *InMemoryBankStorage) FindAllTransactions(accountNumber string) []transaction.Transaction {
	var transactions []transaction.Transaction

	for _, trxn := range this.transactionData {
		if trxn.GetAccountNumber() == accountNumber {
			transactions = append(transactions, trxn)
		}
	}

	return transactions
}
