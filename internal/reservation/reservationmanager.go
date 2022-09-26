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
	var reservations []Reservation
	err = m.database.Model(&Reservation{}).Find(&reservations).Error
	if err != nil {
		resp.Error = err.Error()
		resp.ErrorCode = 3
		return
	}
	resp.Reservations = reservations
	return
}

func (m *ReservationManagerImpl) GetReservationWithID(req GetWithIDRequest) (resp GetWithIDResponse, err error) {
	var reservation Reservation
	err = m.database.Model(&Reservation{}).First(&reservation, req.ID).Error
	if err != nil {
		resp.Error = err.Error()
		resp.ErrorCode = 3
		return
	}

	resp.Reservation = &reservation
	return
}

func (m *ReservationManagerImpl) CreateReservation(req CreateRequest) (resp CreateResponse, err error) {
	newReservation := &Reservation{
		CustomerName:  req.CustomerName,
		CustomerPhone: req.CustomerPhone,
		TableNumber:   req.TableNumber,
		IsCompleted:   false,
		Pax:           req.Pax,
	}

	err = m.database.Create(&newReservation).Error
	if err != nil {
		resp.Error = err.Error()
		resp.ErrorCode = 3
	}
	return
}

func (m *ReservationManagerImpl) FulfillReservation(req FulfillRequest) (resp FulfillResponse, err error) {
	var reservation Reservation
	err = m.database.Model(&Reservation{}).First(&reservation, req.ID).Error
	if err != nil {
		resp.Error = err.Error()
		resp.ErrorCode = 3
		return
	}

	err = m.database.Model(&reservation).Updates(Reservation{IsCompleted: true}).Error
	if err != nil {
		resp.Error = err.Error()
		resp.ErrorCode = 3
		return
	}
	return
}

func (m *ReservationManagerImpl) DeleteReservation(req DeleteRequest) (resp DeleteResponse, err error) {
	var reservation Reservation
	err = m.database.First(&reservation, req.ID).Error
	if err != nil {
		resp.Error = err.Error()
		resp.ErrorCode = 3
		return
	}

	err = m.database.Delete(&reservation, req.ID).Error
	if err != nil {
		resp.Error = err.Error()
		resp.ErrorCode = 3
		return
	}
	return
}
