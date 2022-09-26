package menuitem

import "time"

type Image struct {
	ID        uint      `json:"-" gorm:"primaryKey"`
	ImageUrl  string    `json:"imageUrl" validate:"required" gorm:"type:TEXT"`
	CreatedAt time.Time `json:"-" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time `json:"-" gorm:"type:timestamp;default:current_timestamp ON update current_timestamp"`
}
