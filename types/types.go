package types

import "github.com/golang-jwt/jwt/v5"

type Status string

const (
	StatusPending   Status = "pending"
	StatusCanceled  Status = "canceled"
	StatusCompleted Status = "completed"
	StatusProcessed Status = "processed"
)

type UserStore interface {
	CreateUser(user User) error
	GetUserByEmail(email string) (User, error)
	GetUserByID(id string) (User, error)
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
	GetAddressesByUserID(id string) ([]Address, error)
}

type LaundryStore interface {
	CreateLaundryType(laundryType LaundryType) error
	GetLaundryTypes() ([]LaundryType, error)
	GetLaundryRequestByID(id string) (LaundryRequest, error)
	CreateLaundryRequest(laundryRequest LaundryRequest) error
	GetLaundryTypeByID(id string) (LaundryType, error)
	GetLaundryRequestsByUserID(id string) ([]LaundryRequestResponse, error)
	GetLaundryRequests() ([]LaundryRequestResponse, error)
	UpdateLaundryRequestStatus(status string, rId string, uId string) error
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

type LaundryRequest struct {
	ID             string  `json:"id"`
	UserID         string  `json:"user_id"`
	AdminID        *string `json:"admin_id"`
	LaundryTypeID  string  `json:"laundry_type_id"`
	AddressID      string  `json:"address_id"`
	Weight         float64 `json:"weight"`
	Notes          string  `json:"notes"`
	Status         string  `json:"status"`
	CompletionDate string  `json:"completion_date"`
}

type StatusHistory struct {
	ID               string `json:"id"`
	LaundryRequestID string `json:"laundry_request_id"`
	Status           string `json:"status"`
	UpdatedAt        string `json:"updated_at"`
	UpdatedBy        string `json:"updated_by"`
}

type LaundryType struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	EstimatedDays int     `json:"estimated_days"`
	IsActive      bool    `json:"is_active"`
	CreatedAt     string  `json:"created_at"`
}

type LaundryRequestPayload struct {
	LaundryTypeID string  `json:"laundry_type_id" validate:"required"`
	AddressID     string  `json:"address_id" validate:"required"`
	Weight        float64 `json:"weight" validate:"required"`
	Notes         string  `json:"notes"`
}

type StatusHistoryPayload struct {
	LaundryRequestID string `json:"laundry_request_id" validate:"required"`
	Status           string `json:"status" validate:"required"`
	UpdatedBy        string `json:"updated_by" validate:"required"`
	UpdatedAt        string `json:"updated_at" validate:"required"`
}

type UserRegisterPayload struct {
	Username    string `json:"username" validate:"required"`
	FullName    string `json:"full_name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type UpdateLaundryRequestPayload struct {
	Status string `json:"status" validate:"required,oneof=pending canceled completed processed"`
}

type UserAddressPayload struct {
	StreedAddress string `json:"street_address" validate:"required"`
	City          string `json:"city" validate:"required"`
	State         string `json:"state" validate:"required"`
}

type LaundryTypePayload struct {
	Name          string  `json:"name" validate:"required"`
	Description   string  `json:"description" validate:"required"`
	Price         float64 `json:"price" validate:"required"`
	EstimatedDays int     `json:"estimated_days" validate:"required"`
}

type StatusHistoryResponse struct {
	ID               string `json:"id"`
	LaundryRequestID string `json:"laundry_request_id"`
	Status           string `json:"status"`
	UpdatedAt        string `json:"updated_at"`
	UpdatedBy        string `json:"updated_by"`
}

type LaundryRequestResponse struct {
	ID              string                  `json:"id"`
	Username        string                  `json:"username"`
	Weight          float64                 `json:"weight"`
	LaundryType     string                  `json:"laundry_type"`
	Notes           string                  `json:"notes"`
	CurrentStatus   string                  `json:"current_status"`
	CompletionDate  string                  `json:"completion_date"`
	TotalPrice      float64                 `json:"total_price"`
	StatusHistories []StatusHistoryResponse `json:"status_histories" gorm:"foreignKey:LaundryRequestID"`
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
