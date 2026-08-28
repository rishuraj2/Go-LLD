package account

import (
	"atm/internal/bank/domain/enum"
	"fmt"
	"strings"
	"sync/atomic"
)

type AccountNumberGenerator struct {
	count atomic.Uint64
}

func NewAccountNumberGenerator() *AccountNumberGenerator {
	return &AccountNumberGenerator{}
}

func (this *AccountNumberGenerator) Generate(accountType enum.AccountType) string {
	count := this.count.Add(1)

	prefix := strings.ToUpper(accountType.String()[:4])

	return fmt.Sprintf("%s%06d", prefix, count)
}
