package user

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,max=32"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginResponse struct {
	ErrorCode int    `json:"errorCode"`
	Error     string `json:"error"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required,max=32"`
	Password string `json:"password" validate:"required,min=8"`
	Name     string `json:"name" validate:"required"`
	Role     string `json:"role" validate:""`
}

type RegisterResponse struct {
	ErrorCode int    `json:"errorCode"`
	Error     string `json:"error"`
}

func (user *User) TableName() string {
	return "user_tab"
}
