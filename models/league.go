package models

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type League struct {
	ID        string          `json:"id" gorm:"primary_key"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	DeletedAt *time.Time      `json:"deletedAt,omitempty"`
	Name      string          `json:"name"`
	Manual    string          `json:"manual"`
	StartAt   string          `json:"startAt"`
	FinishAt  string          `json:"finishAt"`
	Rules     *Rules          `json:"rules,omitempty"`
	UID       string          `json:"uid" gorm:"-:migration;-"`
	UIDS      []AdminsLeagues `json:"uids"`
}

type AdminsLeagues struct {
	LeagueID string `gorm:"primaryKey;autoIncrement:false"`
	UID      string `gorm:"primaryKey;autoIncrement:false"`
}

func (l *League) BeforeCreate(tx *gorm.DB) (err error) {
	// uuidをstringに変換し‐を除去
	uid := uuid.New()
	l.ID = strings.Replace(uid.String(), "-", "", -1)
	l.Name = strings.TrimSpace(l.Name)
	l.Manual = strings.TrimSpace(l.Manual)
	return
}
