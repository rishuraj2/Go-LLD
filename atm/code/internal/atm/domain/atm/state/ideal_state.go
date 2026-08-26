package state

import (
	"atm/internal/atm/domain/atm"
	"atm/internal/atm/domain/enum"
	"errors"
)

type IdealState struct{}

func NewIdealState() *IdealState {
	return &IdealState{}
}

func (this *IdealState) InsertCard(ctx *atm.ATM, cardNumber string) error {
	if err := ctx.GetCardReader().ReadCard(cardNumber); err != nil {
		return err
	}

	ctx.SetState(NewCardInsertedState())
	return nil
}

func (this *IdealState) Authenticate(ctx *atm.ATM, pin string) error {
	return errors.New("card not inserted")
}

func (this *IdealState) Transact(ctx *atm.ATM, transactionType enum.TransactionType) error {
	return errors.New("card not inserted")
}

func (this *IdealState) Process(ctx *atm.ATM) error {
	return errors.New("card not inserted")
}

func (this *IdealState) Dispense(ctx *atm.ATM) error {
	return errors.New("card not inserted")
}

func (this *IdealState) EjectCard(ctx *atm.ATM) error {
	return errors.New("card not inserted")
}
