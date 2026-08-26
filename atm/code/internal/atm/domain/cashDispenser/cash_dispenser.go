package cashdispenser

import (
	denominationhandler "atm/internal/atm/domain/cashDispenser/denomination_handler"
	cashreservoir "atm/internal/atm/domain/cashReservoir"
	"errors"
)

type CashDispenser struct {
	chainHead denominationhandler.DenominationHandler
	reservoir *cashreservoir.CashReservoir
}

func NewCashDispenser(reservoir *cashreservoir.CashReservoir) *CashDispenser {
	twoThousandHandler := denominationhandler.NewTwoThousandDenominationHandler()
	fiveHundredHandler := denominationhandler.NewFiveHundredDenominationHandler()
	twoHundredHandler := denominationhandler.NewTwoHundredDenominationHandler()
	oneHundredHandler := denominationhandler.NewOneHundredDenominationHandler()
	fiftyHandler := denominationhandler.NewFiftyDenominationHandler()
	twentyHandler := denominationhandler.NewTwentyDenominationHandler()
	tenHandler := denominationhandler.NewTenDenominationHandler()

	twoThousandHandler.SetNext(fiveHundredHandler)
	fiveHundredHandler.SetNext(twoHundredHandler)
	twoHundredHandler.SetNext(oneHundredHandler)
	oneHundredHandler.SetNext(fiftyHandler)
	fiftyHandler.SetNext(twentyHandler)
	twentyHandler.SetNext(tenHandler)

	return &CashDispenser{
		chainHead: twoThousandHandler,
		reservoir: reservoir,
	}

}

func (this *CashDispenser) Dispense(amount int) error {
	remainingAmount := this.chainHead.Handle(this.reservoir.GetCopy(), amount)
	if remainingAmount == 0 {
		this.chainHead.Handle(this.reservoir, amount)
		return nil
	}

	return errors.New("cannot dispense")
}
