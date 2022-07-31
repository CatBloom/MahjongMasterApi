package service

import (
	"github.com/CatBloom/MahjongMasterApi/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GameService interface {
	GetGameList(lid string) ([]models.Game, error)
	GetGame(id string) (models.Game, error)
	CreateGame(c *gin.Context) (models.Game, error)
	UpdateGame(id string, c *gin.Context) (models.Game, error)
	DeleteGame(id string) error
}

type gameService struct {
	db *gorm.DB
}

func NewGameService(db *gorm.DB) GameService {
	return &gameService{db}
}

type Game = models.Game

func (s gameService) GetGameList(lid string) ([]models.Game, error) {
	var g []Game

	if err := s.db.Debug().Preload("Results", func(db *gorm.DB) *gorm.DB {
		return db.Select("*").Order("Results.rank ASC").Joins("left join players on players.id = results.player_id")
	}).Where("league_id = ?", lid).Order("created_at desc").Find(&g).Error; err != nil {
		return g, err
	}

	return g, nil
}

func (s gameService) GetGame(id string) (models.Game, error) {
	var g Game

	if err := s.db.Debug().Preload("Results").Find(&g, "id = ?", id).Error; err != nil {
		return g, err
	}

	return g, nil
}

func (s gameService) CreateGame(c *gin.Context) (models.Game, error) {
	var g Game

	if err := c.BindJSON(&g); err != nil {
		return g, err
	}

	if err := g.CreateCalcPoint(); err != nil {
		return g, err
	}

	if err := s.db.Debug().Omit("Players.*", "Rules").Create(&g).Error; err != nil {
		return g, err
	}

	// 作成時playerNameを取得する
	// player := game.Players
	// if err := s.db.Debug().Preload("Players").Find(&game).Error; err != nil {
	// 	return game, err
	// }

	return g, nil
}

func (s gameService) UpdateGame(id string, c *gin.Context) (models.Game, error) {
	var g Game

	if err := c.BindJSON(&g); err != nil {
		return g, err
	}

	s.db.Debug().Unscoped().Table("games_players").Where("game_id = ?", id).Delete("")

	if err := s.db.Debug().Session(&gorm.Session{FullSaveAssociations: true}).Omit("Players.*", "Rules").Updates(&g).Error; err != nil {
		return g, err
	}
	return g, nil
}

func (s gameService) DeleteGame(id string) error {
	// var g Game

	// if err := s.db.Debug().Model(&g).Association("Results").Clear(); err != nil {
	// 	return err
	// }

	return nil
}
