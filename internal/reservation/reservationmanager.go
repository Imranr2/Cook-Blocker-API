package reservation

import "gorm.io/gorm"

type ReservationManager interface {
	GetReservations() (GetResponse, error)
	GetReservationWithID(GetWithIDRequest) (GetWithIDResponse, error)
	CreateReservation(CreateRequest) (CreateResponse, error)
	FulfillReservation(FulfillRequest) (FulfillResponse, error)
	DeleteReservation(DeleteRequest) (DeleteResponse, error)
}

type ReservationManagerImpl struct {
	database *gorm.DB
}

func NewReservationManager(database *gorm.DB) ReservationManager {
	return &ReservationManagerImpl{
		database: database,
	}
}

func (m *ReservationManagerImpl) GetReservations() (resp GetResponse, err error) {
	return
}

func (m *ReservationManagerImpl) GetReservationWithID(GetWithIDRequest) (resp GetWithIDResponse, err error) {
	return
}

func (m *ReservationManagerImpl) CreateReservation(CreateRequest) (resp CreateResponse, err error) {
	return
}

func (m *ReservationManagerImpl) FulfillReservation(req FulfillRequest) (resp FulfillResponse, err error) {
	return
}

func (m *ReservationManagerImpl) DeleteReservation(req DeleteRequest) (resp DeleteResponse, err error) {
	return
}
