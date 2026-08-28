package denominationhandler

import (
	cashreservoir "atm/internal/atm/domain/cashReservoir"
	"atm/internal/atm/domain/enum"
)

type TenDenominationHandler struct {
	BaseDenominationHandler
}

func NewTenDenominationHandler() *TenDenominationHandler {
	return &TenDenominationHandler{
		BaseDenominationHandler: NewBaseDenominationHandler(),
	}
}

func (this *TenDenominationHandler) Handle(reservoir *cashreservoir.CashReservoir, amount int) int {
	noteCount := 0

	if amount >= 10 {
		noteCount = int(amount / 10)
	}

	countInReservoir := reservoir.GetDenominationCount(enum.TEN)

	if countInReservoir < noteCount {
		noteCount = countInReservoir
	}

	amountToRemove := map[enum.Denomination]int{
		enum.TEN: noteCount,
	}

	reservoir.RemoveCash(amountToRemove)

	remaining := amount - (noteCount * 10)

	return this.Forward(reservoir, remaining)
}
