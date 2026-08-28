package denominationhandler

import (
	cashreservoir "atm/internal/atm/domain/cashReservoir"
	"atm/internal/atm/domain/enum"
)

type FiveHundredDenominationHandler struct {
	BaseDenominationHandler
}

func NewFiveHundredDenominationHandler() *FiveHundredDenominationHandler {
	return &FiveHundredDenominationHandler{
		BaseDenominationHandler: NewBaseDenominationHandler(),
	}
}

func (this *FiveHundredDenominationHandler) Handle(reservoir *cashreservoir.CashReservoir, amount int) int {
	noteCount := 0

	if amount >= 500 {
		noteCount = int(amount / 500)
	}

	countInReservoir := reservoir.GetDenominationCount(enum.FIVE_HUNDRED)

	if countInReservoir < noteCount {
		noteCount = countInReservoir
	}

	amountToRemove := map[enum.Denomination]int{
		enum.FIVE_HUNDRED: noteCount,
	}

	reservoir.RemoveCash(amountToRemove)

	remaining := amount - (noteCount * 500)

	return this.Forward(reservoir, remaining)
}
