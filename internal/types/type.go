package types

import (
	"time"
)

type RegisterUserRequest struct {
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Email           string    `json:"email"`
	PhoneNumber     string    `json:"phone_number"`
	Password        string    `json:"password"`
	ConfirmPassword string    `json:"confirm_password"`
	CreatedAt       time.Time `json:"created_at"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	Id          int       `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	AccessToken string    `json:"access_token"`
	CreatedAt   time.Time `json:"created_at"`
}

type User struct {
	Id          int       `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Password    []byte    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
}
type Product struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
}
type CreateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

type GetUserByEmailResponse struct {
	Id          int       `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
}

type UserStore interface {
	GetUserByEmailWithPassword(email string) (*User, error)
	GetUserByEmail(email string) (*GetUserByEmailResponse, error)
	GetUserById(id int) (*GetUserByEmailResponse, error)
	RegisterUser(payload RegisterUserRequest) error
	LoginUser(payload LoginUserRequest) (*LoginUserResponse, error)
}

type ProductStore interface {
	GetProductById(id int) (*Product, error)
	GetProductByName(name string) ([]Product, error)
	GetProducts() ([]Product, error)
	CreateProduct(payload CreateProductRequest) error
}
