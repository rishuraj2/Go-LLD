package model

import "tictactoe/internal/domain/enum"

type Cell struct {
	symbol enum.Symbol
}

func (this *Cell) SetSymbol(symbol enum.Symbol) {
	this.symbol = symbol
}

func (this Cell) GetSymbol() enum.Symbol {
	return this.symbol
}

func (this Cell) IsEmpty() bool {
	return this.symbol == enum.EMPTY
}
