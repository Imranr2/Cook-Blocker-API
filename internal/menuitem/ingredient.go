package menuitem

import (
	"time"
)

type Ingredient struct {
	ID        uint        `json:"-" gorm:"primaryKey"`
	Description string `json:"description" validate:"required" gorm:"not null"`
	MenuItemID uint `json:"-" gorm:"not null"`
	CreatedAt time.Time   `json:"-" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time   `json:"-" gorm:"type:timestamp;default:current_timestamp ON update current_timestamp"`
}
