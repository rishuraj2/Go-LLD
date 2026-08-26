package enum

type Denomination int

const (
	TWO_THOUSAND Denomination = iota
	FIVE_HUNDRED
	TWO_HUNDRED
	ONE_HUNDRED
	FIFTY
	TWENTY
	TEN
)

func (this Denomination) String() string {
	val := []string{"two thousand", "five hundred", "two hundred", "one hundred", "fifty", "twenty", "ten"}

	if int(this) < len(val) {
		return val[int(this)]
	}

	return "unknown"
}
