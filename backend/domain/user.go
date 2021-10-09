package domain

import "time"

type User struct {
	ID            int            `json:"id"`
	Name          string         `json:"name"`
	Email         string         `json:"email" gorm:"unique"`
	Image         string         `json:"image"`
	Password      []byte         `json:"password"`
	Presentations []Presentation `gorm:"foreignKey:UserID"`
	Likes         []Like         `gorm:"foreignKey:UserID"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

type UserForGet struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *User) BuildForGet() UserForGet {
	user := UserForGet{}
	user.ID = u.ID
	user.Name = u.Name
	user.Email = u.Email
	return user
}
