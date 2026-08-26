package usecase

import (
	"atm/internal/bank/domain/model/card"
	"atm/internal/bank/repository"
	"errors"
)

type CreateCardUseCase struct {
	repository repository.BankRepository
}

func NewCreateCardUseCase(repo repository.BankRepository) *CreateCardUseCase {
	return &CreateCardUseCase{
		repository: repo,
	}
}

func (this *CreateCardUseCase) Execute(accountNumber, pin string) (string, error) {
	account, err := this.repository.FindAccountByNumber(accountNumber)
	if err != nil {
		return "", err
	}

	if _, err := this.repository.FindCardByAccountNumber(accountNumber); err == nil {
		return "", errors.New("card already exists")
	}

	crd := card.NewCard(card.NewCardNumberGenerator().Generate(account.GetType()), accountNumber, pin)

	return crd.GetCardNumber(), nil
}
