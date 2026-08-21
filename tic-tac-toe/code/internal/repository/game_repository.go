package repository

import "tictactoe/internal/domain/model"

type GameRepository interface {
	SaveGame(data model.Game) error
	UpdateGame(data model.Game) error
	GetGameByID(id string) (model.Game, error)
}
