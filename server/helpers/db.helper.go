package helpers

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

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

func initFlyway() error {
	checkFlyway := exec.Command("which", "flyway")
	err := checkFlyway.Run()
	if err != nil {
		downloadFlyway := exec.Command(
			"wget",
			"-qO-",
			"https://repo1.maven.org/maven2/org/flywaydb/flyway-commandline/7.15.0/flyway-commandline-7.15.0-linux-x64.tar.gz",
			"|",
			"tar",
			"xvz",
			"&&",
			"ln", //don't need sudo in a typical linux contianer
			"-s",
			"`pwd`/flyway-7.15.0/flyway",
			"/usr/local/bin ",
		)
		err := downloadFlyway.Run()
		if err != nil {
			return err
		}
		recheckFlyway := exec.Command("which", "flyway")
		err = recheckFlyway.Run()
		if err != nil {
			return errors.New("flyway cannot be initialized")
		}
	}
	return nil
}

func MigrateTestDatabase() error {
	err := initFlyway()
	if err != nil {
		return err
	}

	driver := viper.GetString("SQL_TEST_DRIVER")
	if driver == "postgres" {
		driver = "postgresql"
	} else {
		return fmt.Errorf("unsupported driver %s", viper.GetString("SQL_TEST_DRIVER"))
	}
	migare := exec.Command(
		"flyway",
		fmt.Sprintf("-url=jdbc:%s://%s:%s/%s", driver, viper.GetString("SQL_TEST_HOST"), viper.GetString("SQL_TEST_PORT"), viper.GetString("SQL_TEST_DB_NAME")),
		fmt.Sprintf("-user=%s", viper.GetString("SQL_TEST_USER")),
		fmt.Sprintf("-password=%s", viper.GetString("SQL_TEST_PASSWORD")),
		fmt.Sprintf("-locations=filesystem:%s", viper.GetString("DATABASE_SCHEMA_MIGRATION_LOCATION")),
		"migrate",
	)

	err = migare.Run()
	if err != nil {
		return err
	}

	return nil
}

func CleanTestDatabase() error {
	err := initFlyway()
	if err != nil {
		return err
	}

	driver := viper.GetString("SQL_TEST_DRIVER")
	if driver == "postgres" {
		driver = "postgresql"
	} else {
		return fmt.Errorf("unsupported driver %s", viper.GetString("SQL_TEST_DRIVER"))
	}
	clean := exec.Command(
		"flyway",
		fmt.Sprintf("-url=jdbc:%s://%s:%s/%s", driver, viper.GetString("SQL_TEST_HOST"), viper.GetString("SQL_TEST_PORT"), viper.GetString("SQL_TEST_DB_NAME")),
		fmt.Sprintf("-user=%s", viper.GetString("SQL_TEST_USER")),
		fmt.Sprintf("-password=%s", viper.GetString("SQL_TEST_PASSWORD")),
		fmt.Sprintf("-locations=filesystem:%s", viper.GetString("DATABASE_SCHEMA_MIGRATION_LOCATION")),
		"clean",
	)
	err = clean.Run()
	if err != nil {
		return err
	}

	return nil
}
