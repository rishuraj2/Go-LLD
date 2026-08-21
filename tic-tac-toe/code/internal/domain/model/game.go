package model

import (
	"errors"
	"tictactoe/internal/domain/enum"

	"github.com/google/uuid"
)

type Game struct {
	id                string
	players           [2]Player
	activePlayerIndex int
	board             *Board
	gameState         enum.GameState
}

var (
	errorGameEnded = errors.New("game has ended")
)

func NewGame(players [2]Player, boardSize int) *Game {
	return &Game{
		id:                uuid.NewString(),
		players:           players,
		activePlayerIndex: 0,
		board:             NewBoard(boardSize),
		gameState:         enum.IN_PROGRESS,
	}
}

func (this *Game) GetID() string {
	return this.id
}

func (this *Game) GetPlayers() [2]Player {
	return this.players
}

func (this *Game) GetActivePlayerIndex() int {
	return this.activePlayerIndex
}

func (this *Game) GetBoard() [][]string {
	return this.board.GetBoard()
}

func (this *Game) GetGameState() enum.GameState {
	return this.gameState
}

func (this *Game) MakeMove(x, y int) error {
	if this.gameState != enum.IN_PROGRESS {
		return errorGameEnded
	}

	currentPlayerSymbol := this.players[this.activePlayerIndex].GetSymbol()
	if err := this.board.SetSymbol(x, y, currentPlayerSymbol); err != nil {
		return err
	}

	this.gameState = this.board.CalculateGameState(currentPlayerSymbol)

	if this.gameState == enum.IN_PROGRESS {
		this.switchActivePlayer()
	}

	return nil
}

func (this *Game) switchActivePlayer() {
	this.activePlayerIndex = (this.activePlayerIndex + 1) % len(this.players)
}
