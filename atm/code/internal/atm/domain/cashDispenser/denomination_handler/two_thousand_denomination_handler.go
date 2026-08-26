package denominationhandler

import (
	cashreservoir "atm/internal/atm/domain/cashReservoir"
	"atm/internal/atm/domain/enum"
)

type TwoThousandDenominationHandler struct {
	BaseDenominationHandler
}

func NewTwoThousandDenominationHandler() *TwoThousandDenominationHandler {
	return &TwoThousandDenominationHandler{
		BaseDenominationHandler: NewBaseDenominationHandler(),
	}
}

func (this *TwoThousandDenominationHandler) Handle(reservoir *cashreservoir.CashReservoir, amount int) int {
	noteCount := 0

	if amount >= 2000 {
		noteCount = int(amount / 2000)
	}

	countInReservoir := reservoir.GetDenominationCount(enum.TWO_THOUSAND)

	if countInReservoir < noteCount {
		noteCount = countInReservoir
	}

	amountToRemove := map[enum.Denomination]int{
		enum.TWO_THOUSAND: noteCount,
	}

	reservoir.RemoveCash(amountToRemove)

	remaining := amount - (noteCount * 2000)

	return this.Forward(reservoir, remaining)
}
