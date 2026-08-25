package usecase

import (
	"atm/internal/bank/domain/enum"
	"atm/internal/bank/domain/model/transaction"
	"atm/internal/bank/repository"
)

type CreditAmountUseCase struct {
	repository repository.BankRepository
}

func NewCreditAmountUseCase(repo repository.BankRepository) *CreditAmountUseCase {
	return &CreditAmountUseCase{
		repository: repo,
	}
}

func (this *CreditAmountUseCase) Execute(accountNumber string, amount float64) (string, error) {
	account, err := this.repository.FindAccountByNumber(accountNumber)
	if err != nil {
		return "", err
	}

	err = account.Credit(amount)
	if err != nil {
		return "", err
	}

	err = this.repository.UpdateAccount(accountNumber, account)
	if err != nil {
		return "", err
	}

	trxn := transaction.NewTransaction(transaction.NewTransactionIDGenerator().Generate(), enum.DEPOSIT, accountNumber, amount)
	this.repository.CreateTransaction(trxn)
	
	return trxn.GetID(), nil
}
