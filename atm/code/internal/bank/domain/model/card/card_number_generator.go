package card

import (
	"atm/internal/bank/domain/enum"
	"fmt"
	"strings"
	"sync/atomic"
)

type CardNumberGenerator struct {
	count atomic.Uint64
}

func NewCardNumberGenerator() *CardNumberGenerator {
	return &CardNumberGenerator{}
}

func (this *CardNumberGenerator) Generate(accountType enum.AccountType) string {
	count := this.count.Add(1)

	prefix := "4388"

	mid := strings.ToUpper(accountType.String()[:4])

	return fmt.Sprintf("%s%s%08d", prefix, mid, count)
}
