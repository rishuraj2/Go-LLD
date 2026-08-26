package state

import "atm/internal/atm/domain/atm"

type CardEjectedState struct{}

func NewCardEjectedState() *CardEjectedState {
	return &CardEjectedState{}
}

func (this *CardEjectedState) InsertCard(ctx *atm.ATM, cardNumber string) error {
	
}

func (this *CardEjectedState) Authenticate(ctx *atm.ATM, pin string) error {

}

func (this *CardEjectedState) Transact(ctx *atm.ATM, amount int) error {

}

func (this *CardEjectedState) Process(ctx *atm.ATM) error {

}

func (this *CardEjectedState) Dispense(ctx *atm.ATM) error {

}

func (this *CardEjectedState) EjectCard(ctx *atm.ATM) error {

}
