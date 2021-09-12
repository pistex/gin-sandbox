package models

import (
	"time"

	"github.com/google/uuid"
)

// User model.
type User struct {
	ID        uuid.UUID  `db:"id" json:"id"`
	Email     string     `db:"email" json:"email"`
	Password  string     `db:"password" json:"password"`
	IsActive  bool       `db:"is_active" json:"isActive"`
	CreatedAt *time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
}

func NewUser(email string, password string) *User {
	return &User{
		ID:       uuid.New(),
		Email:    email,
		Password: password,
	}
}
