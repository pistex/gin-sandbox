package interfaces

import (
	"kwanjai/types"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type IContext interface {
	Config() *types.Config
	Server() *gin.Engine
	DB() *sqlx.DB
	Cache() *redis.Client
}

func NewContext(config *types.Config, gin *gin.Engine, db *sqlx.DB, cache *redis.Client) IContext {
	return &defaultContext{
		config: config,
		gin:    gin,
		db:     db,
		cache:  cache,
	}
}

type defaultContext struct {
	config *types.Config
	gin    *gin.Engine
	db     *sqlx.DB
	cache  *redis.Client
}

func (c *defaultContext) Config() *types.Config {
	return c.config
}

func (c *defaultContext) Server() *gin.Engine {
	return c.gin
}

func (c *defaultContext) DB() *sqlx.DB {
	return c.db
}

func (c *defaultContext) Cache() *redis.Client {
	return c.cache
}

type mockContext struct {
	config *types.Config
	gin    *gin.Engine
	db     *sqlx.DB
	cache  *redis.Client
	mock   sqlmock.Sqlmock
}

type IMockContext interface {
	IContext
	SQLMock() sqlmock.Sqlmock
}

func NewMockContext(config *types.Config, gin *gin.Engine, mockDBDriver string) IMockContext {
	if mockDBDriver == "postgres" {
		mockDBDriver = "postgresql"
	}
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	return &mockContext{
		config: config,
		gin:    gin,
		db:     sqlx.NewDb(db, mockDBDriver),
		cache:  nil,
		mock:   mock,
	}
}

func (c *mockContext) Config() *types.Config {
	return c.config
}

func (c *mockContext) Server() *gin.Engine {
	return c.gin
}

func (c *mockContext) DB() *sqlx.DB {
	return c.db
}

func (c *mockContext) Cache() *redis.Client {
	return c.cache
}

func (c *mockContext) SQLMock() sqlmock.Sqlmock {
	return c.mock
}
