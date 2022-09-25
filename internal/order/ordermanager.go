package order

import (
	"gorm.io/gorm"
)

type OrderManager interface {
	GetOrders() (GetResponse, error)
	GetOrderWithID(GetWithIDRequest) (GetWithIDResponse, error)
	CreateOrder(CreateRequest) (CreateResponse, error)
	CompleteOrder(CompleteRequest) (CompleteResponse, error)
}

type OrderManagerImpl struct {
	database *gorm.DB
}

func NewOrderManager(database *gorm.DB) OrderManager {
	return &OrderManagerImpl{
		database: database,
	}
}

func (m *OrderManagerImpl) GetOrders() (resp GetResponse, err error) {
	return
}

func (m *OrderManagerImpl) GetOrderWithID(req GetWithIDRequest) (resp GetWithIDResponse, err error) {
	return
}

func (m *OrderManagerImpl) CreateOrder(req CreateRequest) (resp CreateResponse, err error) {
	return
}

func (m *OrderManagerImpl) CompleteOrder(req CompleteRequest) (resp CompleteResponse, err error) {
	return
}
