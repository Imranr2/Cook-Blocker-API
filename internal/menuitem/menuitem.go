package menuitem

import (
	"time"

	"github.com/Imanr2/Restaurant_API/internal/user"
)

type MenuItem struct {
	ID          uint         `json:"-" gorm:"primaryKey"`
	Name        string       `gorm:"index;unique;not null"`
	Description string       `gorm:"not null"`
	Price       float32      `gorm:"not null"`
	CreatedBy   uint         `gorm:"not null"`
	Ingredients []Ingredient `gorm:"many2many:menu_item_ingredients"`
	User        user.User    `gorm:"foreignKey:CreatedBy;not null"`
	CreatedAt   time.Time    `json:"-" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt   time.Time    `json:"-" gorm:"type:timestamp;default:current_timestamp ON update current_timestamp"`
}

type GetWithIDRequest struct {
	Id int `json:"id" validate:"required"`
}

type GetWithIDResponse struct {
	Item      MenuItem `json:"item"`
	ErrorCode int      `json:"errorCode"`
	Error     string   `json:"error"`
}

type GetResponse struct {
	MenuItems []MenuItem `json:"menuItems"`
	ErrorCode int        `json:"errorCode"`
	Error     string     `json:"error"`
}

type CreateRequest struct {
	UserId      uint
	Name        string       `json:"name" validate:"required"`
	Description string       `json:"desc" validate:"required"`
	Price       float32      `json:"price" validate:"required"`
	Ingredients []Ingredient `json:"ingredients" validate:"required"`
}

type CreateResponse struct {
	ErrorCode int    `json:"errorCode"`
	Error     string `json:"error"`
}

type DeleteRequest struct {
	UserId uint
	Id     int `json:"id" validate:"required"`
}

type DeleteResponse struct {
	ErrorCode int    `json:"errorCode"`
	Error     string `json:"error"`
}
