package main

import (
	"database/sql"

	"github.com/dmc0001/auth-jwt-project/internal/types"
)

type Config struct {
	addr                   string
	db                     *sql.DB
	userModel              types.UserStore
	productModel           types.ProductStore
	JwtExpiretionInSeconds int
	JwtSecret              string
}

type Application struct {
	config *Config
}
