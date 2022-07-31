package models

import (
	"time"

	"gorm.io/gorm"
)

type Rules struct {
	ID          uint       `json:"id" gorm:"primary_key"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt"`
	LeagueId    string     `json:"league_id" gorm:"foreignkey"`
	GameType    string     `json:"gameType"`
	GameName    string     `json:"gameName"`
	PlayerCount int        `json:"playerCount"`
	Dora        string     `json:"dora"`
	Tanyao      string     `json:"tanyao"`
	StartPoint  int        `json:"startPoint"`
	FinishPoint int        `json:"finishPoint"`
	ReturnPoint int        `json:"returnPoint"`
	CalledPoint int        `json:"calledPoint"`
	ReachPoint  int        `json:"reachPoint"`
	Deposit     int        `json:"deposit"`
	Penalty1    int        `json:"penalty1"`
	Penalty2    int        `json:"penalty2"`
	Penalty3    int        `json:"penalty3"`
	Uma1        int        `json:"uma1"`
	Uma2        int        `json:"uma2"`
	Uma3        int        `json:"uma3"`
	Uma4        int        `json:"uma4"`
}

func (rules *Rules) BeforeCreate(tx *gorm.DB) (err error) {
	// if rules.PlayerCount == 3 {
	// 	rules.Uma4 = nil
	// }
	return
}
