package denominationhandler

import (
	cashreservoir "atm/internal/atm/domain/cashReservoir"
	"atm/internal/atm/domain/enum"
)

type FiftyDenominationHandler struct {
	BaseDenominationHandler
}

func NewFiftyDenominationHandler() *FiftyDenominationHandler {
	return &FiftyDenominationHandler{
		BaseDenominationHandler: NewBaseDenominationHandler(),
	}
}

func (this *FiftyDenominationHandler) Handle(reservoir *cashreservoir.CashReservoir, amount int) int {
	noteCount := 0

	if amount >= 50 {
		noteCount = int(amount / 50)
	}

	countInReservoir := reservoir.GetDenominationCount(enum.FIFTY)

	if countInReservoir < noteCount {
		noteCount = countInReservoir
	}

	amountToRemove := map[enum.Denomination]int{
		enum.FIFTY: noteCount,
	}

	reservoir.RemoveCash(amountToRemove)

	remaining := amount - (noteCount * 50)

	return this.Forward(reservoir, remaining)
}
