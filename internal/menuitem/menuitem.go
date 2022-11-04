package menuitem

import (
	"time"

	"github.com/Imanr2/Restaurant_API/internal/user"
)

type MenuItem struct {
	ID          uint          `json:"id" gorm:"primaryKey"`
	Name        string        `json:"name" gorm:"index;unique;not null"`
	Description string        `json:"desc" gorm:"not null"`
	Steps       string        `json:"steps" gorm:"not null"`
	Price       float32       `json:"price" gorm:"not null"`
	CreatedBy   uint          `json:"createdBy" gorm:"not null"`
	ImageID     uint          `json:"-" gorm:"not null"`
	Ingredients []*Ingredient `json:"ingredients" gorm:"many2many:menu_item_ingredients"`
	Image       Image         `json:"image" gorm:"not null"`
	User        user.User     `json:"-" gorm:"foreignKey:CreatedBy;not null"`
	CreatedAt   time.Time     `json:"-" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt   time.Time     `json:"-" gorm:"type:timestamp;default:current_timestamp ON update current_timestamp"`
}

type GetWithIDRequest struct {
	ID string `json:"id" validate:"required"`
}

type GetWithIDResponse struct {
	MenuItem  *MenuItem `json:"menuItem"`
	ErrorCode int       `json:"errorCode"`
	Error     string    `json:"error"`
}

type GetResponse struct {
	MenuItems []MenuItem `json:"menuItems"`
	ErrorCode int        `json:"errorCode"`
	Error     string     `json:"error"`
}

type CreateRequest struct {
	UserID      uint
	Name        string       `json:"name" validate:"required"`
	Description string       `json:"desc" validate:"required"`
	Price       float32      `json:"price" validate:"required"`
	Steps       string       `json:"steps" validate:"required"`
	Image       Image        `json:"image" validate:"required,dive"`
	Ingredients []Ingredient `json:"ingredients" validate:"required,dive"`
}

type CreateResponse struct {
	ErrorCode int    `json:"errorCode"`
	Error     string `json:"error"`
}

type SaveImageResponse struct {
	ImageUrl string `json:"imageUrl"`
	ErrorCode int    `json:"errorCode"`
	Error     string `json:"error"`
}

type DeleteRequest struct {
	ID string `json:"id" validate:"required"`
}

type DeleteResponse struct {
	ErrorCode int    `json:"errorCode"`
	Error     string `json:"error"`
}
