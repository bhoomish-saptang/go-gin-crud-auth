package constants

import "time"

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Place string `json:"place"`
	Age   int64  `json:"age"`
}

type AuthInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthUser struct {
	Username  string    `json:"username" validate:"required,min=2,max=100"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}
