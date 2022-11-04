package menuitem

import (
	"time"
)

type Ingredient struct {
	ID        uint        `json:"-" gorm:"primaryKey"`
	Description string `json:"description" validate:"required" gorm:"not null"`
	// Name      string      `json:"name" validate:"required" gorm:"not null"`
	// Qty       uint        `json:"qty" validate:"required" gorm:"not null"`
	// Unit      string      `json:"unit" validate:"required" gorm:"not null"`
	MenuItems []*MenuItem `json:"-" gorm:"many2many:menu_item_ingredients"`
	CreatedAt time.Time   `json:"-" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time   `json:"-" gorm:"type:timestamp;default:current_timestamp ON update current_timestamp"`
}
