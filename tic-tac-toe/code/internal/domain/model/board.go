package model

import (
	"errors"
	"tictactoe/internal/domain/enum"
)

type Board struct {
	size int
	grid [][]Cell
}

var (
	errorIllegalMove = errors.New("illegal move")
	errorInvalidCell = errors.New("invalid cell")
)

func NewBoard(size int) *Board {
	grid := make([][]Cell, size)

	for i := 0; i < size; i++ {
		grid[i] = make([]Cell, size)
	}

	return &Board{
		size: size,
		grid: grid,
	}
}

func (this *Board) SetSymbol(x, y int, symbol enum.Symbol) error {

	if x >= this.size || y >= this.size || x < 0 || y < 0 {
		return errorInvalidCell
	}

	if !this.grid[x][y].IsEmpty() {
		return errorIllegalMove
	}

	this.grid[x][y].SetSymbol(symbol)
	return nil
}

func (this Board) GetSymbol(x, y int) (enum.Symbol, error) {
	if x < this.size && y < this.size && x >= 0 && y >= 0 {
		return this.grid[x][y].GetSymbol(), nil
	}

	return enum.EMPTY, errorInvalidCell
}

func (this Board) GetBoard() [][]string {
	res := make([][]string, this.size)
	for i := 0; i < this.size; i++ {
		res[i] = make([]string, this.size)
	}

	for i := 0; i < this.size; i++ {
		for j := 0; j < this.size; j++ {
			res[i][j] = this.grid[i][j].GetSymbol().String()
		}
	}

	return res
}

func (this Board) CalculateGameState(symbol enum.Symbol) enum.GameState {
	filledCellCount := 0
	counterMainDiagonal := 0
	counterReverseDiagonal := 0

	for i := 0; i < this.size; i++ {
		counterRow := 0
		counterColumn := 0

		for j := 0; j < this.size; j++ {
			if this.grid[i][j].GetSymbol() != enum.EMPTY {
				filledCellCount++
			}

			if this.grid[i][j].GetSymbol() == symbol {
				counterRow++
			}

			if this.grid[j][i].GetSymbol() == symbol {
				counterColumn++
			}
		}

		if this.grid[i][i].GetSymbol() == symbol {
			counterMainDiagonal++
		}

		if this.grid[i][this.size-1-i].GetSymbol() == symbol {
			counterReverseDiagonal++
		}

		if counterRow == this.size ||
			counterColumn == this.size ||
			counterMainDiagonal == this.size ||
			counterReverseDiagonal == this.size {
			if symbol == enum.O {
				return enum.WINNER_O
			} else {
				return enum.WINNER_X
			}
		}
	}

	if filledCellCount == this.size*this.size {
		return enum.DRAW
	}

	return enum.IN_PROGRESS
}
