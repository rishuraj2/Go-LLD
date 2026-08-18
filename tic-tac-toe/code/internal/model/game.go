package model

import (
	"fmt"
	"tictactoe/internal/enum"
)

type Game struct {
	Players      []Player
	ActivePlayer *Player
	Board        *Board
	GameState    enum.GameState
}

func NewGame(players []Player, boardSize int) *Game {
	return &Game{
		Players: players,
		ActivePlayer: &players[0],
		Board: NewBoard(boardSize),
		GameState: enum.IN_PROGRESS,
	}
}

func (this *Game) MakeMove(x, y int) error {
	currentPlayerSymbol := this.ActivePlayer.GetSymbol()
	err := this.Board.SetSymbol(x, y, currentPlayerSymbol)
	if err != nil {
		return err
	}

	this.GameState = this.Board.CalculateGameState(currentPlayerSymbol)
	fmt.Printf("[Game Status] %s| Last move: %d, %d\n", this.GameState.String(), x, y)
	return nil
}
