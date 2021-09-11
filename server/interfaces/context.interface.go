package interfaces

import (
	"kwanjai/types"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type IContext interface {
	Config() *types.Config
	Server() *gin.Engine
	DB() *sqlx.DB
}

func NewContext(config *types.Config, gin *gin.Engine, db *sqlx.DB) IContext {
	return &defaultContext{
		config: config,
		gin:    gin,
		db:     db,
	}
}

type defaultContext struct {
	config *types.Config
	gin    *gin.Engine
	db     *sqlx.DB
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
