package usecase

import (
	"tictactoe/internal/repository"
)

type PlayerResult struct {
	Name   string
	Symbol string
}

type GetGameResult struct {
	GameID       string
	Players      [2]PlayerResult
	Board        [][]string
	ActivePlayer string
	GameState    string
}

type GetGameUseCase struct {
	repository repository.GameRepository
}

func NewGetGameUseCase(repo repository.GameRepository) *GetGameUseCase {
	return &GetGameUseCase{
		repository: repo,
	}
}

func (this *GetGameUseCase) Execute(gameID string) (GetGameResult, error) {
	game, err := this.repository.GetGameByID(gameID)
	if err != nil {
		return GetGameResult{}, err
	}

	players := game.GetPlayers()

	playerResponse := [2]PlayerResult{
		{
			Name:   players[0].GetName(),
			Symbol: players[0].GetSymbol().String(),
		},
		{
			Name:   players[1].GetName(),
			Symbol: players[1].GetSymbol().String(),
		},
	}

	return GetGameResult{
		GameID:       game.GetID(),
		Players:      playerResponse,
		Board:        game.GetBoard(),
		ActivePlayer: players[game.GetActivePlayerIndex()].GetName(),
		GameState:    game.GetGameState().String(),
	}, nil
}
