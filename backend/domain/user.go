package domain

import "time"

type Users struct {
	User User `json:"users"`
}

type User struct {
	ID            int            `json:"id"`
	Name          string         `json:"name"`
	Email         string         `json:"email" gorm:"unique"`
	CreadtedAt    time.Time      `json:"created_at"`
	Image         string         `json:"image"`
	Password      []byte         `json:"password"`
	Presentations []Presentation `gorm:"foreignKey:UserID"`
	// Likes         []Like         `gorm:"foreignKey:UserID"`
}

type UserForGet struct {
	ID    int    `json:"id"`
	Name  string `json:"Name"`
	Email string `json:"email"`
}

func (u *User) GetUser() UserForGet {
	user := UserForGet{}
	user.ID = u.ID
	user.Name = u.Name
	user.Email = u.Email

	return user
}

type UsersForGet struct {
	User User `json:"users"`
}

func (u *User) GetUsers() UsersForGet {
	users := UsersForGet{}
	return users
}
