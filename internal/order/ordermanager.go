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
	var orders []Order
	err = m.database.Model(&Order{}).Preload("OrderItems.MenuItem").Find(&orders).Error
	if err != nil {
		resp.Error = err.Error()
		resp.ErrorCode = 3
		return
	}
	resp.Orders = orders
	return
}

func (m *OrderManagerImpl) GetOrderWithID(req GetWithIDRequest) (resp GetWithIDResponse, err error) {
	var order Order
	err = m.database.Model(&Order{}).Preload("OrderItems.MenuItem").First(&order, req.ID).Error
	if err != nil {
		resp.Error = err.Error()
		resp.ErrorCode = 3
		return
	}

	resp.Order = &order
	return
}

func (m *OrderManagerImpl) CreateOrder(req CreateRequest) (resp CreateResponse, err error) {
	newOrder := &Order{
		TableID:     req.TableID,
		UserID:      req.UserID,
		Price:       req.Price,
		IsCompleted: false,
		OrderItems:  req.OrderItems,
	}

	err = m.database.Create(&newOrder).Error
	if err != nil {
		resp.Error = err.Error()
		resp.ErrorCode = 3
	}
	return
}

func (m *OrderManagerImpl) CompleteOrder(req CompleteRequest) (resp CompleteResponse, err error) {
	return
}
