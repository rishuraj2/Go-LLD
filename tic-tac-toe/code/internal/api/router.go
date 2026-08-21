package api

import (
	"tictactoe/internal/api/handler"
	"tictactoe/internal/app/services"
	"tictactoe/internal/repository"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine  *gin.Engine
	handler *handler.GameHandler
}

func NewRouter(engine *gin.Engine, handler *handler.GameHandler) *Router {
	return &Router{
		engine:  engine,
		handler: handler,
	}
}

func SetupRouter() *gin.Engine {
	repo := repository.NewInMemoryStore()
	service := services.NewGameService(repo)
	gameHandler := handler.NewGameHandler(service)
	engine := gin.Default()
	router := NewRouter(engine, gameHandler)
	router.route()

	return engine
}

func (this *Router) route() {
	this.engine.POST("/games", this.handler.CreateGame)
	this.engine.POST("/games/:game_id/moves", this.handler.MakeMove)
	this.engine.GET("/games/:game_id", this.handler.GetGame)
}
