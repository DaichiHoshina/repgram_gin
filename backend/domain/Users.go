package domain

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
    gorm.Model
    ID            int            `json:"id"`
	Name          string         `json:"name"`
	Email         string         `json:"email" gorm:"unique"`
	CreatedAt     time.Time      `json:"created_at"`
	Image         string         `json:"image"`
	Password      []byte         `json:"password"`
	// Presentations []Presentation `gorm:"foreignKey:UserID"`
	// Likes         []Like         `gorm:"foreignKey:UserID"`
}

type UsersForGet struct {
    ID int `json:"id"`
    Name string `json:"screenName"`
    Email *string `json:"email"`
}

func (u *Users) BuildForGet() UsersForGet {
    user := UsersForGet{}
    user.ID = u.ID
    user.Name = u.Name
    if u.Email != nil {
        user.Email = u.Email
    } else {
        empty := ""
        user.Email = &empty
    }
    return user
}
