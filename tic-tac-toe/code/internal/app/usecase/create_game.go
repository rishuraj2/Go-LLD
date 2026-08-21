package usecase

import (
	"fmt"
	"strings"
	"tictactoe/internal/domain/enum"
	"tictactoe/internal/domain/model"
	"tictactoe/internal/repository"
)

type CreateGameUseCase struct {
	repository repository.GameRepository
}

func NewCreateGameUseCase(repo repository.GameRepository) *CreateGameUseCase {
	return &CreateGameUseCase{
		repository: repo,
	}
}

func (this *CreateGameUseCase) Execute(player1Name, player2Name string, boardSize int) (string, error) {
	player1Name = strings.TrimSpace(player1Name)
	player2Name = strings.TrimSpace(player2Name)

	if player1Name == "" ||
		player2Name == "" ||
		boardSize <= 2 ||
		boardSize >= 7 {
		return "", fmt.Errorf("Invalid parameters")
	}

	p1 := model.NewPlayer(player1Name, enum.X)
	p2 := model.NewPlayer(player2Name, enum.O)
	players := [2]model.Player{p1, p2}

	game := model.NewGame(players, boardSize)

	if err := this.repository.SaveGame(*game); err != nil {
		return "", err
	}

	return game.GetID(), nil
}