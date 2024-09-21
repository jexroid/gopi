package api

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UUID      uint32 `gorm:"unique"`
	Phone     int    `json:"phone" gorm:"unique"` //nolint:tagalign
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
}

type Userdetail struct {
	Firstname string
	Lastname  string
	Phone     int
}

type LoginRequest struct {
	Phone    int    `json:"phone"`
	Password string `json:"password"`
}

type ValidateRequest struct {
	Token string `json:"token"`
}

type Argonparams struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

type RegisterResponse struct {
	Ok        bool
	Token     string
	UserExist bool
}

type LoginResponse struct {
	Ok    bool
	Valid bool
	Token string
}

type ValidateResponse struct {
	Ok      bool
	Payload string
	Exp     bool
	User    Userdetail
}

type CreateCar struct {
}
