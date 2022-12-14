package user

import (
	"github.com/Imanr2/Restaurant_API/internal/session"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserManager interface {
	Register(RegisterRequest) RegisterResponse
	Login(LoginRequest) (session.Session, LoginResponse)
}

type UserManagerImpl struct {
	database *gorm.DB
}

func NewUserManager(database *gorm.DB) UserManager {
	return &UserManagerImpl{
		database: database,
	}
}

func (m *UserManagerImpl) Register(req RegisterRequest) (resp RegisterResponse) {
	pwHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		resp.Error = err.Error()
		resp.ErrorCode = 4
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
		resp.ErrorCode = 3
		return
	}
	return
}

func (m *UserManagerImpl) Login(req LoginRequest) (jwt session.Session, resp LoginResponse) {
	var user User
	dbc := m.database.First(&user, User{Username: req.Username})

	if dbc.Error != nil {
		resp.Error = dbc.Error.Error()
		resp.ErrorCode = 3
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		resp.Error = err.Error()
		resp.ErrorCode = 4
		return
	}
	jwt, err = session.GenerateToken(user.ID)

	if err != nil {
		resp.Error = err.Error()
		resp.ErrorCode = 4
		return
	}
	resp.User = user
	return
}
