package web

import "time"

type UserUpdateRequest struct {
	Id        string    `json:"id"`
	Firstname string    `json:"first_name" binding:"required,min=3"`
	Lastname  string    `json:"last_name"`
	Email     string    `json:"email" binding:"required"`
	UpdatedAt time.Time `json:"updated_at"`
}
