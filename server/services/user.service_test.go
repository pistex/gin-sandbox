package services

import (
	"database/sql"
	"kwanjai/interfaces"
	"time"

	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestFindUserByEmailSuccess(t *testing.T) {
	email := "jonh@example.com"
	ctx := interfaces.NewMockContext(nil, nil)
	userService := NewUserService(ctx)

	ctx.SQLMock().
		ExpectQuery(`SELECT * FROM users WHERE email = $1;`).
		WithArgs(email).
		WillReturnRows(
			[]*sqlmock.Rows{
				sqlmock.
					NewRows([]string{"id", "password", "email", "is_active", "created_at", "updated_at"}).
					AddRow(uuid.New(), "$2a$10$TWLWCIalpONROKEgJ3nzieogc67Wf7f1QY5MeRVvDm67tYGjn6jS6", email, false, time.Now(), nil),
			}...,
		)
	user, err := userService.FindByEmail(email)
	assert.NotEmpty(t, user)
	assert.NoError(t, err)
}

func TestFindUserByEmailNotFound(t *testing.T) {
	email := "jonh@example.com"
	ctx := interfaces.NewMockContext(nil, nil)
	userService := NewUserService(ctx)

	ctx.SQLMock().
		ExpectQuery(`SELECT * FROM users WHERE email = $1;`).
		WithArgs(email).
		WillReturnError(sql.ErrNoRows)
	user, err := userService.FindByEmail(email)
	assert.Empty(t, user)
	assert.ErrorIs(t, err, sql.ErrNoRows)
}
