package main

import (
	"tictactoe/internal/api"
)

func main() {
	router := api.SetupRouter()
	router.Run(":8080")
}
