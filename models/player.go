package models

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

type Player struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	Name      string     `json:"name"`
	LeagueId  string     `json:"leagueId"`
	Games     []Game     `json:"games" gorm:"many2many:games_players"`
	// Name      string     `json:"name" gorm:"index:,unique,composite:myname"`
	// LeagueId  uuid.UUID  `json:"leagueId" gorm:"index:,unique,composite:myname"`
}

func (p *Player) BeforeCreate(tx *gorm.DB) (err error) {
	p.Name = strings.TrimSpace(p.Name)
	return
}

func (p *Player) BeforeUpdate(tx *gorm.DB) (err error) {
	p.Name = strings.TrimSpace(p.Name)
	return
}
