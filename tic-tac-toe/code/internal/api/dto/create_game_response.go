package dto

type CreateGameResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	GameId  string `json:"game_id"`
}
