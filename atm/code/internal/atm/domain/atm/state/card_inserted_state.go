package state

import (
	"atm/internal/atm/domain/atm"
	"atm/internal/atm/domain/enum"
	"errors"
)

type CardInsertedState struct{}

func NewCardInsertedState() *CardInsertedState {
	return &CardInsertedState{}
}

func (this *CardInsertedState) InsertCard(ctx *atm.ATM, cardNumber string) error {
	return errors.New("card already inserted")
}

func (this *CardInsertedState) Authenticate(ctx *atm.ATM, pin string) error {
	acc, err := ctx.GetBankServer().AuthenticateCard(ctx.GetCardReader().GetCardNumber(), pin)
	if err != nil {
		return err
	}

	ctx.SetAccountNumber(acc)
	ctx.SetState(NewAuthenticatedState())
	return nil
}

func (this *CardInsertedState) Transact(ctx *atm.ATM, transactionType enum.TransactionType) error {
	return errors.New("card not authenticated")
}

func (this *CardInsertedState) Process(ctx *atm.ATM) error {
	return errors.New("card not authenticated")
}

func (this *CardInsertedState) Dispense(ctx *atm.ATM) error {
	return errors.New("card not authenticated")
}

func (this *CardInsertedState) EjectCard(ctx *atm.ATM) error {
	ctx.GetCardReader().EjectCard()
	ctx.SetState(NewIdealState())
	return nil
}
