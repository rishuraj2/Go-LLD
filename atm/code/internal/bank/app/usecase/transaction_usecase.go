package usecase

import (
	"atm/internal/bank/domain/model/transaction"
	"atm/internal/bank/repository"
)

type TransactionUseCase struct {
	repository repository.BankRepository
}

func NewTransactionUseCase(repo repository.BankRepository) *TransactionUseCase {
	return &TransactionUseCase{
		repository: repo,
	}
}

func (this *TransactionUseCase) Execute(id string) (transaction.Transaction, error) {
	return this.repository.FindTransactionByID(id)
}
