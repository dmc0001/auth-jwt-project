package main

import (
	"database/sql"

	"github.com/dmc0001/auth-jwt-project/internal/types"
)

type Config struct {
	addr      string
	db        *sql.DB
	userModel types.UserStore
}

type Application struct {
	config *Config
}
