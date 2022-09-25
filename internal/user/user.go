package user

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"index;unique;not null"`
	Password  string    `gorm:"not null"`
	Name      string    `gorm:"not null"`
	Role      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:current_timestamp ON update current_timestamp"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,max=32"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginResponse struct {
	ErrorCode int    `json:"errorCode"`
	Error     string `json:"error"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required,max=32"`
	Password string `json:"password" validate:"required,min=8"`
	Name     string `json:"name" validate:"required"`
	Role     string `json:"role" validate:""`
}

type RegisterResponse struct {
	ErrorCode int    `json:"errorCode"`
	Error     string `json:"error"`
}
