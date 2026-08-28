package denominationhandler

import (
	cashreservoir "atm/internal/atm/domain/cashReservoir"
	"atm/internal/atm/domain/enum"
)

type OneHundredDenominationHandler struct {
	BaseDenominationHandler
}

func NewOneHundredDenominationHandler() *OneHundredDenominationHandler {
	return &OneHundredDenominationHandler{
		BaseDenominationHandler: NewBaseDenominationHandler(),
	}
}

func (this *OneHundredDenominationHandler) Handle(reservoir *cashreservoir.CashReservoir, amount int) int {
	noteCount := 0

	if amount >= 100 {
		noteCount = int(amount / 100)
	}

	countInReservoir := reservoir.GetDenominationCount(enum.ONE_HUNDRED)

	if countInReservoir < noteCount {
		noteCount = countInReservoir
	}

	amountToRemove := map[enum.Denomination]int{
		enum.ONE_HUNDRED: noteCount,
	}

	reservoir.RemoveCash(amountToRemove)

	remaining := amount - (noteCount * 100)

	return this.Forward(reservoir, remaining)
}
