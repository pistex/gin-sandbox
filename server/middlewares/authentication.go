package middlewares

import (
	"kwanjai/helpers"
	"kwanjai/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JWTAuthorization middleware.
// Base authentication which always stores user object in Gin context.
// If token verification failed, anonymous user object is stored.
func JWTAuthorization() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		panic("implement me")
	}
}

// AuthenticatedOnly disallows "anonymous" user.
func AuthenticatedOnly(ctx interfaces.IContext) gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		username := helpers.GetUsername(ginContext)
		if username == "anonymous" {
			ginContext.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
