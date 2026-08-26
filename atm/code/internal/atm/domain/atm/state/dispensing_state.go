package state

import "atm/internal/atm/domain/atm"

type DispensingState struct{}

func NewDispensingState() *DispensingState {
	return &DispensingState{}
}

func (this *DispensingState) InsertCard(ctx *atm.ATM, cardNumber string) error {
	
}

func (this *DispensingState) Authenticate(ctx *atm.ATM, pin string) error {

}

func (this *DispensingState) Transact(ctx *atm.ATM, amount int) error {

}

func (this *DispensingState) Process(ctx *atm.ATM) error {

}

func (this *DispensingState) Dispense(ctx *atm.ATM) error {

}

func (this *DispensingState) EjectCard(ctx *atm.ATM) error {

}
