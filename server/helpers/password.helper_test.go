package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidPassword(t *testing.T) {
	password := "password"
	hashedpassword, _ := HashPassword(password)
	assert.Equal(t, CheckPasswordHash("wrongpassword", hashedpassword), false)
}
