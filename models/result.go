package models

import (
	"time"

	"gorm.io/gorm"
)

type Result struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	PlayerId  uint       `json:"playerId"`
	Name      string     `json:"playerName" gorm:"->;-:migration"`
	Rank      uint       `json:"rank"`
	Point     int        `json:"point"`
	CalcPoint float64    `json:"calcPoint"`
	GameID    uint       `json:"gameId"`
}

func (result *Result) BeforeCreate(tx *gorm.DB) (err error) {
	return
}
