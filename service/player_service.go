package service

import (
	"github.com/CatBloom/MahjongMasterApi/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PlayerService interface {
	GetPlayerList(id string) ([]models.Player, error)
	CreatePlayer(c *gin.Context) (models.Player, error)
	UpdatePlayer(id string, c *gin.Context) (Player, error)
	DeletePlayer(id string) error
}

type playerService struct {
	db *gorm.DB
}

func NewPlayerService(db *gorm.DB) PlayerService {
	return &playerService{db}
}

type Player = models.Player

func (s playerService) GetPlayerList(id string) ([]models.Player, error) {
	var p []Player

	if err := s.db.Debug().Table("players").Where("league_id = ?", id).Order("created_at").Find(&p).Error; err != nil {
		return nil, err
	}

	return p, nil
}

func (s playerService) CreatePlayer(c *gin.Context) (models.Player, error) {
	var p Player

	if err := c.BindJSON(&p); err != nil {
		return p, err
	}

	if err := s.db.Table("players").Create(&p).Error; err != nil {
		return p, err
	}

	return p, nil
}

func (s playerService) UpdatePlayer(id string, c *gin.Context) (Player, error) {
	var p Player

	if err := s.db.Where("id = ?", id).First(&p).Error; err != nil {
		return p, err
	}

	if err := c.BindJSON(&p); err != nil {
		return p, err
	}

	s.db.Save(&p)

	return p, nil
}

func (s playerService) DeletePlayer(id string) error {
	var p Player

	//Todo resultのplayer存在チェック

	if err := s.db.Where("id = ?", id).Delete(&p).Error; err != nil {
		return err
	}

	return nil
}
