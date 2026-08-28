package usecase

import (
	"atm/internal/bank/domain/model/transaction"
	"atm/internal/bank/repository"
)

type TransactionHistoryUseCase struct {
	repository repository.BankRepository
}

func NewTransactionHistoryUseCase(repo repository.BankRepository) *TransactionHistoryUseCase {
	return &TransactionHistoryUseCase{
		repository: repo,
	}
}

func (this *TransactionHistoryUseCase) Execute(accountNumber string) []transaction.Transaction {
	return this.repository.FindAllTransactions(accountNumber)
}
