package table

import "time"

type Table struct {
	Number    uint      `json:"number" gorm:"primaryKey;not null"`
	Pax       uint      `json:"pax" gorm:"not null"`
	CreatedAt time.Time `json:"-" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time `json:"-" gorm:"type:timestamp;default:current_timestamp ON update current_timestamp"`
}

type GetResponse struct {
	Tables    []Table `json:"tables"`
	ErrorCode int     `json:"errorCode"`
	Error     string  `json:"error"`
}
