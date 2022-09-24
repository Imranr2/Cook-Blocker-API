package user

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserManager interface {
	Register(RegisterRequest) (RegisterResponse, error)
}

type UserManagerImpl struct {
	database *gorm.DB
}

func NewUserManager(database *gorm.DB) UserManager {
	return &UserManagerImpl{
		database: database,
	}
}

func (m *UserManagerImpl) Register(req RegisterRequest) (resp RegisterResponse, err error) {
	pwHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		resp.Error = err.Error()
		resp.ErrorCode = 2
	}

	newUser := &User{
		Username: req.Username,
		Password: string(pwHash),
		Name:     req.Name,
		Role:     req.Role,
	}

	dbc := m.database.Create(&newUser)

	if dbc.Error != nil {
		resp.Error = dbc.Error.Error()
		resp.ErrorCode = 1
		return resp, dbc.Error
	}
	return
}
