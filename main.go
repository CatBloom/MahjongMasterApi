package main

import (
	"github.com/CatBloom/MahjongMasterApi/controllers"
	"github.com/CatBloom/MahjongMasterApi/db"
	"github.com/CatBloom/MahjongMasterApi/firebase"
	"github.com/CatBloom/MahjongMasterApi/logger"
	"github.com/CatBloom/MahjongMasterApi/server"
	"github.com/CatBloom/MahjongMasterApi/service"
)

func main() {
	db.Init()
	defer db.Close()

	logger.LoggerInit()
	defer logger.LoggerClose()

	firebase.Init()

	leagueService := service.NewLeagueService(db.GetDB())
	leagueController := controllers.NewLeagueController(leagueService)

	playerService := service.NewPlayerService(db.GetDB())
	playerController := controllers.NewPlayerController(playerService)

	gameService := service.NewGameService(db.GetDB())
	gameController := controllers.NewGameController(gameService)

	resultService := service.NewResultService(db.GetDB())
	resultController := controllers.NewResultController(resultService)

	serve := server.NewServer(leagueController, playerController, gameController, resultController)
	serve.Init()
}
