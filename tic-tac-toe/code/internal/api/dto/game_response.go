package dto

type PlayerResponse struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type GameResponseData struct {
	GameID       string            `json:"game_id"`
	Players      [2]PlayerResponse `json:"players"`
	Board        [][]string        `json:"board"`
	ActivePlayer string            `json:"active_player"`
	GameState    string            `json:"game_state"`
}

type GameResponse struct {
	Status  string           `json:"status"`
	Message string           `json:"message"`
	Game    GameResponseData `json:"game"`
}
