package domain

import "time"

type Likes struct {
	Like []Like
}

type Like struct {
	ID             int       `json:"id"`
	UserID         uint      `json:"user_id"`
	PresentationID uint      `json:"presentation_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
