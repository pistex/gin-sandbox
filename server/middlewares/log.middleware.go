package middlewares

import (
	"kwanjai/helpers"
	"kwanjai/interfaces"
	"log"

	"github.com/gin-gonic/gin"
)

func RequestLogger(ctx interfaces.IContext) gin.HandlerFunc {
	return func(g *gin.Context) {
		log.Println("ContentType", g.ContentType())
		log.Println("ClientIP", g.ClientIP())
		log.Println("Method", g.Request.Method)
		helpers.LogHTTPHeader(g.Request.Header)
	}
}
