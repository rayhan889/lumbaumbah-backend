package types

import "github.com/golang-jwt/jwt/v5"

type UserStore interface {
	CreateUser(user User) error
	GetUserByEmail(email string) (User, error)
}

type JWTClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.MapClaims
}

type AdminStore interface {
	CreateAdmin(admin Admin) error
	GetAdminByID(id string) (Admin, error)
	GetAdminByEmail(email string) (Admin, error)
}

type AddressStore interface {
	CreateAddress(address Address) error
	GetAddressByUserID(id string) (Address, error)
}

type User struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
}

type Admin struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

type Address struct {
	ID            string `json:"id"`
	UserID        string `json:"user_id"`
	StreetAddress string `json:"street_address"`
	City          string `json:"city"`
	State         string `json:"state"`
	IsDefault     bool   `json:"is_default"`
	CreatedAt     string `json:"created_at"`
}

type UserRegisterPayload struct {
	Username    string `json:"username" validate:"required"`
	FullName    string `json:"full_name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type UserAddressPayload struct {
	StreedAddress string `json:"street_address" validate:"required"`
	City          string `json:"city" validate:"required"`
	State         string `json:"state" validate:"required"`
	IsDefault     bool   `json:"is_default" validate:"required"`
}

type AdminRegisterPayload struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SigninPayload struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}