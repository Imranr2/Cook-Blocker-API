package order

import (
	"gorm.io/gorm"
)

type OrderManager interface {
}

type OrderManagerImpl struct {
	database *gorm.DB
}

func NewOrderManager(database *gorm.DB) OrderManager {
	return &OrderManagerImpl{
		database: database,
	}
}
