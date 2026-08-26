package usecase

import (
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

func (this *AuthenticateCardUseCase) Execute(cardNumber, pin string) (string, error) {
	card, err := this.repository.FindCardByCardNumber(cardNumber)
	if err != nil {
		return "", err
	}

	if card.GetPin() == pin {
		acc, err := this.repository.FindAccountByNumber(card.GetAccountNumber())
		if err != nil {
			return "", err
		}
		return acc.GetAccountNumber(), nil
	}

	return "", errors.New("invalid card pin")

}
