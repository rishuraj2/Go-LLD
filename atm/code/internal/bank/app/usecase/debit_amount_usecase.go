package usecase

import (
	"atm/internal/bank/domain/enum"
	"atm/internal/bank/domain/model/transaction"
	"atm/internal/bank/repository"
)

type DebitAmountUseCase struct {
	repository repository.BankRepository
}

func NewDebitAmountUseCase(repo repository.BankRepository) *DebitAmountUseCase {
	return &DebitAmountUseCase{
		repository: repo,
	}
}

func (this *DebitAmountUseCase) Execute(accountNumber string, amount float64) (string, error) {
	account, err := this.repository.FindAccountByNumber(accountNumber)
	if err != nil {
		return "", err
	}

	err = account.Debit(amount)
	if err != nil {
		return "", err
	}

	err = this.repository.UpdateAccount(accountNumber, account)
	if err != nil {
		return "", err
	}

	trxn := transaction.NewTransaction(transaction.NewTransactionIDGenerator().Generate(), enum.WITHDRAW, accountNumber, amount)
	this.repository.CreateTransaction(trxn)

	return trxn.GetID(), nil
}
