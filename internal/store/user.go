package store

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/dmc0001/auth-jwt-project/internal/types"
	"github.com/dmc0001/auth-jwt-project/internal/validation"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	Db *sql.DB
}

func NewUserModel(db *sql.DB) (userModel *UserModel) {
	return &UserModel{
		Db: db,
	}
}

func (u *UserModel) GetUserByEmailWithPassword(email string) (*types.User, error) {
	user := &types.User{}
	const q = `
		SELECT id, first_name, last_name, email, phone_number, password, created_at
		FROM users
		WHERE email = ?
		LIMIT 1;
	`
	row := u.Db.QueryRow(q, email)
	err := row.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.PhoneNumber,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return user, nil
}

func (u *UserModel) GetUserByEmail(email string) (*types.GetUserByEmailResponse, error) {
	user := &types.GetUserByEmailResponse{}
	const q = `
		SELECT id, first_name, last_name, email, phone_number, created_at
		FROM users
		WHERE email = ?
		LIMIT 1;
	`
	row := u.Db.QueryRow(q, email)
	err := row.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.PhoneNumber,
		&user.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return user, nil
}

func (u *UserModel) GetUserById(id int) (*types.GetUserByEmailResponse, error) {
	user := &types.GetUserByEmailResponse{}
	const q = `
		SELECT id, first_name, last_name, email, phone_number, created_at
		FROM users
		WHERE id = ?
		LIMIT 1;
	`
	row := u.Db.QueryRow(q, id)
	err := row.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.PhoneNumber,
		&user.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return user, nil
}

func (u *UserModel) RegisterUser(payload types.RegisterUserRequest) error {
	_, err := u.GetUserByEmail(payload.Email)
	if err == nil {
		return fmt.Errorf("User already exists")
	}

	validUser, err := validation.ValidateRegisterUser(&payload)
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(validUser.Password), 12)
	if err != nil {
		return err
	}
	q := "INSERT INTO users (first_name, last_name, email, phone_number, password,created_at) VALUES (?, ?, ?, ?, ?, ?)"

	u.Db.Exec(q,
		validUser.FirstName,
		validUser.LastName,
		validUser.Email,
		validUser.PhoneNumber,
		hashedPassword,
		validUser.CreatedAt,
	)

	return nil
}

func (u *UserModel) LoginUser(payload types.LoginUserRequest) (*types.LoginUserResponse, error) {

	_, err := validation.ValidateLoginUser(payload)
	if err != nil {
		return nil, err
	}
	user, err := u.GetUserByEmailWithPassword(payload.Email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(payload.Password))
	if err != nil {
		return nil, fmt.Errorf("Invalid email or password")
	}

	return &types.LoginUserResponse{
		Id:          user.Id,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		AccessToken: "",
		CreatedAt:   user.CreatedAt,
	}, nil
}
