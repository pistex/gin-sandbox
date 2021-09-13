package middlewares

import (
	"errors"
	"kwanjai/helpers"
	"kwanjai/interfaces"
	"kwanjai/messages"
	"kwanjai/services"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// JWT middleware.
// Base authentication which always stores user object in Gin context.
// If token verification failed, anonymous user object is stored.
func JWT(ctx interfaces.IContext) gin.HandlerFunc {
	return func(g *gin.Context) {
		authHeader := g.GetHeader("Authorization")
		if authHeader == "" {
			g.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		split := strings.Split(authHeader, " ")
		if len(split) != 2 {
			g.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if split[0] != "Bearer" {
			g.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		auhtService := services.NewAuthService(ctx)
		token, err := auhtService.VerifyToken(split[1])
		if errors.Is(err, messages.ErrLoadPrivateKey) {
			g.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if err != nil {
			g.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claim := &jwt.StandardClaims{}
		err = helpers.JsonMapper(token.Claims, claim)
		if err != nil {
			g.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		err = ctx.Cache().Get(ctx.Config().Context, claim.Id).Err()
		if err != nil {
			g.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		userService := services.NewUserService(ctx)
		id, err := uuid.Parse(claim.Subject)
		if err != nil {
			g.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		user, err := userService.Find(id)
		if err != nil {
			g.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		g.Set("user", user)
		g.Set("tokenID", claim.Id)
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
