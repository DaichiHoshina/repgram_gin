package domain

import (
	"time"

	"gorm.io/gorm"
)

type Presentations struct {
	Presentation Presentation `json:"users"`
}

type Presentation struct {
	gorm.Model
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	UserID      string    `json:"user_id"`
	Discription string    `json:"discription"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	User        User      `json:"user"`
	// Likes       []Like    `json:"likes" gorm:"foreignKey:PresentationID"`
}

type PresentationForGet struct {
	ID    int    `json:"id"`
	Name  string `json:"Name"`
	Email string `json:"email"`
}

func (u *Presentation) BuildForGet() PresentationForGet {
	presentaion := PresentationForGet{}
	presentaion.ID = u.ID
	// presentaion.Name = u.Name
	// presentaion.Email = u.Email

	return presentaion
}
