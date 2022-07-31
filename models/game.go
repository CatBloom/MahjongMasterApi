package models

import (
	"math"
	"time"
)

type Game struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	LeagueId  string     `json:"leagueId"`
	Rules     Rules      `json:"rules" gorm:"-:all"`
	Players   []Player   `json:"players" gorm:"many2many:games_players"`
	Results   []Result   `json:"results" gorm:"foreignKey:GameID;-:migration"`
}

//calcPointをrulesの値で計算する関数
func (g *Game) CreateCalcPoint() error {
	pcnt := g.Rules.PlayerCount
	//ウマの配列を作成
	uma := []int{g.Rules.Uma1, g.Rules.Uma2, g.Rules.Uma3}
	if pcnt == 4 {
		uma = append(uma, g.Rules.Uma4)
	}

	for i := 0; i < pcnt; i++ {
		//点数計算
		calc := float64((g.Results[i].Point - g.Rules.ReturnPoint) + uma[i]*1000)
		//rank1位用のオカ
		topPrize := float64((g.Rules.ReturnPoint - g.Rules.StartPoint) * (g.Rules.PlayerCount))
		//誤差対策をするか？
		if g.Results[i].Rank == 1 {
			g.Results[i].CalcPoint = math.Round(((calc+topPrize)/1000)*10) / 10
		} else {
			g.Results[i].CalcPoint = math.Round((calc/1000)*10) / 10
		}
	}
	return nil
}
