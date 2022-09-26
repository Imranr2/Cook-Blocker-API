package order

import (
	"time"

	"github.com/Imanr2/Restaurant_API/internal/menuitem"
	"github.com/Imanr2/Restaurant_API/internal/table"
	"github.com/Imanr2/Restaurant_API/internal/user"
)

type Order struct {
	ID          uint        `json:"-" gorm:"primaryKey"`
	TableNumber uint        `json:"tableNumber" gorm:"not null"`
	UserID      uint        `gorm:"not null"`
	Price       float64     `json:"price" gorm:"not null"`
	IsCompleted bool        `json:"isCompleted" gorm:"default:false"`
	OrderItems  []OrderItem `json:"orderItems" gorm:"foreignKey:OrderID;not null"`
	Table       table.Table `json:"-" gorm:"foreignKey:TableNumber;references:Number;not null"`
	User        user.User   `json:"-" gorm:"foreignKey:UserID;not null"`
	CreatedAt   time.Time   `json:"-" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt   time.Time   `json:"-" gorm:"type:timestamp;default:current_timestamp ON update current_timestamp"`
}

type OrderItem struct {
	ID         uint              `json:"-" gorm:"primaryKey"`
	Qty        uint              `json:"qty" validate:"required" gorm:"not null"`
	OrderID    uint              `json:"-"`
	MenuItemID uint              `json:"-" gorm:"not null"`
	MenuItem   menuitem.MenuItem `json:"menuItem" gorm:"foreignKey:MenuItemID;not null"`
	CreatedAt  time.Time         `json:"-" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt  time.Time         `json:"-" gorm:"type:timestamp;default:current_timestamp ON update current_timestamp"`
}

type GetWithIDRequest struct {
	ID string `json:"id" validate:"required"`
}

type GetWithIDResponse struct {
	Order     *Order `json:"order"`
	ErrorCode int    `json:"errorCode"`
	Error     string `json:"error"`
}

type GetResponse struct {
	Orders    []Order `json:"orders"`
	ErrorCode int     `json:"errorCode"`
	Error     string  `json:"error"`
}

type CreateRequest struct {
	UserID      uint
	TableNumber uint        `json:"tableNumber" validate:"required"`
	Price       float64     `json:"price" validate:"required"`
	OrderItems  []OrderItem `json:"orderItems" validate:"required,dive"`
}

type CreateResponse struct {
	ErrorCode int    `json:"errorCode"`
	Error     string `json:"error"`
}

type CompleteRequest struct {
	ID string `json:"id" validate:"required"`
}

type CompleteResponse struct {
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
