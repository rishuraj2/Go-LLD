package denominationhandler

import (
	cashreservoir "atm/internal/atm/domain/cashReservoir"
	"atm/internal/atm/domain/enum"
)

type TwoHundredDenominationHandler struct {
	BaseDenominationHandler
}

func NewTwoHundredDenominationHandler() *TwoHundredDenominationHandler {
	return &TwoHundredDenominationHandler{
		BaseDenominationHandler: NewBaseDenominationHandler(),
	}
}

func (this *TwoHundredDenominationHandler) Handle(reservoir *cashreservoir.CashReservoir, amount int) int {
	noteCount := 0

	if amount >= 200 {
		noteCount = int(amount / 200)
	}

	countInReservoir := reservoir.GetDenominationCount(enum.TWO_HUNDRED)

	if countInReservoir < noteCount {
		noteCount = countInReservoir
	}

	amountToRemove := map[enum.Denomination]int{
		enum.TWO_HUNDRED: noteCount,
	}

	reservoir.RemoveCash(amountToRemove)

	remaining := amount - (noteCount * 200)

	return this.Forward(reservoir, remaining)
}
