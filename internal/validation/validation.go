package validation

import (
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/dmc0001/auth-jwt-project/internal/types"
)

var emailRx = regexp.MustCompile(`^[A-Za-z0-9._%+\-']+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,}$`)

// var passwordRx = regexp.MustCompile(`^[A-Za-z\d@$!%*#?&]{8,}$`)
// var passwordRx = regexp.MustCompile(`^(?=.*[A-Za-z])(?=.*\d)(?=.*[@$!%*#?&])[A-Za-z\d@$!%*#?&]{8,}$`)
var phoneNumberRx = regexp.MustCompile(`^\s*(?:\+?(\d{1,3}))?[-. (]*(\d{3})[-. )]*(\d{3})[-. ]*(\d{4})(?: *x(\d+))?\s*$`)

func ValidateRegisterUser(req *types.RegisterUserRequest) (*types.RegisterUserRequest, error) {

	first := strings.TrimSpace(req.FirstName)
	last := strings.TrimSpace(req.LastName)
	if first == "" {
		return nil, fmt.Errorf("First name is required")
	}
	if last == "" {
		return nil, fmt.Errorf("Last name is required")
	}

	email, err := NormalizeAndValidateEmail(req.Email)
	if err != nil {
		return nil, fmt.Errorf("Invalid email")
	}

	phone := strings.TrimSpace(req.PhoneNumber)
	if phone != "" {
		if _, err := ValidatePhoneNumber(phone); err != nil {
			return nil, fmt.Errorf("Invalid phone number")
		}
	}

	if err := ValidatePassword(req.Password); err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}
	if err := ValidateConfirmPassword(req.Password, req.ConfirmPassword); err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	clean := &types.RegisterUserRequest{
		FirstName:       first,
		LastName:        last,
		Email:           email,
		PhoneNumber:     phone,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
		CreatedAt:       time.Now(),
	}

	return clean, err
}
func ValidateCreateProduct(req *types.CreateProductRequest) (*types.Product, error) {
	name := strings.TrimSpace(req.Name)
	description := strings.TrimSpace(req.Description)
	image := strings.TrimSpace(req.Image)
	price := req.Price
	quantity := req.Quantity

	if name == "" {
		return nil, fmt.Errorf("Product name is required")
	}
	if description == "" {
		return nil, fmt.Errorf("Description is required")
	}

	if req.Image == "" {
		return nil, fmt.Errorf("Image cannot be empty — set a valid image URL or null")
	}
	if price <= 0 {
		price = 0.0
	}

	if quantity <= 0 {
		quantity = 0
	}

	product := &types.Product{
		Name:        name,
		Description: description,
		Image:       image,
		Price:       price,
		Quantity:    quantity,
		CreatedAt:   time.Now(),
	}

	return product, nil
}

func ValidateLoginUser(req types.LoginUserRequest) (*types.LoginUserRequest, error) {
	// Email normalize + validate
	email, err := NormalizeAndValidateEmail(req.Email)
	if err != nil {
		//errs["email"] = "Invalid email"
		return nil, fmt.Errorf("Invalid email")
	}
	if err := ValidatePassword(req.Password); err != nil {
		return nil, fmt.Errorf("%s", err.Error())

	}
	clean := &types.LoginUserRequest{
		Email:    email,
		Password: req.Password,
	}

	return clean, err

}

// normalize and validate
func NormalizeAndValidateEmail(s string) (string, error) {
	e := strings.ToLower(strings.TrimSpace(s))
	if e == "" || !emailRx.MatchString(e) {
		return "", fmt.Errorf("Invalid email")
	}
	return e, nil
}

func ValidatePhoneNumber(s string) (string, error) {
	if s == "" || !phoneNumberRx.MatchString(s) {
		return "", fmt.Errorf("Invalid phone number")
	}
	return s, nil
}

// For []byte password fields
func ValidatePassword(s string) error {
	var (
		hasMinLen = len(s) >= 8
		hasLetter = false
		hasDigit  = false
		hasSymbol = false
	)

	for _, ch := range s {
		switch {
		case unicode.IsLetter(ch):
			hasLetter = true
		case unicode.IsDigit(ch):
			hasDigit = true
		case strings.ContainsRune("@$!%*#?&", ch):
			hasSymbol = true
		}
	}

	if !hasMinLen || !hasLetter || !hasDigit || !hasSymbol {
		return fmt.Errorf("Password must be ≥8 chars, include a letter, a digit, and a symbol (@$!%*#?&)")
	}
	return nil
}

func ValidateConfirmPassword(password, confirm string) error {
	if password != confirm {
		return fmt.Errorf("Passwords do not match")
	}
	return nil
}
