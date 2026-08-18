package model

import "tictactoe/internal/enum"

type Player struct {
	Name   string
	Symbol enum.Symbol
}

func NewPlayer(name string, symbol enum.Symbol) Player {
	return Player{
		Name:   name,
		Symbol: symbol,
	}
}

func (this Player) GetName() string {
	return this.Name
}

func (this Player) GetSymbol() enum.Symbol {
	return this.Symbol
}
