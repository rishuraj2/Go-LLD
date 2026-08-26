package state

import "atm/internal/atm/domain/atm"

type ProcessingState struct{}

func NewProcessingState() *ProcessingState {
	return &ProcessingState{}
}

func (this *ProcessingState) InsertCard(ctx *atm.ATM, cardNumber string) error {
	
}

func (this *ProcessingState) Authenticate(ctx *atm.ATM, pin string) error {

}

func (this *ProcessingState) Transact(ctx *atm.ATM, amount int) error {

}

func (this *ProcessingState) Process(ctx *atm.ATM) error {

}

func (this *ProcessingState) Dispense(ctx *atm.ATM) error {

}

func (this *ProcessingState) EjectCard(ctx *atm.ATM) error {

}
