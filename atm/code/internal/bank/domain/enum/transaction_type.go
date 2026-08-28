package enum

type TransactionType int

const (
	WITHDRAW TransactionType = iota
	DEPOSIT
	BALANCE_INQUIRY
)

func (this TransactionType) String() string {
	val := []string {"withdraw", "deposit", "balance inquiry"}

	if int(this) < len(val) {
		return val[int(this)]
	}

	return "unknown"
}
