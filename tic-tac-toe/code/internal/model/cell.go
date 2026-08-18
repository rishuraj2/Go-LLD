package model

import "tictactoe/internal/enum"

type Cell struct {
	Symbol enum.Symbol
}

func NewCell() Cell {
	return Cell{
		Symbol: enum.EMPTY,
	}
}

func (this *Cell) SetSymbol(symbol enum.Symbol) {
	this.Symbol = symbol
}

func (this Cell) GetSymbol() enum.Symbol {
	return this.Symbol
}

func (this Cell) IsEmpty() bool {
	return this.Symbol == enum.EMPTY
}
