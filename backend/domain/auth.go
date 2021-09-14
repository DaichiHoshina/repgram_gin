package domain

import "time"

type Auth struct {
	ID            int            `json:"id"`
	Name          string         `json:"name"`
	Email         string         `json:"email" gorm:"unique"`
	CreadtedAt    time.Time      `json:"created_at"`
	Image         string         `json:"image"`
	Password      []byte         `json:"password"`
	Presentations []Presentation `gorm:"foreignKey:UserID"`
	Likes         []Like         `gorm:"foreignKey:UserID"`
}

type AuthForGet struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *Auth) BuildForGet() AuthForGet {
	user := AuthForGet{}
	user.ID = u.ID
	user.Name = u.Name
	user.Email = u.Email
	return user
}
