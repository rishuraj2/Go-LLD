package state

import (
	"atm/internal/atm/domain/atm"
	"atm/internal/atm/domain/enum"
	"errors"
)

type AuthenticatedState struct{}

func NewAuthenticatedState() *AuthenticatedState {
	return &AuthenticatedState{}
}

func (this *AuthenticatedState) InsertCard(ctx *atm.ATM, cardNumber string) error {
	return errors.New("card already inserted")
}

func (this *AuthenticatedState) Authenticate(ctx *atm.ATM, pin string) error {
	return errors.New("card already authenticated")
}

func (this *AuthenticatedState) Transact(ctx *atm.ATM, transactionType enum.TransactionType) error {
	ctx.SetTransactionType(transactionType)
	ctx.SetState(NewTransactedState())
	return nil
}

func (this *AuthenticatedState) Process(ctx *atm.ATM) error {
	return errors.New("yet to transact")
}

func (this *AuthenticatedState) Dispense(ctx *atm.ATM) error {
	return errors.New("yet to transact")
}

func (this *AuthenticatedState) EjectCard(ctx *atm.ATM) error {
	ctx.SetAccountNumber("")
	ctx.GetCardReader().EjectCard()
	ctx.SetState(NewIdealState())
	return nil
}
