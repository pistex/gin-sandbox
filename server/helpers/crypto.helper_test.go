package helpers

import (
	"crypto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashSHA256(t *testing.T) {
	password := "foobar"
	expectedHash := "c3ab8ff13720e8ad9047dd39466b3c8974e592c2fa383d4a3960714caef0c4f2"
	actualHash := HashString(password, crypto.SHA256)
	assert.Equal(t, expectedHash, actualHash)
}
