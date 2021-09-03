package interfaces

import (
	"kwanjai/types"

	"github.com/gin-gonic/gin"
)

type IContext interface {
	GetConfig() *types.Config
	GetServer() *gin.Engine
}

func NewContext(config *types.Config, gin *gin.Engine) IContext {
	return &defaultContext{config: config, gin: gin}
}

type defaultContext struct {
	config *types.Config
	gin    *gin.Engine
}

func (c *defaultContext) GetConfig() *types.Config {
	return c.config
}

func (c *defaultContext) GetServer() *gin.Engine {
	return c.gin
}
