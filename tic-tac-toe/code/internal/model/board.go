package model

import (
	"fmt"
	"tictactoe/internal/enum"
)

type Board struct {
	Size      int
	Grid      [][]Cell
}

func NewBoard(size int) *Board {
	grid := make([][]Cell, size)
	for i := 0; i < size; i++ {
		grid[i] = make([]Cell, size)
	}

	return &Board{
		Size:      size,
		Grid:      grid,
	}
}

func (this *Board) SetSymbol(x, y int, symbol enum.Symbol) error {
	if this.Grid[x][y].IsEmpty() {
		this.Grid[x][y].SetSymbol(symbol)
		return nil
	}

	return fmt.Errorf("[Error] illegal move")
}

func (this Board) GetSymbol(x, y int) enum.Symbol {
	return this.Grid[x][y].GetSymbol()
}

func (this Board) CalculateGameState(symbol enum.Symbol) enum.GameState {
	filledCellCount := 0
	counterMainDiagonal := 0
	counterReverseDiagonal := 0

	for i := 0; i < this.Size; i++ {
		counterRow := 0
		counterColumn := 0

		for j := 0; j < this.Size; j++ {
			if this.Grid[i][j].GetSymbol() != enum.EMPTY {
				filledCellCount++
			}

			if this.Grid[i][j].GetSymbol() == symbol {
				counterRow++
			}

			if this.Grid[j][i].GetSymbol() == symbol {
				counterColumn++
			}
		}

		if this.Grid[i][i].GetSymbol() == symbol {
			counterMainDiagonal++
		}

		if this.Grid[i][this.Size-1-i].GetSymbol() == symbol {
			counterReverseDiagonal++
		}

		if counterRow == this.Size ||
			counterColumn == this.Size ||
			counterMainDiagonal == this.Size ||
			counterReverseDiagonal == this.Size {
			if symbol == enum.O {
				return enum.WINNER_O
			} else {
				return enum.WINNER_X
			}
		}
	}

	if filledCellCount == this.Size*this.Size {
		return enum.DRAW
	}

	return enum.IN_PROGRESS
}
