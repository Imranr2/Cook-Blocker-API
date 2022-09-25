package menuitem

import (
	"time"
)

type Ingredient struct {
	ID        uint       `json:"-" gorm:"primaryKey"`
	Name      string     `gorm:"index;unique;not null"`
	MenuItems []MenuItem `json:"-" gorm:"many2many:menu_item_ingredients"`
	CreatedAt time.Time  `json:"-" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time  `json:"-" gorm:"type:timestamp;default:current_timestamp ON update current_timestamp"`
}
