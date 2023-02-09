package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/haikalvidya/football-league/internal/delivery/payload"
	"gorm.io/gorm"
)

type LeagueModel struct {
	ID        string         `db:"id"`
	Clubname  string         `db:"clubname"`
	Points    int            `db:"points"`
	Played    int            `db:"played"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt *time.Time     `db:"updated_at"`
	DeletedAt gorm.DeletedAt `db:"deleted_at"`
}

func (LeagueModel) TableName() string {
	return "league"
}

func (a *LeagueModel) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New().String()
	a.CreatedAt = time.Now()
	return
}

func (a *LeagueModel) PublicInfo() *payload.LeagueInfo {
	res := &payload.LeagueInfo{
		ClubName: a.Clubname,
		Points:   a.Points,
		Played:   a.Played,
	}

	return res
}
