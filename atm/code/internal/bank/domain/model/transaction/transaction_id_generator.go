package transaction

import (
	"fmt"
	"sync/atomic"
)

type TransactionIDGenerator struct {
	count atomic.Uint64
}

func NewTransactionIDGenerator() *TransactionIDGenerator {
	return &TransactionIDGenerator{}
}

func (this *TransactionIDGenerator) Generate() string {
	count := this.count.Add(1)
	return fmt.Sprintf("TRXN-%d", count)
}
