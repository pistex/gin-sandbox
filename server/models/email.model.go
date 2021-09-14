package models

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// VerificationEmail model.
type VerificationEmail struct {
	ID        uuid.UUID  `db:"id" json:"id"`
	UserID    uuid.UUID  `db:"user_id" json:"userID"`
	Key       string     `db:"key_id" json:"string"`
	ExpiredAt *time.Time `db:"expired_at" json:"expiredAt"`
}

func NewVerificationEmail(user *User, expiredAt *time.Time) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	email := &VerificationEmail{}
	email.UserID = user.ID
	email.Key = fmt.Sprintf("%06d", random.Intn(999999))
	email.ExpiredAt = expiredAt
}
