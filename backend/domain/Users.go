package domain

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string    `json:"email" gorm:"unique"`
	// CedAt timeme `json:"created_at"`
	// Ima   str    `json:"image"`
	// Password  []byte    `json:"password"`
	// Presentations []Presentation `gorm:"foreignKey:UserID"`
	// Likes         []Like         `gorm:"foreignKey:UserID"`
}

type UserForGet struct {
	ID    int    `json:"id"`
	Name  string `json:"Name"`
	Email string `json:"email"`
}

func (u *User) BuildForGet() UserForGet {
	user := UserForGet{}
	user.ID = u.ID
	user.Name = u.Name
	user.Email = u.Email

	return user
}
