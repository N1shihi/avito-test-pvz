package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegisterResponse struct {
	ID    uuid.UUID `json:"id" db:"id"`
	Email string    `json:"email" db:"email"`
	Role  string    `json:"role" db:"role"`
}

func (u *User) CreateRegisterResponse() *RegisterResponse {
	return &RegisterResponse{
		ID:    u.ID,
		Email: u.Email,
		Role:  u.Role,
	}
}

type User struct {
	ID       uuid.UUID `json:"id" db:"id"`
	Email    string    `json:"email" db:"email"`
	Password string    `json:"-" db:"password"`
	Role     string    `json:"role" db:"role"`
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Role     string `json:"role" validate:"required,oneof=employee moderator"`
}

func (u *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (u *User) HashPassword() error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashed)
	return nil

}
