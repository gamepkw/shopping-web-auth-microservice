package model

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterMysql struct {
	Username       string `gorm:"username"`
	HashedPassword string `gorm:"hashed_password"`
}

type RegisterResponse struct {
	Code    string `json:"code"`
	Message string `json:"message,omitempty"`
}

type LoginRequest struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	HashedPassword string
}

type LoginResponse struct {
	Token string `json:"token"`
}
