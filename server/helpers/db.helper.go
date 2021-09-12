package helpers

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func NewDatabase() (*sqlx.DB, error) {
	if viper.GetString("SQL_DRIVER") == "postgres" {
		conectionInfo := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			viper.GetString("SQL_HOST"),
			viper.GetString("SQL_PORT"),
			viper.GetString("SQL_USER"),
			viper.GetString("SQL_PASSWORD"),
			viper.GetString("SQL_DB_NAME"),
			viper.GetString("POSTGRES_SSL_MODE"),
		)
		db, err := sqlx.Connect("postgres", conectionInfo)
		return db, err
	}

	return nil, fmt.Errorf("unsupported driver %s", os.Getenv("SQL_DRIVER"))
}

func NewTestDatabase() (*sqlx.DB, error) {
	if viper.GetString("SQL_TEST_DRIVER") == "postgres" {
		conectionInfo := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			viper.GetString("SQL_TEST_HOST"),
			viper.GetString("SQL_TEST_PORT"),
			viper.GetString("SQL_TEST_USER"),
			viper.GetString("SQL_TEST_PASSWORD"),
			viper.GetString("SQL_TEST_DB_NAME"),
			viper.GetString("POSTGRES_TEST_SSL_MODE"),
		)
		db, err := sqlx.Connect("postgres", conectionInfo)
		return db, err
	}

	return nil, fmt.Errorf("unsupported driver %s", viper.GetString("SQL_TEST_DRIVER"))
}
