package usecase

import "tictactoe/internal/repository"

type MakeMoveUseCase struct {
	repository repository.GameRepository
}

func NewMakeMoveUseCase(repo repository.GameRepository) *MakeMoveUseCase {
	return &MakeMoveUseCase{
		repository: repo,
	}
}

func (this *MakeMoveUseCase) Execute(gameID string, x, y int) error {
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
