package enum

type AccountType int

const (
	SAVINGS AccountType = iota
	CURRENT
)

func (this AccountType) String() string {
	val := []string{"savings", "current"}
	
	if int(this) < len(val) {
		return val[int(this)]
	}

	return "unknown"
}
