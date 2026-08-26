package denominationhandler

import (
	cashreservoir "atm/internal/atm/domain/cashReservoir"
	"atm/internal/atm/domain/enum"
)

type TwentyDenominationHandler struct {
	BaseDenominationHandler
}

func NewTwentyDenominationHandler() *TwentyDenominationHandler {
	return &TwentyDenominationHandler{
		BaseDenominationHandler: NewBaseDenominationHandler(),
	}
}

func (this *TwentyDenominationHandler) Handle(reservoir *cashreservoir.CashReservoir, amount int) int {
	noteCount := 0

	if amount >= 20 {
		noteCount = int(amount / 20)
	}

	countInReservoir := reservoir.GetDenominationCount(enum.TWENTY)

	if countInReservoir < noteCount {
		noteCount = countInReservoir
	}

	amountToRemove := map[enum.Denomination]int{
		enum.TWENTY: noteCount,
	}

	reservoir.RemoveCash(amountToRemove)

	remaining := amount - (noteCount * 20)

	return this.Forward(reservoir, remaining)
}
