package cashreservoir

import "atm/internal/atm/domain/enum"

type CashReservoir struct {
	cash map[enum.Denomination]int
}

func NewCashReservoir() *CashReservoir {
	return &CashReservoir{
		cash: make(map[enum.Denomination]int),
	}
}

func (this *CashReservoir) GetCopy() *CashReservoir {
	cashCopy := make(map[enum.Denomination]int)

	for denomination, count := range this.cash {
		cashCopy[denomination] = count
	}

	return &CashReservoir{
		cash: cashCopy,
	}
}

func (this *CashReservoir) GetCash() map[enum.Denomination]int {
	result := make(map[enum.Denomination]int, len(this.cash))

	for denomination, count := range this.cash {
		result[denomination] = count
	}

	return result
}

func (this *CashReservoir) GetDenominationCount(denomination enum.Denomination) int {
	return this.cash[denomination]
}

func (this *CashReservoir) GetTotalBalance() int {
	total := 0

	for denomination, count := range this.cash {
		switch denomination {
		case enum.TWO_THOUSAND:
			total += count * 2000
		case enum.FIVE_HUNDRED:
			total += count * 500
		case enum.TWO_HUNDRED:
			total += count * 200
		case enum.ONE_HUNDRED:
			total += count * 100
		case enum.FIFTY:
			total += count * 50
		case enum.TWENTY:
			total += count * 20
		case enum.TEN:
			total += count * 10
		}
	}

	return total
}

func (this *CashReservoir) AddCash(cash map[enum.Denomination]int) {
	for denomination, count := range cash {
		this.cash[denomination] += count
	}
}

func (this *CashReservoir) RemoveCash(cash map[enum.Denomination]int) {
	for denomination, count := range cash {
		this.cash[denomination] -= count
	}
}
