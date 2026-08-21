package enum

type GameState int

const (
	IN_PROGRESS GameState = iota
	WINNER_X
	WINNER_O
	DRAW
)

func (this GameState) String() string {
	val := []string{"in progress", "winner x", "winner o", "draw"}
	if int(this) < len(val) {
		return val[int(this)]
	}
	
	return "unknown"
}
