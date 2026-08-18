package enum

type Symbol int

const (
	X Symbol = iota
	O
	EMPTY
)

func (this Symbol) String() string {
	val := []string{"X", "O", "_"}
	if int(this) < len(val) {
		return val[int(this)]
	}

	return "unknown"
}
