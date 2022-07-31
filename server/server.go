package server

import (
	"time"

	"github.com/CatBloom/MahjongMasterApi/controllers"
	"github.com/CatBloom/MahjongMasterApi/db"
	"github.com/CatBloom/MahjongMasterApi/firebase"
	"github.com/CatBloom/MahjongMasterApi/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	league controllers.LeagueController
	player controllers.PlayerController
	game   controllers.GameController
	result controllers.ResultController
}

func NewServer(
	league controllers.LeagueController,
	player controllers.PlayerController,
	game controllers.GameController,
	result controllers.ResultController,
) Server {
	return Server{
		league: league,
		player: player,
		game:   game,
		result: result,
	}
}

func (s Server) Init() {
	r := s.router()
	r.Run()
}

func (s Server) router() *gin.Engine {
	r := gin.New()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:4200",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Accept", "Authorization", "Content-Type",
		},
		AllowCredentials: false,
		MaxAge:           24 * time.Hour,
	}))

	r.Use(firebase.APIAuthWrap(db.GetDB()))
	r.Use(logger.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		l := v1.Group("/league")
		{
			l.GET("/search/:value", s.league.Search)
			l.GET("/list/:uid", s.league.List)
			l.GET("/:id", s.league.Get)
			l.POST("", s.league.Create)
			l.PUT("/:id", s.league.Update)
			l.DELETE("/:id", s.league.Delete)
		}

		p := v1.Group("/player")
		{
			//league毎のplayerlist取得
			p.GET("/list/:lid", s.player.List)
			p.POST("", s.player.Create)
			p.PUT("/:id", s.player.Update)
			p.DELETE("/:id", s.player.Delete)
		}

		g := v1.Group("/game")
		{
			g.GET("/list/:lid", s.game.List)
			g.GET("/:id", s.game.Get)
			g.POST("", s.game.Create)
			g.PUT("/:id", s.game.Update)
			g.DELETE("/:id", s.game.Delete)
		}

		result := v1.Group("/result")
		{
			result.GET("/player/:pid", s.result.GetPlayerResults)
			result.GET("/player/pie/:pid", s.result.GetPlayerPie)
			result.GET("/player/line/:pid", s.result.GetPlayerLine)
			result.GET("/league/:lid", s.result.GetLeagueResults)
		}
	}
	return r
}
