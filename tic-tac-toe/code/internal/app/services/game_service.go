package services

import (
	"tictactoe/internal/app/usecase"
	"tictactoe/internal/repository"
)

type GameService struct {
	createGameUsecase *usecase.CreateGameUseCase
	makeMoveUseCase   *usecase.MakeMoveUseCase
	getGameUseCase    *usecase.GetGameUseCase
}

func NewGameService(repo repository.GameRepository) *GameService {
	return &GameService{
		createGameUsecase: usecase.NewCreateGameUseCase(repo),
		makeMoveUseCase:   usecase.NewMakeMoveUseCase(repo),
		getGameUseCase:    usecase.NewGetGameUseCase(repo),
	}
}

func (this *GameService) CreateGame(player1Name, player2Name string, boardSize int) (string, error) {
	return this.createGameUsecase.Execute(player1Name, player2Name, boardSize)
}

func (this *GameService) MakeMove(gameID string, x, y int) error {
	return this.makeMoveUseCase.Execute(gameID, x, y)
}

func (this *GameService) GetGame(gameID string) (usecase.GetGameResult, error) {
	return this.getGameUseCase.Execute(gameID)
}
