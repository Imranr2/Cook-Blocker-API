package reservation

import (
	"time"

	"github.com/Imanr2/Restaurant_API/internal/table"
)

type Reservation struct {
	ID            uint        `json:"-" gorm:"primaryKey"`
	CustomerName  string      `json:"customerName" gorm:"not null"`
	CustomerPhone string      `json:"customerPhone" gorm:"not null"`
	TableNumber   uint        `json:"tableNumber" gorm:"not null"`
	IsCompleted   bool        `json:"isCompleted" gorm:"default:false"`
	Pax           uint        `json:"pax" gorm:"not null"`
	Table         table.Table `json:"-" gorm:"foreignKey:TableNumber;references:Number;not null"`
	CreatedAt     time.Time   `json:"-" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt     time.Time   `json:"-" gorm:"type:timestamp;default:current_timestamp ON update current_timestamp"`
}

type GetWithIDRequest struct {
	ID string `json:"id" validate:"required"`
}

type GetWithIDResponse struct {
	Reservation *Reservation `json:"reservation"`
	ErrorCode   int          `json:"errorCode"`
	Error       string       `json:"error"`
}

type GetResponse struct {
	Reservations []Reservation `json:"reservations"`
	ErrorCode    int           `json:"errorCode"`
	Error        string        `json:"error"`
}

type CreateRequest struct {
	CustomerName  string `json:"customerName" validate:"required"`
	CustomerPhone string `json:"customerPhone" validate:"required"`
	TableNumber   uint   `json:"tableNumber" validate:"required"`
	Pax           uint   `json:"pax" validate:"required"`
}

type CreateResponse struct {
	ErrorCode int    `json:"errorCode"`
	Error     string `json:"error"`
}

type FulfillRequest struct {
	ID string `json:"id" validate:"required"`
}

type FulfillResponse struct {
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
