package menuitem

import (
	"time"
)

type Ingredient struct {
	ID        uint        `gorm:"primaryKey"`
	Name      string      `gorm:"index;unique;not null"`
	MenuItems []*MenuItem `gorm:"many2many:menu_item_ingredients"`
	CreatedAt time.Time   `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time   `gorm:"type:timestamp;default:current_timestamp ON update current_timestamp"`
}
