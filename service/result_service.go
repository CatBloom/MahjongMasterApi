package service

import (
	"github.com/CatBloom/MahjongMasterApi/models"
	"github.com/CatBloom/MahjongMasterApi/models/response"
	"gorm.io/gorm"
)

type ResultService interface {
	GetPlayerResults(pid string) (models.Player, error)
	GetLeagueResults(id string) ([]response.LeagueResultResponce, error)
	GetPlayerAgg(id string) (response.PlayerAggResponse, error)
	GetPlayerPie(id string) ([]response.PieResponse, error)
	GetPlayerLine(id string) ([]response.LineResponse, error)
}

type resultService struct {
	db *gorm.DB
}

func NewResultService(db *gorm.DB) ResultService {
	return &resultService{db}
}

func (s resultService) GetPlayerResults(pid string) (models.Player, error) {
	var p models.Player

	if err := s.db.Debug().Preload("Games.Results", func(db *gorm.DB) *gorm.DB {
		return db.Select("*").Order("Results.rank ASC").Joins("left join players on players.id = results.player_id")
	}).Preload("Games", func(db *gorm.DB) *gorm.DB {
		return db.Order("Games.created_at DESC")
	}).Where("id = ?", pid).Find(&p).Error; err != nil {
		return p, err
	}

	return p, nil
}

func (s resultService) GetLeagueResults(id string) ([]response.LeagueResultResponce, error) {
	var results []response.LeagueResultResponce

	if err := s.db.
		Debug().
		Table("results").
		Select("ROW_NUMBER() OVER(ORDER BY SUM(results.calc_point) DESC )rank,players.id AS player_Id,players.name AS player_name ,COUNT(*) AS total_game,SUM(results.point) AS total_point ,SUM(results.calc_point) AS total_calc_point ,league_id").
		Joins("LEFT JOIN players ON players.id = results.player_id").
		Group("players.id").
		Order("total_calc_point desc").
		Where("league_id = ?", id).
		Find(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}

func (s resultService) GetPlayerAgg(id string) (response.PlayerAggResponse, error) {
	var results response.PlayerAggResponse

	if err := s.db.
		Debug().
		Table("results").
		Select("players.id AS player_Id,players.name AS player_name ,COUNT(*) AS total_game,SUM(results.point) AS total_point ,SUM(results.calc_point) AS total_calc_point,Round(avg(results.rank),2) AS average_rank").
		Joins("LEFT JOIN players ON players.id = results.player_id").
		Group("players.id").
		Where("player_id = ?", id).
		Find(&results).Error; err != nil {
		return results, err
	}

	return results, nil
}

func (s resultService) GetPlayerPie(id string) ([]response.PieResponse, error) {
	var pie []response.PieResponse

	if err := s.db.
		Debug().
		Table("results").
		Select("rank, COUNT(rank) AS count_rank ,player_id").
		Group("rank,player_id").
		Where("player_id = ?", id).
		Find(&pie).Error; err != nil {
		return pie, err
	}

	return pie, nil
}

func (s resultService) GetPlayerLine(id string) ([]response.LineResponse, error) {
	var line []response.LineResponse

	if err := s.db.
		Debug().
		Table("results").
		Limit(10).
		Select("created_at,rank").
		Where("player_id = ?", id).
		Order("created_at DESC").
		Find(&line).Error; err != nil {
		return line, err
	}

	return line, nil
}
