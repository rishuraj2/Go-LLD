package handler

import (
	"net/http"
	"tictactoe/internal/api/dto"
	"tictactoe/internal/app/services"

	"github.com/gin-gonic/gin"
)

type GameHandler struct {
	service *services.GameService
}

func NewGameHandler(service *services.GameService) *GameHandler {
	return &GameHandler{
		service: service,
	}
}

func (this *GameHandler) CreateGame(ctx *gin.Context) {
	var req dto.CreateGameRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Status:  "bad request",
			Message: err.Error(),
		})
		return
	}

	gameID, err := this.service.CreateGame(req.Player1Name, req.Player2Name, req.BoardSize)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Status:  "bad request",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.CreateGameResponse{
		Status:  "success",
		Message: "game created successfully",
		GameId:  gameID,
	})
}

func (this *GameHandler) MakeMove(ctx *gin.Context) {
	var req dto.MakeMoveRequest

	gameID, exists := ctx.Params.Get("game_id")

	if !exists {
		ctx.JSON(http.StatusNotFound, dto.ErrorResponse{
			Status:  "not found",
			Message: "invalid game id",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Status:  "bad request",
			Message: err.Error(),
		})
		return
	}

	err := this.service.MakeMove(gameID, req.X, req.Y)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Status:  "bad request",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.MakeMoveResponse{
		Status:  "success",
		Message: "move has been accepted",
	})
}

func (this *GameHandler) GetGame(ctx *gin.Context) {
	gameID, exists := ctx.Params.Get("game_id")

	if !exists {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Status:  "bad request",
			Message: "invalid game id",
		})
		return
	}

	game, err := this.service.GetGame(gameID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.ErrorResponse{
			Status:  "not found",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.GameResponse{
		Status:  "success",
		Message: "game fetched successfully",
		Game:    game,
	})
}
