package enum

type ATMState int

const (
	IDEAL ATMState = iota
	CARD_INSERTED
	AUTHENTICATED
	TRASACTING
	PROCESSING
	DISPENSING
	CARD_EJECTED
)

func (this ATMState) String() string {
	val := []string{"ideal", "card inserted", "authenticated", "transacting", "processing", "dispensing", "card ejected"}

	if int(this) < len(val) {
		return val[int(this)]
	}

	return "unknown"
}
