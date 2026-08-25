package usecase

import (
	"atm/internal/bank/domain/enum"
	"atm/internal/bank/domain/model/account"
	"atm/internal/bank/repository"
)

type CreateAccountUseCase struct {
	repository repository.BankRepository
}

func NewCreateAccountUseCase(repo repository.BankRepository) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		repository: repo,
	}
}

func (this *CreateAccountUseCase) Execute(accountType enum.AccountType) account.Account {
	account := account.NewAccount(account.NewAccountNumberGenerator().Generate(accountType), accountType)
	this.repository.CreateAccount(account)
	return account
}
