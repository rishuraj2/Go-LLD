package denominationhandler

import cashreservoir "atm/internal/atm/domain/cashReservoir"

type DenominationHandler interface {
	SetNext(next DenominationHandler)
	Handle(reservoir *cashreservoir.CashReservoir, amount int) int
}

type BaseDenominationHandler struct {
	next DenominationHandler
}

func NewBaseDenominationHandler() BaseDenominationHandler {
	return BaseDenominationHandler{}
}

func (this *BaseDenominationHandler) SetNext(next DenominationHandler) {
	this.next = next
}

func (this *BaseDenominationHandler) Forward(reservoir *cashreservoir.CashReservoir, amount int) int {
	if this.next != nil {
		return this.next.Handle(reservoir, amount)
	}

	return amount
}
