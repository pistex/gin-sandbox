package helpers

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func NewDatabase() (*sqlx.DB, error) {
	if ENVGetString("SQL_DRIVER") == "postgres" {
		conectionInfo := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			ENVGetString("SQL_HOST"),
			ENVGetString("SQL_PORT"),
			ENVGetString("SQL_USER"),
			ENVGetString("SQL_PASSWORD"),
			ENVGetString("SQL_DB_NAME"),
			ENVGetString("POSTGRES_SSL_MODE"),
		)
		db, err := sqlx.Connect("postgres", conectionInfo)
		return db, err
	}

	return nil, fmt.Errorf("unsupported driver %s", ENVGetString("SQL_DRIVER"))
}

func NewTestDatabase() (*sqlx.DB, error) {
	if ENVGetString("SQL_TEST_DRIVER") == "postgres" {
		conectionInfo := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			ENVGetString("SQL_TEST_HOST"),
			ENVGetString("SQL_TEST_PORT"),
			ENVGetString("SQL_TEST_USER"),
			ENVGetString("SQL_TEST_PASSWORD"),
			ENVGetString("SQL_TEST_DB_NAME"),
			ENVGetString("POSTGRES_TEST_SSL_MODE"),
		)
		db, err := sqlx.Connect("postgres", conectionInfo)
		return db, err
	}

	return nil, fmt.Errorf("unsupported driver %s", ENVGetString("SQL_TEST_DRIVER"))
}
