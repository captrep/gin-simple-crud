package user

import (
	"time"
)

type User struct {
	Id        string    `json:"id"`
	Firstname string    `json:"first_name" binding:"required,min=3"`
	Lastname  string    `json:"last_name"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
