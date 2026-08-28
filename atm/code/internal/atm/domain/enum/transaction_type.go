package enum

type TransactionType int

const (
	CHECK_BALANCE TransactionType = iota
	WITHDRAW
	DEPOSIT
	CHECK_HISTORY
)


func (this TransactionType) String() string {
	val := []string{"check balance", "withdraw", "deposit", "transfer", "check history"}

	if int(this) < len(val) {
		return val[int(this)]
	}

	return "unknown"
}