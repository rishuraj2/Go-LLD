package enum

type Symbol int

const (
	EMPTY Symbol = iota
	X
	O
)

func (this Symbol) String() string {
	val := []string{"_", "x", "o"}
	if int(this) < len(val) {
		return val[int(this)]
	}

	return "unknown"
}
