package state

import (
	"atm/internal/atm/domain/atm"
	"atm/internal/atm/domain/enum"
	"errors"
)

type TransactedState struct{}

func NewTransactedState() *TransactedState {
	return &TransactedState{}
}

func (this *TransactedState) InsertCard(ctx *atm.ATM, cardNumber string) error {
	return errors.New("card already inserted")
}

func (this *TransactedState) Authenticate(ctx *atm.ATM, pin string) error {
	return errors.New("card already authenticated")
}

func (this *TransactedState) Transact(ctx *atm.ATM, transactionType enum.TransactionType) error {
	return errors.New("transaction already seletced")
}

func (this *TransactedState) Process(ctx *atm.ATM) error {

}

func (this *TransactedState) Dispense(ctx *atm.ATM) error {

}

func (this *TransactedState) EjectCard(ctx *atm.ATM) error {
	ctx.SetTransactionType(-1)
	ctx.SetAccountNumber("")
	ctx.GetCardReader().EjectCard()
	ctx.SetState(NewIdealState())
	return nil
}
