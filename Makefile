test-e2e:
	cd server/services && SQL_TEST_DRIVER=postgres \
	SQL_TEST_HOST=0.0.0.0 \
	SQL_TEST_PORT=54321 \
	SQL_TEST_DB_NAME=my_test_database \
	SQL_TEST_USER=my_test_user \
	SQL_TEST_PASSWORD=my_test_password \
	POSTGRES_TEST_SSL_MODE=disable \
	DATABASE_SCHEMA_MIGRATION_LOCATION=../database/migration go test . -v \
	&& cd ../..

