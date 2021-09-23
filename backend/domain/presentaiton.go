package domain

import (
	"time"
)

type Presentations []Presentation

type Presentation struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	UserID      string    `json:"user_id"`
	Discription string    `json:"discription"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	DeltedAt    time.Time `json:"deleted_at"`
	User        User      `json:"user"`
	Likes       []Like    `json:"likes" gorm:"foreignKey:PresentationID"`
}

type PresentationForGet struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Discription string `json:"discription"`
}

func (p *Presentation) BuildForGet() PresentationForGet {
	presentation := PresentationForGet{}
	presentation.ID = p.ID
	presentation.Title = p.Title
	presentation.Discription = p.Discription

	return presentation
}
