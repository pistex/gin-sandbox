package models

import (
	"time"

	"github.com/google/uuid"
)

// VerificationEmail model.
type VerificationEmail struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Key       string    `json:"string"`
	ExpiredAt time.Time `json:"expiredAt"`
}
