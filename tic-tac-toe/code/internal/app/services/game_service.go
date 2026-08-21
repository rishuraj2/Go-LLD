package services

import (
	"fmt"
	"strings"
	"tictactoe/internal/api/dto"
	"tictactoe/internal/domain/enum"
	"tictactoe/internal/domain/model"
	"tictactoe/internal/repository"
)

type GameService struct {
	repository repository.GameRepository
}

func NewGameService(repo repository.GameRepository) *GameService {
	return &GameService{
		repository: repo,
	}
}

func (this *GameService) CreateGame(player1Name, player2Name string, boardSize int) (string, error) {
	player1Name = strings.TrimSpace(player1Name)
	player2Name = strings.TrimSpace(player2Name)

	if player1Name == "" ||
		player2Name == "" ||
		boardSize <= 2 ||
		boardSize >= 7 {
		return "", fmt.Errorf("[GameService] error in creating game. Invalid parameters")
	}

	p1 := model.NewPlayer(player1Name, enum.X)
	p2 := model.NewPlayer(player2Name, enum.O)
	players := [2]model.Player{p1, p2}

	game := model.NewGame(players, boardSize)

	fmt.Println(game.GetBoard())

	if err := this.repository.SaveGame(*game); err != nil {
		return "", err
	}

	return game.GetID(), nil
}

func (this *GameService) MakeMove(gameID string, x, y int) error {
	game, err := this.repository.GetGameByID(gameID)
	if err != nil {
		return err
	}

	err = game.MakeMove(x, y)
	if err != nil {
		return err
	}

	err = this.repository.UpdateGame(game)
	if err != nil {
		return err
	}

	return nil
}

func (this *GameService) GetGame(gameID string) (dto.GameResponseData, error) {
	game, err := this.repository.GetGameByID(gameID)
	if err != nil {
		return dto.GameResponseData{}, err
	}

	players := game.GetPlayers()

	playerResponse := [2]dto.PlayerResponse{
		{
			Name:   players[0].GetName(),
			Symbol: players[0].GetSymbol().String(),
		},
		{
			Name:   players[1].GetName(),
			Symbol: players[1].GetSymbol().String(),
		},
	}

	return dto.GameResponseData{
		GameID:       game.GetID(),
		Players:      playerResponse,
		Board:        game.GetBoard(),
		ActivePlayer: players[game.GetActivePlayerIndex()].GetName(),
		GameState:    game.GetGameState().String(),
	}, nil

}
