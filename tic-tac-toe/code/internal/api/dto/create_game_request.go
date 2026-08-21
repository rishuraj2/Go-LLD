package dto

type CreateGameRequest struct {
	Player1Name string `json:"player1_name"`
	Player2Name string `json:"player2_name"`
	BoardSize   int    `json:"board_size"`
}
