package usecase

import (
	"atm/internal/bank/domain/model/account"
	"atm/internal/bank/repository"
	"errors"
)

type AuthenticateCardUseCase struct {
	repository repository.BankRepository
}

func NewAuthenticateCardUseCase(repo repository.BankRepository) *AuthenticateCardUseCase {
	return &AuthenticateCardUseCase{
		repository: repo,
	}
}

func (this *AuthenticateCardUseCase) Execute(cardNumber, pin string) (account.Account, error) {
	card, err := this.repository.FindCardByCardNumber(cardNumber)
	if err != nil {
		return account.Account{}, err
	}

	if card.GetPin() == pin {
		return this.repository.FindAccountByNumber(card.GetAccountNumber())
	}

	return account.Account{}, errors.New("invalid card pin")

}
