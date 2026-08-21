package model

import "tictactoe/internal/domain/enum"

type Player struct {
	name   string
	symbol enum.Symbol
}

func NewPlayer(name string, symbol enum.Symbol) Player {
	return Player{
		name:   name,
		symbol: symbol,
	}
}

func (this Player) GetName() string {
	return this.name
}

func (this Player) GetSymbol() enum.Symbol {
	return this.symbol
}
